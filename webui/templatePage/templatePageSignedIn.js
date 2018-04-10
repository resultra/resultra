function initMyTemplateList() {
	
	var $templateList = $("#myTemplateList")
	var $showInactiveCheckbox = $('#showInactiveTemplates')
	
	function reloadTemplateList(includeInactive) {
		var getTemplateListParams = {
			includeInactive:includeInactive
		}
		jsonAPIRequest("database/getUserTemplateList",getTemplateListParams,function(templateList) {
			$templateList.empty()
			for (var templIndex=0; templIndex<templateList.length; templIndex++) {	
				var templateInfo = templateList[templIndex]
				addTemplateListItem(templateInfo)
			}
		})
	}
	

	function addTemplateListItem(templateInfo) {

		var $listItem = $('#templateListItemTemplate').clone()
		$listItem.attr("id","")

		var $nameLabel = $listItem.find(".nameLabel")
		$nameLabel.text(templateInfo.databaseName)
		if (!templateInfo.isActive) {
			$nameLabel.addClass("disabledTemplateName")
		}
	
		// Only enable the link to open the tracker if the tracker is  active.	
		var $settingsLink = $listItem.find(".templateSettingsButton")
		
		$settingsLink.click(function() {
			
			function propertiesDialogHidden() {
				var includeInactive = $showInactiveCheckbox.prop("checked")
				reloadTemplateList(includeInactive)
			}
			
			openTemplatePropertiesDialog(templateInfo,propertiesDialogHidden)
		})
		
		$templateList.append($listItem)
	
	}
		
	reloadTemplateList(false)
	
		
	initCheckboxControlChangeHandler($showInactiveCheckbox, false, function(includeInactive) {
		reloadTemplateList(includeInactive)
	})
	
	var $backToTrackersButton = $('#backToTrackersButton')
	$backToTrackersButton.click(function() {
		navigateToMainWindowContent("workspaceHome")
	})
	
}

function initTemplatePageSignedInPageContent() {	
	initMyTemplateList()
}

function navigateToTemplatesPage() {
	
	theMainWindowLayout.disableRHSSidebar()
	theMainWindowLayout.disableLHSSidebar()
	clearMainWindowHeaderButtonsContent()
	resetWorkspaceBreadcrumbHeader()
		
	setMainWindowContent('/templatePage/mainContent',function() {
		initTemplatePageSignedInPageContent()
	})
		
	setMainWindowOffPageContent('/templatePage/offPageContent',function() {
	})
	
	
	
}
