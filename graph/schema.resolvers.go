package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/Dagetby/bramti/graph/utils"
	"time"

	"github.com/Dagetby/bramti/graph/generated"
	"github.com/Dagetby/bramti/graph/model"
)

func (r *mutationResolver) CreateTwit(ctx context.Context, input model.NewTwit) (*model.Twit, error) {
	user, ok := r.checkUser(input.UserID)
	if !ok {
		return nil, errors.New("We don't have user with this Id\n Please create new user")
	}

	newTwit := &model.Twit{
		ID:              utils.RandStringRunes(5),
		ContentText:     input.ContextText,
		PublicationDate: time.Now(),
		Author:          user,
	}
	user.Twits = append(user.Twits, newTwit)

	m, ok := r.checkListener(input.UserID)
	if ok {
		m <- newTwit
	}
	return newTwit, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.addUser(&model.User{
		ID:   input.UserID,
		Name: input.Name,
	}), nil
}

func (r *queryResolver) Twits(ctx context.Context, id string, limit *int, offset *int) ([]*model.Twit, error) {
	result, err := r.getTwits(id)
	if err != nil {
		return nil, err
	}

	if limit != nil && offset != nil {
		start := *offset
		end := *limit + *offset

		if end > len(result) {
			end = len(result)
		}

		return result[start:end], nil
	}

	return result, nil
}

func (r *subscriptionResolver) TwitPublished(ctx context.Context, id string) (<-chan *model.Twit, error) {
	_, ok := r.checkUser(id)
	if !ok {
		return nil, errors.New("Sorry, we cant find user with this id")
	}
	result := r.addListener(id)
	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(r.listener, id)
		r.mu.Unlock()
	}()

	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
