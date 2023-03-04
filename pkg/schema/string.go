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

type StringRegexValidatorSchema struct {
	Expression string `hcl:"expression,attr"`
}

type StringLengthValidatorSchema struct {
	Minimum *uint `hcl:"min,optional"`
	Maximum *uint `hcl:"max,optional"`
}

type StringCaseModifierSchema struct {
	Kind string `hcl:"kind,attr"`
}

type StringSchema struct {
	Name            string                       `hcl:"name,label"`
	Default         string                       `hcl:"default,attr"`
	Accessor        bool                         `hcl:"accessor,optional"`
	RegexValidator  *StringRegexValidatorSchema  `hcl:"regexValidator,block"`
	LengthValidator *StringLengthValidatorSchema `hcl:"lengthValidator,block"`
	CaseModifier    *StringCaseModifierSchema    `hcl:"caseModifier,block"`
}

type StringArraySchema struct {
	Name            string                       `hcl:"name,label"`
	Accessor        bool                         `hcl:"accessor,optional"`
	RegexValidator  *StringRegexValidatorSchema  `hcl:"regexValidator,block"`
	LengthValidator *StringLengthValidatorSchema `hcl:"lengthValidator,block"`
	CaseModifier    *StringCaseModifierSchema    `hcl:"caseModifier,block"`
}

type StringMapSchema struct {
	Name            string                       `hcl:"name,label"`
	Value           string                       `hcl:"value,attr"`
	Accessor        bool                         `hcl:"accessor,optional"`
	RegexValidator  *StringRegexValidatorSchema  `hcl:"regexValidator,block"`
	LengthValidator *StringLengthValidatorSchema `hcl:"lengthValidator,block"`
	CaseModifier    *StringCaseModifierSchema    `hcl:"caseModifier,block"`
}
