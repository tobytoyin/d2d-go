package contracts

type Source struct {
	Path  string
	Type  string
	Extra *map[string]string // optional map for accessing Source
}
