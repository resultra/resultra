package global

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"time"
)

const GlobalTypeText string = "text"
const GlobalTypeNumber string = "number"
const GlobalTypeTime string = "time"
const GlobalTypeBool string = "bool"
const GlobalTypeLongText string = "longText"
const GlobalTypeFile string = "file"
const GlobalTypeImage string = "image"

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
	case GlobalTypeImage:
		return true
	default:
		return false
	}
}

type TextValue struct {
	Val string `json:"val"`
}

type ImageValue struct {
	CloudFileName string `json:"cloudFileName"`
	OrigFileName  string `json:"origFileName"`
}

type BoolValue struct {
	Val bool `json:"val"`
}

type TimeValue struct {
	Val time.Time `json:"val"`
}

type NumberValue struct {
	Val float64 `json:'val'`
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
	case GlobalTypeImage:
		var imageVal ImageValue
		if err := generic.DecodeJSONString(encodedVal, &imageVal); err != nil {
			return nil, fmt.Errorf("decodeGlobalValue: failure decoding image value: %v", err)
		}
		return imageVal, nil
	case GlobalTypeBool:
		var boolVal BoolValue
		if err := generic.DecodeJSONString(encodedVal, &boolVal); err != nil {
			return nil, fmt.Errorf("decodeGlobalValue: failure decoding boolean value: %v", err)
		}
		return boolVal.Val, nil
	case GlobalTypeTime:
		var timeVal TimeValue
		if err := generic.DecodeJSONString(encodedVal, &timeVal); err != nil {
			return nil, fmt.Errorf("decodeGlobalValue: failure decoding boolean value: %v", err)
		}
		return timeVal.Val, nil
	case GlobalTypeNumber:
		var numVal NumberValue
		if err := generic.DecodeJSONString(encodedVal, &numVal); err != nil {
			return nil, fmt.Errorf("decodeGlobalValue: failure decoding boolean value: %v", err)
		}
		return numVal.Val, nil
	default:
		return nil, fmt.Errorf("decodeGlobalValue: Unrecognized field type: %v", valueType)
	}
}
