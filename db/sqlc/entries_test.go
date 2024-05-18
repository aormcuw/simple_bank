package simplebank

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T) Entry {
	account := createRandomAccount(t)
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    20,
	}
	entry, err := testQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}
func TestCreateEntry(t *testing.T) {
	createRandomEntries(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntries(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}
