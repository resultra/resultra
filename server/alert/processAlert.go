package alert

import (
	"fmt"
	"log"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
	"time"
)

type AlertProcessingContext struct {
	CalcFieldConfig *calcField.CalcFieldUpdateConfig
	RecordID        string
	UpdateTimestamp time.Time
	PrevFieldVals   record.RecFieldValues
	CurrFieldVals   record.RecFieldValues
	LatestFieldVals record.RecFieldValues
	ItemSummary     string
	ProcessedAlert  Alert
}

const alertCondChange string = "changed"
const alertCondFalse string = "false"
const alertCondTrue string = "true"
const alertCondCleared string = "cleared"
const alertCondIncreased string = "increased"
const alertCondDecreased string = "decreased"

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

	captionMsg, err := generateAlertNotificationCaption(context)
	if err != nil {
		return nil, fmt.Errorf("processTimeFieldAlert: %v", err)
	}

	switch cond.ConditionID {
	case alertCondChange:
		valBefore, foundValBefore := context.PrevFieldVals.GetTimeFieldValue(cond.FieldID)
		valAfter, foundValAfter := context.CurrFieldVals.GetTimeFieldValue(cond.FieldID)
		if (foundValBefore == false) && (foundValAfter == false) {
			log.Printf("processTimeFieldAlert: no change - both vals undefined")
			return nil, nil
		} else if (foundValBefore == false) && (foundValAfter == true) {
			log.Printf("processTimeFieldAlert: change - value defined: %v", valAfter)

			alertNofify := AlertNotification{
				AlertID:          context.ProcessedAlert.AlertID,
				RecordID:         context.RecordID,
				Timestamp:        context.UpdateTimestamp,
				ItemSummary:      context.ItemSummary,
				Caption:          captionMsg,
				TriggerCondition: cond,
				DateBefore:       &valBefore,
				DateAfter:        &valAfter}

			return &alertNofify, nil
		} else if (foundValBefore == true) && (foundValAfter == false) {
			log.Printf("processTimeFieldAlert: change - time value cleared: cleared val = %v", valBefore)
			alertNofify := AlertNotification{
				AlertID:          context.ProcessedAlert.AlertID,
				RecordID:         context.RecordID,
				Timestamp:        context.UpdateTimestamp,
				ItemSummary:      context.ItemSummary,
				Caption:          captionMsg,
				TriggerCondition: cond,
				DateBefore:       &valBefore,
				DateAfter:        &valAfter}
			return &alertNofify, nil
		} else { // value found bother before and after update => compare the actual dates
			if valBefore.Equal(valAfter) {
				log.Printf("processTimeFieldAlert: no change - values equal %v", valBefore)
				return nil, nil // no change
			} else {
				log.Printf("processTimeFieldAlert: change - time values changed: %v -> %v", valBefore, valAfter)
				alertNofify := AlertNotification{
					AlertID:          context.ProcessedAlert.AlertID,
					RecordID:         context.RecordID,
					Timestamp:        context.UpdateTimestamp,
					ItemSummary:      context.ItemSummary,
					Caption:          captionMsg,
					TriggerCondition: cond,
					DateBefore:       &valBefore,
					DateAfter:        &valAfter}
				return &alertNofify, nil
			}
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
			ItemSummary:      context.ItemSummary,
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
			ItemSummary:      context.ItemSummary,
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

// processAlert processs a single alert for a single set of previous and current (before and after)
// field values (including calculated fields).
func processAlert(context AlertProcessingContext) (*AlertNotification, error) {
	for _, currAlertCond := range context.ProcessedAlert.Properties.Conditions {

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
		default:
			return nil, fmt.Errorf("processAlert: Unsupported field type for alert: %v", fieldInfo.Type)
		} // switch

	}
	return nil, nil
}
