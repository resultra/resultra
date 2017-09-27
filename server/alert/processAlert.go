package alert

import (
	"fmt"
	"log"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
	"sort"
	"time"
)

type AlertProcessingContext struct {
	CalcFieldConfig *calcField.CalcFieldUpdateConfig
	RecordID        string
	UpdateTimestamp time.Time
	PrevFieldVals   record.RecFieldValues
	CurrFieldVals   record.RecFieldValues
	LatestFieldVals record.RecFieldValues
	ProcessedAlert  Alert
}

const alertCondChange string = "changed"
const alertCondFalse string = "false"
const alertCondTrue string = "true"
const alertCondCleared string = "cleared"
const alertCondIncreased string = "increased"
const alertCondDecreased string = "decreased"
const alertCondAdded string = "added"
const alertCondRemoved string = "removed"
const alertCondCurrUserAdded string = "currUserAdded"

func generateAlertNotificationCaption(context AlertProcessingContext) (string, error) {

	preprocessedCaption := []byte(context.ProcessedAlert.Properties.CaptionMessage)

	formatFieldValue := func(fieldIDRef string) (string, bool) {
		for _, currFieldInfo := range context.CalcFieldConfig.Fields {
			if fieldIDRef == currFieldInfo.FieldID {
				switch currFieldInfo.Type {
				case field.FieldTypeText:
					val, foundVal := context.CurrFieldVals.GetTextFieldValue(currFieldInfo.FieldID)
					if !foundVal {
						return "[no value]", true
					} else {
						return val, true
					}
				case field.FieldTypeNumber:
					val, foundVal := context.CurrFieldVals.GetNumberFieldValue(currFieldInfo.FieldID)
					if !foundVal {
						return "[no value]", true
					} else {
						return fmt.Sprintf("%v", val), true
					}
				case field.FieldTypeBool:
					val, foundVal := context.CurrFieldVals.GetBoolFieldValue(currFieldInfo.FieldID)
					if !foundVal {
						return "[no value]", true
					} else {
						if val == true {
							return "true", true
						} else {
							return "false", true
						}
					}
				case field.FieldTypeTime:
					val, foundVal := context.CurrFieldVals.GetTimeFieldValue(currFieldInfo.FieldID)
					if !foundVal {
						return "[no date]", true
					} else {
						return val.Format("01/02/2006"), true
					}
				default:
					return "", false
				} // switch

			}
		}
		return "", false
	}

	replaceFieldRefWithFieldVal := func(s []byte) []byte {

		fieldIDRef := string(s[1 : len(s)-1])
		fieldRefVal, valFound := formatFieldValue(fieldIDRef)
		if valFound {
			return []byte(fieldRefVal)
		} else {
			return s // no replacement
		}
	}

	mergedCaption := msgTemplateFieldRefPattern.ReplaceAllFunc(preprocessedCaption, replaceFieldRefWithFieldVal)

	return string(mergedCaption), nil
}

func processTimeFieldAlert(context AlertProcessingContext, cond AlertCondition) (*AlertNotification, error) {

	log.Printf("Processing time field alert: %+v", cond)

	valBefore, foundValBefore := context.PrevFieldVals.GetTimeFieldValue(cond.FieldID)
	valAfter, foundValAfter := context.CurrFieldVals.GetTimeFieldValue(cond.FieldID)

	valChanged := func() bool {
		if (foundValBefore == false) && (foundValAfter == false) {
			return false
		} else if (foundValBefore == false) && (foundValAfter == true) {
			return true
		} else if (foundValBefore == true) && (foundValAfter == false) {
			return true
		} else { // value found bothe before and after update => compare the actual values
			if valBefore.Equal(valAfter) {
				return false
			} else {
				return true
			}
		}

	}

	generateNotification := func() (*AlertNotification, error) {

		captionMsg, err := generateAlertNotificationCaption(context)
		if err != nil {
			return nil, fmt.Errorf("processTimeFieldAlert: %v", err)
		}

		alertNofify := AlertNotification{
			AlertID:          context.ProcessedAlert.AlertID,
			RecordID:         context.RecordID,
			Timestamp:        context.UpdateTimestamp,
			Caption:          captionMsg,
			TriggerCondition: cond}

		return &alertNofify, nil

	}

	switch cond.ConditionID {
	case alertCondChange:
		if valChanged() {
			return generateNotification()
		}
	case alertCondCleared:
		if valChanged() && (!foundValAfter) {
			return generateNotification()
		}
	case alertCondIncreased:
		if foundValBefore && foundValAfter && (valAfter.After(valBefore)) {
			return generateNotification()
		}
	case alertCondDecreased:
		if foundValBefore && foundValAfter && (valAfter.Before(valBefore)) {
			return generateNotification()
		}
	default:
		return nil, nil
	}

	return nil, nil
}

