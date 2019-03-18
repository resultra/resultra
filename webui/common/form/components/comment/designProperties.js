// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function loadCommentComponentProperties($comment,commentRef) {
	console.log("Loading comment component properties")
	
	var elemPrefix = "comment_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for comment box")
		var formatParams = {
			parentFormID: commentRef.parentFormID,
			commentID: commentRef.commentID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/comment/setLabelFormat", formatParams, function(updatedComment) {
			setCommentComponentLabel($comment,updatedComment)
			setContainerComponentInfo($comment,updatedComment,commentRef.commentID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: commentRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)

	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: commentRef.parentFormID,
			commentID: commentRef.commentID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/comment/setVisibility",params,function(updatedComment) {
			setContainerComponentInfo($comment,updatedComment,commentRef.commentID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: commentRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	
	var permissionParams = {
		elemPrefix: elemPrefix,
		initialVal: commentRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: commentRef.parentFormID,
				commentID: commentRef.commentID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/comment/setPermissions",params,function(updatedComment) {
				setContainerComponentInfo($comment,updatedComment,commentRef.commentID)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(permissionParams)

	var helpPopupParams = {
		initialMsg: commentRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: commentRef.parentFormID,
				commentID: commentRef.commentID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/comment/setHelpPopupMsg",params,function(updatedComment) {
				setContainerComponentInfo($comment,updatedComment,commentRef.commentID)
				updateComponentHelpPopupMsg($comment, updatedComment)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)


	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: commentRef.parentFormID,
		componentID: commentRef.commentID,
		componentLabel: 'comment box',
		$componentContainer: $comment
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#commentComponentProps')
	
	toggleFormulaEditorForField(commentRef.properties.fieldID)
	
}