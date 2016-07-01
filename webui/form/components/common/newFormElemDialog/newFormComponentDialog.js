function slideToNextDialogPanel(currPanelSelector, nextPanelSelector) {
	function showNextPanel() {
		$(nextPanelSelector).show("slide",{direction:"right"},200);
	}
	$(currPanelSelector).hide("slide",{direction:"left"},200,showNextPanel);
	
}

function slideToPrevDialogPanel(currPanelSelector, prevPanelSelector) {
	function showNextPanel() {
		$(prevPanelSelector).show("slide",{direction:"left"},200);
	}
	$(currPanelSelector).hide("slide",{direction:"right"},200,showNextPanel);
	
}

function updateWizardDialogProgress(elemPrefix,progressPerc) {
	
	var progressSelector = '#' + elemPrefix + 'WizardDialogProgress'
	
	$(progressSelector).css('width', progressPerc+'%').attr('aria-valuenow', progressPerc);
	
}


function openNewFormComponentDialog(elemPrefix) {
	
	var dlgSelector = '#' + elemPrefix + 'NewFormComponentDialog'
	var selectFieldSelector = '#' + elemPrefix + 'SelectExistingOrNewFieldPanel'
	var newFieldSelector =  '#' + elemPrefix + 'NewFieldPanel'
	
	$('.newFormComponentPanel').hide()
	$(selectFieldSelector).show()
	updateWizardDialogProgress(elemPrefix,10)
		
	$(dlgSelector).modal('show')
	
	var nextButtonSelector = '#' + elemPrefix + 'NewFormComponentNewFieldNextButton'
	$(nextButtonSelector).unbind("click")
	$(nextButtonSelector).click(function(e) {
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
		
		slideToNextDialogPanel(selectFieldSelector,newFieldSelector)
		updateWizardDialogProgress(elemPrefix,60)
				
	})
	
	
	var prevButtonSelector = '#' + elemPrefix + 'NewFormComponentSelectFieldPrevButton'
	$(prevButtonSelector).unbind("click")
	$(prevButtonSelector).click(function(e) {
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality

		slideToPrevDialogPanel(newFieldSelector,selectFieldSelector)
		updateWizardDialogProgress(elemPrefix,40)
	})
	
	
}