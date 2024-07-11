package repository

// Operation is some repository operation that can be executed.
type Operation[Repo any, Res any] interface {
	// ExecuteOperation executes this operation and returns it's result.
	ExecuteOperation(repo Repo) (Res, error)
}
