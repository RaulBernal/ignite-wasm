package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "ignite-wasm/testutil/keeper"
	"ignite-wasm/x/ignitewasm/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.IgnitewasmKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
