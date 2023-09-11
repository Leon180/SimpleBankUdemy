package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, claims)

	require.NotZero(t, claims.ID)
	require.Equal(t, username, claims.Username)
	require.WithinDuration(t, issuedAt, claims.IssuedAt.Time, time.Second)
	require.WithinDuration(t, expiredAt, claims.ExpiresAt.Time, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, "token has invalid claims: token is expired")
	require.Nil(t, claims)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	claims, err := NewPayload(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	claims, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, "token is unverifiable: error while executing keyfunc: token is invalid")
	require.Nil(t, claims)
}
