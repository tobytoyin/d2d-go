package contracts

// SourceTask is an interface storing different
// handling service for Source
type SourceTask interface {
	// this allow string reading from Source
	SourceIO(source Source) []byte
}
