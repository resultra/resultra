// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initSubmitFormUILayoutPanes()
{
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	
	// Initialize the page layout
	var mainLayout = $('#layoutPage').layout({
		inset: zeroPaddingInset,
		north: fixedUILayoutPaneParams(40),
	})
				
}



