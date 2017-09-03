

function labelControlContainerHTML() {
		
	return '<div class="input-group">'+
				'<div class="formLabelControl">' + 
					'<select class="form-control labelCompSelectionControl" multiple="multiple"></select>' +
				'</div>' + 
			'</div>'
	
}

function labelContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class=" layoutContainer labelFormContainer">' +
			'<div class="form-group marginBottom0">'+
				'<label>New Label</label>' + componentHelpPopupButtonHTML() +
				labelControlContainerHTML() +
			'</div>'+
		'</div>';
										
	return containerHTML
}

function labelControlFromLabelComponentContainer($labelContainer) {
	return $labelContainer.find(".labelCompSelectionControl")
}


function labelTableCellContainerHTML() {
	var containerHTML = ''+
		'<div class=" layoutContainer labelTableCellContainer">' +
			labelControlContainerHTML() +
		'</div>';									
	return containerHTML
	
}

function initLabelSelectionControl($container, labelRef) {
		
	var labelWidth = labelRef.properties.geometry.sizeWidth - 15
	var $labelControl = labelControlFromLabelComponentContainer($container)
	
	
	$labelControl.select2({
		placeholder: "Enter labels", // TODO - Allow a property to configure the placeholder.
		width: labelWidth,
		tags:true,
		tokenSeparators: [',']
		/*
		ajax: {
			dataType: 'json',
			url: '/auth/searchUsers',
			delay: 250,
			data: function (params) {
				var queryParams = {
				  searchTerm: params.term, // search term
				  page: params.page
				}
	      	  return queryParams
		  	},
			processResults: function (data, params) {
			      // parse the results into the format expected by Select2
			      // since we are using custom formatting functions we do not need to
			      // alter the remote JSON data, except to indicate that infinite
			      // scrolling can be used
			      params.page = params.page || 1;
				  
				  var select2results = []
				  for(var matchIndex = 0; matchIndex < data.matchedUserInfo.length; matchIndex++) {
					  var currMatch = data.matchedUserInfo[matchIndex]
					  var select2result = {
						  id:currMatch.userID,
						  text:'@'+currMatch.userName
					  }
					  select2results.push(select2result)
				  }

			      return {
			        results: select2results,
			        pagination: {
			          more: (params.page * 30) < data.matchedUserInfo.length
			        }
			      };
			  },
			cache: true	
		}
		*/
	});

}	
	
	

function setLabelComponentLabel($label,label) {
	var $label = $label.find('label')
	
	setFormComponentLabel($label,label.properties.fieldID,
			label.properties.labelFormat)	
	
}

function initLabelFormComponentContainer($container,label) {
	setLabelComponentLabel($container,label)
	initComponentHelpPopupButton($container, label)	
	initLabelSelectionControl($container, label)
	
	setElemFixedWidthFlexibleHeight($container,
				label.properties.geometry.sizeWidth)
	
}


