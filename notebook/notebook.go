// Copyright 2014 The notebook Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package notebook provides a notebook for present .article/.slide file.
package notebook

import "fmt"

// Note represents a note.
// IDs are set only for Note that are saved by a NoteBook.
type Note struct {
	ID   int64  // Unique identifier
	Text string // Description
}

// NewNote creates a new Note given a string, that can't be empty.
func NewNote(text string) (*Note, error) {
	if text == "" {
		return nil, fmt.Errorf("empty note")
	}
	return &Note{0, text}, nil
}

// NoteBook manages a list of Note in memory.
type NoteBook struct {
	notes  []*Note
	lastID int64
}

// NewNoteBook returns an empty NoteBook.
func NewNoteBook() *NoteBook {
	return &NoteBook{}
}

// Save saves the given Note in the NoteBook.
func (m *NoteBook) Save(note *Note) error {
	if note.ID == 0 {
		m.lastID++
		note.ID = m.lastID
		m.notes = append(m.notes, cloneNote(note))
		return nil
	}

	for i, t := range m.notes {
		if t.ID == note.ID {
			m.notes[i] = cloneNote(note)
			return nil
		}
	}
	return fmt.Errorf("unknown note")
}

// cloneNote creates and returns a deep copy of the given Notes.
func cloneNote(t *Note) *Note {
	c := *t
	return &c
}

// All returns the list of all the Notes in the NoteBook.
func (m *NoteBook) All() []*Note {
	return m.notes
}
