package store

type Store interface {
	Save(val map[string]interface{}) error
	Load(params map[string]interface{}) (interface{}, error)
	Initialize(params map[string]interface{}) error
}
