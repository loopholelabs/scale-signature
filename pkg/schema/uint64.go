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

type Uint64LimitValidatorSchema struct {
	Maximum *uint64 `hcl:"maximum,optional"`
	Minimum *uint64 `hcl:"minimum,optional"`
}

type Uint64Schema struct {
	Name           string                      `hcl:"name,label"`
	Default        uint64                      `hcl:"default,attr"`
	Accessor       *bool                       `hcl:"accessor,optional"`
	LimitValidator *Uint64LimitValidatorSchema `hcl:"limitValidator,block"`
}

type Uint64ArraySchema struct {
	Name           string                      `hcl:"name,label"`
	Accessor       *bool                       `hcl:"accessor,optional"`
	LimitValidator *Uint64LimitValidatorSchema `hcl:"limitValidator,block"`
}

type Uint64MapSchema struct {
	Name           string                      `hcl:"name,label"`
	Value          string                      `hcl:"value,attr"`
	Accessor       *bool                       `hcl:"accessor,optional"`
	LimitValidator *Uint64LimitValidatorSchema `hcl:"limitValidator,block"`
}
