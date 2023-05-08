/*
	Copyright 2023 Loophole Labs
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

package rust

import (
	"fmt"
	"github.com/loopholelabs/scale-signature/pkg/schema"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerator(t *testing.T) {
	g, err := New()
	require.NoError(t, err)

	s := new(schema.Schema)
	err = s.Decode([]byte(schema.MasterTestingSchema))
	require.NoError(t, err)

	require.NoError(t, s.Validate())

	formatted, err := g.Generate(s, "types", "v0.1.0")
	if err != nil {
		fmt.Printf("%s", err)
		t.Fatal(err)
	}
	t.Log(formatted)

}
