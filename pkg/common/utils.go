/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package common

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ToCamelCase transforms k8s resource Name with camel case, for examples:
// - make-food to MakeFood
// - make.food to MakeFood
func ToCamelCase(input string) string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == '.' || r == '-'
	})
	titleCase := cases.Title(language.English)
	for i, word := range words {
		words[i] = titleCase.String(word)
	}
	return strings.Join(words, "")
}
