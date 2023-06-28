package network

import (
	"fmt"
	"testing"
	"time"

	dbm "github.com/cometbft/cometbft-db"
	tmrand "github.com/cometbft/cometbft/libs/rand"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	pruningtypes "github.com/cosmos/cosmos-sdk/store/pruning/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	genutil "github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"

	"github.com/duality-labs/duality/app"
	appparams "github.com/duality-labs/duality/app/params"
)

type (
	Network = network.Network
	Config  = network.Config
)

// New creates instance with fully configured cosmos network.
// Accepts optional config, that will be used in place of the DefaultConfig() if provided.
func New(t *testing.T, configs ...network.Config) *network.Network {
	if len(configs) > 1 {
		panic("at most one config should be provided")
	}
	var cfg network.Config
	if len(configs) == 0 {
		cfg = DefaultConfig()
	} else {
		cfg = configs[0]
	}
	net, err := network.New(t, t.TempDir(), cfg)
	require.NoError(t, err)
	t.Cleanup(net.Cleanup)

	return net
}

// DefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
func DefaultConfig() network.Config {
	// app doesn't have this modules anymore, but we need them for test setup, which uses gentx and MsgCreateValidator
	app.ModuleBasics[genutiltypes.ModuleName] = genutil.AppModuleBasic{}
	app.ModuleBasics[stakingtypes.ModuleName] = staking.AppModuleBasic{}

	encoding := appparams.MakeTestEncodingConfig()

	return network.Config{
		Codec:             encoding.Marshaler,
		TxConfig:          encoding.TxConfig,
		LegacyAmino:       encoding.Amino,
		InterfaceRegistry: encoding.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val network.ValidatorI) servertypes.Application {
			// err := modifyConsumerGenesis(val)
			// if err != nil {
			// 	panic(err)
			// }

			return app.New(
				val.GetCtx().Logger, dbm.NewMemDB(), nil, true, map[int64]bool{}, val.GetCtx().Config.RootDir, 0,
				encoding,
				app.EmptyAppOptions{},
				baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.GetAppConfig().Pruning)),
				baseapp.SetMinGasPrices(val.GetAppConfig().MinGasPrices),
			)
		},
		GenesisState:  app.ModuleBasics.DefaultGenesis(encoding.Marshaler),
		TimeoutCommit: 2 * time.Second,
		ChainID:       "chain-" + tmrand.NewRand().Str(6),
		// Some changes are introduced to make the tests run as if Duality is a standalone chain.
		// This will only work if NumValidators is set to 1.
		NumValidators:   1,
		BondDenom:       sdk.DefaultBondDenom,
		MinGasPrices:    fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
		AccountTokens:   sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
		StakingTokens:   sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
		BondedTokens:    sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy: pruningtypes.PruningOptionNothing,
		CleanupDir:      true,
		SigningAlgo:     string(hd.Secp256k1Type),
		KeyringOptions:  []keyring.Option{},
	}
}

// TODO: safe to delete???
// func modifyConsumerGenesis(val network.ValidatorI) error {
// 	genFile := val.GetCtx().Config.GenesisFile()
// 	appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
// 	if err != nil {
// 		return sdkerrors.Wrap(err, "failed to read genesis from the file")
// 	}

// 	tmProtoPublicKey, err := cryptocodec.ToTmProtoPublicKey(val.GetCtx().)
// 	if err != nil {
// 		return sdkerrors.Wrap(err, "invalid public key")
// 	}

// 	initialValset := []types1.ValidatorUpdate{{PubKey: tmProtoPublicKey, Power: 100}}
// 	vals, err := tmtypes.PB2TM.ValidatorUpdates(initialValset)
// 	if err != nil {
// 		return sdkerrors.Wrap(err, "could not convert val updates to validator set")
// 	}

// 	consumerGenesisState := testutil.CreateMinimalConsumerTestGenesis()
// 	consumerGenesisState.InitialValSet = initialValset
// 	consumerGenesisState.ProviderConsensusState.NextValidatorsHash = tmtypes.NewValidatorSet(vals).Hash()

// 	if err := consumerGenesisState.Validate(); err != nil {
// 		return sdkerrors.Wrap(err, "invalid consumer genesis")
// 	}

// 	consumerGenStateBz, err := val.ClientCtx.Codec.MarshalJSON(consumerGenesisState)
// 	if err != nil {
// 		return sdkerrors.Wrap(err, "failed to marshal consumer genesis state into JSON")
// 	}

// 	appState[ccvconsumertypes.ModuleName] = consumerGenStateBz
// 	appStateJSON, err := json.Marshal(appState)
// 	if err != nil {
// 		return sdkerrors.Wrap(err, "failed to marshal application genesis state into JSON")
// 	}

