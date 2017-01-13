
var linkedComponentValTypeField = "field"
var linkedComponentValTypeGlobal = "global"

function fieldComponentValType(fieldID) {
	// a ComponentLink can be synthesized with just a field ID.
	var componentLink = {
		linkedValType: linkedComponentValTypeField,
		fieldID: fieldID,
		globalID: ""
	}
	
	return componentLink
}