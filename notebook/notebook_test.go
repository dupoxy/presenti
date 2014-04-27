// Copyright 2014 The notebook Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notebook

import "testing"

func newNoteOrFatal(t *testing.T, text string) *Note {
	note, err := NewNote("text")
	if err != nil {
		t.Fatalf("new note: %v", err)
	}
	return note
}

func TestNewNote(t *testing.T) {
	text := "text"
	note := newNoteOrFatal(t, text)
	if note.Text != text {
		t.Errorf("expected text %q, got %q", text, note.Text)
	}
}

func TestNewNoteEmptyText(t *testing.T) {
	_, err := NewNote("")
	if err == nil {
		t.Errorf("note with empty text created")
	}
}

func TestSaveNoteAndRetrieve(t *testing.T) {
	note := newNoteOrFatal(t, "test")
	m := NewNoteBook()
	m.Save(note)
	all := m.All()
	if len(all) != 1 {
		t.Errorf("expected 1 note, got %v", len(all))
	}
	if *all[0] != *note {
		t.Errorf("expected %v, got %v", note, all)
	}
}

func TestSaveAndRetrieveTwoNotes(t *testing.T) {
	test := newNoteOrFatal(t, "test")
	test2 := newNoteOrFatal(t, "test2")
	m := NewNoteBook()
	m.Save(test)
	m.Save(test2)
	all := m.All()
	if len(all) != 2 {
		t.Errorf("expected 2 notes, got %v", len(all))
	}
	if *all[0] != *test && *all[1] != *test {
		t.Errorf("missing note: %v", test)
	}
	if *all[0] != *test2 && *all[1] != *test2 {
		t.Errorf("missing note: %v", test2)
	}
}
