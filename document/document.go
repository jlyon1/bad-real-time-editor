package document

import (

)
type Document struct {
  docValue string
}


func (document Document) GetDocumentValue() (string){
  return document.docValue;
}

func New(val string) (Document){
  d := Document{
    docValue: "asdf",
  }
  return d
}
