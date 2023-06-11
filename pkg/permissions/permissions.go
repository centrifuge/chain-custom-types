package permissions

import (
	"errors"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type EventPermissionsAdded struct {
	Phase  types.Phase
	To     types.AccountID
	Scope  PermissionScope
	Role   Role
	Topics []types.Hash
}

type EventPermissionsRemoved struct {
	Phase  types.Phase
	To     types.AccountID
	Scope  PermissionScope
	Role   Role
	Topics []types.Hash
}

type EventPermissionsPurged struct {
	Phase  types.Phase
	From   types.AccountID
	Scope  PermissionScope
	Topics []types.Hash
}

type PermissionScope struct {
	IsPool bool
	AsPool types.U64

	IsCurrency bool
	AsCurrency types.CurrencyID
}

func (p *PermissionScope) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsPool = true

		return decoder.Decode(&p.AsPool)
	case 1:
		p.IsCurrency = true

		return decoder.Decode(&p.AsCurrency)
	default:
		return errors.New("unsupported permission scope")
	}
}

func (p PermissionScope) Encode(encoder scale.Encoder) error {
	switch {
	case p.IsPool:
		if err := encoder.PushByte(0); err != nil {
			return err
		}

		return encoder.Encode(p.AsPool)
	case p.IsCurrency:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(p.AsCurrency)

	default:
		return errors.New("unsupported permission scope")
	}
}

type Role struct {
	IsPoolRole bool
	AsPoolRole PoolRole

	IsPermissionedCurrencyRole bool
	AsPermissionedCurrencyRole PermissionedCurrencyRole
}

func (r *Role) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		r.IsPoolRole = true

		return decoder.Decode(&r.AsPoolRole)
	case 1:
		r.IsPermissionedCurrencyRole = true

		return decoder.Decode(&r.AsPermissionedCurrencyRole)
	default:
		return errors.New("unsupported role")
	}
}

func (r Role) Encode(encoder scale.Encoder) error {
	switch {
	case r.IsPoolRole:
		if err := encoder.PushByte(0); err != nil {
			return err
		}

		return encoder.Encode(r.AsPoolRole)
	case r.IsPermissionedCurrencyRole:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(r.AsPermissionedCurrencyRole)
	default:
		return errors.New("unsupported role")
	}
}

type PoolRole struct {
	IsPoolAdmin bool

	IsBorrower bool

	IsPricingAdmin bool

	IsLiquidityAdmin bool

	IsMemberListAdmin bool

	IsLoanAdmin bool

	IsTrancheInvestor bool
	AsTrancheInvestor TrancheInvestor

	IsPODReadAccess bool
}

func (p *PoolRole) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsPoolAdmin = true

		return nil
	case 1:
		p.IsBorrower = true

		return nil
	case 2:
		p.IsPricingAdmin = true

		return nil
	case 3:
		p.IsLiquidityAdmin = true

		return nil
	case 4:
		p.IsMemberListAdmin = true

		return nil
	case 5:
		p.IsLoanAdmin = true

		return nil
	case 6:
		p.IsTrancheInvestor = true

		return decoder.Decode(&p.AsTrancheInvestor)
	case 7:
		p.IsPODReadAccess = true

		return nil
	default:
		return errors.New("unsupported pool role")
	}
}

func (p PoolRole) Encode(encoder scale.Encoder) error {
	switch {
	case p.IsPoolAdmin:
		return encoder.PushByte(0)
	case p.IsBorrower:
		return encoder.PushByte(1)
	case p.IsPricingAdmin:
		return encoder.PushByte(2)
	case p.IsLiquidityAdmin:
		return encoder.PushByte(3)
	case p.IsMemberListAdmin:
		return encoder.PushByte(4)
	case p.IsLoanAdmin:
		return encoder.PushByte(5)
	case p.IsTrancheInvestor:
		if err := encoder.PushByte(6); err != nil {
			return err
		}

		return encoder.Encode(p.AsTrancheInvestor)
	case p.IsPODReadAccess:
		return encoder.PushByte(7)
	default:
		return errors.New("unsupported pool role")
	}
}

type TrancheInvestor struct {
	TrancheID [16]types.U8
	Moment    types.U64
}

type PermissionedCurrencyRole struct {
	IsHolder bool
	AsHolder types.U64

	IsManager bool

	IsIssuer bool
}

func (p *PermissionedCurrencyRole) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsHolder = true

		return decoder.Decode(&p.AsHolder)
	case 1:
		p.IsManager = true

		return nil
	case 2:
		p.IsIssuer = true

		return nil
	default:
		return errors.New("unsupported permissioned currency role")
	}
}

func (p PermissionedCurrencyRole) PermissionedCurrencyRole(encoder scale.Encoder) error {
	switch {
	case p.IsHolder:
		if err := encoder.PushByte(0); err != nil {
			return err
		}

		return encoder.Encode(p.AsHolder)
	case p.IsManager:
		return encoder.PushByte(1)
	case p.IsIssuer:
		return encoder.PushByte(2)
	default:
		return errors.New("unsupported permissioned currency role")
	}
}
