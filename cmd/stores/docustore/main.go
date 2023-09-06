package main

import (
	"fmt"

	"github.com/johnmwood/distributed-learning/internal/api/stores/docustore"
)

func main() {
	store := docustore.NewDocuStore()
	doc, err := store.Read("7475f809-baea-4c99-8d20-6c458aad6f5d")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(doc)
}
