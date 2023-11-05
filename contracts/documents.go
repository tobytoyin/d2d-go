package contracts

type DocumentComponent interface{}

type DocContent struct {
	DocumentComponent

	Contents string
	Byte     []byte
}

type DocUID string

type DocSummary string

// type DocUID struct {
// 	string
// 	DocumentComponent
// }

type DocMetadata struct {
	// A Struct for the metadata in a document
	DocType    string
	Properties map[string]interface{}
}
