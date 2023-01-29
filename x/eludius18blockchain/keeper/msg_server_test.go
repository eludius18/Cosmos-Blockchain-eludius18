package keeper_test

import (
	"context"
	"testing"

	keepertest "eludius18-blockchain/testutil/keeper"
	"eludius18-blockchain/x/eludius18blockchain/keeper"
	"eludius18-blockchain/x/eludius18blockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Eludius18blockchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