func processBoolFieldAlert(context AlertProcessingContext, cond AlertCondition) (*AlertNotification, error) {

	log.Printf("Processing bool field alert: %+v", cond)

	valBefore, foundValBefore := context.PrevFieldVals.GetBoolFieldValue(cond.FieldID)
	valAfter, foundValAfter := context.CurrFieldVals.GetBoolFieldValue(cond.FieldID)

	valChanged := func() bool {
		if (foundValBefore == false) && (foundValAfter == false) {
			return false
		} else if (foundValBefore == false) && (foundValAfter == true) {
			return true
		} else if (foundValBefore == true) && (foundValAfter == false) {
			return true
		} else { // value found bothe before and after update => compare the actual values
			if valBefore == valAfter {
				return false
			} else {
				return true
			}
		}

	}

	generateNotification := func() (*AlertNotification, error) {

		captionMsg, err := generateAlertNotificationCaption(context)
		if err != nil {
			return nil, fmt.Errorf("processTimeFieldAlert: %v", err)
		}

		alertNofify := AlertNotification{
			AlertID:          context.ProcessedAlert.AlertID,
			RecordID:         context.RecordID,
			Timestamp:        context.UpdateTimestamp,
			Caption:          captionMsg,
			TriggerCondition: cond}

		return &alertNofify, nil

	}

	switch cond.ConditionID {
	case alertCondChange:
		if valChanged() {
			return generateNotification()
		}
	case alertCondFalse:
		if valChanged() && foundValAfter && (valAfter == false) {
			return generateNotification()
		}
	case alertCondTrue:
		if valChanged() && foundValAfter && (valAfter == true) {
			return generateNotification()
		}
	case alertCondCleared:
		if valChanged() && (!foundValAfter) {
			return generateNotification()
		}
	default:
		return nil, nil
	}

	return nil, nil
}

func processNumberFieldAlert(context AlertProcessingContext, cond AlertCondition) (*AlertNotification, error) {

	log.Printf("Processing number field alert: %+v", cond)

	valBefore, foundValBefore := context.PrevFieldVals.GetNumberFieldValue(cond.FieldID)
	valAfter, foundValAfter := context.CurrFieldVals.GetNumberFieldValue(cond.FieldID)

	valChanged := func() bool {
		if (foundValBefore == false) && (foundValAfter == false) {
			return false
		} else if (foundValBefore == false) && (foundValAfter == true) {
			return true
		} else if (foundValBefore == true) && (foundValAfter == false) {
			return true
		} else { // value found both before and after update => compare the actual values
			if valBefore == valAfter {
				return false
			} else {
				return true
			}
		}

	}

	generateNotification := func() (*AlertNotification, error) {

		captionMsg, err := generateAlertNotificationCaption(context)
		if err != nil {
			return nil, fmt.Errorf("processTimeFieldAlert: %v", err)
		}

		alertNofify := AlertNotification{
			AlertID:          context.ProcessedAlert.AlertID,
			RecordID:         context.RecordID,
			Timestamp:        context.UpdateTimestamp,
			Caption:          captionMsg,
			TriggerCondition: cond}

		return &alertNofify, nil

	}

	switch cond.ConditionID {
	case alertCondChange:
		if valChanged() {
			return generateNotification()
		}
	case alertCondIncreased:
		if foundValBefore && foundValAfter && (valAfter > valBefore) {
			return generateNotification()
		}
	case alertCondDecreased:
		if foundValBefore && foundValAfter && (valAfter < valBefore) {
			return generateNotification()
		}
	case alertCondCleared:
		if valChanged() && (!foundValAfter) {
			return generateNotification()
		}
	default:
		return nil, nil
	}

	return nil, nil
}

