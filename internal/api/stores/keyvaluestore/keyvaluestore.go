package keyvaluestore

type KeyValueStore struct {
	cache map[string]string
}

func (d *KeyValueStore) Create(data string) error {
	return nil
}

func (d *KeyValueStore) Read(key string) (string, error) {
	return "", nil
}

func (d *KeyValueStore) Update(key string, data string) error {
	return nil
}

func (d *KeyValueStore) Delete(key string) error {
	return nil
}
