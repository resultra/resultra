
var paletteItemsEditConfig = {
	paletteItemTextBox: textBoxDesignFormConfig,
	paletteItemCheckBox: checkBoxDesignFormConfig,
	paletteItemToggle: toggleDesignFormConfig,
	paletteItemDatePicker: datePickerDesignFormConfig,
	paletteItemHtmlEditor: htmlEditorDesignFormConfig,
	paletteItemAttachment: attachmentDesignFormConfig,
	paletteItemImage: imageDesignFormConfig,
	paletteItemHeader: formHeaderDesignFormConfig,
	paletteItemRating: ratingDesignFormConfig,
	paletteItemSelection: selectionDesignFormConfig,
	paletteItemUserSelection: userSelectionDesignFormConfig,
	paletteItemUserTag: userTagDesignFormConfig,
	paletteItemComment: commentDesignFormConfig,
	paletteItemButton: formButtonDesignFormConfig,
	paletteItemProgress: progressDesignFormConfig,
	paletteItemGauge: gaugeDesignFormConfig,
	paletteItemCaption: formCaptionDesignFormConfig,
	paletteItemNumberInput: numberInputDesignFormConfig,
	paletteItemSocialButton: socialButtonDesignFormConfig,
	paletteItemLabel:labelDesignFormConfig,
	paletteItemEmailAddr: emailAddrDesignFormConfig,
	paletteItemUrlLink: urlLinkDesignFormConfig,
	paletteItemFile: fileDesignFormConfig
}

var designFormContext

