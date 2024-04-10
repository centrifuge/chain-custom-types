package loans

import (
	"errors"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type EventLoansCreated struct {
	Phase    types.Phase
	PoolID   types.U64
	LoanID   types.U64
	LoanInfo LoanInfo
	Topics   []types.Hash
}

type EventLoansBorrowed struct {
	Phase  types.Phase
	PoolID types.U64
	LoanID types.U64
	Amount types.U128
	Topics []types.Hash
}

type EventLoansRepaid struct {
	Phase           types.Phase
	PoolID          types.U64
	LoanID          types.U64
	Amount          types.U128
	UncheckedAmount types.U128
	Topics          []types.Hash
}

type EventLoansWrittenOff struct {
	Phase  types.Phase
	PoolID types.U64
	LoanID types.U64
	Status WriteOffStatus
	Topics []types.Hash
}

type EventLoansMutated struct {
	Phase    types.Phase
	PoolID   types.U64
	LoanID   types.U64
	Mutation LoanMutation
	Topics   []types.Hash
}

type EventLoansClosed struct {
	Phase      types.Phase
	PoolID     types.U64
	LoanID     types.U64
	Collateral Asset
	Topics     []types.Hash
}

type EventLoansPortfolioValuationUpdated struct {
	Phase      types.Phase
	PoolID     types.U64
	Valuation  types.U128
	UpdateType PortfolioValuationUpdateType
	Topics     []types.Hash
}

type EventLoansWriteOffPolicyUpdated struct {
	Phase  types.Phase
	PoolID types.U64
	Policy []WriteOffRule
	Topics []types.Hash
}

type ActiveLoan struct {
	Schedule                  RepaymentSchedule
	Collateral                Asset
	Restrictions              LoanRestrictions
	Borrower                  types.AccountID
	WriteOffPercentage        types.U128
	OriginationDate           types.U64
	Pricing                   Pricing
	TotalBorrowed             types.U128
	TotalRepaid               RepaidAmount
	RepaymentsOnScheduleUntil types.U64
}

type RepaidAmount struct {
	Principal   types.U128
	Interest    types.U128
	Unscheduled types.U128
}

type ClosedLoan struct {
	ClosedAt      types.U32
	Info          LoanInfo
	TotalBorrowed types.U128
	TotalRepaid   RepaidAmount
}

type WriteOffRule struct {
	Triggers []UniqueWriteOffTrigger
	Status   WriteOffStatus
}

type UniqueWriteOffTrigger struct {
	IsPrincipalOverdueDays bool
	AsPrincipalOverdueDays types.U32

	IsPriceOutdated bool
	AsPriceOutdated types.U64
}

func (u *UniqueWriteOffTrigger) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		u.IsPrincipalOverdueDays = true

		return decoder.Decode(&u.AsPrincipalOverdueDays)
	case 1:
		u.IsPriceOutdated = true

		return decoder.Decode(&u.AsPriceOutdated)
	default:
		return errors.New("unsupported unique writeoff trigger")
	}
}

func (u UniqueWriteOffTrigger) Encode(encoder scale.Encoder) error {
	switch {
	case u.IsPrincipalOverdueDays:
		if err := encoder.PushByte(0); err != nil {
			return err
		}

		return encoder.Encode(u.AsPrincipalOverdueDays)
	case u.IsPriceOutdated:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(u.AsPriceOutdated)
	default:
		return errors.New("unsupported unique writeoff trigger")
	}
}

type PortfolioValuationUpdateType struct {
	IsExact   bool
	IsInexact bool
}

func (p *PortfolioValuationUpdateType) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsExact = true

		return nil
	case 1:
		p.IsInexact = true

		return nil
	default:
		return errors.New("unsupported portfolio valuation update type")
	}
}

func (p PortfolioValuationUpdateType) Encode(encoder scale.Encoder) error {
	switch {
	case p.IsExact:
		return encoder.PushByte(0)
	case p.IsInexact:
		return encoder.PushByte(1)
	default:
		return errors.New("unsupported portfolio valuation update type")
	}
}

