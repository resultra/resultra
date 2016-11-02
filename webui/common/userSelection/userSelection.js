function initUserSelection(selectionParams) {
	
	var configParams = {
		width: '250px'
	}
	$.extend(configParams,selectionParams)
	
	configParams.selectionInput.select2({
		placeholder: "Select a user",
		dropdownParent: configParams.dropdownParent,
		minimumInputLength: 2,
		width: configParams.width,
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
	});
	
	
	
}