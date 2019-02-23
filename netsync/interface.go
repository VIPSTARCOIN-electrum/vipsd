// Copyright (c) 2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netsync

import (
	"github.com/VIPSTARCOIN-electrum/vipsd/blockchain"
	"github.com/VIPSTARCOIN-electrum/vipsd/chaincfg"
	"github.com/VIPSTARCOIN-electrum/vipsd/chaincfg/chainhash"
	"github.com/VIPSTARCOIN-electrum/vipsd/mempool"
	"github.com/VIPSTARCOIN-electrum/vipsd/peer"
	"github.com/VIPSTARCOIN-electrum/vipsd/wire"
	"github.com/VIPSTARCOIN-electrum/vipsutil"
)

// PeerNotifier exposes methods to notify peers of status changes to
// transactions, blocks, etc. Currently server (in the main package) implements
// this interface.
type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)

	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	TransactionConfirmed(tx *vipsutil.Tx)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	PeerNotifier PeerNotifier
	Chain        *blockchain.BlockChain
	TxMemPool    *mempool.TxPool
	ChainParams  *chaincfg.Params

	DisableCheckpoints bool
	MaxPeers           int

	FeeEstimator *mempool.FeeEstimator
}
