//go:build !tinygo && !js && !wasm

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

package signature

import (
	"strings"
)

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