function initDesignFormAdminPageContent(pageContext,formInfo) {
	
	
	designFormContext = { databaseID:formInfo.parentDatabaseID,
		formID:formInfo.formID,
		formName: formInfo.formName,
		isSingleUserWorkspace: pageContext.isSingleWorkspace }
	GlobalFormPagePrivs = "edit" 
	
	
		
	var designFormPaletteLayoutConfig =  {
		parentLayoutSelector: formDesignCanvasSelector,
		saveLayoutFunc: function(updatedLayout) { } // no-op: layout gets saved after placeholder replaced with real component.
	}
	
	
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) {
			return paletteItemsEditConfig[paletteItemID].draggableHTMLFunc(placeholderID)
		},
		
		initDummyDragAndDropComponentContainer: function(paletteItemID, $paletteItemContainer) {
			// If a palette item needs to initialize the dragged item after it's been
			// inserted into the DOM, then this is done in the initDummyDragAndDropComponentContainer function
			return paletteItemsEditConfig[paletteItemID].initDummyDragAndDropComponentContainer($paletteItemContainer)			
		},
				
		dropComplete: function(droppedItemInfo) {			
			// After the drag operation is complete, the resizable
			// properties need to be initialized.
			//
			// At this point, the placholder div for the bar chart will have just been inserted. However, the DOM may 
			// not be completely updated at this point. To ensure this, a small delay is needed before
			// drawing the dummy bar charts. See http://goo.gl/IloNM for more.
			var objEditConfig = paletteItemsEditConfig[droppedItemInfo.paletteItemID]
			
			setTimeout(function() {
				initObjectGridEditBehavior(droppedItemInfo.droppedElem,objEditConfig,designFormPaletteLayoutConfig) 
			}, 50);
					
			// "repackage" the dropped item paramaters for creating a new layout element. Also add the formID
			// to the parameters.
			var containerParams = {
				parentFormID: formInfo.formID,
				geometry: droppedItemInfo.geometry,
				containerID: droppedItemInfo.placeholderID,
				containerObj: droppedItemInfo.droppedElem,
				finalizeLayoutIncludingNewComponentFunc: droppedItemInfo.finalizeLayoutIncludingNewComponentFunc
				};
				
			objEditConfig.createNewItemAfterDropFunc(formInfo.parentDatabaseID,formInfo.formID,containerParams)
		},
		
		dropDestSelector: formDesignCanvasSelector,
		paletteSelector: "#paletteSidebar",
	}
	
	
	
	initDesignPalette(paletteConfig,designFormPaletteLayoutConfig)			
	
	// Initialize all the different plug-ins/configurations for text boxes, check boxes, etc.
	console.log("designForm: Initializing form design plug-ins/configurations ...")
	$.each(paletteItemsEditConfig, function (i, designFormConfig) {
		designFormConfig.initFunc()
	})
	console.log("designForm: Done initializing form design plug-ins/configurations.")
			
	// Initialize the page layout
	var formDesignLayout = $('#designFormLayoutContent').layout({
		north: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		south: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		// Important: The 'showOverflowOnHover' options give a higher
		// z-index to sidebars and other panels with controls, etc. Otherwise
		// popups and other controlls will not be shown on top of the rest
		// of the layout.
		north__showOverflowOnHover:	true,
		south__showOverflowOnHover:	true 
	})
	
	function showFormulaEditPane() { 
		hideSiblingsShowOne('#formDesignerFormulaEditor')
		formDesignLayout.resizeAll()
		formDesignLayout.open("south") 
	}
	
	function showDesignFormPalette() {
		hideSiblingsShowOne('#designFormPaletteItems')
		formDesignLayout.resizeAll()
		formDesignLayout.open("south")
	}
	
	function hideFormulaEditPanel() { 
		formDesignLayout.close("south")
	}
	
	var formulaEditorParams = {
		databaseID: formInfo.parentDatabaseID,
		showEditorFunc:showFormulaEditPane,
		hideEditorFunc:hideFormulaEditPanel
	}
	
		
	var designFormLayoutConfig =  createFormLayoutDesignConfig(formInfo.formID)
	var $parentFormLayout = $(formDesignCanvasSelector)
	
	var loadFormConfig = {
		$parentFormLayout: $parentFormLayout,
		formContext: designFormContext,
		initTextBoxFunc: function(componentContext,$textBox,textBoxObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: textBoxObjectRef.textBoxID }
			initFormComponentDesignBehavior($textBox,componentIDs,textBoxObjectRef,textBoxDesignFormConfig,designFormLayoutConfig)
		},
		initEmailAddrFunc: function(componentContext,$emailAddr,emailAddrObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: emailAddrObjectRef.emailAddrID }
			initFormComponentDesignBehavior($emailAddr,componentIDs,emailAddrObjectRef,emailAddrDesignFormConfig,designFormLayoutConfig)
		},
		initFileFunc: function(componentContext,$file,fileObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: fileObjectRef.fileID }
			initFormComponentDesignBehavior($file,componentIDs,fileObjectRef,fileDesignFormConfig,designFormLayoutConfig)
		},
		initUrlLinkFunc: function(componentContext,$urlLink,urlLinkObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: urlLinkObjectRef.urlLinkID }
			initFormComponentDesignBehavior($urlLink,componentIDs,urlLinkObjectRef,urlLinkDesignFormConfig,designFormLayoutConfig)
		},
		initNumberInputFunc: function(componentContext,$numberInput,numberInputObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: numberInputObjectRef.numberInputID }
			initFormComponentDesignBehavior($numberInput,componentIDs,numberInputObjectRef,numberInputDesignFormConfig,designFormLayoutConfig)
		},
		initSelectionFunc: function(componentContext,$selection,selectionObjectRef,initDoneCallback) {
			var componentIDs = { formID: formInfo.formID, componentID: selectionObjectRef.selectionID }
			initFormComponentDesignBehavior($selection,componentIDs,
						selectionObjectRef,selectionDesignFormConfig,designFormLayoutConfig)
			initDoneCallback()
		},
		initCheckBoxFunc: function(componentContext,$checkBox,checkBoxObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: checkBoxObjectRef.checkBoxID }
			initFormComponentDesignBehavior($checkBox,componentIDs,checkBoxObjectRef,checkBoxDesignFormConfig,designFormLayoutConfig)
		},
		initToggleFunc: function(componentContext,$toggle,toggleObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: toggleObjectRef.toggleID }
			initFormComponentDesignBehavior($toggle,componentIDs,
					toggleObjectRef,toggleDesignFormConfig,designFormLayoutConfig)
		},
		initProgressFunc: function(componentContext,$progress,progressObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: progressObjectRef.progressID }
			initFormComponentDesignBehavior($progress,componentIDs,progressObjectRef,progressDesignFormConfig,designFormLayoutConfig)
		},	
		initGaugeFunc: function(componentContext,$gauge,gaugeObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: gaugeObjectRef.gaugeID }
			initFormComponentDesignBehavior($gauge,componentIDs,gaugeObjectRef,gaugeDesignFormConfig,designFormLayoutConfig)
		},	
		initCommentFunc: function(componentContext,$comment,commentObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: commentObjectRef.commentID }
			initFormComponentDesignBehavior($comment,componentIDs,commentObjectRef,commentDesignFormConfig,designFormLayoutConfig)
		},
		initRatingFunc: function(componentContext,$rating,ratingObjectRef) {
			initRatingDesignControlBehavior($rating,ratingObjectRef)
			var componentIDs = { formID: formInfo.formID, componentID: ratingObjectRef.ratingID }
			initFormComponentDesignBehavior($rating,componentIDs,ratingObjectRef,ratingDesignFormConfig,designFormLayoutConfig)
		},
		initSocialButtonFunc: function(componentContext,$socialButton,socialButtonObjectRef) {
			initSocialButtonDesignControlBehavior($socialButton,socialButtonObjectRef)
			var componentIDs = { formID: formInfo.formID, componentID: socialButtonObjectRef.socialButtonID }
			initFormComponentDesignBehavior($socialButton,componentIDs,socialButtonObjectRef,socialButtonDesignFormConfig,designFormLayoutConfig)
		},
		initLabelFunc: function(componentContext,$label,labelRef) {
			initLabelDesignControlBehavior($label,labelRef)
			var componentIDs = { formID: formInfo.formID, componentID: labelRef.labelID }
			initFormComponentDesignBehavior($label,componentIDs,labelRef,labelDesignFormConfig,designFormLayoutConfig)
		},
		initUserSelectionFunc: function(componentContext,$userSelection,userSelectionObjectRef) {
			initUserSelectionDesignControlBehavior(userSelectionObjectRef)
			var componentIDs = { formID: formInfo.formID, componentID: userSelectionObjectRef.userSelectionID }
			initFormComponentDesignBehavior($userSelection,componentIDs,userSelectionObjectRef,
						userSelectionDesignFormConfig,designFormLayoutConfig)
		},
		initUserTagFunc: function(componentContext,$userTag,userTagObjectRef) {
			initUserTagDesignControlBehavior(userTagObjectRef)
			var componentIDs = { formID: formInfo.formID, componentID: userTagObjectRef.userTagID }
			initFormComponentDesignBehavior($userTag,componentIDs,userTagObjectRef,
						userTagDesignFormConfig,designFormLayoutConfig)
		},
		initDatePickerFunc: function(componentContext,$datePicker,datePickerObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: datePickerObjectRef.datePickerID }
			initFormComponentDesignBehavior($datePicker,componentIDs,datePickerObjectRef,datePickerDesignFormConfig,designFormLayoutConfig)
		},
		initHtmlEditorFunc: function(componentContext,$htmlEditor,htmlEditorObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: htmlEditorObjectRef.htmlEditorID }
			initFormComponentDesignBehavior($htmlEditor,componentIDs,htmlEditorObjectRef,htmlEditorDesignFormConfig,designFormLayoutConfig)
		},
		initAttachmentFunc: function(componentContext,$image,imageObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: imageObjectRef.imageID }
			initFormComponentDesignBehavior($image,componentIDs,imageObjectRef,attachmentDesignFormConfig,designFormLayoutConfig)
		},
		initImageFunc: function(componentContext,$image,imageObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: imageObjectRef.imageID }
			initFormComponentDesignBehavior($image,componentIDs,imageObjectRef,imageDesignFormConfig,designFormLayoutConfig)
		},
		initHeaderFunc: function($header,componentContext,headerObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: headerObjectRef.headerID }
			initFormComponentDesignBehavior($header,componentIDs,headerObjectRef,formHeaderDesignFormConfig,designFormLayoutConfig)
		},
		initCaptionFunc: function($caption,componentContext,captionObjectRef) {
			initCaptionDesignControlBehavior($caption,captionObjectRef)			
		},
		initFormButtonFunc: function(componentContext,$button,buttonObjectRef) {
			var componentIDs = { formID: formInfo.formID, componentID: buttonObjectRef.buttonID }
			initFormComponentDesignBehavior($button,componentIDs,buttonObjectRef,formButtonDesignFormConfig,designFormLayoutConfig)
		}
	}
	
	function showFormPropertySidebar() {
		hideSiblingsShowOne('#formProps')
		closeFormulaEditor()
		initDesignFormProperties(formInfo.formID)
	}	
	
	
	function doneLoadingFormData() {
			// The formula editor depends on the field information first being initialized.
			initFormulaEditor(formulaEditorParams)
		
			showFormPropertySidebar() // initially show the overall form properties.
			showDesignFormPalette()
	}
	
	loadFormComponentsIntoSingleLayout(loadFormConfig,doneLoadingFormData); 
	
	
	initObjectCanvasSelectionBehavior(formDesignCanvasSelector, function() {
		showFormPropertySidebar()
		showDesignFormPalette()
	})
	
	
	initSettingsPageButtonLink('#formPropsBackToFormListButton',"forms")

	initButtonClickHandler("#viewTableOfContentsMenuButton", function() {
		theMainWindowLayout.toggleLHSSidebar()
	})
	
		
}

function navigateToFormDesignerPageContent(pageContext,formInfo) {
	console.log("navigating to form designer")
	
	function initFormDesignerContent(initDoneCallback) {
		
		var contentSectionsRemaining = 3
		function processOneSection() {
			contentSectionsRemaining--
			if (contentSectionsRemaining <=0) {
				initDoneCallback()
			}
		}
		
		const sidebarContentURL = '/admin/frm/sidebarContent/' + formInfo.formID
		setRHSSidebarContent(sidebarContentURL, function() {
			processOneSection()
		})

		const settingsPageURL = '/admin/frm/designPageContent/' + formInfo.formID
		setSettingsPageContent(settingsPageURL,function() {
			processOneSection()
		})
		
		const offPageContentURL = '/admin/frm/offPageContent/' + formInfo.formID
		setSettingsPageOffPageContent(offPageContentURL,function() {
			processOneSection()
		})
		
	}
	
	initFormDesignerContent(function() {
		theMainWindowLayout.showRHSSidebar()
		theMainWindowLayout.openRHSSidebar()
		initDesignFormAdminPageContent(pageContext,formInfo)
	})		
	
	
}
