package kvstore

type KVStore interface {
	Set(k string, v string) error
	Get(k string) (string, error)
}
