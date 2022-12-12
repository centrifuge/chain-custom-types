package proxy

import (
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
)

type CentrifugeProxyType uint8

const (
	Any CentrifugeProxyType = iota
	NonTransfer
	Governance
	Staking
	NonProxy
	Borrow
	Price
	Invest
	ProxyManagement
	KeystoreManagement
	PodOperation
	PodAuth
)

var (
	proxyTypeMap = map[CentrifugeProxyType]struct{}{
		Any:                {},
		NonTransfer:        {},
		Governance:         {},
		Staking:            {},
		NonProxy:           {},
		Borrow:             {},
		Price:              {},
		Invest:             {},
		ProxyManagement:    {},
		KeystoreManagement: {},
		PodOperation:       {},
		PodAuth:            {},
	}

	ProxyTypeName = map[CentrifugeProxyType]string{
		Any:                "Any",
		NonTransfer:        "NonTransfer",
		Governance:         "Governance",
		Staking:            "Staking",
		NonProxy:           "NonProxy",
		Borrow:             "Borrow",
		Price:              "Price",
		Invest:             "Invest",
		ProxyManagement:    "ProxyManagement",
		KeystoreManagement: "KeystoreManagement",
		PodOperation:       "PodOperation",
		PodAuth:            "PodAuth",
	}

	ProxyTypeValue = map[string]CentrifugeProxyType{
		"Any":                Any,
		"NonTransfer":        NonTransfer,
		"Governance":         Governance,
		"Staking":            Staking,
		"NonProxy":           NonProxy,
		"Borrow":             Borrow,
		"Price":              Price,
		"Invest":             Invest,
		"ProxyManagement":    ProxyManagement,
		"KeystoreManagement": KeystoreManagement,
		"PodOperation":       PodOperation,
		"PodAuth":            PodAuth,
	}
)

func (pt *CentrifugeProxyType) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	pb := CentrifugeProxyType(b)

	if _, ok := proxyTypeMap[pb]; !ok {
		return fmt.Errorf("unknown ProxyType enum: %v", pb)
	}

	*pt = pb

	return nil
}

func (pt CentrifugeProxyType) Encode(encoder scale.Encoder) error {
	return encoder.PushByte(byte(pt))
}
