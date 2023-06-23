package rewards

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

// These are the events for the `rewards` pallet.

type EventLiquidityRewardsBaseGroupAwarded struct {
	Phase   types.Phase
	GroupID types.U32
	Amount  types.U128
	Topics  []types.Hash
}

type EventLiquidityRewardsBaseStakeDeposited struct {
	Phase      types.Phase
	GroupID    types.U32
	CurrencyID types.CurrencyID
	AccountID  types.AccountID
	Amount     types.U128
	Topics     []types.Hash
}

type EventLiquidityRewardsBaseStakeWithdrawn struct {
	Phase      types.Phase
	GroupID    types.U32
	CurrencyID types.CurrencyID
	AccountID  types.AccountID
	Amount     types.U128
	Topics     []types.Hash
}

type EventLiquidityRewardsBaseRewardClaimed struct {
	Phase      types.Phase
	GroupID    types.U32
	CurrencyID types.CurrencyID
	AccountID  types.AccountID
	Amount     types.U128
	Topics     []types.Hash
}

type EventLiquidityRewardsBaseCurrencyAttached struct {
	Phase      types.Phase
	CurrencyID types.CurrencyID
	From       types.Option[types.U32]
	To         types.U32
	Topics     []types.Hash
}
