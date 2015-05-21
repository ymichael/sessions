package sessions

import "errors"

var (
	NotFoundError = errors.New("Session Object Not Found.")
	SaveError     = errors.New("Error Saving Session Object.")
)

type Store interface {
	/**
	 * Given the session id, returns the session object associated with it, if exists.
	 * Otherwise, returns NotFoundError.
	 */
	Get(key string) (map[string]interface{}, error)

	/**
	 * Persists the session object to the store.
	 */
	Save(key string, object map[string]interface{}) error
}

/**
 * Simple in-memory session.Store
 */
type MemoryStore map[string]map[string]interface{}

func (m MemoryStore) Get(key string) (map[string]interface{}, error) {
	if val, ok := m[key]; ok {
		return val, nil
	} else {
		return nil, NotFoundError
	}
}

func (m MemoryStore) Save(key string, object map[string]interface{}) error {
	m[key] = object
	return nil
}
