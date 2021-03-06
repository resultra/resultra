// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function getFormComponentContext(formContext, contextLoadCompleteCallback) {
	var contextPartsRemaining = 5;
	var context = {}
	
	context.formID = formContext.formID
	context.databaseID = formContext.databaseID
	
	function completeOneContextPart() {
		contextPartsRemaining -= 1
		if(contextPartsRemaining <= 0) {
			contextLoadCompleteCallback(context)
		}
	}
	
	loadFieldInfo(context.databaseID, [fieldTypeAll],function(fieldsByID) {
		context.fieldsByID = fieldsByID
		completeOneContextPart()
	})
	
	initFieldInfo( formContext.databaseID, function () {
		completeOneContextPart()
	})
	
	getGlobalInfoIndexedByID(context.databaseID,function(globalsByID) {
		context.globalsByID = globalsByID
		completeOneContextPart()
	})
	
	jsonAPIRequest("frm/getFormInfo", { formID: context.formID }, function(formInfo) {
		context.formInfo = formInfo
		completeOneContextPart()
	})
	
	var getUserInfoParams = {}
	jsonRequest("/auth/getCurrentUserInfo",getUserInfoParams,function(currUserInfo) {
		context.currUserInfo = currUserInfo
		completeOneContextPart()
	})	
	
	
}

