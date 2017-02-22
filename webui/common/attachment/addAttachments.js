
function initAddAttachmentControl(params) {
	
	var uploadImageURL = "/api/attachment/upload"
	
	var uploadedFileResults = []
		
	params.$addAttachmentInput.fileupload({
		dataType: 'json',
		// paramName corresponds to the name given to the file when it is sent to the server. 
		// This name needs to match the name given to the FormFile() function on the server.
		paramName: "uploadFile",
		url: uploadImageURL,

		start: function(e) {
			uploadedFileResults = []
			console.log("Attachment: Starting file upload operation")
		},

		// The done callback is invoked every time an individual file complete's its upload, not when the 
		// entire upload operation is done. This is the default behavior, unless the 'singleFileUploads' option
		// is set to false, in which case multiple files may be uploaded.
		done: function(e, data) {
			console.log("Attachment:  upload request done.")
			$.each(data.result.files, function(index, file) {
				uploadedFileResults.push(file.attachmentInfo)
			})
		},

		progress: function(e, data) {
			var uploadProgress = parseInt(data.loaded / data.total * 100, 10);
			console.log("Attachment: File upload progress: " + uploadProgress)
		},

		progressall: function(e, data) {
			var uploadProgress = parseInt(data.loaded / data.total * 100, 10);
			console.log("Attachment: All file upload progress: " + uploadProgress)
		},

		stop: function(e) {
			console.log("Attachment: All file uploads complete: " + JSON.stringify(uploadedFileResults))

			params.attachDoneCallback(uploadedFileResults)

			// Reset/clear the list of current attachments being added.
			uploadedFileResults = []
		}
	});

		
	// Wait until actually starting the upload to initialize the form data with the record ID and
	// field ID. The reason this can't happen at the same time as the initial upload button initialization is
	// that the records haven't been loaded when the initial initialization takes place, so the current
	// record is unknown.
	params.$addAttachmentInput.bind('fileuploadsubmit', function(e, data) {
		// The example input, doesn't have to be part of the upload form:

		var fileUploadParams = {
			parentDatabaseID: params.parentDatabaseID
		}

		console.log("Attachment: Starting file upload: params = " + JSON.stringify(fileUploadParams))

		data.formData = fileUploadParams
	});
	
}