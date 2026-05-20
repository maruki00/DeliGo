package cache

// can be redis or in-memory cache for lightwieght
type Cache interface {
	Get() any
	Set(any) error
	Delete(any) error
	Clear() error
}
