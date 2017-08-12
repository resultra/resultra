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

func processTimeFieldAlert(context AlertProcessingContext, cond AlertCondition) (bool, error) {

	log.Printf("Processing time field alert: %+v", cond)

	switch cond.ConditionID {
	case alertCondChange:
		valBefore, foundValBefore := context.PrevFieldVals.GetTimeFieldValue(cond.FieldID)
		valAfter, foundValAfter := context.CurrFieldVals.GetTimeFieldValue(cond.FieldID)
		if (foundValBefore == false) && (foundValAfter == false) {
			log.Printf("processTimeFieldAlert: no change - both vals undefined")
			return false, nil
		} else if (foundValBefore == false) && (foundValAfter == true) {
			log.Printf("processTimeFieldAlert: change - value defined: %v", valAfter)
			return true, nil
		} else if (foundValBefore == true) && (foundValAfter == false) {
			log.Printf("processTimeFieldAlert: change - time value cleared: cleared val = %v", valBefore)
			return true, nil
		} else { // value found bother before and after update => compare the actual dates
			if valBefore.Equal(valAfter) {
				log.Printf("processTimeFieldAlert: no change - values equal %v", valBefore)
				return false, nil // no change
			} else {
				log.Printf("processTimeFieldAlert: change - time values changed: %v -> %v", valBefore, valAfter)
				return true, nil
			}
		}
	default:
		return false, nil
	}

	return false, nil
}

func processAlert(context AlertProcessingContext) error {
	for _, currAlertCond := range context.ProcessedAlert.Properties.Conditions {

		condFieldID := currAlertCond.FieldID
		fieldInfo, fieldInfoFound := context.CalcFieldConfig.FieldsByID[condFieldID]
		if !fieldInfoFound {
			return fmt.Errorf("GenerateRecordAlerts: missing field info for field id = %v", condFieldID)
		}
		switch fieldInfo.Type {
		//			case field.FieldTypeText:
		//			case field.FieldTypeNumber:
		//			case field.FieldTypeBool:
		case field.FieldTypeTime:
			alertGenerated, genErr := processTimeFieldAlert(context, currAlertCond)
			if genErr != nil {
				return fmt.Errorf("processAlert:  %v", genErr)
			} else if alertGenerated == true {
				log.Printf("Alert generated: alert = %v, field = %v, condition = %v",
					context.ProcessedAlert.Name, fieldInfo.Name, currAlertCond.ConditionID)
				return nil // No need to process after matching the first condition
			}
		default:
			return fmt.Errorf("processAlert: Unsupported field type for alert: %v", fieldInfo.Type)
		} // switch

	}
	return nil
}
