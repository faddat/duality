package keeper_test

import (
	"context"
	"fmt"
	"testing"

	dualityapp "github.com/NicholasDotSol/duality/app"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type MsgServerTestSuite struct {
	suite.Suite
	app         *dualityapp.App
	msgServer   types.MsgServer
	ctx         sdk.Context
	queryClient types.QueryClient
	alice       sdk.AccAddress
	bob         sdk.AccAddress
	carol       sdk.AccAddress
	dan         sdk.AccAddress
	goCtx       context.Context
	feeTiers    []types.FeeList
}

func TestMsgServerTestSuite(t *testing.T) {
	suite.Run(t, new(MsgServerTestSuite))
}

func (s *MsgServerTestSuite) SetupTest() {
	app := dualityapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.DexKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, s.alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accBob := app.AccountKeeper.NewAccountWithAddress(ctx, s.bob)
	app.AccountKeeper.SetAccount(ctx, accBob)
	accCarol := app.AccountKeeper.NewAccountWithAddress(ctx, s.carol)
	app.AccountKeeper.SetAccount(ctx, accCarol)
	accDan := app.AccountKeeper.NewAccountWithAddress(ctx, s.dan)
	app.AccountKeeper.SetAccount(ctx, accDan)

	// add the fee tiers of 1, 3, 5, 10 ticks
	feeTiers := []types.FeeList{
		{Id: 0, Fee: 1},
		{Id: 1, Fee: 3},
		{Id: 2, Fee: 5},
		{Id: 3, Fee: 10},
	}

	// Set Fee List
	app.DexKeeper.AppendFeeList(ctx, feeTiers[0])
	app.DexKeeper.AppendFeeList(ctx, feeTiers[1])
	app.DexKeeper.AppendFeeList(ctx, feeTiers[2])
	app.DexKeeper.AppendFeeList(ctx, feeTiers[3])

	s.app = app
	s.msgServer = keeper.NewMsgServerImpl(app.DexKeeper)
	s.ctx = ctx
	s.goCtx = sdk.WrapSDKContext(ctx)
	s.queryClient = queryClient
	s.alice = sdk.AccAddress([]byte("alice"))
	s.bob = sdk.AccAddress([]byte("bob"))
	s.carol = sdk.AccAddress([]byte("carol"))
	s.dan = sdk.AccAddress([]byte("dan"))
	s.feeTiers = feeTiers
}

func (s *MsgServerTestSuite) fundAccountBalancesDec(account sdk.AccAddress, aBalance sdk.Dec, bBalance sdk.Dec) {
	aBalanceInt := sdk.NewIntFromBigInt(aBalance.BigInt())
	bBalanceInt := sdk.NewIntFromBigInt(bBalance.BigInt())
	balances := sdk.NewCoins(NewACoin(aBalanceInt), NewBCoin(bBalanceInt))
	err := FundAccount(s.app.BankKeeper, s.ctx, account, balances)
	s.Assert().NoError(err)
	s.assertAccountBalancesDec(account, aBalance, bBalance)
}

func (s *MsgServerTestSuite) fundAccountBalances(account sdk.AccAddress, aBalance int, bBalance int) {
	s.fundAccountBalancesDec(account, NewDec(aBalance), NewDec(bBalance))
}

func (s *MsgServerTestSuite) fundAliceBalances(a int, b int) {
	s.fundAccountBalances(s.alice, a, b)
}

func (s *MsgServerTestSuite) fundAliceBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.fundAccountBalancesDec(s.alice, a, b)
}

func (s *MsgServerTestSuite) fundBobBalances(a int, b int) {
	s.fundAccountBalances(s.bob, a, b)
}

func (s *MsgServerTestSuite) fundBobBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.fundAccountBalancesDec(s.bob, a, b)
}

func (s *MsgServerTestSuite) fundCarolBalances(a int, b int) {
	s.fundAccountBalances(s.carol, a, b)
}

func (s *MsgServerTestSuite) fundCarolBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.fundAccountBalancesDec(s.carol, a, b)
}

