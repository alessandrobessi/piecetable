package piecetable

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Node struct {
	BufferIndex int
	Offset      int
	Length      int
	LineStarts  []int
}

type PieceTable struct {
	Buffers []string
	Nodes   []Node
}

func (p *PieceTable) GetIndices(index int) (int, int, int) {
	var piece int
	var offset int
	var bufferIndex int

	remainingOffset := index
	for i, node := range p.Nodes {
		if remainingOffset <= node.Length {
			piece = i
			offset = node.Offset + remainingOffset
			bufferIndex = node.BufferIndex
			break
		}
		remainingOffset -= node.Length
	}
	return piece, offset, bufferIndex
}

func (p *PieceTable) GetText() string {
	text := ""
	for _, node := range p.Nodes {
		text += p.Buffers[node.BufferIndex][node.Offset : node.Offset+node.Length]
	}
	return text
}

func (p *PieceTable) Insert(text string, index int) {
	piece, offset, bufferIndex := p.GetIndices(index)
	newBufferIndex := len(p.Buffers)
	p.Buffers = append(p.Buffers, text)
	insertNodes := []Node{}
	if offset-p.Nodes[piece].Offset > 0 {
		insertNodes = append(insertNodes, Node{bufferIndex,
			p.Nodes[piece].Offset,
			offset - p.Nodes[piece].Offset,
			FindLineStarts(p.Buffers[bufferIndex][p.Nodes[piece].Offset : offset-p.Nodes[piece].Offset])})
	}

	insertNodes = append(insertNodes, Node{newBufferIndex, 0, len(text), FindLineStarts(text)})

	if p.Nodes[piece].Length-offset+p.Nodes[piece].Offset > 0 {
		insertNodes = append(insertNodes, Node{bufferIndex,
			offset,
			p.Nodes[piece].Length - offset + p.Nodes[piece].Offset,
			FindLineStarts(p.Buffers[bufferIndex][offset : p.Nodes[piece].Length-offset+p.Nodes[piece].Offset])})
	}

	p.Nodes = mergeNodes(p.Nodes[:piece], insertNodes, p.Nodes[piece+1:])
}

func (p *PieceTable) Delete(index int, length int) {
	startPiece, startOffset, startBufferIndex := p.GetIndices(index)
	stopPiece, stopOffset, stopBufferIndex := p.GetIndices(index + length)

	if startPiece == stopPiece {
		if p.Nodes[startPiece].Offset == startOffset {
			p.Nodes[startPiece].Offset += length
			p.Nodes[startPiece].Length -= length
			return
		} else if stopOffset == p.Nodes[startPiece].Offset+p.Nodes[startPiece].Length {
			p.Nodes[startPiece].Length -= length
			return
		}
	}

	deleteNodes := []Node{
		Node{startBufferIndex,
			p.Nodes[startPiece].Offset,
			startOffset - p.Nodes[startPiece].Offset,
			FindLineStarts(p.Buffers[startBufferIndex][p.Nodes[startPiece].Offset : startOffset-p.Nodes[startPiece].Offset]),
		},
		Node{stopBufferIndex,
			stopOffset,
			p.Nodes[stopPiece].Length - stopOffset + p.Nodes[stopPiece].Offset,
			FindLineStarts(p.Buffers[stopBufferIndex][stopOffset : p.Nodes[stopPiece].Length-stopOffset+p.Nodes[stopPiece].Offset]),
		},
	}

	deleteCount := stopPiece - startPiece + 1
	p.Nodes = mergeNodes(p.Nodes[:startPiece], deleteNodes, p.Nodes[startPiece+deleteCount:])

}

func (p *PieceTable) GetLine(index int) string {
	line := ""
	startNode := 0
	startOffset := 0
	stopNode := 0
	stopOffset := 0

	lineCount := 0
	for i, node := range p.Nodes {
		for _, lineStart := range node.LineStarts {
			lineCount += 1

			if lineCount == index {
				startNode = i
				startOffset = lineStart + 1
			}

			if lineCount == index+1 {
				stopNode = i
				stopOffset = lineStart
				break
			}
		}

	}

	if startNode == stopNode {
		node := p.Nodes[startNode]
		line += p.Buffers[startNode][node.Offset+startOffset : node.Offset+stopOffset]
		return line
	}

	for i, node := range p.Nodes[startNode : stopNode+1] {
		if startNode+i == startNode {
			line += p.Buffers[node.BufferIndex][node.Offset+startOffset : node.Offset+node.Length]
		} else if startNode+i == stopNode {
			line += p.Buffers[node.BufferIndex][node.Offset : node.Offset+stopOffset]
		} else {
			line += p.Buffers[node.BufferIndex][node.Offset : node.Offset+node.Length]
		}
	}
	return line
}

func ReadFromFile(path string, bufferSize int) PieceTable {

	p := PieceTable{
		Buffers: []string{},
		Nodes:   []Node{},
	}

	f, err := os.Open(path)

	if err != nil {
		fmt.Println("cannot read the file", err)
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	bufferIndex := 0
	for {
		buf := make([]byte, bufferSize)
		numReadBytes, err := r.Read(buf)

		if numReadBytes == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				break
			}
		}

		buf = buf[:numReadBytes]
		text := string(buf)
		p.Buffers = append(p.Buffers, text)
		p.Nodes = append(p.Nodes, Node{
			bufferIndex,
			0,
			len(text),
			FindLineStarts(text),
		})
		bufferIndex += 1

	}

	if p.Buffers[bufferIndex-1][len(p.Buffers[bufferIndex-1])-1] != '\n' {
		p.Buffers[bufferIndex-1] += "\n"
	}

	return p
}

func mergeNodes(args ...[]Node) []Node {
	mergedNodes := make([]Node, 0)
	for _, node := range args {
		mergedNodes = append(mergedNodes, node...)
	}
	return mergedNodes
}

func FindLineStarts(s string) []int {
	lineStarts := make([]int, 0)

	for i, c := range s {
		if c == '\n' {
			lineStarts = append(lineStarts, i)
		}
	}

	return lineStarts
}
