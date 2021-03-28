package variables

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	variableAdapterBuilder := variable.NewAdapterBuilder()
	builder := NewBuilder()
	return createAdapter(variableAdapterBuilder, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the variables adapter
type Adapter interface {
	FromVariables(section parsers.VariableSection) (Variables, error)
}

// Builder represents the variables builder
type Builder interface {
	Create() Builder
	WithVariables(vrs []variable.Variable) Builder
	Now() (Variables, error)
}

// Variables represents variables
type Variables interface {
	All() []variable.Variable
	Merge(vr Variables) Variables
}
