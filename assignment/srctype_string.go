// Code generated by "stringer -type SrcType"; DO NOT EDIT.

package assignment

import "fmt"

const _SrcType_name = "SrcTypeUnspecifiedSrcTypeAlertSrcTypeEscalationPolicyStepSrcTypeRotationParticipantSrcTypeScheduleRuleSrcTypeServiceSrcTypeUser"

var _SrcType_index = [...]uint8{0, 18, 30, 57, 83, 102, 116, 127}

func (i SrcType) String() string {
	if i < 0 || i >= SrcType(len(_SrcType_index)-1) {
		return fmt.Sprintf("SrcType(%d)", i)
	}
	return _SrcType_name[_SrcType_index[i]:_SrcType_index[i+1]]
}