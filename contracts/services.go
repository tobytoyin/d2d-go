package contracts

// SourceTask is an interface storing different
// handling service for Source
type SourceTask interface {
	// this allow string reading from Source
	SourceIO(source Source) []byte
	UiD(soure Source) string

	Contents(source Source) DocContent
	Metadata(source Source) DocMetadata
	Summary(source Source) DocSummary
}
