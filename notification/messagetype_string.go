// Code generated by "stringer -type MessageType"; DO NOT EDIT.

package notification

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MessageTypeAlert-0]
	_ = x[MessageTypeAlertStatus-1]
	_ = x[MessageTypeTest-2]
	_ = x[MessageTypeVerification-3]
}

const _MessageType_name = "MessageTypeAlertMessageTypeAlertStatusMessageTypeTestMessageTypeVerification"

var _MessageType_index = [...]uint8{0, 16, 38, 53, 76}

func (i MessageType) String() string {
	if i < 0 || i >= MessageType(len(_MessageType_index)-1) {
		return "MessageType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}