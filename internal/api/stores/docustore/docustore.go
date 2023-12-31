package docustore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Document struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

type DocuStore struct {
	cache map[string]Document
}

func (d *DocuStore) Create(data Document) error {
	return nil
}

func (d *DocuStore) Read(key string) (Document, error) {
	if doc, found := d.cache[key]; found {
		return doc, nil
	}
	return Document{}, &ErrDocumentNotFound{key}
}

func (d *DocuStore) Update(key string, data Document) error {
	if _, found := d.cache[key]; found {
		d.cache[key] = data
	}
	return &ErrDocumentNotFound{key}
}

func (d *DocuStore) Delete(key string) error {
	if _, found := d.cache[key]; found {
		delete(d.cache, key)
	}
	return &ErrDocumentNotFound{key}
}

func NewDocuStore() DocuStore {
	store := DocuStore{
		cache: make(map[string]Document),
	}
	store.loadLocalCache("../../../internal/api/stores/docustore/examples")
	return store
}

// load local files into DocuStore from directory
func (d *DocuStore) loadLocalCache(dir string) {
	fmt.Println(os.Getwd())
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("loadLocalCache failed with error:", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filepath := filepath.Join(dir, file.Name())
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Println("error reading file:", err)
			continue
		}

		doc := Document{}
		if err := json.Unmarshal(data, &doc); err != nil {
			log.Println("error unmarshaling json:", err)
		}
		if doc.ID == "" {
			generateDocumentID(&doc, filepath)
		}
		d.addDocumentToCache(doc)
	}
}

func (d *DocuStore) addDocumentToCache(doc Document) error {
	if _, found := d.cache[doc.ID]; found {
		return errors.New("doc id already exists in cache")
	}
	d.cache[doc.ID] = doc

	return nil
}

func generateDocumentID(doc *Document, filepath string) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println("error generating new uuid:", err)
		return
	}
	doc.ID = uuid.String()

	updatedData, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		log.Println("error marshaling new uuid:", err)
		return
	}
	err = ioutil.WriteFile(filepath, updatedData, 0644)
	if err != nil {
		log.Println("error writing new uuid back to json file:", err)
		return
	}
}
