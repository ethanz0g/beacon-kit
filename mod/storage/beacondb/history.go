// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package beacondb

import (
	"github.com/berachain/beacon-kit/mod/primitives"
)

// UpdateBlockRootAtIndex sets a block root in the BeaconStore.
func (kv *KVStore[
	DepositT, ForkT, BeaconBlockHeaderT,
	ExecutionPayloadT, Eth1DataT, ValidatorT,
]) UpdateBlockRootAtIndex(
	index uint64,
	root primitives.Root,
) error {
	return kv.blockRoots.Set(kv.ctx, index, root)
}

// GetBlockRoot retrieves the block root from the BeaconStore.
func (kv *KVStore[
	DepositT, ForkT, BeaconBlockHeaderT,
	ExecutionPayloadT, Eth1DataT, ValidatorT,
]) GetBlockRootAtIndex(
	index uint64,
) (primitives.Root, error) {
	return kv.blockRoots.Get(kv.ctx, index)
}

// SetLatestBlockHeader sets the latest block header in the BeaconStore.
func (kv *KVStore[
	DepositT, ForkT, BeaconBlockHeaderT,
	ExecutionPayloadT, Eth1DataT, ValidatorT,
]) SetLatestBlockHeader(
	header BeaconBlockHeaderT,
) error {
	return kv.latestBlockHeader.Set(kv.ctx, header)
}

// GetLatestBlockHeader retrieves the latest block header from the BeaconStore.
func (kv *KVStore[
	DepositT, ForkT, BeaconBlockHeaderT,
	ExecutionPayloadT, Eth1DataT, ValidatorT,
]) GetLatestBlockHeader() (
	BeaconBlockHeaderT, error,
) {
	return kv.latestBlockHeader.Get(kv.ctx)
}

// UpdateStateRootAtIndex updates the state root at the given slot.
func (kv *KVStore[
	DepositT, ForkT, BeaconBlockHeaderT,
	ExecutionPayloadT, Eth1DataT, ValidatorT,
]) UpdateStateRootAtIndex(
	idx uint64,
	stateRoot primitives.Root,
) error {
	return kv.stateRoots.Set(kv.ctx, idx, stateRoot)
}

// StateRootAtIndex returns the state root at the given slot.
func (kv *KVStore[
	DepositT, ForkT, BeaconBlockHeaderT,
	ExecutionPayloadT, Eth1DataT, ValidatorT,
]) StateRootAtIndex(
	idx uint64,
) (primitives.Root, error) {
	return kv.stateRoots.Get(kv.ctx, idx)
}