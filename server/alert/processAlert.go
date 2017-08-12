package alert

import (
	"fmt"
	"log"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
)

type AlertProcessingContext struct {
	CalcFieldConfig *calcField.CalcFieldUpdateConfig
	PrevFieldVals   record.RecFieldValues
	CurrFieldVals   record.RecFieldValues
	ProcessedAlert  Alert
}

const alertCondChange string = "changed"

type AlertNotification struct {
	AlertInfo Alert       `json:"alertInfo"`
	ValBefore interface{} `json:"valBefore"`
	ValAfter  interface{} `json:"valAfter"`
}

func processTimeFieldAlert(context AlertProcessingContext, cond AlertCondition) (*AlertNotification, error) {

	log.Printf("Processing time field alert: %+v", cond)

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
				AlertInfo: context.ProcessedAlert,
				ValBefore: valBefore,
				ValAfter:  valAfter}

			return &alertNofify, nil
		} else if (foundValBefore == true) && (foundValAfter == false) {
			log.Printf("processTimeFieldAlert: change - time value cleared: cleared val = %v", valBefore)
			alertNofify := AlertNotification{
				AlertInfo: context.ProcessedAlert,
				ValBefore: valBefore,
				ValAfter:  valAfter}
			return &alertNofify, nil
		} else { // value found bother before and after update => compare the actual dates
			if valBefore.Equal(valAfter) {
				log.Printf("processTimeFieldAlert: no change - values equal %v", valBefore)
				return nil, nil // no change
			} else {
				log.Printf("processTimeFieldAlert: change - time values changed: %v -> %v", valBefore, valAfter)
				alertNofify := AlertNotification{
					AlertInfo: context.ProcessedAlert,
					ValBefore: valBefore,
					ValAfter:  valAfter}
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
