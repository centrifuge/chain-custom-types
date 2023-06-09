package pools

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

type EventPoolRegistryRegistered struct {
	Phase  types.Phase
	PoolID types.U64
	Topics []types.Hash
}

type EventPoolRegistryUpdateRegistered struct {
	Phase  types.Phase
	PoolID types.U64
	Topics []types.Hash
}

type EventPoolRegistryUpdateExecuted struct {
	Phase  types.Phase
	PoolID types.U64
	Topics []types.Hash
}

type EventPoolRegistryUpdateStored struct {
	Phase  types.Phase
	PoolID types.U64
	Topics []types.Hash
}

type EventPoolRegistryMetadataSet struct {
	Phase    types.Phase
	PoolID   types.U64
	Metadata []types.U8
	Topics   []types.Hash
}
