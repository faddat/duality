package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

type Liquidity interface {
	Swap(maxAmount sdk.Int) (inAmount sdk.Int, outAmount sdk.Int)
	Price() *types.Price
}

func NewLiquidityIterator(
	keeper Keeper,
	ctx sdk.Context,
	tradingPair types.DirectionalTradingPair,
) *LiquidityIterator {

	return &LiquidityIterator{
		iter:   keeper.NewTickIterator(ctx, tradingPair.PairId, tradingPair.TokenOut),
		keeper: keeper,
		ctx:    ctx,
		pairId: tradingPair.PairId,
		is0To1: tradingPair.IsTokenInToken0(),
	}

}

type LiquidityIterator struct {
	keeper Keeper
	pairId *types.PairId
	ctx    sdk.Context
	iter   TickIterator
	is0To1 bool
}

func (s *LiquidityIterator) Next() Liquidity {
	// Move iterator to the next tick after each call
	// iter must be in valid state to call next
	defer func() {
		if s.iter.Valid() {
			s.iter.Next()
		}
	}()

	for ; s.iter.Valid(); s.iter.Next() {
		tick := s.iter.Value()
		switch liquidity := tick.Liquidity.(type) {
		case *types.TickLiquidity_PoolReserves:
			var err error
			var pool Liquidity
			poolReserves := *liquidity.PoolReserves
			if s.is0To1 {
				//Pool Reserves is upperTick
				pool, err = s.createPool0To1(poolReserves)
			} else {
				//Pool Reserves is is lowerTick
				pool, err = s.createPool1To0(poolReserves)
			}
			// TODO: we are not actually handling the error here we're just stopping iteration
			// Should be a very rare edge case where the opposing tick is initialized
			// above/below the Min/Max tick limit
			if err != nil {
				return nil
			}
			return pool

		case *types.TickLiquidity_LimitOrderTranche:
			tranche := liquidity.LimitOrderTranche
			// If we hit a tranche with an expired goodTil date keep iterating
			if tranche.IsExpired(s.ctx) {
				continue
			} else {
				return tranche
			}

		default:
			panic("Tick does not have liquidity")

		}
	}
	return nil
}

func (s *LiquidityIterator) createPool0To1(upperTick types.PoolReserves) (Liquidity, error) {
	upperTickIndex := upperTick.TickIndex
	centerTickIndex := upperTickIndex - utils.MustSafeUint64(upperTick.Fee)
	lowerTickIndex := centerTickIndex - utils.MustSafeUint64(upperTick.Fee)
	lowerTick, err := s.keeper.GetOrInitPoolReserves(s.ctx, s.pairId, s.pairId.Token0, lowerTickIndex, upperTick.Fee)
	if err != nil {
		return nil, err
	}
	pool := NewPool(
		centerTickIndex,
		lowerTick,
		&upperTick,
	)
	return NewLiquidityFromPool0To1(&pool), nil
}

func (s *LiquidityIterator) createPool1To0(lowerTick types.PoolReserves) (Liquidity, error) {
	lowerTickIndex := lowerTick.TickIndex
	centerTickIndex := lowerTickIndex + utils.MustSafeUint64(lowerTick.Fee)
	upperTickIndex := centerTickIndex + utils.MustSafeUint64(lowerTick.Fee)
	upperTick, err := s.keeper.GetOrInitPoolReserves(s.ctx, s.pairId, s.pairId.Token1, upperTickIndex, lowerTick.Fee)
	if err != nil {
		return nil, err
	}

	pool := NewPool(
		centerTickIndex,
		&lowerTick,
		upperTick,
	)
	return NewLiquidityFromPool1To0(&pool), nil
}

func (s *LiquidityIterator) Close() {
	s.iter.Close()
}

func (k Keeper) SaveLiquidity(sdkCtx sdk.Context, liquidityI Liquidity) {
	switch liquidity := liquidityI.(type) {
	case *types.LimitOrderTranche:
		k.SaveTranche(sdkCtx, *liquidity)

	case *PoolLiquidity:
		k.SavePool(sdkCtx, *liquidity.pool)
	default:
		panic("Invalid liquidity type")
	}

}

func (k Keeper) Swap(ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tokenOut string,
	amountIn sdk.Int,
	limitPrice *sdk.Dec) (totalIn sdk.Int, totalOut sdk.Int, error error) {
	cacheCtx, writeCache := ctx.CacheContext()
	pair := types.NewDirectionalTradingPair(pairId, tokenIn, tokenOut)

	remainingIn := amountIn
	totalOut = sdk.ZeroInt()

	// verify that amount left is not zero and that there are additional valid ticks to check
	liqIter := NewLiquidityIterator(k, ctx, pair)
	defer liqIter.Close()
	for remainingIn.GT(sdk.ZeroInt()) {
		liq := liqIter.Next()
		if liq == nil {
			break
		}

		// break as soon as we iterated past limitPrice
		if limitPrice != nil && liq.Price().ToDec().LT(*limitPrice) {
			break

		}

		inAmount, outAmount := liq.Swap(remainingIn)

		remainingIn = remainingIn.Sub(inAmount)
		totalOut = totalOut.Add(outAmount)

		k.SaveLiquidity(cacheCtx, liq)
	}

	writeCache()
	totalIn = amountIn.Sub(remainingIn)
	return totalIn, totalOut, nil
}