type LoanMutation struct {
	IsMaturity bool
	AsMaturity Maturity

	IsInterestPayments bool
	AsInterestPayments InterestPayments

	IsPayDownSchedule bool
	AsPayDownSchedule PayDownSchedule

	IsInternal bool
	AsInternal InternalMutation
}

func (l *LoanMutation) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		l.IsMaturity = true

		return decoder.Decode(&l.AsMaturity)
	case 1:
		l.IsInterestPayments = true

		return decoder.Decode(&l.AsInterestPayments)
	case 2:
		l.IsPayDownSchedule = true

		return decoder.Decode(&l.AsPayDownSchedule)
	case 3:
		l.IsInternal = true

		return decoder.Decode(&l.AsInternal)
	default:
		return errors.New("unsupported loan mutation")
	}
}

func (l LoanMutation) Encode(encoder scale.Encoder) error {
	switch {
	case l.IsMaturity:
		if err := encoder.PushByte(0); err != nil {
			return err
		}

		return encoder.Encode(l.AsMaturity)
	case l.IsInterestPayments:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(l.AsInterestPayments)
	case l.IsPayDownSchedule:
		if err := encoder.PushByte(2); err != nil {
			return err
		}

		return encoder.Encode(l.AsPayDownSchedule)
	case l.IsInternal:
		if err := encoder.PushByte(3); err != nil {
			return err
		}

		return encoder.Encode(l.AsInternal)
	default:
		return errors.New("unsupported loan mutation")
	}
}

type InternalMutation struct {
	IsInterestRate bool
	AsInterestRate types.U128

	IsValuationMethod bool
	AsValuationMethod ValuationMethod

	IsProbabilityOfDefault bool
	AsProbabilityOfDefault types.U128

	IsLossGivenDefault bool
	AsLossGivenDefault types.U128

	IsDiscountRate bool
	AsDiscountRate types.U128
}

func (i *InternalMutation) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		i.IsInterestRate = true

		return decoder.Decode(&i.AsInterestRate)
	case 1:
		i.IsValuationMethod = true

		return decoder.Decode(&i.AsValuationMethod)
	case 2:
		i.IsProbabilityOfDefault = true

		return decoder.Decode(&i.AsProbabilityOfDefault)
	case 3:
		i.IsLossGivenDefault = true

		return decoder.Decode(&i.AsLossGivenDefault)
	case 4:
		i.IsDiscountRate = true

		return decoder.Decode(&i.AsDiscountRate)
	default:
		return errors.New("unsupported internal mutation")
	}
}

func (i InternalMutation) Encode(encoder scale.Encoder) error {
	switch {
	case i.IsInterestRate:
		if err := encoder.PushByte(0); err != nil {
			return err
		}

		return encoder.Encode(i.AsInterestRate)
	case i.IsValuationMethod:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(i.AsValuationMethod)
	case i.IsProbabilityOfDefault:
		if err := encoder.PushByte(2); err != nil {
			return err
		}

		return encoder.Encode(i.AsProbabilityOfDefault)
	case i.IsLossGivenDefault:
		if err := encoder.PushByte(3); err != nil {
			return err
		}

		return encoder.Encode(i.IsLossGivenDefault)
	case i.IsDiscountRate:
		if err := encoder.PushByte(4); err != nil {
			return err
		}

		return encoder.Encode(i.AsDiscountRate)
	default:
		return errors.New("unsupported internal mutation")
	}
}

type WriteOffStatus struct {
	Percentage types.U128
	Penalty    types.U128
}

type LoanInfo struct {
	Schedule     RepaymentSchedule
	Collateral   Asset
	InterestRate InterestRate
	Pricing      Pricing
	Restrictions LoanRestrictions
}

type InterestRate struct {
	IsFixed bool
	AsFixed FixedInterestRate
}

func (i *InterestRate) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		i.IsFixed = true

		return decoder.Decode(&i.AsFixed)
	default:
		return errors.New("unsupported interest rate")
	}
}

func (i InterestRate) Encode(encoder scale.Encoder) error {
	switch {
	case i.IsFixed:
		if err := encoder.PushByte(0); err != nil {
			return err
		}

		return encoder.Encode(i.AsFixed)
	default:
		return errors.New("unsupported interest rate")
	}
}

