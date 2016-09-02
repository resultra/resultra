package global

import (
	"fmt"
	"resultra/datasheet/server/generic"
)

const GlobalTypeText string = "text"
const GlobalTypeNumber string = "number"
const GlobalTypeTime string = "time"
const GlobalTypeBool string = "bool"
const GlobalTypeLongText string = "longText"
const GlobalTypeFile string = "file"

func validGlobalType(globalType string) bool {
	switch globalType {
	case GlobalTypeText:
		return true
	case GlobalTypeNumber:
		return true
	case GlobalTypeTime:
		return true
	case GlobalTypeBool:
		return true
	case GlobalTypeLongText:
		return true
	case GlobalTypeFile:
		return true
	default:
		return false
	}
}

type TextValue struct {
	Val string `json:"val"`
}

func decodeGlobalValue(valueType string, encodedVal string) (interface{}, error) {

	if !validGlobalType(valueType) {
		return nil, fmt.Errorf("decodeGlobalValue: unrecognized value type: %v", valueType)
	}

	switch valueType {
	case GlobalTypeText:
		var textVal TextValue
		if err := generic.DecodeJSONString(encodedVal, &textVal); err != nil {
			return nil, fmt.Errorf("decodeGlobalValue: failure decoding text value: %v", err)
		}
		return textVal.Val, nil
	default:
		return nil, fmt.Errorf("decodeGlobalValue: Unrecognized field type: %v", valueType)
	}
}
