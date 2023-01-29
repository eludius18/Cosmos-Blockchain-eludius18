package keeper_test

import (
	"testing"

	testkeeper "eludius18-blockchain/testutil/keeper"
	"eludius18-blockchain/x/eludius18blockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Eludius18blockchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
