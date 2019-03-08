// Copyright 2019 The inb-go Authors
// This file is part of the inb-go library.
//
// The inb-go library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The inb-go library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the inb-go library. If not, see <http://www.gnu.org/licenses/>.

// Package vdpos implements the delegated-proof-of-stake consensus engine.
package vdpos

import (
	"github.com/insight-chain/inb-go/common"
	"github.com/insight-chain/inb-go/consensus"
)

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the delegated-proof-of-stake scheme.
type API struct {
	chain consensus.ChainReader
	vdpos *Vdpos
}

// GetSnapshot retrieves the state snapshot at a given block.
func (api *API) GetSnapshot(number uint64) (*Snapshot, error) {
	header := api.chain.GetHeaderByNumber(number)
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.vdpos.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil, nil, defaultLoopCntRecalculateSigners)
}

// GetSnapshotAtHash retrieves the state snapshot at a given block.
func (api *API) GetSnapshotAtHash(hash common.Hash) (*Snapshot, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.vdpos.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil, nil, defaultLoopCntRecalculateSigners)
}

func (api *API) GetSigners(number uint64) ([]common.Address, error) {
	header := api.chain.GetHeaderByNumber(number)
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := api.vdpos.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil, nil, defaultLoopCntRecalculateSigners)
	if err != nil {
		return nil, err
	}
	return snap.signers(), nil
}

// GetSignersAtHash retrieves the list of authorized signers at the specified block.
func (api *API) GetSignersAtHash(hash common.Hash) ([]common.Address, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := api.vdpos.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil, nil, defaultLoopCntRecalculateSigners)
	if err != nil {
		return nil, err
	}
	return snap.signers(), nil
}
