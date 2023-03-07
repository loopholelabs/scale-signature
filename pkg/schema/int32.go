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

type Int32LimitValidatorSchema struct {
	Maximum *int32 `hcl:"maximum,optional"`
	Minimum *int32 `hcl:"minimum,optional"`
}

type Int32Schema struct {
	Name           string                     `hcl:"name,label"`
	Default        int32                      `hcl:"default,attr"`
	Accessor       *bool                      `hcl:"accessor,optional"`
	LimitValidator *Int32LimitValidatorSchema `hcl:"limitValidator,block"`
}

func (s *Int32Schema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(s.Name) {
		return fmt.Errorf("invalid %s.int32 name: %s", model.Name, s.Name)
	}

	if s.LimitValidator != nil {
		if s.LimitValidator.Maximum != nil {
			if s.LimitValidator.Minimum != nil {
				if *s.LimitValidator.Minimum > *s.LimitValidator.Maximum {
					return fmt.Errorf("invalid %s.%s.limitValidator: minimum cannot be greater than maximum", model.Name, s.Name)
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

type Int32ArraySchema struct {
	Name           string                     `hcl:"name,label"`
	Accessor       *bool                      `hcl:"accessor,optional"`
	LimitValidator *Int32LimitValidatorSchema `hcl:"limitValidator,block"`
}

func (s *Int32ArraySchema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(s.Name) {
		return fmt.Errorf("invalid %s.int32Array name: %s", model.Name, s.Name)
	}

	if s.LimitValidator != nil {
		if s.LimitValidator.Maximum != nil {
			if s.LimitValidator.Minimum != nil {
				if *s.LimitValidator.Minimum > *s.LimitValidator.Maximum {
					return fmt.Errorf("invalid %s.%s.limitValidator: minimum cannot be greater than maximum", model.Name, s.Name)
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

type Int32MapSchema struct {
	Name           string                     `hcl:"name,label"`
	Value          string                     `hcl:"value,attr"`
	Accessor       *bool                      `hcl:"accessor,optional"`
	LimitValidator *Int32LimitValidatorSchema `hcl:"limitValidator,block"`
}

func (s *Int32MapSchema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(s.Name) {
		return fmt.Errorf("invalid %s.int32Map name: %s", model.Name, s.Name)
	}

	if s.LimitValidator != nil {
		if s.LimitValidator.Maximum != nil {
			if s.LimitValidator.Minimum != nil {
				if *s.LimitValidator.Minimum > *s.LimitValidator.Maximum {
					return fmt.Errorf("invalid %s.%s.limitValidator: minimum cannot be greater than maximum", model.Name, s.Name)
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
