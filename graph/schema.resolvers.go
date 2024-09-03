package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/df-wg/graphql-test-server/graph/model"
	"math/big"
	"sort"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	//TODO: Find better way of making unique id
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	id := fmt.Sprintf("U%d", randNumber)
	user := &model.User{
		ID:        id,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		FullName:  fmt.Sprintf("%s %s", input.FirstName, input.LastName),
	}
	r.UserList = append(r.UserList, user)
	return user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	for i, user := range r.UserList {
		if user.ID == id {
			r.UserList = append(r.UserList[:i], r.UserList[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New(fmt.Sprintf("Failed to find user %s", id))
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context, firstName *string, orderByEmail *bool) ([]*model.User, error) {
	users := r.UserList
	if firstName != nil {
		var filteredUsers []*model.User
		for _, user := range users {
			if user.FirstName == *firstName {
				filteredUsers = append(filteredUsers, user)
			}
		}
		users = filteredUsers
	}

	if orderByEmail != nil && *orderByEmail {
		sort.Slice(users, func(i, j int) bool {
			return users[i].Email < users[j].Email
		})
	}

	return users, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
