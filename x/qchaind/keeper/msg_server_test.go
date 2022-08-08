package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qchainnetwork/qchaind/testutil/keeper"
	"github.com/qchainnetwork/qchaind/x/qchaind/keeper"
	"github.com/qchainnetwork/qchaind/x/qchaind/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.QchaindKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
