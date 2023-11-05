package contracts

import (
	"d2d/contracts"
	"fmt"
	"testing"
)

func TestDocuContentCreation(t *testing.T) {
	docContent := new(contracts.DocContent)

	docContent.Contents = "anc"

	fmt.Println(docContent)

}

func TestDocUID(t *testing.T) {
	DocUID := contracts.DocUID("asd")

	fmt.Println(DocUID)
}