func (s *MsgServerTestSuite) fundDanBalances(a int, b int) {
	s.fundAccountBalances(s.dan, a, b)
}

func (s *MsgServerTestSuite) fundDanBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.fundAccountBalancesDec(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertAccountBalances(account sdk.AccAddress, aBalance int, bBalance int) {
	s.assertAccountBalancesDec(account, NewDec(aBalance), NewDec(bBalance))
}

func (s *MsgServerTestSuite) assertAccountBalancesDec(
	account sdk.AccAddress,
	aBalance sdk.Dec,
	bBalance sdk.Dec,
) {
	aActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenA")
	aDec := sdk.NewDecFromBigIntWithPrec(aActual.Amount.BigInt(), 18)
	s.Assert().Equal(aBalance, aDec)

	bActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenB")
	bDec := sdk.NewDecFromBigIntWithPrec(bActual.Amount.BigInt(), 18)
	s.Assert().Equal(bBalance, bDec)
}

func (s *MsgServerTestSuite) assertAliceBalances(a int, b int) {
	s.assertAccountBalances(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertAliceBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertBobBalances(a int, b int) {
	s.assertAccountBalances(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertBobBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertCarolBalances(a int, b int) {
	s.assertAccountBalances(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertCarolBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertDanBalances(a int, b int) {
	s.assertAccountBalances(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDanBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDexBalances(a int, b int) {
	s.assertAccountBalances(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
}

func (s *MsgServerTestSuite) assertDexBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
}

func (s *MsgServerTestSuite) aliceLimitBuys(buying string, tick int, amountIn int) {
	s.limitBuys(s.alice, buying, tick, amountIn)
}

func (s *MsgServerTestSuite) bobLimitBuys(buying string, tick int, amountIn int) {
	s.limitBuys(s.bob, buying, tick, amountIn)
}

func (s *MsgServerTestSuite) carolLimitBuys(buying string, tick int, amountIn int) {
	s.limitBuys(s.carol, buying, tick, amountIn)
}

func (s *MsgServerTestSuite) danLimitBuys(buying string, tick int, amountIn int) {
	s.limitBuys(s.dan, buying, tick, amountIn)
}

func (s *MsgServerTestSuite) aliceLimitSells(selling string, tick int, amountIn int) {
	s.limitSells(s.alice, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) bobLimitSells(selling string, tick int, amountIn int) {
	s.limitSells(s.bob, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) carolLimitSells(selling string, tick int, amountIn int) {
	s.limitSells(s.carol, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) danLimitSells(selling string, tick int, amountIn int) {
	s.limitSells(s.dan, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) limitBuys(account sdk.AccAddress, tokenOut string, tick int, amountIn int) {
	var tokenIn string
	if tokenOut == "TokenA" {
		tokenIn = "TokenB"
	} else {
		tokenIn = "TokenA"
	}
	s.limitSells(account, tokenIn, tick, amountIn)
}

func (s *MsgServerTestSuite) limitSells(account sdk.AccAddress, tokenIn string, tick int, amountIn int) {
	amountInDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(amountIn)))
	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		TokenIn:   tokenIn,
		AmountIn:  amountInDec,
	})
	s.Assert().Nil(err)
}

type Deposit struct {
	AmountA   sdk.Dec
	AmountB   sdk.Dec
	TickIndex int64
	FeeIndex  uint64
}

func NewDeposit(amountA int, amountB int, tickIndex int, feeIndex int) *Deposit {
	return &Deposit{
		AmountA:   sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(amountA))),
		AmountB:   sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(amountB))),
		TickIndex: int64(tickIndex),
		FeeIndex:  uint64(feeIndex),
	}
}

func (s *MsgServerTestSuite) aliceDeposits(deposits ...*Deposit) {
	s.deposits(s.alice, deposits...)
}

func (s *MsgServerTestSuite) bobDeposits(deposits ...*Deposit) {
	s.deposits(s.bob, deposits...)
}

func (s *MsgServerTestSuite) carolDeposits(deposits ...*Deposit) {
	s.deposits(s.carol, deposits...)
}

func (s *MsgServerTestSuite) danDeposits(deposits ...*Deposit) {
	s.deposits(s.dan, deposits...)
}

func (s *MsgServerTestSuite) deposits(account sdk.AccAddress, deposits ...*Deposit) {
	amountsA := make([]sdk.Dec, len(deposits))
	amountsB := make([]sdk.Dec, len(deposits))
	tickIndexes := make([]int64, len(deposits))
	feeIndexes := make([]uint64, len(deposits))
	for i, e := range deposits {
		amountsA[i] = e.AmountA
		amountsB[i] = e.AmountB
		tickIndexes[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
	}

	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     account.String(),
		Receiver:    account.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    amountsA,
		AmountsB:    amountsB,
		TickIndexes: tickIndexes,
		FeeIndexes:  feeIndexes,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) assertDepositReponse(depositResponse DepositReponse, expectedDepositResponse DepositReponse) {

	for i, _ := range expectedDepositResponse.amountsA {
		s.Assert().Equal(depositResponse.amountsA[i], expectedDepositResponse.amountsA[i])
	}

	for i, _ := range expectedDepositResponse.amountsB {
		s.Assert().Equal(depositResponse.amountsB[i], expectedDepositResponse.amountsB[i])
	}
}

type DepositReponse struct {
	amountsA []sdk.Dec
	amountsB []sdk.Dec
}

type Withdrawl struct {
	TickIndex      int64
	FeeIndex       uint64
	SharesToRemove sdk.Dec
}

func NewWithdrawl(tickIndex int, feeIndex int, sharesToRemove int) *Withdrawl {
	return &Withdrawl{
		TickIndex:      int64(tickIndex),
		FeeIndex:       uint64(feeIndex),
		SharesToRemove: sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(sharesToRemove))),
	}
}

func (s *MsgServerTestSuite) withdraws(account sdk.AccAddress, withdrawls ...*Withdrawl) {
	tickIndexes := make([]int64, len(withdrawls))
	feeIndexes := make([]uint64, len(withdrawls))
	sharesToRemove := make([]sdk.Dec, len(withdrawls))
	for i, e := range withdrawls {
		tickIndexes[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
		sharesToRemove[i] = e.SharesToRemove
	}

	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        account.String(),
		Receiver:       account.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: sharesToRemove,
		TickIndexes:    tickIndexes,
		FeeIndexes:     feeIndexes,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceWithdraws(withdrawals ...*Withdrawl) {
	s.withdraws(s.alice, withdrawals...)
}

func (s *MsgServerTestSuite) bobWithdraws(withdrawals ...*Withdrawl) {
	s.withdraws(s.bob, withdrawals...)
}

func (s *MsgServerTestSuite) carolWithdraws(withdrawals ...*Withdrawl) {
	s.withdraws(s.carol, withdrawals...)
}

func (s *MsgServerTestSuite) danWithdraws(withdrawals ...*Withdrawl) {
	s.withdraws(s.dan, withdrawals...)
}

func (s *MsgServerTestSuite) aliceCancelsLimitSell(keyToken string, tick int, key int, sharesOut int) {
	s.cancelsLimitSell(s.alice, keyToken, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) bobCancelsLimitSell(keyToken string, tick int, key int, sharesOut int) {
	s.cancelsLimitSell(s.bob, keyToken, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) carolCancelsLimitSell(keyToken string, tick int, key int, sharesOut int) {
	s.cancelsLimitSell(s.carol, keyToken, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) danCancelsLimitSell(keyToken string, tick int, key int, sharesOut int) {
	s.cancelsLimitSell(s.dan, keyToken, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) cancelsLimitSell(account sdk.AccAddress, selling string, tick int, key int, sharesOut int) {
	sharesOutDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(sharesOut)))
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		KeyToken:  selling,
		Key:       uint64(key),
		SharesOut: sharesOutDec,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) cancelsLimitBuy(account sdk.AccAddress, buying string, tick int, key int, sharesOut int) {
	var selling string
	if buying == "TokenA" {
		selling = "TokenB"
	} else {
		selling = "TokenA"
	}
	s.cancelsLimitSell(account, selling, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) aliceCancelsLimitBuy(buying string, tick int, key int, sharesOut int) {
	s.cancelsLimitBuy(s.alice, buying, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) bobCancelsLimitBuy(buying string, tick int, key int, sharesOut int) {
	s.cancelsLimitBuy(s.bob, buying, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) carolCancelsLimitBuy(buying string, tick int, key int, sharesOut int) {
	s.cancelsLimitBuy(s.carol, buying, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) danCancelsLimitBuy(buying string, tick int, key int, sharesOut int) {
	s.cancelsLimitBuy(s.dan, buying, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) aliceBuys(buying string, amountIn int, minOut int) {
	s.buys(s.alice, buying, amountIn, minOut)
}

func (s *MsgServerTestSuite) bobBuys(buying string, amountIn int, minOut int) {
	s.buys(s.bob, buying, amountIn, minOut)
}

func (s *MsgServerTestSuite) carolBuys(buying string, amountIn int, minOut int) {
	s.buys(s.bob, buying, amountIn, minOut)
}

func (s *MsgServerTestSuite) danBuys(buying string, amountIn int, minOut int) {
	s.buys(s.bob, buying, amountIn, minOut)
}

func (s *MsgServerTestSuite) aliceSells(selling string, amountIn int, minOut int) {
	s.sells(s.alice, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) bobSells(selling string, amountIn int, minOut int) {
	s.sells(s.bob, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) carolSells(selling string, amountIn int, minOut int) {
	s.sells(s.bob, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) danSells(selling string, amountIn int, minOut int) {
	s.sells(s.bob, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) buys(account sdk.AccAddress, buying string, amountIn int, minOut int) {
	var selling string
	if buying == "TokenA" {
		selling = "TokenB"
	} else {
		selling = "TokenA"
	}
	s.sells(account, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) sells(account sdk.AccAddress, selling string, amountIn int, minOut int) {
	amountInDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(amountIn)))
	minOutDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(minOut)))
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  account.String(),
		Receiver: account.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		TokenIn:  selling,
		AmountIn: amountInDec,
		MinOut:   minOutDec,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceWithdrawsLimitBuy(buying string, tick int) {
	s.withdrawsLimitBuy(s.alice, buying, tick)
}

func (s *MsgServerTestSuite) bobWithdrawsLimitBuy(buying string, tick int) {
	s.withdrawsLimitBuy(s.bob, buying, tick)
}

func (s *MsgServerTestSuite) carolWithdrawsLimitBuy(buying string, tick int) {
	s.withdrawsLimitBuy(s.carol, buying, tick)
}

func (s *MsgServerTestSuite) danWithdrawsLimitBuy(buying string, tick int) {
	s.withdrawsLimitBuy(s.dan, buying, tick)
}

func (s *MsgServerTestSuite) aliceWithdrawsLimitSell(selling string, tick int) {
	s.withdrawsLimitSell(s.alice, selling, tick)
}

func (s *MsgServerTestSuite) bobWithdrawsLimitSell(selling string, tick int) {
	s.withdrawsLimitSell(s.bob, selling, tick)
}

func (s *MsgServerTestSuite) carolWithdrawsLimitSell(selling string, tick int) {
	s.withdrawsLimitSell(s.carol, selling, tick)
}

func (s *MsgServerTestSuite) danWithdrawsLimitSell(selling string, tick int) {
	s.withdrawsLimitSell(s.dan, selling, tick)
}

func (s *MsgServerTestSuite) withdrawsLimitSell(account sdk.AccAddress, selling string, tick int) {
	var buying string
	if selling == "TokenA" {
		buying = "TokenB"
	} else {
		buying = "TokenA"
	}
	s.withdrawsLimitBuy(account, buying, tick)
}

func (s *MsgServerTestSuite) withdrawsLimitBuy(account sdk.AccAddress, buying string, tick int) {
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		KeyToken:  buying,
		Key:       0,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) traceBalances() {
	aliceA := s.app.BankKeeper.GetBalance(s.ctx, s.alice, "TokenA")
	aliceB := s.app.BankKeeper.GetBalance(s.ctx, s.alice, "TokenB")
	bobA := s.app.BankKeeper.GetBalance(s.ctx, s.bob, "TokenA")
	bobB := s.app.BankKeeper.GetBalance(s.ctx, s.bob, "TokenB")
	carolA := s.app.BankKeeper.GetBalance(s.ctx, s.carol, "TokenA")
	carolB := s.app.BankKeeper.GetBalance(s.ctx, s.carol, "TokenB")
	danA := s.app.BankKeeper.GetBalance(s.ctx, s.dan, "TokenA")
	danB := s.app.BankKeeper.GetBalance(s.ctx, s.dan, "TokenB")
	fmt.Printf(
		"Alice: %+v %+v\nBob: %+v %+v\nCarol: %+v %+v\nDan: %+v %+v",
		aliceA, aliceB,
		bobA, bobB,
		carolA, carolB,
		danA, danB,
	)
}

type Withdrawl struct {
	TickIndex int64
	FeeIndex  uint64
	Shares    sdk.Dec
}

func NewWithdrawl(shares int64, tick int64, fee uint64) *Withdrawl {
	return &Withdrawl{
		Shares:    sdk.NewDec(shares),
		FeeIndex:  fee,
		TickIndex: tick,
	}
}

func (s *MsgServerTestSuite) aliceWithdraws(withdrawls ...*Withdrawl) error {
	sharesToRemove := make([]sdk.Dec, len(withdrawls))
	tickIndicies := make([]int64, len(withdrawls))
	feeIndexes := make([]uint64, len(withdrawls))
	for i, e := range withdrawls {
		sharesToRemove[i] = e.Shares
		tickIndicies[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
	}

	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: sharesToRemove,
		TickIndexes:    tickIndicies,
		FeeIndexes:     feeIndexes,
	})

	return err
}

func (s *MsgServerTestSuite) getShares(
	account sdk.AccAddress,
	pairId string,
	tick int64,
	fee uint64,
) (shares sdk.Dec) {

	sharesData, sharesFound := s.app.DexKeeper.GetShares(s.ctx, account.String(), pairId, tick, fee)
	s.Assert().True(sharesFound)
	return sharesData.SharesOwned
}

func (s *MsgServerTestSuite) assertAccountShares(
	account sdk.AccAddress,
	pairId string,
	tick int64,
	fee uint64,
	sharesExpected sdk.Dec,
) {
	sharesOwned := s.getShares(account, pairId, tick, fee)
	s.Assert().Equal(sharesExpected, sharesOwned)
}

func (s *MsgServerTestSuite) assertAliceShares(
	tick int64,
	fee uint64,
	sharesExpected sdk.Dec) {

	s.assertAccountShares(s.alice, "TokenA/TokenB", tick, fee, sharesExpected)
}

func (s *MsgServerTestSuite) assertCurrentTicks(
	expected0To1 int64,
	expected1To0 int64,
) {
	tickMap, found := s.app.DexKeeper.GetPairMap(s.ctx, "TokenA/TokenB")
	s.Assert().NotNil(found)
	s.Assert().Equal(expected0To1, tickMap.TokenPair.CurrentTick0To1)
	s.Assert().Equal(expected1To0, tickMap.TokenPair.CurrentTick1To0)
}

func (s *MsgServerTestSuite) printTicks() {
	tickMap, _ := s.app.DexKeeper.GetPairMap(s.ctx, "TokenA/TokenB")
	fmt.Printf("\nTick0To1: %v, Tick1To0: %v", tickMap.TokenPair.CurrentTick0To1, tickMap.TokenPair.CurrentTick1To0)
}
