// Copyright 2021 chopnico All rights reserved.
// Use of this source code is governed by a GNU GPLv3
// license that can be found in the LICENSE file.

package output

import (
	"encoding/json"
	"reflect"
	"strings"
)

// exported function to print an item as json
func FormatItemAsJson(t interface{}) string {
	b := strings.Builder{}

	j, _ := json.Marshal(t)
	b.WriteString(string(j))

	return b.String()
}

// exported function to print an item as pretty json
func FormatItemAsPrettyJson(t interface{}) string {
	b := strings.Builder{}

	j, _ := json.MarshalIndent(t, "", "  ")
	b.WriteString(string(j))

	return b.String()
}

// exported function to print items as pretty json
func FormatItemsAsPrettyJson(t interface{}) string {
	a := value(reflect.ValueOf(t))
	b := strings.Builder{}

	for i := 0; i < a.Len(); i++ {
		j, _ := json.MarshalIndent(a.Index(i).Interface(), "", "  ")

		b.WriteString(string(j) + "\n")
	}
	return b.String()
}

// exported function to print items as json
func FormatItemsAsJson(t interface{}) string {
	a := value(reflect.ValueOf(t))
	b := strings.Builder{}

	for i := 0; i < a.Len(); i++ {
		j, _ := json.Marshal(a.Index(i).Interface())

		b.WriteString(string(j) + "\n")
	}

	return b.String()
}
