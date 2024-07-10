package pkg

type Query[T any, Svc any] interface {
	// ExecuteQuery executes a Query on service.
	ExecuteQuery(svc Svc) (T, error)
}
