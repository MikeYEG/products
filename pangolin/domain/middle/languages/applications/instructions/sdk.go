package instructions

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an instructions builder
type Builder interface {
	Create() Builder
	WithList(list []instruction.Instruction) Builder
	Now() (Instructions, error)
}

// Instructions represents a language instructions
type Instructions interface {
	All() []instruction.Instruction
}