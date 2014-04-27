// Copyright 2014 The presenti Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"code.google.com/p/go.tools/present"
	"github.com/dupoxy/presenti/notebook"
)

// readInput reads file from input and return it's name and data.
func readInput(input string) (name string, data []byte) {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return input, data
}

// createAndOpen creates a present file and returns it or an error.
// You must call f.Close() wen your done using f.
func createAndOpen(arg, suffix string, nb *notebook.NoteBook) (f *os.File, err error) {
	f, err = os.Create(output(arg, suffix))
	if err != nil {
		return nil, fmt.Errorf("Error Creating %s: %s\n", f.Name(), err)
	}
	for _, v := range nb.All() {
		fmt.Fprintf(f, v.Text)
	}
	f.Close()
	f, err = os.Open(output(arg, suffix))
	if err != nil {
		return nil, fmt.Errorf("Error Opening %s: %s\n", f.Name(), err)
	}
	return f, nil
}

// output trims .go suffix from arg and returns arg with the specified suffix.
func output(arg, suffix string) (out string) {
	out = strings.TrimSuffix(arg, ".go")
	out = out + suffix
	return out
}

// check f is a parsable present file and return error if any.
func check(f io.Reader, name string) (err error) {
	_, err = present.Parse(f, name, 0)
	if err != nil {
		return err
	}
	return nil
}

// topLevelComment returns true if column position is 1.
func topLevelComment(fset *token.FileSet, pos token.Pos) bool {
	p := strings.Split(fset.Position(pos).String(), ":")
	if p[2] == "1" {
		return true
	}
	return false

}

// lineNum returns the line number from fset and pos.
func lineNum(fset *token.FileSet, pos token.Pos) int {
	p := strings.Split(fset.Position(pos).String(), ":")
	i, err := strconv.Atoi(p[1])
	if err != nil {
		log.Fatal(err)
	}
	return i
}
