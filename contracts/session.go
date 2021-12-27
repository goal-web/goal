package contracts

type Session interface {
	GetName() string
	SetName(name string)
	GetId() string
	SetId(id string)
	Start() bool
	Save()
	All() map[string]string
	Exists(key string) bool
	Has(key string) bool
	Get(key, defaultValue string) string
	Pull(key, defaultValue string) string
	Put(key, value string)
	Token() string
	RegenerateToken()
	Remove(key string) string
	Forget(keys ...string)
	Flush()
	Invalidate() bool
	Regenerate(destroy bool) bool
	Migrate(destroy bool) bool
	IsStarted() bool
	PreviousUrl() string
	SetPreviousUrl(url string)
}
