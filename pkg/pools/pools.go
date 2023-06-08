package pools

import (
	"errors"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type EventPoolSystemRebalanced struct {
	Phase  types.Phase
	PoolID types.U64
	Topics []types.Hash
}

type EventPoolSystemMaxReserveSet struct {
	Phase  types.Phase
	PoolID types.U64
	Topics []types.Hash
}

type EventPoolSystemEpochClosed struct {
	Phase   types.Phase
	PoolID  types.U64
	EpochID types.U32
	Topics  []types.Hash
}

type EventPoolSystemSolutionSubmitted struct {
	Phase    types.Phase
	PoolID   types.U64
	EpochID  types.U32
	Solution EpochSolution
	Topics   []types.Hash
}

type EventPoolSystemEpochExecuted struct {
	Phase   types.Phase
	PoolID  types.U64
	EpochID types.U32
	Topics  []types.Hash
}

type EventPoolSystemCreated struct {
	Phase     types.Phase
	Admin     types.AccountID
	Depositor types.AccountID
	PoolID    types.U64
	Essence   PoolEssence
	Topics    []types.Hash
}

type EventPoolSystemUpdated struct {
	Phase  types.Phase
	PoolID types.U64
	Old    PoolEssence
	New    PoolEssence
	Topics []types.Hash
}

type PoolEssence struct {
	Currency     types.CurrencyID
	MaxReserve   types.U128
	MaxNavAge    types.U64
	MinEpochTime types.U64
	Tranches     []TrancheEssence
}

type TrancheEssence struct {
	Currency TrancheCurrency
	Ty       TrancheType
	Metadata TrancheMetadata
}

type TrancheMetadata struct {
	TokenName   []types.U8
	TokenSymbol []types.U8
}

type TrancheType struct {
	IsResidual bool

	IsNonResidual bool
	AsNonResidual NonResidual
}

func (t *TrancheType) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		t.IsResidual = true

		return nil
	case 1:
		t.IsNonResidual = true

		return decoder.Decode(&t.AsNonResidual)
	default:
		return errors.New("unsupported tranche type")
	}
}

func (t TrancheType) Encode(encoder scale.Encoder) error {
	switch {
	case t.IsResidual:
		return encoder.PushByte(0)
	case t.IsNonResidual:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(t.AsNonResidual)

	default:
		return errors.New("unsupported tranche type")
	}
}

type NonResidual struct {
	InterestRatePerSec types.UCompact
	MinRiskBuffer      types.U64
}

type TrancheCurrency struct {
	PoolID    types.U64
	TrancheID [16]types.U8
}

type EpochSolution struct {
	IsHealthy bool
	AsHealthy HealthySolution

	IsUnhealthy bool
	AsUnhealthy UnhealthySolution
}

func (e *EpochSolution) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		e.IsHealthy = true

		return decoder.Decode(&e.AsHealthy)
	case 1:
		e.IsUnhealthy = true

		return decoder.Decode(&e.AsUnhealthy)
	default:
		return errors.New("unsupported epoch solution")
	}
}

func (e EpochSolution) Encode(encoder scale.Encoder) error {
	switch {
	case e.IsHealthy:
		if err := encoder.PushByte(0); err != nil {
			return err
		}

		return encoder.Encode(e.AsHealthy)
	case e.IsUnhealthy:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(e.AsUnhealthy)

	default:
		return errors.New("unsupported epoch solution")
	}
}

type HealthySolution struct {
	Solution []TrancheSolution
	Balance  types.U128
}

type UnhealthySolution struct {
	State                       []UnhealthyState
	Solution                    []TrancheSolution ``
	RiskBufferImprovementScores types.Option[[]types.U128]
	ReserveImprovementScore     types.Option[types.U128]
}

type UnhealthyState struct {
	IsMaxReserveViolated    bool
	IsMinRiskBufferViolated bool
}

func (u *UnhealthyState) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		u.IsMaxReserveViolated = true

		return nil
	case 1:
		u.IsMinRiskBufferViolated = true

		return nil
	default:
		return errors.New("unsupported unhealthy state")
	}
}

func (u UnhealthyState) Encode(encoder scale.Encoder) error {
	switch {
	case u.IsMaxReserveViolated:

		return encoder.PushByte(0)
	case u.IsMinRiskBufferViolated:

		return encoder.PushByte(1)
	default:
		return errors.New("unsupported unhealthy state")
	}
}

type TrancheSolution struct {
	InvestFulfillment types.U64
	RedeemFulfillment types.U64
}
