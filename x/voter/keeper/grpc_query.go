package keeper

import (
	"github.com/ryandines/voter/x/voter/types"
)

var _ types.QueryServer = Keeper{}
