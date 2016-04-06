

function initBarChartPropsEditHandling(barChartPropsArgs)
{
	
	var args = barChartPropsArgs
	
	$( "#barChartProps" ).accordion();
	
	$( "#barChartProps" ).form({
    	fields: {
	        barChartTitleProp: {
	          rules: [
	            {
	              type   : 'empty',
	              prompt : 'Please enter a title'
	            }
	          ]
	        }, // newFieldName validation
     	},
  	})
	
	// Change title property
	$('#barChartTitleProp').unbind("focusout") // unbind any previous handlers
	$('#barChartTitleProp').focusout(function () {
		console.log("Done editing title for dashboard id = " + 
				args.dashboardID + " barChartID=" + args.barChartID)
		
		var newTitle = $("#barChartProps").form('get value','barChartTitleProp')
		var setTitleParams = {
			uniqueID: {
				parentDashboardID: barChartPropsArgs.dashboardID,
				barChartID: barChartPropsArgs.barChartID
			},
			title: newTitle
		}
		
		jsonAPIRequest("setBarChartTitle",setTitleParams,function(updatedBarChartRef) {
			console.log("Bar Chart Title updated on server: bar chart id= " + updatedBarChartRef.barChartID)
			args.propertyUpdateComplete(updatedBarChartRef)
		})
		
	})
}

function loadBarChartProperties(barChartPropsArgs) {
	
	var barChartContainer = $('#'+barChartPropsArgs.barChartID)
	var barChartRef = barChartContainer.data("barChartRef")
	
	$("#barChartProps").form('set value','barChartTitleProp',barChartRef.title)
	
	// Toggle to the bar chart properties, hiding the other property panels
	hideSiblingsShowOne('#barChartProps')
			
	initBarChartPropsEditHandling(barChartPropsArgs)
}