func processCommentFieldAlert(context AlertProcessingContext, cond AlertCondition) (*AlertNotification, error) {

	log.Printf("Processing comment field alert: %+v", cond)

	valBefore, foundValBefore := context.PrevFieldVals.GetCommentFieldValue(cond.FieldID)
	valAfter, foundValAfter := context.CurrFieldVals.GetCommentFieldValue(cond.FieldID)

	log.Printf("Processing comment field alert: found before=%v before=%+v found after=%v after=%+v",
		foundValBefore, valBefore, foundValAfter, valAfter)

	valAdded := func() bool {
		if (foundValBefore == false) && (foundValAfter == false) {
			return false
		} else if (foundValBefore == false) && (foundValAfter == true) {
			return true
		} else if (foundValBefore == true) && (foundValAfter == false) {
			return false
		} else { // value found both before and after update => compare the actual values

			// TBD - The follownig isn't really a good measure if a comment has been added. In particular,
			// it doesn't handle the case where the same comment is added over and over. For example,
			// If someone posts a comment like 'reply with "yes" if you agree' What is needed is to
			// also check the timestamps on the 2 values.
			if valAfter.CommentText != valBefore.CommentText {
				return true
			} else {
				return false
			}
		}

	}

	generateNotification := func() (*AlertNotification, error) {

		captionMsg, err := generateAlertNotificationCaption(context)
		if err != nil {
			return nil, fmt.Errorf("processTimeFieldAlert: %v", err)
		}

		alertNofify := AlertNotification{
			AlertID:          context.ProcessedAlert.AlertID,
			RecordID:         context.RecordID,
			Timestamp:        context.UpdateTimestamp,
			Caption:          captionMsg,
			TriggerCondition: cond}

		return &alertNofify, nil

	}

	switch cond.ConditionID {
	case alertCondAdded:
		if valAdded() {
			return generateNotification()
		}
	default:
		return nil, nil
	}

	return nil, nil
}

func sameStringVals(vals1, vals2 []string) bool {
	if vals1 == nil && vals2 == nil {
		return true
	}
	if vals1 == nil || vals2 == nil {
		return false
	}
	if len(vals1) != len(vals2) {
		return false
	}

	sort.Strings(vals1)
	sort.Strings(vals2)

	for i := range vals1 {
		if vals1[i] != vals2[i] {
			return false
		}
	}
	return true
}

func valMemberOfValSet(val string, vals []string) bool {
	for _, currVal := range vals {
		if currVal == val {
			return true
		}
	}
	return false
}

func processUserFieldAlert(context AlertProcessingContext, cond AlertCondition) (*AlertNotification, error) {

	log.Printf("Processing user field alert: %+v", cond)

	valBefore, foundValBefore := context.PrevFieldVals.GetUsersFieldValue(cond.FieldID)
	valAfter, foundValAfter := context.CurrFieldVals.GetUsersFieldValue(cond.FieldID)
	log.Printf("Processing user field alert: found before=%v before=%+v found after=%v after=%+v",
		foundValBefore, valBefore, foundValAfter, valAfter)

	valChanged := func() bool {
		if (foundValBefore == false) && (foundValAfter == false) {
			return false
		} else if (foundValBefore == false) && (foundValAfter == true) {
			return true
		} else if (foundValBefore == true) && (foundValAfter == false) {
			return true
		} else { // value found both before and after update => compare the actual values
			if sameStringVals(valBefore, valAfter) {
				return true
			} else {
				return false
			}
		}

	}

	generateNotification := func() (*AlertNotification, error) {

		captionMsg, err := generateAlertNotificationCaption(context)
		if err != nil {
			return nil, fmt.Errorf("processTimeFieldAlert: %v", err)
		}

		alertNofify := AlertNotification{
			AlertID:          context.ProcessedAlert.AlertID,
			RecordID:         context.RecordID,
			Timestamp:        context.UpdateTimestamp,
			Caption:          captionMsg,
			TriggerCondition: cond}

		return &alertNofify, nil

	}

	switch cond.ConditionID {
	case alertCondChange:
		if valChanged() {
			return generateNotification()
		}
	case alertCondCleared:
		if valChanged() && (len(valAfter) == 0) {
			return generateNotification()
		}
	case alertCondIncreased:
		if valChanged() && (len(valAfter) > len(valBefore)) {
			return generateNotification()
		}
	case alertCondDecreased:
		if valChanged() && (len(valAfter) < len(valBefore)) {
			return generateNotification()
		}
	case alertCondCurrUserAdded:
		currUserID := context.CalcFieldConfig.CurrUserID
		if valChanged() &&
			!valMemberOfValSet(currUserID, valBefore) &&
			valMemberOfValSet(currUserID, valAfter) {
			return generateNotification()
		}
	default:
		return nil, nil
	}

	return nil, nil
}