type FixedInterestRate struct {
	RatePerYear types.U128
	Compounding CompoundingSchedule
}

type CompoundingSchedule struct {
	IsSecondly bool
}

func (c *CompoundingSchedule) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		c.IsSecondly = true

		return nil
	default:
		return errors.New("unsupported compounding schedule")
	}
}

func (c CompoundingSchedule) Encode(encoder scale.Encoder) error {
	switch {
	case c.IsSecondly:
		return encoder.PushByte(0)
	default:
		return errors.New("unsupported compounding schedule")

	}
}

type Pricing struct {
	IsInternal bool
	AsInternal InternalPricing

	IsExternal bool
	AsExternal ExternalPricing
}

func (p *Pricing) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsInternal = true

		return decoder.Decode(&p.AsInternal)
	case 1:
		p.IsExternal = true

		return decoder.Decode(&p.AsExternal)
	default:
		return errors.New("unsupported pricing")
	}
}

func (p Pricing) Encode(encoder scale.Encoder) error {
	switch {
	case p.IsInternal:
		if err := encoder.PushByte(0); err != nil {
			return err
		}

		return encoder.Encode(p.AsInternal)
	case p.IsExternal:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(p.AsExternal)
	default:
		return errors.New("unsupported pricing")
	}
}

type InternalPricing struct {
	CollateralValue types.U128
	ValuationMethod ValuationMethod
	MaxBorrowAmount InternalPricingMaxBorrowAmount
}

type ExternalPricing struct {
	PriceID         PriceID
	MaxBorrowAmount ExternalPricingMaxBorrowAmount
	Notional        types.U128
}

type ExternalPricingMaxBorrowAmount struct {
	IsNoLimit bool

	IsQuantity bool
	AsQuantity types.U128
}

func (e *ExternalPricingMaxBorrowAmount) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		e.IsNoLimit = true

		return nil
	case 1:
		e.IsQuantity = true

		return decoder.Decode(&e.AsQuantity)
	default:
		return errors.New("unsupported external pricing max borrow amount")
	}
}

func (e ExternalPricingMaxBorrowAmount) Encode(encoder scale.Encoder) error {
	switch {
	case e.IsNoLimit:
		return encoder.PushByte(0)
	case e.IsQuantity:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(e.AsQuantity)
	default:
		return errors.New("unsupported external pricing max borrow amount")
	}
}

type PriceID struct {
	IsIsin bool
	AsIsin [12]types.U8
}

func (p *PriceID) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsIsin = true

		return decoder.Decode(&p.AsIsin)
	default:
		return errors.New("unsupported price ID")
	}
}

func (p PriceID) Encode(encoder scale.Encoder) error {
	switch {
	case p.IsIsin:
		if err := encoder.PushByte(0); err != nil {
			return err
		}

		return encoder.Encode(p.AsIsin)
	default:
		return errors.New("unsupported price ID")
	}
}

type LoanRestrictions struct {
	Borrows    BorrowRestrictions
	Repayments RepayRestrictions
}

type BorrowRestrictions struct {
	IsNotWrittenOff bool
	IsFullOnce      bool
}

func (r *BorrowRestrictions) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		r.IsNotWrittenOff = true

		return nil
	case 1:
		r.IsFullOnce = true

		return nil
	default:
		return errors.New("unsupported borrowed restrictions")
	}
}

func (r BorrowRestrictions) Encode(encoder scale.Encoder) error {
	switch {
	case r.IsNotWrittenOff:
		return encoder.PushByte(0)
	case r.IsFullOnce:
		return encoder.PushByte(1)
	default:
		return errors.New("unsupported borrowed restrictions")
	}
}

type RepayRestrictions struct {
	IsNone     bool
	IsFullOnce bool
}

func (r *RepayRestrictions) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		r.IsNone = true

		return nil
	case 1:
		r.IsFullOnce = true

		return nil
	default:
		return errors.New("unsupported repay restrictions")
	}
}

func (r RepayRestrictions) Encode(encoder scale.Encoder) error {
	switch {
	case r.IsNone:
		return encoder.PushByte(0)
	case r.IsFullOnce:
		return encoder.PushByte(1)
	default:
		return errors.New("unsupported repay restrictions")
	}
}

