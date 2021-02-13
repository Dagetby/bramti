package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"time"

	"github.com/Dagetby/bramti/graph/generated"
	"github.com/Dagetby/bramti/graph/model"
)

func (r *mutationResolver) CreateTwit(ctx context.Context, input model.NewTwit) (*model.Twit, error) {
	if _, ok := users[input.UserID]; !ok {
		newUser := &model.User{
			ID:   input.UserID,
			Name: string(rand.Int()),
		}
		users[input.UserID] = newUser
	}
	user := users[input.UserID]
	newTwit := &model.Twit{
		ContentText:     input.ContextText,
		PublicationDate: time.Now(),
		AuthorID:        user,
	}

	user.Twits = append(user.Twits, newTwit)

	for _, observer := range twitPublishedChannel {
		observer <- newTwit
	}

	return newTwit, nil
}

func (r *queryResolver) Twits(ctx context.Context, id string, limit *int, offset *int) ([]*model.Twit, error) {
	var result []*model.Twit
	var allTodos = users[id].Twits

	result = allTodos

	if limit != nil && offset != nil {
		start := *offset
		end := *limit + *offset

		if end > len(allTodos) {
			end = len(allTodos)
		}

		return result[start:end], nil
	}

	return result, nil
}

func (r *subscriptionResolver) TwitPublished(ctx context.Context, id string) (<-chan *model.Twit, error) {

	twitEvent := make(chan *model.Twit, 1)
	go func() {
		<-ctx.Done()
	}()
	twitPublishedChannel[id] = twitEvent
	return twitEvent, nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var users map[string]*model.User
var twitPublishedChannel map[string]chan *model.Twit

func init() {
	twitPublishedChannel = map[string]chan *model.Twit{}
	users = map[string]*model.User{}
}
