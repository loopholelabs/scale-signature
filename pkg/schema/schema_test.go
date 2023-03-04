/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package schema

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSchema(t *testing.T) {
	s := new(Schema)
	err := s.Decode([]byte(`
name = "testName"
tag = "1testTag"
model testModel {
	description = "this is a test model"
    string testString {
		default = "asdfsa"
	    regexValidator {
			expression = ".*"
		}
		lengthValidator {
			min = 1
			max = 3
		}
	}
}

model testModel2 {
	model "myTest" {
		reference = "testModel"
	}
}
`))
	require.NoError(t, err)

	require.NoError(t, s.Validate())
}
