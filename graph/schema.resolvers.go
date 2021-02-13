package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/Dagetby/bramti/api"

	"github.com/Dagetby/bramti/graph/generated"
	"github.com/Dagetby/bramti/graph/model"
)

func (r *mutationResolver) CreateTwit(ctx context.Context, input model.NewTwit) (*api.Twit, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Twits(ctx context.Context) ([]*api.Twit, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
