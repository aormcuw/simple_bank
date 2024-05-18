package simplebank

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	entry1 := CreateEntriesParams{
		AccountID: account1.ID,
		Amount:    20,
	}
	entry2, err := testQueries.CreateEntries(context.Background(), entry1)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
}
