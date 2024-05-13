package keeper

import (
	"ignite-wasm/x/ignitewasm/types"
)

var _ types.QueryServer = Keeper{}
