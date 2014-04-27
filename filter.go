// Copyright 2014 The presenti Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strings"

	"github.com/dupoxy/presenti/notebook"
)

// newFilteredNote creats a new note with aplied filters or returns an error.
func newFilteredNote(s string) (*notebook.Note, error) {
	switch {
	// NOTE(): Since we uses /* and */ as multiple lines comments delimiter,
	// we can not use them in comments.

	// check that string only got one /* or */
	case strings.Count(s, "/*") == 1 && strings.Count(s, "*/") == 1:
		text := strings.TrimPrefix(s, noteBookMarker) // remove noteBookMarker prefix
		text = strings.TrimPrefix(text, "\n")
		text = strings.TrimLeft(text, " ")
		text = strings.Replace(text, " */", "\n", 1)
		text = strings.Replace(text, "*/", "\n", 1)
		text = strings.TrimRight(text, " ")
		note, err := notebook.NewNote(text) // create new note
		if err != nil {
			return nil, fmt.Errorf("notebook.newNote: %v", err)
		}
		return note, nil
	default:
		return nil, fmt.Errorf("newFilteredNote as more then one /* or */")
	}
}
