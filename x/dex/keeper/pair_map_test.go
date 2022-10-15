package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
	"github.com/NicholasDotSol/duality/testutil/nullify"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPairMap(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PairMap {
	items := make([]types.PairMap, n)
	for i := range items {
		items[i].PairId = strconv.Itoa(i)

		keeper.SetPairMap(ctx, items[i])
	}
	return items
}

func TestPairMapGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPairMap(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPairMap(ctx,
			item.PairId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPairMapRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPairMap(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePairMap(ctx,
			item.PairId,
		)
		_, found := keeper.GetPairMap(ctx,
			item.PairId,
		)
		require.False(t, found)
	}
}

func TestPairMapGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPairMap(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPairMap(ctx)),
	)
}
