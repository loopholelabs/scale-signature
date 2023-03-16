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

type Uint64LimitValidatorSchema struct {
	Maximum *uint64 `hcl:"maximum,optional"`
	Minimum *uint64 `hcl:"minimum,optional"`
}

type Uint64Schema struct {
	Name           string                      `hcl:"name,label"`
	Default        uint64                      `hcl:"default,attr"`
	Accessor       *bool                       `hcl:"accessor,optional"`
	LimitValidator *Uint64LimitValidatorSchema `hcl:"limit_validator,block"`
}

func (s *Uint64Schema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(s.Name) {
		return fmt.Errorf("invalid %s.uint64 name: %s", model.Name, s.Name)
	}

	if s.LimitValidator != nil {
		if s.LimitValidator.Maximum != nil {
			if s.LimitValidator.Minimum != nil {
				if *s.LimitValidator.Minimum > *s.LimitValidator.Maximum {
					return fmt.Errorf("invalid %s.%s.limit_validator: minimum cannot be greater than maximum", model.Name, s.Name)
				}
			}
		}
	}

	if s.Accessor != nil {
		if *s.Accessor == false && s.LimitValidator != nil {
			return fmt.Errorf("invalid %s.%s.accessor: cannot be false while using validators or modifiers", model.Name, s.Name)
		}
	} else {
		if s.LimitValidator != nil {
			s.Accessor = new(bool)
			*s.Accessor = true
		} else {
			s.Accessor = new(bool)
			*s.Accessor = false
		}
	}

	return nil
}

type Uint64ArraySchema struct {
	Name           string                      `hcl:"name,label"`
	InitialSize    uint64                      `hcl:"initial_size,attr"`
	Accessor       *bool                       `hcl:"accessor,optional"`
	LimitValidator *Uint64LimitValidatorSchema `hcl:"limit_validator,block"`
}

func (s *Uint64ArraySchema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(s.Name) {
		return fmt.Errorf("invalid %s.uint64_array name: %s", model.Name, s.Name)
	}

	if s.LimitValidator != nil {
		if s.LimitValidator.Maximum != nil {
			if s.LimitValidator.Minimum != nil {
				if *s.LimitValidator.Minimum > *s.LimitValidator.Maximum {
					return fmt.Errorf("invalid %s.%s.limit_validator: minimum cannot be greater than maximum", model.Name, s.Name)
				}
			}
		}
	}

	if s.Accessor != nil {
		if *s.Accessor == false && s.LimitValidator != nil {
			return fmt.Errorf("invalid %s.%s.accessor: cannot be false while using validators or modifiers", model.Name, s.Name)
		}
	} else {
		if s.LimitValidator != nil {
			s.Accessor = new(bool)
			*s.Accessor = true
		} else {
			s.Accessor = new(bool)
			*s.Accessor = false
		}
	}

	return nil
}

type Uint64MapSchema struct {
	Name           string                      `hcl:"name,label"`
	Value          string                      `hcl:"value,attr"`
	Accessor       *bool                       `hcl:"accessor,optional"`
	LimitValidator *Uint64LimitValidatorSchema `hcl:"limit_validator,block"`
}

func (s *Uint64MapSchema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(s.Name) {
		return fmt.Errorf("invalid %s.uint64_map name: %s", model.Name, s.Name)
	}

	if s.LimitValidator != nil {
		if s.LimitValidator.Maximum != nil {
			if s.LimitValidator.Minimum != nil {
				if *s.LimitValidator.Minimum > *s.LimitValidator.Maximum {
					return fmt.Errorf("invalid %s.%s.limit_validator: minimum cannot be greater than maximum", model.Name, s.Name)
				}
			}
		}
	}

	if s.Accessor != nil {
		if *s.Accessor == false && s.LimitValidator != nil {
			return fmt.Errorf("invalid %s.%s.accessor: cannot be false while using validators or modifiers", model.Name, s.Name)
		}
	} else {
		if s.LimitValidator != nil {
			s.Accessor = new(bool)
			*s.Accessor = true
		} else {
			s.Accessor = new(bool)
			*s.Accessor = false
		}
	}
	return nil
}
