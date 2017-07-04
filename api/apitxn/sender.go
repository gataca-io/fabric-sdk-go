/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package apitxn allows SDK users to plugin their own implementations of transaction processing.
package apitxn

import (
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Sender provides the ability for a transaction to be created and sent.
type Sender interface {
	CreateTransaction(resps []*TransactionProposalResponse) (*Transaction, error)
	SendTransaction(tx *Transaction) ([]*TransactionResponse, error)
}

// ProposalSender provides the ability for a transaction proposal to be created and sent.
type ProposalSender interface {
	CreateTransactionProposal(chaincodeName string, channelID string, args []string, sign bool, transientData map[string][]byte) (*TransactionProposal, error)
	SendTransactionProposal(proposal *TransactionProposal, retry int, targets []ProposalProcessor) ([]*TransactionProposalResponse, error)
}

// The Transaction object created from an endorsed proposal.
type Transaction struct {
	Proposal    *TransactionProposal
	Transaction *pb.Transaction
}

// TransactionResponse contains information returned by the orderer.
type TransactionResponse struct {
	Orderer string
	Err     error
}
