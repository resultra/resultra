
function updateAlertConditions(panelParams) {
	
}

function createAlertConditionListItem(panelParams,fieldName) {
	
	var $listItem = $('#alertConditionListItem').clone()
	$listItem.attr("id","")
	
	var $fieldLabel = $listItem.find(".alertConditionItemLabel")
	$fieldLabel.text(fieldName)
	
	var $deleteButton = $listItem.find(".alertConditionDeleteButton")
	initButtonControlClickHandler($deleteButton,function() {
		$listItem.remove()
		updateFilterRules(panelParams)
	})

	return $listItem
}

function dateAlertConditionListItem(propsParams,fieldInfo,defaultConditionInfo) {
	
	var $listItem = createAlertConditionListItem(propsParams,fieldInfo.name)
	
	var $alertProps = $("#alertDateFieldConditionProps").clone()
	$alertProps.attr("id","")
	
	
	$listItem.append($alertProps)
	
	return $listItem
}

function createAlertPropsConditionItem(propsParams,fieldInfo,defaultConditionInfo) {

	switch (fieldInfo.type) {
	case fieldTypeTime: 
		return dateAlertConditionListItem(propsParams,fieldInfo,defaultConditionInfo)
	case fieldTypeText:
	case fieldTypeNumber:
	case fieldTypeBool: 
	default:
		console.log("createFilterRulePanelListItem: Unsupported field type:  " + fieldInfo.type)
		return $("<div>TBD</div>")
	}
	
	
}

function initAlertConditionProps(params) {
			
	var fieldSelectionDropdownParams = {
		elemPrefix: "alertCondition_",
		databaseID: params.databaseID,
		fieldTypes: [fieldTypeTime],
		fieldSelectionCallback: function(fieldInfo) {
			
			var $alertConditionList = $("#alertConditionList")
			
			// Use null to signify no default condition information. This is true when
			// creating new rules, but will not be when re-loading the rules.
			var defaultConditionInfo = null
			$alertConditionList.append(createAlertPropsConditionItem(params,fieldInfo,defaultConditionInfo))
		}
	}
	initFieldSelectionDropdown(fieldSelectionDropdownParams)
	
	
}