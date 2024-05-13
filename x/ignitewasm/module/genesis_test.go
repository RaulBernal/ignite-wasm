package ignitewasm_test

import (
	"testing"

	keepertest "ignite-wasm/testutil/keeper"
	"ignite-wasm/testutil/nullify"
	ignitewasm "ignite-wasm/x/ignitewasm/module"
	"ignite-wasm/x/ignitewasm/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgnitewasmKeeper(t)
	ignitewasm.InitGenesis(ctx, k, genesisState)
	got := ignitewasm.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
