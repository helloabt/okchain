package types

import (
	sdkgenutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
)

// const
const (
	ModuleName = sdkgenutiltypes.ModuleName
)

type (
	// AccountKeeper is the type alias of the one in cmsdk
	AccountKeeper = sdkgenutiltypes.AccountKeeper
	// StakingKeeper is the type alias of the one in cmsdk
	StakingKeeper = sdkgenutiltypes.StakingKeeper
)
