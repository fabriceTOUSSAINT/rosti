package graph

//go:generate go run github.com/99designs/gqlgen generate
import (
	"github.com/fabriceTOUSSAINT/rosti/server/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
// 	ginContext := ctx.Value("GinContextKey")
// 	if ginContext == nil {
// 		err := fmt.Errorf("could not retrieve gin.Context")
// 		return nil, err
// 	}

// 	gc, ok := ginContext.(*gin.Context)
// 	if !ok {
// 		err := fmt.Errorf("gin.Context has wrong type")
// 		return nil, err
// 	}
// 	return gc, nil
// }

// func (r *resolver) Todo(ctx context.Context) (*Todo, error) {
// 	gc, err := GinContextFromContext(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// ...
// }


type Resolver struct{
	todos []*model.Todo
	// todo model.Todo
}
