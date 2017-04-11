function populateNumberFormatSelection($selection) {
	$selection.empty()
	$selection.append(defaultSelectOptionPromptHTML("Select a number format"))
	
	$selection.append(selectOptionHTML("generic","General 1022.00"))
	$selection.append(selectOptionHTML("percent","Percent 5.55%"))
	$selection.append(selectOptionHTML("currency","Currency (USD) $1022.00"))
	$selection.append(selectOptionHTML("currency0prec","Currency (USD) $1022"))
		
}