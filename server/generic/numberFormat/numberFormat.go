// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package numberFormat

const NumberFormatGeneral string = "general"
const NumberFormatCurrency string = "currency"
const NumberFormatCurrency0Prec string = "currency0prec"
const NumberFormatPercent0 string = "percent0"
const NumberFormatPercent1 string = "percent1"
const NumberFormatPercent string = "percent"

type NumberFormatProperties struct {
	Format string `json:"format"`
}

func DefaultNumberFormatProperties() NumberFormatProperties {
	return NumberFormatProperties{Format: "general"}
}
