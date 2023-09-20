package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, claims, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, claims)
	require.NotEmpty(t, token)

	claims, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, claims)

	require.NotZero(t, claims.ID)
	require.Equal(t, username, claims.Username)
	require.WithinDuration(t, issuedAt, claims.IssuedAt.Time, time.Second)
	require.WithinDuration(t, expiredAt, claims.ExpiresAt.Time, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, claims, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, claims)
	require.NotEmpty(t, token)

	claims, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, claims)
}

func TestInvalidPasetoTokenAlgNone(t *testing.T) {
	// to do
}
