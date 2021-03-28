package variable

import (
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapterBuilder creates a new adapter builder
func NewAdapterBuilder() AdapterBuilder {
	valueFactory := var_value.NewFactory()
	valueAdapter := var_value.NewAdapter()
	builder := NewBuilder()
	return createAdapterBuilder(valueFactory, valueAdapter, builder)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	IsGlobal() AdapterBuilder
	Now() Adapter
}

// Adapter represents the variables adapter
type Adapter interface {
	FromVariable(declaration parsers.VariableDeclaration) (Variable, error)
}

// Builder represents the variable builder
type Builder interface {
	Create() Builder
	IsGlobal() Builder
	WithName(name string) Builder
	WithValue(val var_value.Value) Builder
	IsIncoming() Builder
	IsOutgoing() Builder
	IsMandatory() Builder
	Now() (Variable, error)
}

// Variable represents a variable
type Variable interface {
	IsGlobal() bool
	IsMandatory() bool
	IsIncoming() bool
	IsOutgoing() bool
	Name() string
	Value() var_value.Value
}
