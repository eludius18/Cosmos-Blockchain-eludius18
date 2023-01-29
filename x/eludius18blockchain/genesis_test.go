package eludius18blockchain_test

import (
	"testing"

	keepertest "eludius18-blockchain/testutil/keeper"
	"eludius18-blockchain/testutil/nullify"
	"eludius18-blockchain/x/eludius18blockchain"
	"eludius18-blockchain/x/eludius18blockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Eludius18blockchainKeeper(t)
	eludius18blockchain.InitGenesis(ctx, *k, genesisState)
	got := eludius18blockchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
