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

package schema

import "fmt"

type EnumSchema struct {
	Name    string   `hcl:"name,label"`
	Default string   `hcl:"default,attr"`
	Values  []string `hcl:"values,attr"`
}

func (s *EnumSchema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(s.Name) {
		return fmt.Errorf("invalid %s.enum name: %s", model.Name, s.Name)
	}

	for _, value := range s.Values {
		if value == s.Default {
			return nil
		}
	}

	return fmt.Errorf("invalid %s.%s.default: %s is not a valid value", model.Name, s.Name, s.Default)
}

type EnumArraySchema struct {
	Name   string   `hcl:"name,label"`
	Values []string `hcl:"values,attr"`
}

func (s *EnumArraySchema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(s.Name) {
		return fmt.Errorf("invalid %s.enumArray name: %s", model.Name, s.Name)
	}

	return nil
}

type EnumMapSchema struct {
	Name     string   `hcl:"name,label"`
	Values   []string `hcl:"values,attr"`
	Value    string   `hcl:"value,attr"`
	Accessor bool     `hcl:"accessor,optional"`
}

func (s *EnumMapSchema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(s.Name) {
		return fmt.Errorf("invalid %s.enumMap name: %s", model.Name, s.Name)
	}

	return nil
}
