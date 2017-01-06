package record

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"time"
)

type NumberCellValue struct {
	Val float64 `json:'val'`
}

type BoolCellValue struct {
	Val bool `json:"val"`
}

type FileCellValue struct {
	CloudName string `json:"cloudName"`
	OrigName  string `json:"origName"`
}

type UserCellValue struct {
	UserID string `json:"userID"`
}

type TimeCellValue struct {
	Val time.Time `json:"val"`
}

type TextCellValue struct {
	Val string `json:"val"`
}

func DecodeCellValue(fieldType string, encodedVal string) (interface{}, error) {
	switch fieldType {
	case field.FieldTypeText:
		var textVal TextCellValue
		if err := generic.DecodeJSONString(encodedVal, &textVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding text value: %v", err)
		}
		return textVal.Val, nil
	case field.FieldTypeNumber:
		var numberVal NumberCellValue
		if err := generic.DecodeJSONString(encodedVal, &numberVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding number value: %v", err)
		}
		return numberVal.Val, nil
	case field.FieldTypeTime:
		var timeVal TimeCellValue
		if err := generic.DecodeJSONString(encodedVal, &timeVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding time value: %v", err)
		}
		return timeVal.Val, nil
	case field.FieldTypeBool:
		var boolVal BoolCellValue
		if err := generic.DecodeJSONString(encodedVal, &boolVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding boolean value: %v", err)
		}
		return boolVal.Val, nil
	case field.FieldTypeLongText:
		var textVal TextCellValue
		if err := generic.DecodeJSONString(encodedVal, &textVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding long text value: %v", err)
		}
		return textVal.Val, nil
	case field.FieldTypeComment:
		var textVal TextCellValue
		if err := generic.DecodeJSONString(encodedVal, &textVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding long text value: %v", err)
		}
		return textVal.Val, nil
	case field.FieldTypeFile:
		var fileVal FileCellValue
		if err := generic.DecodeJSONString(encodedVal, &fileVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding file value: %v", err)
		}
		return fileVal, nil
	case field.FieldTypeUser:
		var userVal UserCellValue
		if err := generic.DecodeJSONString(encodedVal, &userVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding user value: %v", err)
		}
		return userVal.UserID, nil
	default:
		return nil, fmt.Errorf("DecodeCellValue: Unrecognized field type: %v", fieldType)

	}
}
