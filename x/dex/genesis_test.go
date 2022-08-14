package dex_test

import (
	"testing"

	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
	"github.com/NicholasDotSol/duality/testutil/nullify"
	"github.com/NicholasDotSol/duality/x/dex"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		NodesList: []types.Nodes{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		NodesCount: 2,
		VirtualPriceTickQueueList: []types.VirtualPriceTickQueue{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VirtualPriceTickQueueCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, *k, genesisState)
	got := dex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.NodesList, got.NodesList)
	require.Equal(t, genesisState.NodesCount, got.NodesCount)
	require.ElementsMatch(t, genesisState.VirtualPriceTickQueueList, got.VirtualPriceTickQueueList)
	require.Equal(t, genesisState.VirtualPriceTickQueueCount, got.VirtualPriceTickQueueCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
