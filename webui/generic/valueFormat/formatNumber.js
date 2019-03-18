// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function initNumberFormatSelection(params) {
	
	var selectFormatSelector = createPrefixedSelector(params.elemPrefix,"NumberFormatSelection")
	
	$(selectFormatSelector).val(params.initialFormat)
	initSelectionChangedHandler(selectFormatSelector, params.formatChangedCallback)
}


function formatNumberValue(format, rawVal) {
	
	var numberVal = convertStringToNumber(rawVal)
	if (numberVal == null) {
		return rawVal // Don't format the value if it is not a number
	} 
		
	function isInt(val) {
		return Number(val) % 1 === 0 // remainder non-zero with modulo arithmetic
	}
	
	// Use a custom format for numbers.
	var currencyFormat = { neg:"-%s%v",pos:"%s%v",zero:"%s%v" }

	switch (format) {
		case "percent":
			return (numberVal*100.0).toFixed(2) + "%"
		case "percent0":
			return (numberVal*100.0).toFixed(0) + "%"
		case "percent1":
			return (numberVal*100.0).toFixed(1) + "%"
		case "number2":
			return numberVal.toFixed(2)	
		case "number1":
			return numberVal.toFixed(1)	
		case "integer":
			return Math.round(numberVal)	
		case "currency":
			return accounting.formatMoney(numberVal,{format:currencyFormat})
		case "currency0prec":
			return accounting.formatMoney(numberVal,{precision:0,format:currencyFormat})
		default:
			return rawVal
	}

}