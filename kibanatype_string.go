// Code generated by "stringer -type KibanaType"; DO NOT EDIT.

package kibana

import "strconv"

const _KibanaType_name = "KibanaTypeVanillaKibanaTypeLogzio"

var _KibanaType_index = [...]uint8{0, 17, 33}

func (i KibanaType) String() string {
	i -= 1
	if i < 0 || i >= KibanaType(len(_KibanaType_index)-1) {
		return "KibanaType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _KibanaType_name[_KibanaType_index[i]:_KibanaType_index[i+1]]
}
