package usecases

import "context"

type (
	ManyReader[T any] interface {
		FetchByKey(context.Context) (T, error)
	}
)