// 	genDoc.AppState = appStateJSON
// 	err = genutil.ExportGenesisFile(genDoc, genFile)
// 	if err != nil {
// 		return sdkerrors.Wrap(err, "failed to export genesis state")
// 	}

// 	return nil
// }

// New creates instance with fully configured cosmos network.
// Accepts optional config, that will be used in place of the DefaultConfig() if provided.
func NewCLITest(t *testing.T, configs ...network.Config) *network.Network {
	if len(configs) > 1 {
		panic("at most one config should be provided")
	}
	var cfg network.Config
	if len(configs) == 0 {
		cfg = DefaultConfigCLITest()
	} else {
		cfg = configs[0]
	}
	net, err := network.New(t, t.TempDir(), cfg)
	require.NoError(t, err)
	t.Cleanup(net.Cleanup)

	return net
}

// DefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
func DefaultConfigCLITest() network.Config {
	// app doesn't have this modules anymore, but we need them for test setup, which uses gentx and MsgCreateValidator
	app.ModuleBasics[genutiltypes.ModuleName] = genutil.AppModuleBasic{}
	app.ModuleBasics[stakingtypes.ModuleName] = staking.AppModuleBasic{}

	encoding := appparams.MakeTestEncodingConfig()

	return network.Config{
		Codec:             encoding.Marshaler,
		TxConfig:          encoding.TxConfig,
		LegacyAmino:       encoding.Amino,
		InterfaceRegistry: encoding.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val network.ValidatorI) servertypes.Application {
			// err := modifyConsumerGenesis(val)
			// if err != nil {
			// 	panic(err)
			// }

			// err := modifyConsumerGenesisCLITestSetup(val)
			// if err != nil {
			// 	panic(err)
			// }

			return app.New(
				val.GetCtx().Logger, dbm.NewMemDB(), nil, true, map[int64]bool{}, val.GetCtx().Config.RootDir, 0,
				encoding,
				app.EmptyAppOptions{},
				baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.GetAppConfig().Pruning)),
				baseapp.SetMinGasPrices(val.GetAppConfig().MinGasPrices),
			)
		},
		GenesisState:  app.ModuleBasics.DefaultGenesis(encoding.Marshaler),
		TimeoutCommit: 2 * time.Second,
		ChainID:       "chain-" + tmrand.NewRand().Str(6),
		// Some changes are introduced to make the tests run as if Duality is a standalone chain.
		// This will only work if NumValidators is set to 1.
		NumValidators:   1,
		BondDenom:       sdk.DefaultBondDenom,
		MinGasPrices:    fmt.Sprintf("0.000000006%s", sdk.DefaultBondDenom),
		AccountTokens:   sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
		StakingTokens:   sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
		BondedTokens:    sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy: pruningtypes.PruningOptionNothing,
		CleanupDir:      true,
		SigningAlgo:     string(hd.Secp256k1Type),
		KeyringOptions:  []keyring.Option{},
	}
}

// func modifyConsumerGenesisCLITestSetup(val network.Validator) error {
// 	genFile := val.Ctx.Config.GenesisFile()
// 	appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
// 	if err != nil {
// 		return err
// 	}

// 	// Modify the data structure here
// 	dexData := appState[dextypes.ModuleName]
// 	var dexGenesisState dextypes.GenesisState
// 	json.Unmarshal(dexData, &dexGenesisState)

// 	newRawJSON, _ := json.Marshal(dexGenesisState)
// 	appState[dextypes.ModuleName] = newRawJSON

// 	bankData := appState[banktypes.ModuleName]
// 	var bankGenesisState banktypes.GenesisState
// 	json.Unmarshal(bankData, &bankGenesisState)

// 	bankGenesisState.Balances = []banktypes.Balance{
// 		{
// 			Address: val.Address.String(),
// 			Coins: sdk.Coins{
// 				sdk.NewCoin("TokenA", sdk.NewInt(100000000)),
// 				sdk.NewCoin("TokenB", sdk.NewInt(100000000)),
// 				sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000)),
// 			},
// 		},
// 	}
// 	newRawJSON, _ = json.Marshal(bankGenesisState)
// 	appState[banktypes.ModuleName] = newRawJSON

// 	appStateJSON, _ := json.Marshal(appState)
// 	// Write the modified app state to the genesis file
// 	genDoc.AppState = json.RawMessage(appStateJSON)

// 	err = genutil.ExportGenesisFile(genDoc, genFile)
// 	if err != nil {
// 		return sdkerrors.Wrap(err, "failed to export genesis state")
// 	}

// 	return nil
// }
