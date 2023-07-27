package docustore

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
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
	return Document{}, nil
}

func (d *DocuStore) Update(key string, data Document) error {
	return nil
}

func (d *DocuStore) Delete(key string) error {
	return nil
}

func NewDocuStore() DocuStore {
	store := DocuStore{
		cache: make(map[string]Document),
	}
	store.loadLocalCache("internal/api/stores/docustore/examples")
	return store
}

// load local files into DocuStore from directory
func (d *DocuStore) loadLocalCache(dir string) {
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
		if doc.id == "" {
			generateDocumentID(&doc, filepath)
		}
		d.addDocumentToCache(doc)
	}
}

func (d *DocuStore) addDocumentToCache(doc Document) error {
	if _, found := d.cache[doc.id]; found {
		return errors.New("doc id already exists in cache")
	}
	d.cache[doc.id] = doc

	return nil
}

func generateDocumentID(doc *Document, filepath string) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println("error generating new uuid:", err)
		return
	}
	doc.id = uuid.String()

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
