// Copyright 2021 chopnico All rights reserved.
// Use of this source code is governed by a GNU GPLv3
// license that can be found in the LICENSE file.

package output

import (
	"reflect"
)

func value(t reflect.Value) reflect.Value {
	switch t.Kind() {
	case reflect.Ptr:
		return t.Elem()
	default:
		return t
	}
}
