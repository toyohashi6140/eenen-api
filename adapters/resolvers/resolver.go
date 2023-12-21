package resolvers

import "github.com/toyohashi6140/eenen-api/usecases/lyricusecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	LyricUsecase lyricusecase.LyricUsecase
}

func NewResolver() *Resolver {
	return &Resolver{
		LyricUsecase: lyricusecase.NewLyricUsecase(),
	}
}
