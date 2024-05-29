package app

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/module"

	storetypes "cosmossdk.io/store/types"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	// consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
)

// RegisterUpgradeHandlers registers upgrade handlers.

func (app App) RegisterUpgradeHandlers() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	app.StickyFingers(upgradeInfo)
}
func (app *App) StickyFingers(_ upgradetypes.Plan) {
	planName := "stickyfingers-wasm"
	app.UpgradeKeeper.SetUpgradeHandler(
		planName,
		func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			app.Logger().Info("Cosmos-SDK v0.50.x is here...")

			return nil, nil //app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)

		},
	)
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == planName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				// circuittypes.ModuleName,
				// ibcfeetypes.ModuleName,
				// nft.ModuleName,
				// consensusparamtypes.StoreKey,
				wasmtypes.ModuleName,
			},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

}
