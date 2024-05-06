package application_test

import (
	"testing"
	"vegas/src/application"

	"github.com/stretchr/testify/require"
)

func TestRocket_IsValid(t *testing.T) {
	rocket := application.Rocket{}

	rocket.Amount = 10

	ok, err := rocket.IsValid()
	require.Nil(t, err)
	require.Equal(t, true, ok)
	require.Equal(t, rocket.Status, application.FLYING)

	rocket.Amount = -1
	ok, err = rocket.IsValid()

	require.Equal(t, "the amount should be greater than zero", err.Error())
	require.Equal(t, false, ok)

	rocket.Amount = 0
	rocket.Status = "err"
	ok, err = rocket.IsValid()

	require.Equal(t, "the status should be CRASH or FLYING", err.Error())
	require.Equal(t, false, ok)
}

func TestRocket_Fly(t *testing.T) {
	rocket := application.Rocket{}

	rocket.Amount = 10

	err := rocket.Fly(-1)
	require.Equal(t, "seconds must be greater than zero", err.Error())

	err = rocket.Fly(0)
	require.Equal(t, "seconds must be greater than zero", err.Error())

	err = rocket.Fly(5)
	require.Nil(t, err)

	require.Equal(t, application.CRASH, rocket.Status)
	require.GreaterOrEqual(t, int32(41), rocket.Amount)
}