function populateOneFormLayoutWithComponents(loadFormConfig, componentContext,formDonePopulatingCallback) {
	
	var compenentIDComponentMap = {}	

	function initHeaderLayout($componentRow,header,initDoneCallback) {
		// Create an HTML block for the container
		console.log("loadFormComponents: initializing header: " + JSON.stringify(header))
	
		var containerHTML = formHeaderContainerHTML(header.headerID);
		var $containerObj = $(containerHTML)
		$containerObj.find(".formHeader").text(header.properties.label)
		setHeaderFormComponentHeaderSize($containerObj,header.properties.headerSize)
		setHeaderFormComponentUnderlined($containerObj,header.properties.underlined)
				
		$componentRow.append($containerObj)
		
		setElemFixedWidthFlexibleHeight($containerObj,header.properties.geometry.sizeWidth)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($containerObj,header,header.headerID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initHeaderFunc($containerObj,componentContext,header)
		initDoneCallback()
		
	}


	function initCaptionLayout($componentRow,caption,initDoneCallback) {
		// Create an HTML block for the container
		console.log("loadFormComponents: initializing caption: " + JSON.stringify(caption))
	
		var containerHTML = formCaptionContainerHTML(caption.captionID);
		var $containerObj = $(containerHTML)
		$containerObj.find(".formCaption").text(caption.properties.label)
		setFormCaptionColorScheme($containerObj,caption.properties.colorScheme)
				
		$componentRow.append($containerObj)
		
		setElemDimensions($containerObj,caption.properties.geometry)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($containerObj,caption,caption.captionID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initCaptionFunc($containerObj,componentContext,caption)
		initDoneCallback()
		
	}


	function initFormButtonLayout($componentRow,formButton,initDoneCallback) {
		// Create an HTML block for the container
		console.log("loadFormComponents: initializing form button: " + JSON.stringify(formButton))
	
		var containerHTML = formButtonContainerHTML();
		var $containerObj = $(containerHTML)
		setFormButtonSize($containerObj,formButton.properties.size)
		setFormButtonColorScheme($containerObj,formButton.properties.colorScheme)
		setFormButtonLabel($containerObj,formButton)
		
				
		$componentRow.append($containerObj)
		
		setElemFixedWidthFlexibleHeight($containerObj,formButton.properties.geometry.sizeWidth)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($containerObj,formButton,formButton.buttonID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initFormButtonFunc(componentContext,$containerObj,formButton)
		initDoneCallback()
		
	}


	function initTextBoxLayout($componentRow,textBox,initDoneCallback) {
		// Create an HTML block for the container
		console.log("loadFormComponents: initializing text box: " + JSON.stringify(textBox))
	
		var containerHTML = textBoxContainerHTML(textBox.textBoxID);
		var containerObj = $(containerHTML)
		
		initTextBoxFormComponentContainer(containerObj,textBox)
					
		$componentRow.append(containerObj)
		
		setElemFixedWidthFlexibleHeight(containerObj,
					textBox.properties.geometry.sizeWidth)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,textBox,textBox.textBoxID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initTextBoxFunc(componentContext,containerObj,textBox)
		initDoneCallback()
		
	}

	function initEmailAddrLayout($componentRow,emailAddr,initDoneCallback) {
		// Create an HTML block for the container
	
		var containerHTML = emailAddrContainerHTML(emailAddr.emailAddrID);
		var containerObj = $(containerHTML)
		
		initEmailAddrFormComponentContainer(containerObj,emailAddr)
					
		$componentRow.append(containerObj)
		
		setElemFixedWidthFlexibleHeight(containerObj,
					emailAddr.properties.geometry.sizeWidth)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,emailAddr,emailAddr.emailAddrID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initEmailAddrFunc(componentContext,containerObj,emailAddr)
		initDoneCallback()
		
	}
	
	function initFileLayout($componentRow,fileRef,initDoneCallback) {
		// Create an HTML block for the container
	
		var containerHTML = fileContainerHTML(fileRef.fileID);
		var containerObj = $(containerHTML)
		
		initFileFormComponentContainer(containerObj,fileRef)
					
		$componentRow.append(containerObj)
		
		setElemFixedWidthFlexibleHeight(containerObj,
					fileRef.properties.geometry.sizeWidth)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,fileRef,fileRef.fileID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initFileFunc(componentContext,containerObj,fileRef)
		initDoneCallback()
		
	}

	function initImageLayout($componentRow,imageRef,initDoneCallback) {
		// Create an HTML block for the container
	
		var containerHTML = imageContainerHTML(imageRef.imageID);
		var $containerObj = $(containerHTML)
		
		initImageFormComponentContainer($containerObj,imageRef)
					
		$componentRow.append($containerObj)
		
		// By default the dimensions of the image form component are set to whatever the
		// actual geometry is. However, when viewing images, the size of the container
		// is set to a fixed width and variable height, with the maximum height of the
		// image set to the height in the geometry.
		setElemDimensions($containerObj,imageRef.properties.geometry)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($containerObj,imageRef,imageRef.imageID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initImageFunc(componentContext,$containerObj,imageRef)
		initDoneCallback()
		
	}
	
	
	function initUrlLinkLayout($componentRow,urlLink,initDoneCallback) {
		// Create an HTML block for the container
	
		var containerHTML = urlLinkContainerHTML(urlLink.urlLinkID);
		var containerObj = $(containerHTML)
		
		initUrlLinkFormComponentContainer(containerObj,urlLink)
					
		$componentRow.append(containerObj)
		
		setElemFixedWidthFlexibleHeight(containerObj,
					urlLink.properties.geometry.sizeWidth)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,urlLink,urlLink.urlLinkID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initUrlLinkFunc(componentContext,containerObj,urlLink)
		initDoneCallback()
		
	}

	
	function initNumberInputLayout($componentRow,numberInput,initDoneCallback) {
		// Create an HTML block for the container
		console.log("loadFormComponents: initializing number input: " + JSON.stringify(numberInput))
	
		var containerHTML = numberInputContainerHTML(numberInput.numberInputID);
		var containerObj = $(containerHTML)
		
		initNumberInputFormContainer(containerObj,numberInput)
				
		$componentRow.append(containerObj)
		
		setElemFixedWidthFlexibleHeight(containerObj,
					numberInput.properties.geometry.sizeWidth)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,numberInput,numberInput.numberInputID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initNumberInputFunc(componentContext,containerObj,numberInput)
		initDoneCallback()
		
	}

	function initSelectionLayout($componentRow,selection,initDoneCallback) {
		// Create an HTML block for the container
		console.log("loadFormComponents: initializing selection: " + JSON.stringify(selection))
	
		var containerHTML = selectionContainerHTML(selection.selectionID);
		var containerObj = $(containerHTML)
		
		
		setSelectionComponentLabel(containerObj,selection)
		initSelectionComponentClearValueButton(containerObj,selection)
	
		$componentRow.append(containerObj)
		
		setElemFixedWidthFlexibleHeight(containerObj,
					selection.properties.geometry.sizeWidth)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,selection,selection.selectionID)
	
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initSelectionFunc(componentContext,containerObj,selection,function() {
				initDoneCallback()
		})

		
	}


	function initCommentLayout($componentRow,comment,initDoneCallback) {
		// Create an HTML block for the container
		
		var containerHTML = commentContainerHTML(comment.commentID);
		var containerObj = $(containerHTML)
		
		setCommentComponentLabel(containerObj,comment)
		initComponentHelpPopupButton(containerObj, comment)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append(containerObj)

		setElemDimensions(containerObj,comment.properties.geometry)
				
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,comment,comment.commentID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initCommentFunc(componentContext,containerObj,comment)
		initDoneCallback()
	}



	function initProgressLayout($componentRow,progress,initDoneCallback) {
		// Create an HTML block for the container
		
		var containerHTML = progressContainerHTML();
		var $progressContainer = $(containerHTML)
				
		initProgressFormComponentContainer($progressContainer,progress)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append($progressContainer)
		setElemFixedWidthFlexibleHeight($progressContainer,
					progress.properties.geometry.sizeWidth)
		
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($progressContainer,progress,progress.progressID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initProgressFunc(componentContext,$progressContainer,progress)
		initDoneCallback()
	}


	function initGaugeLayout($componentRow,gaugeRef,initDoneCallback) {
		// Create an HTML block for the container
		
		var containerHTML = gaugeContainerHTML();
		var $gaugeContainer = $(containerHTML)
		
		initGaugeFormComponentContainer($gaugeContainer,gaugeRef)
				
		
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append($gaugeContainer)
		setElemFixedWidthFlexibleHeight($gaugeContainer,
					gaugeRef.properties.geometry.sizeWidth)
		
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($gaugeContainer,gaugeRef,gaugeRef.gaugeID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initGaugeFunc(componentContext,$gaugeContainer,gaugeRef)
		initDoneCallback()
	}


	function initCheckBoxLayout($componentRow,checkBox,initDoneCallback) {
		// Create an HTML block for the container
		
		var containerHTML = checkBoxContainerHTML(checkBox.checkBoxID);
		var $checkboxContainer = $(containerHTML)
				
		initCheckboxComponentFormContainer($checkboxContainer,checkBox)		
		
		var checkboxColorSchemeClass = "checkbox-"+checkBox.properties.colorScheme
		$checkboxContainer.addClass(checkboxColorSchemeClass)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append($checkboxContainer)
		
		setElemFixedWidthFlexibleHeight($checkboxContainer,checkBox.properties.geometry.sizeWidth)
				
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($checkboxContainer,checkBox,checkBox.checkBoxID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initCheckBoxFunc(componentContext,$checkboxContainer,checkBox)
		initDoneCallback()
	}
	
	function initToggleLayout($componentRow,toggle,initDoneCallback) {
		// Create an HTML block for the container
		
		var containerHTML = toggleContainerHTML(toggle.toggleID);
		var $toggleContainer = $(containerHTML)
				
		initToggleComponentFormComponentContainer($toggleContainer,toggle)
		
		var toggleColorSchemeClass = "checkbox-"+toggle.properties.colorScheme
		$toggleContainer.addClass(toggleColorSchemeClass)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append($toggleContainer)
		
		setElemFixedWidthFlexibleHeight($toggleContainer,toggle.properties.geometry.sizeWidth)
				
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($toggleContainer,toggle,toggle.toggleID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initToggleFunc(componentContext,$toggleContainer,toggle)
		initDoneCallback()
	}


	function initRatingLayout($componentRow,rating,initDoneCallback) {
		// Create an HTML block for the container
		
		var containerHTML = ratingContainerHTML(rating.ratingID);	
		var $ratingContainer = $(containerHTML)
		
		initRatingFormComponentContainer($ratingContainer,rating)
				
		// Position the object withing the #layoutCanvas div
		$componentRow.append($ratingContainer)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($ratingContainer,rating,rating.ratingID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initRatingFunc(componentContext,$ratingContainer,rating)
		initDoneCallback()
	}

	function initSocialButtonLayout($componentRow,socialButton,initDoneCallback) {
		
		var containerHTML = socialButtonContainerHTML(socialButton.socialButtonID);	
		var $socialButtonContainer = $(containerHTML)
		
		initSocialButtonFormComponentContainer($socialButtonContainer,socialButton)
				
		// Position the object withing the #layoutCanvas div
		$componentRow.append($socialButtonContainer)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($socialButtonContainer,socialButton,socialButton.socialButtonID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initSocialButtonFunc(componentContext,$socialButtonContainer,socialButton)
		initDoneCallback()
	}
	
	function initLabelLayout($componentRow,label,initDoneCallback) {
		
		var containerHTML = labelContainerHTML(label.labelID);	
		var $labelContainer = $(containerHTML)
		
		initLabelFormComponentContainer($labelContainer,label)
				
		// Position the object withing the #layoutCanvas div
		$componentRow.append($labelContainer)

		setElemFixedWidthFlexibleHeight($labelContainer,
					label.properties.geometry.sizeWidth)

	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($labelContainer,label,label.labelID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initLabelFunc(componentContext,$labelContainer,label)
		initDoneCallback()
	}

	function initUserSelectionLayout($componentRow,userSelection,initDoneCallback) {
		// Create an HTML block for the container
		
		var containerHTML = userSelectionContainerHTML(userSelection.userSelectionID);
		
		var $containerObj = $(containerHTML)
							
		initUserSelectionFormComponentContainer(componentContext,$containerObj,userSelection)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append($containerObj)
		
		setElemFixedWidthFlexibleHeight($containerObj,
					userSelection.properties.geometry.sizeWidth)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($containerObj,userSelection,userSelection.userSelectionID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initUserSelectionFunc(componentContext,$containerObj,userSelection)
		initDoneCallback()
	}
	
	function initUserTagLayout($componentRow,userTag,initDoneCallback) {
		// Create an HTML block for the container
		
		var containerHTML = userTagContainerHTML(userTag.userTagID);
		
		var $containerObj = $(containerHTML)
							
		initUserTagFormComponentContainer(componentContext,$containerObj,userTag)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append($containerObj)
		
		setElemFixedWidthFlexibleHeight($containerObj,
					userTag.properties.geometry.sizeWidth)
	
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo($containerObj,userTag,userTag.userTagID)
		
		// Callback for any specific initialization for either the form design or view mode 
		loadFormConfig.initUserTagFunc(componentContext,$containerObj,userTag)
		initDoneCallback()
	}

	
	function initDatePickerLayout($componentRow,datePicker,initDoneCallback) {
		// Create an HTML block for the container			
		var containerHTML = datePickerContainerHTML(datePicker.datePickerID);
		var containerObj = $(containerHTML)
			
		initDatePickerContainerControls(containerObj,datePicker)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append(containerObj)
		setElemFixedWidthFlexibleHeight(containerObj,datePicker.properties.geometry.sizeWidth)
		
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,datePicker,datePicker.datePickerID)
					
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initDatePickerFunc(componentContext,containerObj,datePicker)
		initDoneCallback()
		
	}

	function initHtmlEditorLayout($componentRow,htmlEditor,initDoneCallback) {
		
		var containerHTML = htmlEditorContainerHTML(htmlEditor.htmlEditorID);
		var containerObj = $(containerHTML)
		
		// Set the label to the field name		
		setEditorComponentLabel(containerObj,htmlEditor)
		initComponentHelpPopupButton(containerObj, htmlEditor)
		
		initEditorFormComponentViewModeGeometry(containerObj,htmlEditor)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append(containerObj)
		
		
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,htmlEditor,htmlEditor.htmlEditorID)
		
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initHtmlEditorFunc(componentContext,containerObj,htmlEditor)
		initDoneCallback()
	}
	
	function initAttachmentEditorLayout($componentRow,image,initDoneCallback) {
		var containerHTML = attachmentContainerHTML(image.imageID);
		var containerObj = $(containerHTML)
		
		// Position the object withing the #layoutCanvas div
		$componentRow.append(containerObj)
		setElemDimensions(containerObj,image.properties.geometry)

		setAttachmentComponentLabel(containerObj,image)
		initComponentHelpPopupButton(containerObj, image)
		
		 // Store the newly created object reference in the DOM element. This is needed for follow-on
		 // property setting, resizing, etc.
		setContainerComponentInfo(containerObj,image,image.imageID)
		
		// Callback for any specific initialization for either the form design or view mode
		loadFormConfig.initAttachmentFunc(componentContext,containerObj,image)
		initDoneCallback()
		
	}

	var formInfo = componentContext.formInfo
	
	// Iterate through each type of component and populate a map/dictonary with 
	// the component ID and method to initialize the compnent. This map is then
	// referenced when populating the layout to put each component in the right place
	// on the layout.
	for (var textBoxIter in formInfo.textBoxes) {			
		var textBoxProps = formInfo.textBoxes[textBoxIter]			
		console.log("Form layout: text box component info=" + JSON.stringify(textBoxProps))
		compenentIDComponentMap[textBoxProps.textBoxID] = {
			componentInfo: textBoxProps,
			initFunc: initTextBoxLayout
		}			

	} // for each text box


	for (var emailAddrIter in formInfo.emailAddrs) {			
		var emailAddrProps = formInfo.emailAddrs[emailAddrIter]			
		compenentIDComponentMap[emailAddrProps.emailAddrID] = {
			componentInfo: emailAddrProps,
			initFunc: initEmailAddrLayout
		}			

	} // for each email address input
	
	for (var urlLinkAddrIter in formInfo.urlLinks) {			
		var urlLinkProps = formInfo.urlLinks[urlLinkAddrIter]			
		compenentIDComponentMap[urlLinkProps.urlLinkID] = {
			componentInfo: urlLinkProps,
			initFunc: initUrlLinkLayout
		}			

	} // for each email address input


	for (var numberInputIter in formInfo.numberInputs) {			
		var numberInputProps = formInfo.numberInputs[numberInputIter]			
		console.log("Form layout: text box component info=" + JSON.stringify(numberInputProps))
		compenentIDComponentMap[numberInputProps.numberInputID] = {
			componentInfo: numberInputProps,
			initFunc: initNumberInputLayout
		}			

	} // for each number input


	for (var selectionIter in formInfo.selections) {			
		var selectionProps = formInfo.selections[selectionIter]			
		console.log("Form layout: selection component info=" + JSON.stringify(selectionProps))
		compenentIDComponentMap[selectionProps.selectionID] = {
			componentInfo: selectionProps,
			initFunc: initSelectionLayout
		}			

	} // for each selection


	for (var checkBoxIter in formInfo.checkBoxes) {
		var checkBoxProps = formInfo.checkBoxes[checkBoxIter]
		console.log("loadFormComponents: initializing check box: " + JSON.stringify(checkBoxProps))
		compenentIDComponentMap[checkBoxProps.checkBoxID] = {
			componentInfo: checkBoxProps,
			initFunc: initCheckBoxLayout
		}		
	}
	
	for (var toggleIter in formInfo.toggles) {
		var toggleProps = formInfo.toggles[toggleIter]
		console.log("loadFormComponents: initializing toggle: " + JSON.stringify(toggleProps))
		compenentIDComponentMap[toggleProps.toggleID] = {
			componentInfo: toggleProps,
			initFunc: initToggleLayout
		}		
	}

	for (var progressIter in formInfo.progressIndicators) {
		var progressProps = formInfo.progressIndicators[progressIter]
		console.log("loadFormComponents: initializing progress indicator: " + JSON.stringify(progressProps))
		compenentIDComponentMap[progressProps.progressID] = {
			componentInfo: progressProps,
			initFunc: initProgressLayout
		}		
	}
	
	for (var gaugeIter in formInfo.gauges) {
		var gaugeProps = formInfo.gauges[gaugeIter]
		console.log("loadFormComponents: initializing gauge: " + JSON.stringify(gaugeProps))
		compenentIDComponentMap[gaugeProps.gaugeID] = {
			componentInfo: gaugeProps,
			initFunc: initGaugeLayout
		}		
	}

	for (var ratingIter in formInfo.ratings) {
		var ratingProps = formInfo.ratings[ratingIter]
		console.log("loadFormComponents: initializing rating: " + JSON.stringify(ratingProps))
		compenentIDComponentMap[ratingProps.ratingID] = {
			componentInfo: ratingProps,
			initFunc: initRatingLayout
		}		
	}
	
	for (var socialButtonIter in formInfo.socialButtons) {
		var socialButtonProps = formInfo.socialButtons[socialButtonIter]
		console.log("loadFormComponents: initializing rating: " + JSON.stringify(socialButtonProps))
		compenentIDComponentMap[socialButtonProps.socialButtonID] = {
			componentInfo: socialButtonProps,
			initFunc: initSocialButtonLayout
		}		
	}

	for (var commentIter in formInfo.comments) {
		var commentProps = formInfo.comments[commentIter]
		console.log("loadFormComponents: initializing comment component: " + JSON.stringify(commentProps))
		compenentIDComponentMap[commentProps.commentID] = {
			componentInfo: commentProps,
			initFunc: initCommentLayout
		}		
	}

	for (var userSelIter in formInfo.userSelections) {
		var userSelProps = formInfo.userSelections[userSelIter]
		console.log("loadFormComponents: initializing user selection: " + JSON.stringify(userSelProps))
		compenentIDComponentMap[userSelProps.userSelectionID] = {
			componentInfo: userSelProps,
			initFunc: initUserSelectionLayout
		}		
	}
	
	for (var userTagIter in formInfo.userTags) {
		var userTagProps = formInfo.userTags[userTagIter]
		console.log("loadFormComponents: initializing user tag: " + JSON.stringify(userTagProps))
		compenentIDComponentMap[userTagProps.userTagID] = {
			componentInfo: userTagProps,
			initFunc: initUserTagLayout
		}		
	}


	for (var headerIter in formInfo.headers) {
		var headerProps = formInfo.headers[headerIter]
		console.log("loadFormComponents: initializing header: " + JSON.stringify(headerProps))
		compenentIDComponentMap[headerProps.headerID] = {
			componentInfo: headerProps,
			initFunc: initHeaderLayout
		}		
	}

	for (var buttonIter in formInfo.formButtons) {
		var buttonProps = formInfo.formButtons[buttonIter]
		console.log("loadFormComponents: initializing form button: " + JSON.stringify(buttonProps))
		compenentIDComponentMap[buttonProps.buttonID] = {
			componentInfo: buttonProps,
			initFunc: initFormButtonLayout
		}		
	}

	
	for (var datePickerIter in formInfo.datePickers) {
		var datePickerProps = formInfo.datePickers[datePickerIter]
		console.log("loadFormComponents: initializing date picker: " + JSON.stringify(datePickerProps))
		compenentIDComponentMap[datePickerProps.datePickerID] = {
			componentInfo: datePickerProps,
			initFunc: initDatePickerLayout
		}			
		
	}

	for (var htmlEditorIter in formInfo.htmlEditors) {
		var htmlEditorProps = formInfo.htmlEditors[htmlEditorIter]
		console.log("loadFormComponents: initializing html editor: " + JSON.stringify(htmlEditorProps))

		compenentIDComponentMap[htmlEditorProps.htmlEditorID] = {
			componentInfo: htmlEditorProps,
			initFunc: initHtmlEditorLayout
		}			
	}

	for (var attachmentIter in formInfo.attachments) {
		var attachmentProps = formInfo.attachments[attachmentIter]
		console.log("loadFormComponents: initializing attachment editor: " + JSON.stringify(attachmentProps))

		compenentIDComponentMap[attachmentProps.imageID] = {
			componentInfo: attachmentProps,
			initFunc: initAttachmentEditorLayout
		}			

	}		

	for (var imageIter in formInfo.images) {
		var imageProps = formInfo.images[imageIter]
		console.log("loadFormComponents: initializing image editor: " + JSON.stringify(imageProps))

		compenentIDComponentMap[imageProps.imageID] = {
			componentInfo: imageProps,
			initFunc: initImageLayout
		}			

	}		



	for (var captionIter in formInfo.captions) {
		var captionProps = formInfo.captions[captionIter]
		console.log("loadFormComponents: initializing caption: " + JSON.stringify(captionProps))

		compenentIDComponentMap[captionProps.captionID] = {
			componentInfo: captionProps,
			initFunc: initCaptionLayout
		}			

	}		
	
	
	for (var labelIter in formInfo.labels) {
		var labelProps = formInfo.labels[labelIter]
		
		console.log("loadFormComponents: initializing caption: " + JSON.stringify(labelProps))

		compenentIDComponentMap[labelProps.labelID] = {
			componentInfo: labelProps,
			initFunc: initLabelLayout
		}			

	}		

	for (var fileIter in formInfo.files) {
		var fileProps = formInfo.files[fileIter]
		
		console.log("loadFormComponents: initializing file component: " + JSON.stringify(fileProps))

		compenentIDComponentMap[fileProps.fileID] = {
			componentInfo: fileProps,
			initFunc: initFileLayout
		}			

	}		

	
	var formLayout = formInfo.form.properties.layout
	populateComponentLayout(formLayout,loadFormConfig.$parentFormLayout,compenentIDComponentMap,function() {
		formDonePopulatingCallback()
	})
	
}


function loadFormComponentsIntoSingleLayout(loadFormConfig, doneLoadingFormDataFunc) {
	
	getFormComponentContext(loadFormConfig.formContext, function(componentContext) {											
		populateOneFormLayoutWithComponents(loadFormConfig,componentContext,function() {
					doneLoadingFormDataFunc()
		})		

	})
	
}