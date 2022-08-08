package qchaind_test

import (
	"testing"

	keepertest "github.com/qchainnetwork/qchaind/testutil/keeper"
	"github.com/qchainnetwork/qchaind/testutil/nullify"
	"github.com/qchainnetwork/qchaind/x/qchaind"
	"github.com/qchainnetwork/qchaind/x/qchaind/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.QchaindKeeper(t)
	qchaind.InitGenesis(ctx, *k, genesisState)
	got := qchaind.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