// processAlert processs a single alert for a single set of previous and current (before and after)
// field values (including calculated fields).
func processAlert(context AlertProcessingContext) (*AlertNotification, error) {

	if context.ProcessedAlert.Properties.Condition == nil {
		return nil, nil
	}
	currAlertCond := *context.ProcessedAlert.Properties.Condition

	condFieldID := currAlertCond.FieldID
	fieldInfo, fieldInfoFound := context.CalcFieldConfig.FieldsByID[condFieldID]
	if !fieldInfoFound {
		return nil, fmt.Errorf("GenerateRecordAlerts: missing field info for field id = %v", condFieldID)
	}
	switch fieldInfo.Type {
	//			case field.FieldTypeText:
	//			case field.FieldTypeNumber:
	case field.FieldTypeTime:
		alertNotification, genErr := processTimeFieldAlert(context, currAlertCond)
		if genErr != nil {
			return nil, fmt.Errorf("processAlert:  %v", genErr)
		} else if alertNotification != nil {
			log.Printf("Alert generated: alert = %v, field = %v, condition = %v",
				context.ProcessedAlert.Name, fieldInfo.Name, currAlertCond.ConditionID)
			return alertNotification, nil // No need to process after matching the first condition
		}
	case field.FieldTypeNumber:
		alertNotification, genErr := processNumberFieldAlert(context, currAlertCond)
		if genErr != nil {
			return nil, fmt.Errorf("processAlert:  %v", genErr)
		} else if alertNotification != nil {
			log.Printf("Alert generated: alert = %v, field = %v, condition = %v",
				context.ProcessedAlert.Name, fieldInfo.Name, currAlertCond.ConditionID)
			return alertNotification, nil // No need to process after matching the first condition
		}
	case field.FieldTypeBool:
		alertNotification, genErr := processBoolFieldAlert(context, currAlertCond)
		if genErr != nil {
			return nil, fmt.Errorf("processAlert:  %v", genErr)
		} else if alertNotification != nil {
			log.Printf("Alert generated: alert = %v, field = %v, condition = %v",
				context.ProcessedAlert.Name, fieldInfo.Name, currAlertCond.ConditionID)
			return alertNotification, nil // No need to process after matching the first condition
		}
	case field.FieldTypeComment:
		alertNotification, genErr := processCommentFieldAlert(context, currAlertCond)
		if genErr != nil {
			return nil, fmt.Errorf("processAlert:  %v", genErr)
		} else if alertNotification != nil {
			log.Printf("Alert generated: alert = %v, field = %v, condition = %v",
				context.ProcessedAlert.Name, fieldInfo.Name, currAlertCond.ConditionID)
			return alertNotification, nil // No need to process after matching the first condition
		}
	case field.FieldTypeUser:
		alertNotification, genErr := processUserFieldAlert(context, currAlertCond)
		if genErr != nil {
			return nil, fmt.Errorf("processAlert:  %v", genErr)
		} else if alertNotification != nil {
			log.Printf("Alert generated: alert = %v, field = %v, condition = %v",
				context.ProcessedAlert.Name, fieldInfo.Name, currAlertCond.ConditionID)
			return alertNotification, nil // No need to process after matching the first condition
		}
	default:
		return nil, fmt.Errorf("processAlert: Unsupported field type for alert: %v", fieldInfo.Type)
	} // switch

	return nil, nil
}
