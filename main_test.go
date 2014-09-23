// Copyright 2014 The presenti Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var gopath = os.Getenv("GOPATH")
var testfilepath = "src/github.com/dupoxy/presenti/test"

var dotests = []struct {
	in  string
	out string
}{
	{"good.go", "good.out"},
	{"badnonewline/bad.go", "badnonewline/bad.out"},
	{"badnonewline/bad.go", "badnonewline/bad.out"},
	{"badnotoplev/bad.go", "badnotoplev/bad.out"},
	{"eof/good.go", "eof/good.out"},
}

func TestDo(t *testing.T) {
	gopathlist := strings.Split(gopath, ":")
	gopath = gopathlist[0]
	for _, tt := range dotests {
		inpath := gopath + "/" + testfilepath + "/" + tt.in
		outpath := gopath + "/" + testfilepath + "/" + tt.out
		b, err := exec.Command("presenti", inpath).CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		_, d := readInput(outpath)
		ok := bytes.Equal(b, d)
		if !ok {
			t.Errorf("presenti %q =>\n%s\nWant:\n%s", inpath, b, d)
		}
	}
}

func TestNewFilteredNote_Bad(t *testing.T) {

	s := `
/* I m just /* bad
*/
`
	_, err := newFilteredNote(s)
	if err == nil {
		t.Errorf("newFilteredNote %q => err == nil Want: err != nil", s)
	}
	if err != nil {
		t.Logf("newFilteredNote %q => err != nil Want: err != nil", s)
	}
}

func TestNewFilteredNote_Bad2(t *testing.T) {
	s := `
/* I m just */ bad
*/
`
	_, err := newFilteredNote(s)
	if err == nil {
		t.Errorf("newFilteredNote %q => err == nil Want: err != nil", s)
	}
	if err != nil {
		t.Logf("newFilteredNote %q => err != nil Want: err != nil", s)
	}
}
