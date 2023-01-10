package db

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
)

func generateRandomUser() CreateUserParams {
	return CreateUserParams{
		Username:  faker.Username(),
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  faker.Password(),
	}
}

func TestCreateUser(t *testing.T) {
	arg := generateRandomUser()

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
}

func TestDeleteUser(t *testing.T) {
	arg := generateRandomUser()

	user, err := testQueries.CreateUser(context.Background(), arg)

	err = testQueries.DeleteUser(context.Background(), user.ID)

	require.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	arg := generateRandomUser()

	user, err := testQueries.CreateUser(context.Background(), arg)

	getUser, err := testQueries.GetUser(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, getUser)

	require.Equal(t, user.Email, getUser.Email)
	require.Equal(t, arg.Username, getUser.Username)
	require.Equal(t, arg.FirstName, getUser.FirstName)
	require.Equal(t, arg.LastName, getUser.LastName)

	require.NotZero(t, getUser.ID)
	require.Equal(t, user.CreatedAt, getUser.CreatedAt)
	require.Equal(t, user.UpdatedAt, getUser.UpdatedAt)
}

func TestListUsers(t *testing.T) {
	_, err := testQueries.CreateUser(context.Background(), generateRandomUser())
	_, err = testQueries.CreateUser(context.Background(), generateRandomUser())
	_, err = testQueries.CreateUser(context.Background(), generateRandomUser())
	_, err = testQueries.CreateUser(context.Background(), generateRandomUser())

	require.NoError(t, err)

	arg := ListUsersParams{
		Limit:  2,
		Offset: 0,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, users)
	require.Len(t, users, int(arg.Limit))
}

func TestUpdateUser(t *testing.T) {
	// create user
	u, err := testQueries.CreateUser(context.Background(), generateRandomUser())
	// update user
	arg := UpdateUserParams{
		ID:        u.ID,
		FirstName: "Test",
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	require.NotEqual(t, u.FirstName, updatedUser.FirstName)
	require.Equal(t, arg.Email, updatedUser.Email)
	require.Equal(t, arg.Username, updatedUser.Username)
	require.Equal(t, arg.FirstName, updatedUser.FirstName)
	require.Equal(t, arg.LastName, updatedUser.LastName)

	require.NotZero(t, updatedUser.ID)
	require.Equal(t, u.CreatedAt, updatedUser.CreatedAt)
	require.NotEqual(t, updatedUser.CreatedAt, updatedUser.UpdatedAt)
}

func TestUpdateUserPassword(t *testing.T) {
	// create user
	u, err := testQueries.CreateUser(context.Background(), generateRandomUser())
	// update user
	arg := UpdateUserPasswordParams{
		ID:       u.ID,
		Password: "Test-Password",
	}

	err = testQueries.UpdateUserPassword(context.Background(), arg)

	require.NoError(t, err)
}
