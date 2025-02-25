package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordShouldMatch(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NoError(t, CheckPassword(password, hashedPassword))

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEqual(t, hashedPassword, hashedPassword2)
}

func TestPasswordShouldNotdMatch(t *testing.T) {
	password := RandomString(6)
	wrongPassword := RandomString(6)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)

	err = CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
