package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 안녕", "녕안 ,olleH"},
		{"", ""},
	}

	ts := httptest.NewServer(http.HandlerFunc(reverseHandler))
	defer ts.Close()

	for _, test := range tests {
		res, err := http.Get(fmt.Sprintf("%s/reverse?in=%s", ts.URL, url.QueryEscape(test.in)))
		if err != nil {
			t.Fatalf("Get error: %v", err)
		}

		got, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Fatalf("ReadAll error: %v", err)
		}

		if string(got) != test.want {
			t.Errorf("Reverse(%q) == %q, want %q", test.in, got, test.want)
		}
	}
}
