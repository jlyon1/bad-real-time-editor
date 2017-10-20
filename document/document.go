package document

import (
	"log"
	"time"
)

const (
	ADD = 1
	DEL = 0
)

type Document struct {
	docValue string
}

type Delta struct {
	Updatetime time.Time
	Operation  int
  StartPosition   int
	Change     string
}

func (document *Document) GetDocumentValue() string {
	return document.docValue
}

func (document *Document) OverwriteText(val string) {
	document.docValue = val
	log.Println(val)
}

func New(val string) Document {
	d := Document{
		docValue: "asdf",
	}
	return d
}
