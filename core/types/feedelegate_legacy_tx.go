// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// FeeDelegateLegacyTx is a transaction that pays fees instead.

type FeeDelegateLegacyTx struct {
	SenderTx    LegacyTx
	MaxFeeLimit *big.Int
	FeePayer    *common.Address `rlp:"nil"`

	// Signature values
	FV *big.Int `json:"fv" gencodec:"required"`
	FR *big.Int `json:"fr" gencodec:"required"`
	FS *big.Int `json:"fs" gencodec:"required"`
}

func (tx *FeeDelegateLegacyTx) getSenderTx() TxData {
	return nil
}

func (tx *FeeDelegateLegacyTx) SetSenderTx(senderTx LegacyTx) {
	tx.SenderTx.Nonce = senderTx.nonce()
	tx.SenderTx.To = senderTx.to()
	tx.SenderTx.Data = senderTx.data()
	tx.SenderTx.Gas = senderTx.gas()
	tx.SenderTx.Value = senderTx.value()
	tx.SenderTx.GasPrice = senderTx.gasPrice()
	v, r, s := senderTx.rawSignatureValues()
	tx.SenderTx.V = v
	tx.SenderTx.R = r
	tx.SenderTx.S = s
}

//func (tx *FeeDelegateLegacyTx) SetSenderTxhash(senderTx TxData) {
//	//tx.SenderTxHash = senderTx.getSenderTxHash()
//}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *FeeDelegateLegacyTx) copy() TxData {
	cpy := &FeeDelegateLegacyTx{
		SenderTx:    tx.copyLegacyTx(),
		MaxFeeLimit: tx.MaxFeeLimit,
		FeePayer:    copyAddressPtr(tx.FeePayer),
		// fee payer Signature values
		FV: new(big.Int),
		FR: new(big.Int),
		FS: new(big.Int),
	}

	if tx.FV != nil {
		cpy.FV.Set(tx.FV)
	}
	if tx.FR != nil {
		cpy.FR.Set(tx.FR)
	}
	if tx.FS != nil {
		cpy.FS.Set(tx.FS)
	}
	return cpy
}

func (tx *FeeDelegateLegacyTx) copyLegacyTx() LegacyTx {
	cpy := LegacyTx{
		Nonce: tx.SenderTx.Nonce,
		To:    copyAddressPtr(tx.SenderTx.To),
		Data:  common.CopyBytes(tx.SenderTx.Data),
		Gas:   tx.SenderTx.Gas,
		// These are initialized below.
		Value:    new(big.Int),
		GasPrice: new(big.Int),
		V:        new(big.Int),
		R:        new(big.Int),
		S:        new(big.Int),
	}
	if tx.SenderTx.Value != nil {
		cpy.Value.Set(tx.SenderTx.Value)
	}
	if tx.SenderTx.GasPrice != nil {
		cpy.GasPrice.Set(tx.SenderTx.GasPrice)
	}
	if tx.SenderTx.V != nil {
		cpy.V.Set(tx.SenderTx.V)
	}
	if tx.SenderTx.R != nil {
		cpy.R.Set(tx.SenderTx.R)
	}
	if tx.SenderTx.S != nil {
		cpy.S.Set(tx.SenderTx.S)
	}
	return cpy
}

// accessors for FeeDelegateTx.
func (tx *FeeDelegateLegacyTx) txType() byte              { return FeeDelegateLegacyTxType }
func (tx *FeeDelegateLegacyTx) chainID() *big.Int         { return tx.SenderTx.chainID() }
func (tx *FeeDelegateLegacyTx) accessList() AccessList    { return tx.SenderTx.accessList() }
func (tx *FeeDelegateLegacyTx) data() []byte              { return tx.SenderTx.data() }
func (tx *FeeDelegateLegacyTx) gas() uint64               { return tx.SenderTx.gas() }
func (tx *FeeDelegateLegacyTx) gasFeeCap() *big.Int       { return tx.SenderTx.gasFeeCap() }
func (tx *FeeDelegateLegacyTx) gasTipCap() *big.Int       { return tx.SenderTx.gasPrice() }
func (tx *FeeDelegateLegacyTx) gasPrice() *big.Int        { return tx.SenderTx.gasPrice() }
func (tx *FeeDelegateLegacyTx) value() *big.Int           { return tx.SenderTx.value() }
func (tx *FeeDelegateLegacyTx) nonce() uint64             { return tx.SenderTx.nonce() }
func (tx *FeeDelegateLegacyTx) to() *common.Address       { return tx.SenderTx.to() }
func (tx *FeeDelegateLegacyTx) maxfeelimit() *big.Int     { return tx.MaxFeeLimit }
func (tx *FeeDelegateLegacyTx) feePayer() *common.Address { return tx.FeePayer }
func (tx *FeeDelegateLegacyTx) rawFeePayerSignatureValues() (v, r, s *big.Int) {
	return tx.FV, tx.FR, tx.FS
}

func (tx *FeeDelegateLegacyTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.SenderTx.rawSignatureValues()
}

func (tx *FeeDelegateLegacyTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.FV, tx.FR, tx.FS = v, r, s
}

//
//
//func (tx *FeeDelegateLegacyTx) rawSenderTxSignatureValues() (v, r, s *big.Int) {
//	return tx.SenderTx.rawSignatureValues()
//}
//
//func (tx *FeeDelegateLegacyTx) SetSenderTxSignatureValues(chainID, v, r, s *big.Int) {
//	tx.SenderTx.setSignatureValues(chainID, v, r, s)
//}
//
//func (tx *FeeDelegateLegacyTx) getFeePayer() *common.Address {
//	return tx.FeePayer
//}
//
//func (tx *FeeDelegateLegacyTx) senderTxcopy() TxData {
//	return tx.SenderTx.copy()
//}
