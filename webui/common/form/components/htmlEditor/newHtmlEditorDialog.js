

function openNewHtmlEditorDialog(databaseID,formID,containerParams)
{
				
	function createNewHtmlEditor($parentDialog, newComponentParams) {
		jsonAPIRequest("frm/htmlEditor/new",newComponentParams,function(newHtmlEditorObjectRef) {
	          console.log("saveNewHtmlEditor: Done getting new ID:response=" + JSON.stringify(newHtmlEditorObjectRef));
		  	  
			  
	  		  var componentLabel = getFieldRef(newHtmlEditorObjectRef.properties.fieldID).name
			  containerParams.containerObj.find('label').text(componentLabel)
			  	  
	  		  var newComponentSetupParams = {
				  parentFormID: formID,
	  		  	  $container: containerParams.containerObj,
				  componentID: newHtmlEditorObjectRef.htmlEditorID,
				  componentObjRef: newHtmlEditorObjectRef,
				  designFormConfig: htmlEditorDesignFormConfig
	  		  }
			  setupNewlyCreatedFormComponentInfo(newComponentSetupParams)
			  		  			  	
			  $parentDialog.modal("hide")

	       }) // newLayoutContainer API request
	}
		
	var newFormComponentDialogParams = {
		elemPrefix: "htmlEditor_",
		databaseID: databaseID,
		formID: formID,
		hideCreateCalcFieldCheckbox: true,
		fieldTypes: [fieldTypeLongText],
		containerParams: containerParams,
		createNewFormComponent: createNewHtmlEditor
	}
	
	openNewFormComponentDialog(newFormComponentDialogParams)

		
} // newLayoutContainer

function initNewHtmlEditorDialog() {
}