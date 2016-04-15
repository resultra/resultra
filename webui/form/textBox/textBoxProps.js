

function resizeTextBox(resizeParams) {
	
	var resizeTextBoxParams = {
		parentLayoutID: resizeParams.formID,
		containerID: resizeParams.formItemID,
		geometry: resizeParams.geometry
	};
	
 	jsonAPIRequest("resizeLayoutContainer",resizeTextBoxParams,function(replyData) {
 		console.log("Done resizing text box")
 	})	
	
}