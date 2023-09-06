package docustore

import "fmt"

type ErrDocumentNotFound struct {
	id string
}

func (e *ErrDocumentNotFound) Error() string {
	return fmt.Sprintf("document id %s not found", e.id)
}

type ErrDocumentAlreadyExists struct {
	id string
}

func (e *ErrDocumentAlreadyExists) Error() string {
	return fmt.Sprintf("document id %s already exists", e.id)
}
