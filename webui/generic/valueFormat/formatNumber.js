
function initNumberFormatSelection(params) {
	
	var selectFormatSelector = createPrefixedSelector(params.elemPrefix,"NumberFormatSelection")
	
	$(selectFormatSelector).val(params.initialFormat)
	initSelectionChangedHandler(selectFormatSelector, params.formatChangedCallback)
}


function formatNumberValue(format, rawVal) {
	
	function isInt(val) {
		return Number(val) % 1 === 0 // remainder non-zero with modulo arithmetic
	}
	
	switch (format) {
		case "percent":
			return (Number(rawVal)*100.0).toFixed(2) + "%"
		case "percent0":
			return (Number(rawVal)*100.0).toFixed(0) + "%"
		case "percent1":
			return (Number(rawVal)*100.0).toFixed(1) + "%"
		case "general":
			if(isInt(rawVal)) {
				return Number(rawVal)
				
			} else {
				return accounting.toFixed(rawVal,2)	
			}
			
		case "currency":
			return accounting.formatMoney(rawVal)
		case "currency0prec":
			return accounting.formatMoney(rawVal,{precision:0})
		default:
			return rawVal
	}

}