package document

import (
  "log"

)
type Document struct {
  docValue string
}


func (document *Document) GetDocumentValue() (string){
  return document.docValue
}

func (document *Document) OverwriteText(val string){
  document.docValue = val
  log.Println(val);
}

func New(val string) (Document){
  d := Document{
    docValue: "asdf",
  }
  return d
}
