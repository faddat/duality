package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/dex/types"
)

func TestPoolQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPools(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryPoolRequest
		response *types.QueryPoolResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryPoolRequest{
				PairID:    "TokenA<>TokenB",
				TickIndex: msgs[0].TickIndex,
				Fee:       msgs[0].Fee,
			},
			response: &types.QueryPoolResponse{Pool: msgs[0].Pool},
		},
		{
			desc: "Second",
			request: &types.QueryPoolRequest{
				PairID:    "TokenA<>TokenB",
				TickIndex: msgs[1].TickIndex,
				Fee:       msgs[1].Fee,
			},
			response: &types.QueryPoolResponse{Pool: msgs[1].Pool},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryPoolRequest{
				PairID:    "TokenA<>TokenB",
				TickIndex: 0,
				Fee:       100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Pool(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
