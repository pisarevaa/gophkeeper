// Code generated by "stringer -type=ContextKeyEnum -linecomment -output contextkey_enum_string.go"; DO NOT EDIT.

package model

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ContextKeyUnknown-0]
	_ = x[ContextUserID-1]
}

const _ContextKeyEnum_name = "unknowntext"

var _ContextKeyEnum_index = [...]uint8{0, 7, 11}

func (i ContextKeyEnum) String() string {
	if i < 0 || i >= ContextKeyEnum(len(_ContextKeyEnum_index)-1) {
		return "ContextKeyEnum(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ContextKeyEnum_name[_ContextKeyEnum_index[i]:_ContextKeyEnum_index[i+1]]
}
