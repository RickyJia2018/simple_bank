package db

import (
	"context"
	"testing"

	"github.com/RickyJia2018/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRanomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	require.Equal(t, entry.AccountID, account.ID)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRanomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRanomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)

}

func TestListEntries(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRanomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
