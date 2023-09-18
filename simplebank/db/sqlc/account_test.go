package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	// Create a new test account with mock(test) data.
	arg := CreateAccountParams{

		Owner:    "Adir",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	// Check that the create account is successful (err == nil).
	require.NoError(t, err)
	//  Check that the account object is not empty.
	require.NotEmpty(t, account)
	// Check that the account arguments (arg) we passed to create account function, match the account object that was created
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	// Check that postgress created an ID for the account
	require.NotZero(t, account.ID)
	// Check that the created_at field is not empty
	require.NotZero(t, account.CreatedAt)

}
