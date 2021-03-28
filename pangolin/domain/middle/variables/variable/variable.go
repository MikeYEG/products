package variable

import (
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value"
)

type variable struct {
	isGlobal    bool
	isMandatory bool
	isIncoming  bool
	isOutgoing  bool
	name        string
	value       var_value.Value
}

func createVariable(
	isGlobal bool,
	isMandatory bool,
	isIncoming bool,
	isOutgoing bool,
	name string,
	value var_value.Value,
) Variable {
	out := variable{
		isGlobal:    isGlobal,
		isMandatory: isMandatory,
		isIncoming:  isIncoming,
		isOutgoing:  isOutgoing,
		name:        name,
		value:       value,
	}

	return &out
}

// IsGlobal returns true if the variable is global, false otherwise
func (obj *variable) IsGlobal() bool {
	return obj.isGlobal
}

// IsMandatory returns true if the variable is mandatory, false otherwise
func (obj *variable) IsMandatory() bool {
	return obj.isMandatory
}

// IsIncoming returns true if the variable is incoming, false otherwise
func (obj *variable) IsIncoming() bool {
	return obj.isIncoming
}

// IsOutgoing returns true if the variable is outgoing, false otherwise
func (obj *variable) IsOutgoing() bool {
	return obj.isOutgoing
}

// Name returns the name
func (obj *variable) Name() string {
	return obj.name
}

// Value returns the value
func (obj *variable) Value() var_value.Value {
	return obj.value
}
