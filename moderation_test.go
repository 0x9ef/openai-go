// Copyright (c) 2022 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package openai

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModerate(t *testing.T) {
	testCases := []struct {
		name                   string
		input                  string
		expectedHate           bool
		expectedHateThreating  bool
		expectedSelfHarm       bool
		expectedSexual         bool
		expectedSexualMinors   bool
		expectedViolence       bool
		expectedVioenceGraphic bool
	}{
		{
			name:         "success:moderate hate",
			input:        "Im going to kill all people!",
			expectedHate: true,
		},
		{
			name:                  "success:moderate hate threating",
			input:                 "I kill every white man",
			expectedHateThreating: true,
		},
		{
			name:             "success:moderate self harm",
			input:            "Imma going to suicide because I hate myself",
			expectedSelfHarm: true,
		},
		{
			name:           "success:moderate sexual",
			input:          "Naked girl",
			expectedSexual: true,
		},
		{
			name:                 "success:moderate sexual minors",
			input:                "Naked girl under 18",
			expectedSexualMinors: true,
		},
		{
			name:             "success:moderate violence",
			input:            "You have to beat up muslims",
			expectedViolence: true,
		},
	}

	e := New(os.Getenv("OPENAI_KEY"))
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := e.Moderate(context.Background(), tc.input)
			if err != nil {
				log.Fatal(err)
			}
			switch {
			case tc.expectedHate:
				assert.Equalf(t, tc.expectedHate, r.Results[0].Categories.Hate, "expected hate")
			case tc.expectedHateThreating:
				assert.Equal(t, tc.expectedHateThreating, r.Results[0].Categories.HateThreatening, "expected hate threating")
			case tc.expectedSelfHarm:
				assert.Equal(t, tc.expectedSelfHarm, r.Results[0].Categories.SelfHarm, "expected self harm")
			case tc.expectedSexual:
				assert.Equal(t, tc.expectedSexual, r.Results[0].Categories.Sexual, "expected sexual")
			case tc.expectedSexualMinors:
				assert.Equal(t, tc.expectedSexualMinors, r.Results[0].Categories.SexualMinors, "expected sexual minors")
			case tc.expectedViolence:
				assert.Equal(t, tc.expectedViolence, r.Results[0].Categories.Violence, "expected violence")
			case tc.expectedVioenceGraphic:
				assert.Equal(t, tc.expectedVioenceGraphic, r.Results[0].Categories.ViolenceGraphic, "expected violence graphic")
			}
		})
	}
}
