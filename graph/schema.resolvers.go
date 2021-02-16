package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/Dagetby/bramti/graph/generated"
	"github.com/Dagetby/bramti/graph/model"
)

func (r *mutationResolver) CreateTwit(ctx context.Context, input model.NewTwit) (*model.Twit, error) {
	if _, ok := users[input.UserID]; !ok {
		newUser := &model.User{
			ID:   input.UserID,
			Name: RandStringRunes(7),
		}
		users[input.UserID] = newUser
	}
	user := users[input.UserID]
	newTwit := &model.Twit{
		ID:              RandStringRunes(5),
		ContentText:     input.ContextText,
		PublicationDate: time.Now(),
		Author:          user,
	}

	user.Twits = append(user.Twits, newTwit)
	if m, ok := manager[input.UserID]; ok {
		m.twitChannel <- newTwit
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
	result := make(chan *model.Twit, 1)
	mu := &sync.Mutex{}

	mu.Lock()
	_, ok := users[id]
	manager[id] = &listener{
		twitChannel: result,
	}
	mu.Unlock()
	if !ok {
		return nil, errors.New("Sorry, we cant find user with this id")
	}

	go func() {
		<-ctx.Done()
		mu.Lock()
		delete(twitPublishedChannel, id)
		mu.Unlock()
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type listener struct {
	twitChannel chan *model.Twit
}

var manager map[string]*listener
var twitPublishedChannel map[string]chan *model.Twit
var users map[string]*model.User

func init() {
	manager = map[string]*listener{}
	twitPublishedChannel = map[string]chan *model.Twit{}
	users = map[string]*model.User{}

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
