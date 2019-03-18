// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package userAuth

type AuthResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

func newAuthResponse(success bool, msg string) *AuthResponse {
	resp := AuthResponse{
		Success: success,
		Msg:     msg}
	return &resp
}
