package piecetable_test

import (
	pt "example.com/alessandrobessi/piecetable/pkg/piecetable"
	"testing"
)


func TestInsertStart(t *testing.T) {
	originalText := "the quick brown fox\njumped over the lazy dog"
	editedText := "Foxy, the quick brown fox\njumped over the lazy dog"
	p := pt.PieceTable{
		Buffers: []string{originalText},
		Nodes:   []pt.Node{pt.Node{BufferIndex: 0, Offset: 0, Length: len(originalText)}},
	}
	p.Insert("Foxy, ", 0)
	if p.GetText() != editedText {
		t.Error()
	}
}

func TestInsertMiddle(t *testing.T) {
	originalText := "the quick brown fox\njumped over the lazy dog"
	editedText := "the quick brown fox\nwent to the park and\njumped over the lazy dog"
	p := pt.PieceTable{
		Buffers: []string{originalText},
		Nodes:   []pt.Node{pt.Node{BufferIndex: 0, Offset: 0, Length: len(originalText)}},
	}
	p.Insert("went to the park and\n", 20)
	if p.GetText() != editedText {
		t.Error()
	}
}

func TestInsertEnd(t *testing.T) {
	originalText := "the quick brown fox\njumped over the lazy dog"
	editedText := "the quick brown fox\njumped over the lazy dog."
	p := pt.PieceTable{
		Buffers: []string{originalText},
		Nodes:   []pt.Node{pt.Node{BufferIndex: 0, Offset: 0, Length: len(originalText)}},
	}
	p.Insert(".", len(p.GetText()))
	if p.GetText() != editedText {
		t.Error()
	}
}

func TestDeleteStart(t *testing.T) {
	originalText := "Foxy, the quick brown fox\njumped over the lazy dog"
	editedText := "the quick brown fox\njumped over the lazy dog"
	p := pt.PieceTable{
		Buffers: []string{originalText},
		Nodes:   []pt.Node{pt.Node{BufferIndex: 0, Offset: 0, Length: len(originalText)}},
	}
	p.Delete(0, 6)
	if p.GetText() != editedText {
		t.Error()
	}
}

func TestDeleteMiddle(t *testing.T) {
	originalText := "the quick brown fox\njumped over the lazy dog"
	editedText := "the brown fox\njumped over the lazy dog"
	p := pt.PieceTable{
		Buffers: []string{originalText},
		Nodes:   []pt.Node{pt.Node{BufferIndex: 0, Offset: 0, Length: len(originalText)}},
	}
	p.Delete(4, 6)
	if p.GetText() != editedText {
		t.Error(p.GetText())
	}
}

func TestDeleteEnd(t *testing.T) {
	originalText := "the quick brown fox\njumped over the lazy dog."
	editedText := "the quick brown fox\njumped over the lazy dog"
	p := pt.PieceTable{
		Buffers: []string{originalText},
		Nodes:   []pt.Node{pt.Node{BufferIndex: 0, Offset: 0, Length: len(originalText)}},
	}
	p.Delete(len(p.GetText())-1, 1)

	if p.GetText() != editedText {
		t.Error()
	}
}

func TestGetLineStart(t *testing.T) {
	originalText := "the quick brown fox\nwent to the park and\njumped over the lazy dog"
	result := "the quick brown fox"
	p := pt.PieceTable{
		Buffers: []string{originalText},
		Nodes: []pt.Node{pt.Node{
			BufferIndex: 0,
			Offset:      0,
			Length:      len(originalText),
			LineStarts:  pt.FindLineStarts(originalText),
		}},
	}

	if p.GetLine(0) != result {
		t.Error()
	}
}

func TestGetLineMiddle(t *testing.T) {
	originalText := "the quick brown fox\nwent to the park and\njumped over the lazy dog"
	result := "went to the park and"
	p := pt.PieceTable{
		Buffers: []string{originalText},
		Nodes: []pt.Node{pt.Node{
			BufferIndex: 0,
			Offset:      0,
			Length:      len(originalText),
			LineStarts:  pt.FindLineStarts(originalText),
		}},
	}

	if p.GetLine(1) != result {
		t.Error()
	}
}

func TestGetLineEnd(t *testing.T) {
	originalText := "the quick brown fox\nwent to the park and\njumped over the lazy dog\n"
	result := "jumped over the lazy dog"
	p := pt.PieceTable{
		Buffers: []string{originalText},
		Nodes: []pt.Node{pt.Node{
			BufferIndex: 0,
			Offset:      0,
			Length:      len(originalText),
			LineStarts:  pt.FindLineStarts(originalText),
		}},
	}

	if p.GetLine(2) != result {
		t.Error()
	}
}
