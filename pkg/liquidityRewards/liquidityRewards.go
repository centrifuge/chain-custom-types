package liquidityRewards

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

type EpochChanges struct {
	Duration   types.Option[types.U32]
	Reward     types.Option[types.U128]
	Weights    []EpochWeights
	Currencies []EpochCurrencies
}

type EpochWeights struct {
	GroupID types.U32
	Weight  types.U64
}

type EpochCurrencies struct {
	CurrencyID types.CurrencyID
	GroupID    types.U32
}

type EventLiquidityRewardsNewEpoch struct {
	Phase       types.Phase
	EndsOn      types.U32
	Reward      types.U128
	LastChanges EpochChanges
	Topics      []types.Hash
}
