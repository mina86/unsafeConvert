// Unsafe string↔[]byte conversion library
// Copyright (c) 2017 by Michał Nazarewicz <mina86@mina86.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unsafeConvert

import (
	"fmt"
	"testing"
)

var strings = []string{
	repeat(' ', 0),
	repeat(' ', 1),
	"foo",
	"żółw",
	repeat(' ', 100),
	repeat(' ', 10000),
	repeat(' ', 1000000),
}

func repeat(ch byte, l int) string {
	var data = make([]byte, l)
	for i := 0; i < l; i++ {
		data[i] = ch
	}
	return string(data)
}

func TestString(t *testing.T) {
	for _, want := range strings {
		arg := []byte(want)
		got := String(arg)
		if got != want {
			t.Errorf("String(%q) = %q but want %q",
				arg, got, want)
		}
	}
}

func eq(a, b []byte) bool {
	i, l := 0, len(a)
	if len(a) != len(b) {
		return false
	}
	for i < l && a[i] == b[i] {
		i++
	}
	return i == l
}

func TestBytes(t *testing.T) {
	for _, arg := range strings {
		want := []byte(arg)
		got := Bytes(arg)
		if !eq(got, want) {
			t.Errorf("Bytes(%q) = %q but want %q",
				arg, got, want)
		}
	}
}

func TestNil(t *testing.T) {
	got := String(nil)
	if got != "" {
		t.Errorf("String(nil) = %q but want %q", got, "")
	}
}

func lengths(f func(l int)) {
	for l := 0; l <= 10000000; {
		f(l)
		if l == 0 {
			l = 10
		} else {
			l *= 100
		}
	}
}

func BenchmarkString(b *testing.B) {
	lengths(func(l int) {
		arg := make([]byte, l)

		name := fmt.Sprintf("String/%d", l)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n += 1 {
				String(arg)
			}
		})

		name = fmt.Sprintf("string/%d", l)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n += 1 {
				_ = string(arg)
			}
		})
	})
}

func BenchmarkBytes(b *testing.B) {
	lengths(func(l int) {
		arg := string(make([]byte, l))

		name := fmt.Sprintf("Bytes/%d", l)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n += 1 {
				Bytes(arg)
			}
		})

		name = fmt.Sprintf("[]byte/%d", l)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n += 1 {
				_ = []byte(arg)
			}
		})
	})
}
