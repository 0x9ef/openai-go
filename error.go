// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package openai

import "encoding/json"

type APIError struct {
	Err struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
		Type       string `json:"type"`
	} `json:"error"`
}

func (e APIError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "undefined error"
	}
	return string(b)
}
