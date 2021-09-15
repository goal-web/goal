package contracts

type Route interface {

}

type Router interface {
	Get(path string) Route
	Post(path string) Route
	Delete(path string) Route
	Put(path string) Route
}
