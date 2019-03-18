// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function populateNumberFormatSelection($selection) {
	$selection.empty()
	$selection.append(defaultSelectOptionPromptHTML("Select a number format"))
	
	$selection.append(selectOptionHTML("number2","Number 1022.00"))
	$selection.append(selectOptionHTML("number1","Number 1022.2"))
	$selection.append(selectOptionHTML("integer","Rounded Integer 1022"))
	$selection.append(selectOptionHTML("percent","Percent 5.55%"))
	$selection.append(selectOptionHTML("percent0","Percent 5%"))
	$selection.append(selectOptionHTML("percent1","Percent 5.5%"))
	$selection.append(selectOptionHTML("currency","Currency (USD) $1022.00"))
	$selection.append(selectOptionHTML("currency0prec","Currency (USD) $1022"))
		
}