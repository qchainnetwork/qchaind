package keeper_test

import (
	"testing"

	testkeeper "github.com/qchainnetwork/qchaind/testutil/keeper"
	"github.com/qchainnetwork/qchaind/x/qchaind/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.QchaindKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
