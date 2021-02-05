package transactions

import (
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type transaction struct {
	hash hash.Hash
	body bodies.Body
	sig  signature.RingSignature
}

func createTransaction(
	hash hash.Hash,
	body bodies.Body,
	sig signature.RingSignature,
) Transaction {
	out := transaction{
		hash: hash,
		body: body,
		sig:  sig,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Body returns the body
func (obj *transaction) Body() bodies.Body {
	return obj.body
}

// Signature returns the ring signature
func (obj *transaction) Signature() signature.RingSignature {
	return obj.sig
}
