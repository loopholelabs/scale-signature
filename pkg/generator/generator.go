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

package generator

import (
	"bytes"
	"github.com/loopholelabs/scale-signature/pkg/generator/templates"
	"github.com/loopholelabs/scale-signature/pkg/schema"
	"go/format"
	"text/template"
)

const (
	packageName = "types"
)

type Generator struct {
	templ *template.Template
}

func New() (*Generator, error) {
	templ, err := template.New("").Funcs(templateFunctions()).ParseFS(templates.FS, "*.templ")
	if err != nil {
		return nil, err
	}

	return &Generator{
		templ: templ,
	}, nil
}

func (g *Generator) Generate(schema *schema.Schema, version string) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := g.templ.ExecuteTemplate(buf, "types.go.templ", map[string]any{
		"schema":  schema,
		"version": version,
		"package": packageName,
	})
	if err != nil {
		return nil, err
	}
	return format.Source(buf.Bytes())
}

func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"Primitive":               schema.Primitive,
		"IsPrimitive":             schema.ValidPrimitiveType,
		"PolyglotPrimitive":       schema.PolyglotPrimitive,
		"PolyglotPrimitiveEncode": schema.PolyglotPrimitiveEncode,
		"PolyglotPrimitiveDecode": schema.PolyglotPrimitiveDecode,
		"Deref":                   func(i *bool) bool { return *i },
		"LowerFirst":              func(s string) string { return string(s[0]+32) + s[1:] },
	}
}
