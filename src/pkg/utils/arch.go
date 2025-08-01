/*
 * Copyright © 2023 – 2025 Red Hat Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

func getDefaultReleaseArch() (string, error) {
	return "latest", nil
}

func getFullyQualifiedImageArch(image, release string) string {
	imageFull := "quay.io/toolbx/" + image
	return imageFull
}

func getP11KitClientPathsArch() []string {
	paths := []string{"/usr/lib/pkcs11/p11-kit-client.so"}
	return paths
}

func parseReleaseArch(release string) (string, error) {
	if release != "latest" && release != "rolling" && release != "" {
		return "", &ParseReleaseError{"The release must be 'latest'."}
	}

	return "latest", nil
}
