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

package golang

import (
	"bytes"
	"errors"
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
	templ, err := template.New("").Funcs(templateFunctions()).ParseFS(templates.FS, "*go.templ")
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
		"Primitive":               Primitive,
		"IsPrimitive":             schema.ValidPrimitiveType,
		"PolyglotPrimitive":       PolyglotPrimitive,
		"PolyglotPrimitiveEncode": PolyglotPrimitiveEncode,
		"PolyglotPrimitiveDecode": PolyglotPrimitiveDecode,
		"Deref":                   func(i *bool) bool { return *i },
		"LowerFirst":              func(s string) string { return string(s[0]+32) + s[1:] },
		"Params":                  Params,
	}
}

func Primitive(t string) string {
	switch t {
	case "string", "int32", "int64", "uint32", "uint64", "float32", "float64", "bool":
		return t
	case "bytes":
		return "[]byte"
	default:
		return ""
	}
}

func PolyglotPrimitive(t string) string {
	switch t {
	case "string":
		return "polyglot.StringKind"
	case "int32":
		return "polyglot.Int32Kind"
	case "int64":
		return "polyglot.Int64Kind"
	case "uint32":
		return "polyglot.Uint32Kind"
	case "uint64":
		return "polyglot.Uint64Kind"
	case "float32":
		return "polyglot.Float32Kind"
	case "float64":
		return "polyglot.Float64Kind"
	case "bool":
		return "polyglot.BoolKind"
	case "bytes":
		return "polyglot.BytesKind"
	default:
		return "polyglot.AnyKind"
	}
}

func PolyglotPrimitiveEncode(t string) string {
	switch t {
	case "string":
		return "String"
	case "int32":
		return "Int32"
	case "int64":
		return "Int64"
	case "uint32":
		return "Uint32"
	case "uint64":
		return "Uint64"
	case "float32":
		return "Float32"
	case "float64":
		return "Float64"
	case "bool":
		return "Bool"
	case "bytes":
		return "Bytes"
	default:
		return ""
	}
}

func PolyglotPrimitiveDecode(t string) string {
	switch t {
	case "string":
		return "String"
	case "int32":
		return "Int32"
	case "int64":
		return "Int64"
	case "uint32":
		return "Uint32"
	case "uint64":
		return "Uint64"
	case "float32":
		return "Float32"
	case "float64":
		return "Float64"
	case "bool":
		return "Bool"
	case "bytes":
		return "Bytes"
	default:
		return ""
	}
}

func Params(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("parameters must be a list of key/value pairs")
	}
	params := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("keys must be strings")
		}
		params[key] = values[i+1]
	}
	return params, nil
}
