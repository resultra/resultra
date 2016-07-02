
var newFieldDialogPanelID = "newField"


function createNewFieldDialogPanelContextBootstrap(elemPrefix) {
	
	var panelSelector = "#" + elemPrefix + "NewFieldPanel"
	var fieldRefNameInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldRefName")
	var fieldNameInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldName")
	var refNameHelpPopup = createPrefixedTemplElemInfo(elemPrefix,"RefNameHelp")
	var isCalcFieldField = createPrefixedTemplElemInfo(elemPrefix,"NewFieldCalcFieldCheckbox")
	var isCalcFieldInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldIsCalcFieldInput")
	var fieldTypeSelection = createPrefixedTemplElemInfo(elemPrefix,"NewFieldValTypeSelection")
	
	var dialogProgressSelector = "#" + elemPrefix + "NewFormElemDialogProgress"
	
	function validateForm() {
		
	}
	
	var newFieldPanelConfig = {
		panelID: newFieldDialogPanelID,
		divID: panelSelector,
		progressPerc:60,
		dlgButtons: null, // todo - initialize buttons with Bootstrap based button handling
		initPanel: function ($parentDialog) {
			
			var prevButtonSelector = '#' + elemPrefix + 'NewFormComponentSelectFieldPrevButton'
			$(prevButtonSelector).unbind("click")
			$(prevButtonSelector).click(function(e) {
				$(this).blur();
			    e.preventDefault();// prevent the default anchor functionality
				
				transitionToPrevWizardDlgPanelByPanelID($parentDialog,createNewOrExistingFieldDialogPanelID)
			})
			
		}, // init panel
		transitionIntoPanel: function ($dialog) { }
	} // wizard dialog configuration for panel to create new field
	
	return newFieldPanelConfig;
	
}

