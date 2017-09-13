package alert

import (
	"fmt"
	"log"
	"regexp"
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

func generateAlertNotificationCaption(context AlertProcessingContext) (string, error) {

	preprocessedCaption := []byte(context.ProcessedAlert.Properties.CaptionMessage)

	fieldReplSearch := regexp.MustCompile(`\[[a-zA-Z0-9_]+\]`)

	formatFieldValue := func(fieldRefName string) string {
		for _, currFieldInfo := range context.CalcFieldConfig.Fields {
			if fieldRefName == currFieldInfo.RefName {
				switch currFieldInfo.Type {
				//			case field.FieldTypeText:
				//			case field.FieldTypeNumber:
				//			case field.FieldTypeBool:
				case field.FieldTypeTime:
					val, foundVal := context.CurrFieldVals.GetTimeFieldValue(currFieldInfo.FieldID)
					if !foundVal {
						return "(no date)"
					} else {
						return val.Format("01/02/2006")
					}
				default:
					return fieldRefName
				} // switch

			}
		}
		return fieldRefName
	}

	replaceFieldRefWithFieldVal := func(s []byte) []byte {
		fieldRefName := string(s[1 : len(s)-1])
		log.Printf("generateAlertNotificationCaption: found field %v", fieldRefName)
		return []byte(formatFieldValue(fieldRefName))
	}

	mergedCaption := fieldReplSearch.ReplaceAllFunc(preprocessedCaption, replaceFieldRefWithFieldVal)

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
		//			case field.FieldTypeBool:
		case field.FieldTypeTime:
			alertNotification, genErr := processTimeFieldAlert(context, currAlertCond)
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
