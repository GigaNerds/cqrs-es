package cqrs_es

// Query is an interface for querying aggregates by some rules.
type Query[T any, Svc any] interface {
	// ExecuteQuery executes a Query on service.
	ExecuteQuery(svc Svc) (T, error)
}
