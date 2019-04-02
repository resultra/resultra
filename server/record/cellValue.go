// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package record

import (
	"fmt"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/generic"
	"time"
)

type NumberCellValue struct {
	Val *float64 `json:'val'`
}

type BoolCellValue struct {
	Val *bool `json:"val"`
}

type AttachmentCellValue struct {
	Attachments []string `json:"attachments"`
}

type FileCellValue struct {
	Attachment *string `json:"attachment"`
}

type ImageCellValue struct {
	Attachment *string `json:"attachment"`
}

type UserCellValue struct {
	UserID string `json:"userID"`
}

type UsersCellValue struct {
	UserIDs []string `json:"userIDs"`
}

type LabelCellValue struct {
	Labels []string `json:"labels"`
}

type TimeCellValue struct {
	Val *time.Time `json:"val"`
}

type TextCellValue struct {
	Val *string `json:"val"`
}

type EmailAddrCellValue struct {
	Val *string `json:"val"`
}

type UrlCellValue struct {
	Val *string `json:"val"`
}

type CommentCellValue struct {
	CommentText string   `json:"commentText"`
	Attachments []string `json:"attachments"`
}

func DecodeCellValue(fieldType string, encodedVal string) (interface{}, error) {
	switch fieldType {
	case field.FieldTypeText:
		var textVal TextCellValue
		if err := generic.DecodeJSONString(encodedVal, &textVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding text value: %v", err)
		}
		if textVal.Val == nil {
			return nil, nil
		} else {
			return *(textVal.Val), nil
		}
	case field.FieldTypeEmail:
		var addrVal EmailAddrCellValue
		if err := generic.DecodeJSONString(encodedVal, &addrVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding text value: %v", err)
		}
		if addrVal.Val == nil {
			return nil, nil
		} else {
			return *(addrVal.Val), nil
		}
	case field.FieldTypeURL:
		var urlVal UrlCellValue
		if err := generic.DecodeJSONString(encodedVal, &urlVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding text value: %v", err)
		}
		if urlVal.Val == nil {
			return nil, nil
		} else {
			return *(urlVal.Val), nil
		}
	case field.FieldTypeNumber:
		var numberVal NumberCellValue
		if err := generic.DecodeJSONString(encodedVal, &numberVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding number value: %v", err)
		}
		// The value is stored using a pointer to a float64. This value format allows NULL
		// values to be set when a number field's value is cleared. When retrieving the value,
		// either a nil(null) or literal value must be returned. The nil value is interpreted
		// as an undefined result in the calculated field evaluation.
		if numberVal.Val == nil {
			return nil, nil
		} else {
			return *(numberVal.Val), nil
		}
	case field.FieldTypeTime:
		var timeVal TimeCellValue
		if err := generic.DecodeJSONString(encodedVal, &timeVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding time value: %v", err)
		}
		if timeVal.Val == nil {
			return nil, nil
		} else {
			return *(timeVal.Val), nil
		}
		return timeVal.Val, nil
	case field.FieldTypeBool:
		var boolVal BoolCellValue
		if err := generic.DecodeJSONString(encodedVal, &boolVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding boolean value: %v", err)
		}
		if boolVal.Val == nil {
			return nil, nil
		} else {
			return *(boolVal.Val), nil
		}
	case field.FieldTypeLongText:
		var textVal TextCellValue
		if err := generic.DecodeJSONString(encodedVal, &textVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding long text value: %v", err)
		}
		if textVal.Val == nil {
			return nil, nil
		} else {
			return *(textVal.Val), nil
		}
	case field.FieldTypeComment:
		var commentVal CommentCellValue
		if err := generic.DecodeJSONString(encodedVal, &commentVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding comment value: %v", err)
		}
		return commentVal, nil
	case field.FieldTypeAttachment:
		var fileVal AttachmentCellValue
		if err := generic.DecodeJSONString(encodedVal, &fileVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding file value: %v", err)
		}
		return fileVal, nil
	case field.FieldTypeFile:
		var fileVal FileCellValue
		if err := generic.DecodeJSONString(encodedVal, &fileVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding file value: %v", err)
		}
		if fileVal.Attachment == nil {
			return nil, nil
		} else {
			return fileVal.Attachment, nil
		}
	case field.FieldTypeImage:
		var imageVal ImageCellValue
		if err := generic.DecodeJSONString(encodedVal, &imageVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding file value: %v", err)
		}
		if imageVal.Attachment == nil {
			return nil, nil
		} else {
			return imageVal.Attachment, nil
		}
	case field.FieldTypeUser:
		var userVal UserCellValue
		if err := generic.DecodeJSONString(encodedVal, &userVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding user value: %v", err)
		}
		if len(userVal.UserID) == 0 {
			return nil, nil
		} else {
			return userVal.UserID, nil
		}
	case field.FieldTypeUsers:
		var userVal UsersCellValue
		if err := generic.DecodeJSONString(encodedVal, &userVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding user value: %v", err)
		}
		if userVal.UserIDs == nil {
			return nil, nil
		} else {
			return userVal.UserIDs, nil
		}
	case field.FieldTypeLabel:
		var labelVal LabelCellValue
		if err := generic.DecodeJSONString(encodedVal, &labelVal); err != nil {
			return nil, fmt.Errorf("DecodeCellValue: failure decoding user value: %v", err)
		}
		if labelVal.Labels == nil {
			return nil, nil
		} else {
			return labelVal.Labels, nil
		}
	default:
		return nil, fmt.Errorf("DecodeCellValue: Unrecognized field type: %v", fieldType)

	}
}
