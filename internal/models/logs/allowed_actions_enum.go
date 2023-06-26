package logs

import (
)

type AllowedAction string

const (
	Insert AllowedAction = "insert"
	Update = "update"
	Delete = "delete"
	VirtualDelete = "virtual_delete"
)

func (action AllowedAction) String() string {
	result := ""

	switch action {
	case Insert:
		result = string(Insert)
	case Update:
		result = string(Update)
	case Delete:
		result = string(Delete)
	case VirtualDelete:
		result = string(VirtualDelete)
	}

	return result
}

func (action AllowedAction) IsValid () bool {
	switch action {
		case Insert, Update, Delete, VirtualDelete:
			return true
		default:
			return false
	}
}

func RequiresData (actionStr string) bool {
	switch AllowedAction(actionStr) {
		case Insert, Update:
			return true
		default:
			return false
	}
}