type InternalPricingMaxBorrowAmount struct {
	IsUpToTotalBorrowed bool
	AsUpToTotalBorrowed AdvanceRate

	IsUpToOutstandingDebt bool
	AsUpToOutstandingDebt AdvanceRate
}

type AdvanceRate struct {
	AdvanceRate types.U128
}

func (m *InternalPricingMaxBorrowAmount) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		m.IsUpToTotalBorrowed = true

		return decoder.Decode(&m.AsUpToTotalBorrowed)
	case 1:
		m.IsUpToOutstandingDebt = true

		return decoder.Decode(&m.AsUpToOutstandingDebt)
	default:
		return errors.New("unsupported max borrow amount")
	}
}

func (m InternalPricingMaxBorrowAmount) Encode(encoder scale.Encoder) error {
	switch {
	case m.IsUpToTotalBorrowed:
		if err := encoder.PushByte(0); err != nil {
			return err
		}

		return encoder.Encode(m.AsUpToTotalBorrowed)
	case m.IsUpToOutstandingDebt:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(m.AsUpToOutstandingDebt)
	default:
		return errors.New("unsupported max borrow amount")
	}
}

type ValuationMethod struct {
	IsDiscountedCashFlow bool
	AsDiscountedCashFlow DiscountedCashFlow

	IsOutstandingDebt bool
}

func (v *ValuationMethod) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		v.IsDiscountedCashFlow = true

		return decoder.Decode(&v.AsDiscountedCashFlow)
	case 1:
		v.IsOutstandingDebt = true

		return nil
	default:
		return errors.New("unsupported valuation method")
	}
}

func (v ValuationMethod) Encode(encoder scale.Encoder) error {
	switch {
	case v.IsDiscountedCashFlow:
		if err := encoder.PushByte(0); err != nil {
			return nil
		}

		return encoder.Encode(v.AsDiscountedCashFlow)
	case v.IsOutstandingDebt:
		return encoder.PushByte(1)
	default:
		return errors.New("unsupported valuation method")
	}
}

type DiscountedCashFlow struct {
	ProbabilityOfDefault types.U128
	LossGivenDefault     types.U128
	DiscountRate         InterestRate
}

type Asset struct {
	CollectionID types.U64
	ItemID       types.U128
}

type RepaymentSchedule struct {
	Maturity         Maturity
	InterestPayments InterestPayments
	PayDownSchedule  PayDownSchedule
}

type PayDownSchedule struct {
	IsNone bool
}

func (p *PayDownSchedule) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsNone = true

		return nil
	default:
		return errors.New("unsupported pay down schedule")
	}
}

func (p PayDownSchedule) Encode(encoder scale.Encoder) error {
	if !p.IsNone {
		return errors.New("invalid pay down schedule")
	}

	return encoder.PushByte(0)
}

type InterestPayments struct {
	IsNone bool
}

func (i *InterestPayments) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		i.IsNone = true

		return nil
	default:
		return errors.New("unsupported interest payments")
	}
}

func (i InterestPayments) Encode(encoder scale.Encoder) error {
	if !i.IsNone {
		return errors.New("invalid interest payments")
	}

	return encoder.PushByte(0)
}

type FixedMaturity struct {
	Date      types.U64
	Extension types.U64
}

func (f *FixedMaturity) Decode(decoder scale.Decoder) error {
	if err := decoder.Decode(&f.Date); err != nil {
		return err
	}

	return decoder.Decode(&f.Extension)
}

func (f FixedMaturity) Encode(encoder scale.Encoder) error {
	if err := encoder.Encode(f.Date); err != nil {
		return err
	}

	return encoder.Encode(f.Extension)
}

type Maturity struct {
	IsFixed bool
	AsFixed FixedMaturity
}

func (m *Maturity) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		m.IsFixed = true

		return decoder.Decode(&m.AsFixed)
	default:
		return errors.New("unsupported maturity")
	}
}

func (m Maturity) Encode(encoder scale.Encoder) error {
	if !m.IsFixed {
		return errors.New("invalid maturity")
	}

	if err := encoder.PushByte(0); err != nil {
		return nil
	}

	return encoder.Encode(m.AsFixed)
}
