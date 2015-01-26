// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stringutil

import (
	"fmt"
	"testing"
)

func ExampleValues() {
	v := Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")

	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
	// Output:
	// Ava
	// Jess
	// [Jess Sarah Zoe]
}

func TestReverse(t *testing.T) {
	// Controlling whether a test runs at all.
	// if runtime.GOARCH == "arm" {
	//        t.Skip("this doesn't work on ARM")
	// }

	var tests = []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 안녕", "녕안 ,olleH"},
		{"", ""},
	}

	for _, test := range tests {
		got := Reverse(test.in)
		if got != test.want {
			t.Errorf("Reverse(%q) == %q, want %q", test.in, got, test.want)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("Hello")
	}
}
