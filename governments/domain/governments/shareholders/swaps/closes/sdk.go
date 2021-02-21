package closes

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/trades"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewContentBuilder creates a new content builder instance
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter)
}

// Builder represents a close builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(sig signature.RingSignature) Builder
	Now() (Close, error)
}

// Close represents a swap close
type Close interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithTrade(trade trades.Trade) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a close content
type Content interface {
	Hash() hash.Hash
	Trade() trades.Trade
	CreatedOn() time.Time
}

// Service represents a close service
type Service interface {
	Insert(close Close) error
}
