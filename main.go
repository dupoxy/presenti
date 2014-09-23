// Copyright 2014 The presenti Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* #nb:
A inline presentation system for Go

The presenti Authors
https://www.github.com/dupoxy/presenti/
*/

package main

import (
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
)

/* #nb: * Install

To Install:

	$ go get -u github.com/dupoxy/presenti
	
build status:

.image https://drone.io/github.com/dupoxy/presenti/status.png

* Readme

.code README


* Be nice

It try to be nice to godoc and other Go tools too.
To do so, it uses a notebook marker:

.code main.go /^const/,/\n/

And also by only using top level comments and checking that they have an umpty \n after the end.

* why

To keep talks in sync with the code.

* Constrains

- keep it as simple as possible.
- A .go file equal an .article/Slide file to group code by subject.
- do not replace doc.go, godoc and flag.Usage().
- use multiple line comments so they can be minimized if your code editor does that.
*/

const noteBookMarker = "/* #nb:"

func usage() {
	txt := `Presenti is an inline present system for Go.

It aim to produce article or slide file by extracting multiples lines
top level comments, from start to end of Go source file.

Usage of %s:

%s [<flags>] input.go

output default to stdout.

Flags:
`
	fmt.Fprintf(os.Stderr, txt, os.Args[0], os.Args[0])
	flag.PrintDefaults()
	end := `
To Displays slide presentations and articles,
use the present command from the go.talks subrepository.
It runs a web server that presents slide and article files
from the current directory.

To Install it:
	$ go get -u code.google.com/p/go.talks/present

for more on presenti:
	
	$ cd $GOPATH/src/github.com/dupoxy/presenti
	$ present
`
	fmt.Fprintln(os.Stderr, end)

}

var (
	article = flag.Bool("a", false, "save output to an article file")
	slide   = flag.Bool("s", false, "save output to a slide file")	
)

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	err := do(args[0])
	if err != nil {
		log.Fatal(err)
	}
}

func do(input string) error {
	// read args
	name, data := readInput(input)
	// check if it is a valid source file
	data, err := format.Source(data)
	if err != nil {
		return fmt.Errorf("format.Source: %v", err)
	}
	// make notebook
	noteBook, err := makeNoteBook(name, data)
	if err != nil {
		return fmt.Errorf("makeNotebook: %v", err)
	}
	file := false
	suf := ""
	if *article {
		file = true
		suf = ".article"
	}
	if *slide {
		file = true
		suf = ".slide"
	}
	switch {
	case file: // write output to file
		f, err := createAndOpen(input, suf, noteBook)
		defer f.Close()
		if err != nil {
			return fmt.Errorf("createAndOpen: %v", err)
		}
		// Parse present file to check errors
		err = check(f, f.Name())
		if err != nil {
			del := os.Remove(f.Name())
			if del != nil {
				return fmt.Errorf("os.Remove: %v", del)
			}
			return fmt.Errorf("check: %v", err)
		}

	default: // write output to stdout
		for _, v := range noteBook.All() {
			fmt.Fprintf(os.Stdout, v.Text)
		}
		suf = ".tmp"
		f, err := createAndOpen(input, suf, noteBook)
		defer f.Close()
		if err != nil {
			return fmt.Errorf("createAndOpen: %v", err)
		}
		// Parse present file to check errors
		err = check(f, f.Name())
		if err != nil {
			del := os.Remove(f.Name())
			if del != nil {
				return fmt.Errorf("os.Remove: %v", del)
			}
			return fmt.Errorf("check: %v", err)
		}
		err = os.Remove(f.Name())
		if err != nil {
			return fmt.Errorf("os.Remove: %v", err)
		}
	}
	return nil
}
