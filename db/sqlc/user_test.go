package db

import (
	"context"
	"testing"

	"github.com/edwins-leonardi/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomOwner())
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		FullName:       util.RandomOwner() + " " + util.RandomOwner(),
		Email:          util.RandomOwner() + "@gmail.com",
		HashedPassword: hashedPassword,
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	u1 := createRandomUser(t)

	u2, err := testQueries.GetUser(context.Background(), u1.Username)

	require.Empty(t, err)
	require.NoError(t, err)

	require.Equal(t, u1.Username, u2.Username)
	require.Equal(t, u1.FullName, u2.FullName)
	require.Equal(t, u1.Email, u2.Email)
	require.Equal(t, u1.HashedPassword, u2.HashedPassword)
	require.Equal(t, u1.CreatedAt, u2.CreatedAt)
	require.Equal(t, u1.PasswordChangedAt, u2.PasswordChangedAt)
}
