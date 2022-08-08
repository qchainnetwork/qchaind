package keeper

import (
	"github.com/qchainnetwork/qchaind/x/qchaind/types"
)

var _ types.QueryServer = Keeper{}
