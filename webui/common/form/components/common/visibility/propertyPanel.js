// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initFormComponentVisibilityPropertyPanel(params) {

	var visibilityFilterConditionPropertyPanelParams = {
		elemPrefix: params.elemPrefix,
		databaseID: params.databaseID,
		defaultFilterRules: params.initialConditions,
		initDone: function () {},
		updateFilterRules: function (updatedFilterRules) {

			console.log("Updating form component filter conditions: " + JSON.stringify(updatedFilterRules))
			params.saveVisibilityConditionsCallback(updatedFilterRules)		
		}

	}
	initFilterPropertyPanel(visibilityFilterConditionPropertyPanelParams)

}
