package kvstore

type memoryStore map[string]string

func NewMemoryKVStore() KVStore {
	return memoryStore{}
}

func (m memoryStore) Get(k string) (string, error) {
	k, ok := m[k]
	if !ok {
		return "", KeyDoesNotExist
	}

	return k, nil
}

func (m memoryStore) Set(k string, v string) error {
	m[k] = v
	return nil
}
