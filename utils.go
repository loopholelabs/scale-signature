//go:build !tinygo && !js && !wasm

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

package signature

import (
	"fmt"
	"github.com/loopholelabs/scale-signature/generator"
	"os"
	"path"
	"strings"
)

func CreateGoSignature(scaleFilePath string, directory string, signaturePath string) error {
	g := generator.New()
	err := os.MkdirAll(path.Join(path.Dir(scaleFilePath), directory), 0755)
	if err != nil {
		if !os.IsExist(err) {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}

	signatureFile, err := os.OpenFile(fmt.Sprintf("%s/signature.go", path.Join(path.Dir(scaleFilePath), directory)), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error creating signature go file: %w", err)
	}

	err = g.ExecuteGoSignatureGeneratorTemplate(signatureFile, "signature", signaturePath)
	if err != nil {
		return fmt.Errorf("error generating signature go file: %w", err)
	}

	return nil
}

func CreateRustSignature(scaleFilePath string, directory string, signaturePath string) error {
	g := generator.New()
	err := os.MkdirAll(path.Join(path.Dir(scaleFilePath), directory), 0755)
	if err != nil {
		if !os.IsExist(err) {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}

	signatureFile, err := os.OpenFile(fmt.Sprintf("%s/signature.rs", path.Join(path.Dir(scaleFilePath), directory)), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error creating signature rust file: %w", err)
	}

	err = g.ExecuteRustSignatureGeneratorTemplate(signatureFile, "http", signaturePath, "HttpContext")
	if err != nil {
		return fmt.Errorf("error generating signature rust file: %w", err)
	}

	return nil
}

// ParseSignature parses and returns the Organization, Name, and Version of a signature string.
// If there is no organization, the organization will be an empty string.
// If there is no tag, the tag will be an empty string.
func ParseSignature(signature string) (string, string, string) {
	signatureOrganizationSplit := strings.Split(signature, "/")
	if len(signatureOrganizationSplit) == 1 {
		signatureOrganizationSplit = []string{"", signature}
	}
	signatureVersionSplit := strings.Split(signatureOrganizationSplit[1], "@")
	if len(signatureVersionSplit) == 1 {
		signatureVersionSplit = []string{signatureVersionSplit[0], ""}
	}
	return signatureOrganizationSplit[0], signatureVersionSplit[0], signatureVersionSplit[1]
}
