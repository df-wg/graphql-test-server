package graph_test

import (
	"context"
	"github.com/df-wg/graphql-test-server/graph"
	"github.com/df-wg/graphql-test-server/graph/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	resolver := &graph.Resolver{
		UserList: []*model.User{},
	}

	input := model.UserInput{
		Email:     "john.doe@example.com",
		FirstName: "John",
		LastName:  "Doe",
	}
	user, err := resolver.Mutation().CreateUser(context.TODO(), input)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, input.Email, user.Email)
	assert.Equal(t, input.FirstName, user.FirstName)
	assert.Equal(t, input.LastName, user.LastName)
	assert.Equal(t, "John Doe", user.FullName)
	assert.NotEqual(t, "", user.ID)
}

func TestUsers(t *testing.T) {
	resolver := &graph.Resolver{
		UserList: []*model.User{
			{
				ID:        "1",
				Email:     "john.doe@example.com",
				FirstName: "John",
				LastName:  "Doe",
				FullName:  "John Doe",
			},
			{
				ID:        "2",
				Email:     "jane.doe@example.com",
				FirstName: "Jane",
				LastName:  "Doe",
				FullName:  "Jane Doe",
			},
		},
	}

	// Test querying users without filters
	users, err := resolver.Query().Users(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	// Test querying users with a first name filter
	firstName := "John"
	users, err = resolver.Query().Users(context.Background(), &firstName, nil)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "John", users[0].FirstName)

	// Test querying users with email ordering
	orderByEmail := true
	users, err = resolver.Query().Users(context.Background(), nil, &orderByEmail)
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "jane.doe@example.com", users[0].Email)
	assert.Equal(t, "john.doe@example.com", users[1].Email)
	assert.NotEqual(t, users[0].ID, users[1].ID)
}

func TestDeleteUser(t *testing.T) {
	resolver := &graph.Resolver{
		UserList: []*model.User{
			{
				ID:        "1",
				Email:     "john.doe@example.com",
				FirstName: "John",
				LastName:  "Doe",
				FullName:  "John Doe",
			},
		},
	}

	// Test deleting a user
	deleted, err := resolver.Mutation().DeleteUser(context.Background(), "1")
	assert.NoError(t, err)
	assert.True(t, deleted)
	assert.Len(t, resolver.UserList, 0)

	// Test deleting a non-existent user
	deleted, err = resolver.Mutation().DeleteUser(context.Background(), "2")
	assert.False(t, deleted)
	assert.Error(t, err)
}
