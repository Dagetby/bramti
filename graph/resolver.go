package graph

import (
	"errors"
	"github.com/Dagetby/bramti/graph/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	user     *model.User
	twits    []*model.Twit
	listener map[string]chan *model.Twit
	users    map[string]*model.User
	mu       *sync.Mutex
}

func NewResolver() *Resolver {
	return &Resolver{
		user:     &model.User{},
		twits:    make([]*model.Twit, 0),
		listener: map[string]chan *model.Twit{},
		users:    map[string]*model.User{},
		mu:       &sync.Mutex{},
	}
}

func (r *Resolver) checkUser(id string) (*model.User, bool) {
	r.mu.Lock()
	user, ok := r.users[id]
	r.mu.Unlock()
	if ok {
		return user, true
	}
	return nil, false
}

func (r *Resolver) addUser(user *model.User) *model.User {
	r.mu.Lock()
	r.users[user.ID] = user
	r.mu.Unlock()
	return user
}

func (r *Resolver) addListener(id string) chan *model.Twit {
	ch := make(chan *model.Twit)
	r.mu.Lock()
	r.listener[id] = ch
	r.mu.Unlock()
	return ch
}

func (r *Resolver) checkListener(id string) (chan *model.Twit, bool) {
	r.mu.Lock()
	l, ok := r.listener[id]
	r.mu.Unlock()
	if ok {
		return l, true
	}
	return nil, false
}

func (r *Resolver) getTwits(id string) ([]*model.Twit, error) {
	r.mu.Lock()
	user, ok := r.users[id]
	r.mu.Unlock()
	if !ok {
		return nil, errors.New("We don't have user with this ID")
	}
	return user.Twits, nil
}
