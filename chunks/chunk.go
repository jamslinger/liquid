// Package chunks parses and evaluates template strings.
//
// A template string is composed of a sequence of chunk.
// Each chunk is either an object {{a.b}}, a tag {%if a>b%}, or a text chunk (anything outside of {{}} and {%%}.)
package chunks

import "fmt"

// Chunk is a chunk of a template source. It is one of an object "{{…}}", a tag "{%…%}", or the text between objects and tags.
type Chunk struct {
	Type       ChunkType
	SourceInfo SourceInfo
	Name       string // Name is the tag name of a tag Chunk. E.g. the tag name of "{% if 1 %}" is "if".
	Parameters string // Parameters is the tag arguments of a tag Chunk. E.g. the tag arguments of "{% if 1 %}" is "1".
	Source     string // Source is the entirety of the chunk, including the "{{", "{%", etc. markers.
}

func (c Chunk) String() string {
	switch c.Type {
	case TextChunkType:
		return fmt.Sprintf("%s{%#v}", c.Type, c.Source)
	case TagChunkType:
		return fmt.Sprintf("%s{Tag:%#v, Args:%#v}", c.Type, c.Name, c.Parameters)
	case ObjChunkType:
		return fmt.Sprintf("%s{%#v}", c.Type, c.Parameters)
	default:
		return fmt.Sprintf("%s{%#v}", c.Type, c.Source)
	}
}

// SourceInfo contains a Chunk's source information
type SourceInfo struct {
	Pathname string
	lineNo   int
}

// ChunkType is the type of a Chunk
type ChunkType int

const (
	// TextChunkType is the type of a text Chunk
	TextChunkType ChunkType = iota
	// TagChunkType is the type of a tag Chunk "{%…%}"
	TagChunkType
	// ObjChunkType is the type of an object Chunk "{{…}}"
	ObjChunkType
)