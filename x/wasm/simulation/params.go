package simulation

import (
	"fmt"
	"math/rand"

	"github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func ParamChanges(r *rand.Rand, cdc codec.Codec) []simtypes.ParamChange {
	params := RandomParams(r)
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.ParamStoreKeyUploadAccess),
			func(r *rand.Rand) string {
				jsonBz, err := cdc.MarshalJSON(&params.CodeUploadAccess)
				if err != nil {
					panic(err)
				}
				return string(jsonBz)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.ParamStoreKeyInstantiateAccess),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%q", params.CodeUploadAccess.Permission.String())
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.ParamStoreKeyMaxWasmCodeSize),
			func(r *rand.Rand) string {
				return fmt.Sprintf(`"%d"`, params.MaxWasmCodeSize)
			},
		),
	}
}

func RandomParams(r *rand.Rand) types.Params {
	permissionType := types.AccessType(simtypes.RandIntBetween(r, 1, 3))
	account, _ := simtypes.RandomAcc(r, simtypes.RandomAccounts(r, 10))
	accessConfig := permissionType.With(account.Address)
	return types.Params{
		CodeUploadAccess:             accessConfig,
		InstantiateDefaultPermission: accessConfig.Permission,
		MaxWasmCodeSize:              uint64(simtypes.RandIntBetween(r, 1, 600) * 1024),
	}
}
