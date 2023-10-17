package db

import (
	"context"
	"testing"
	"time"

	"github.com/afrianska/go-try/simplebank/pkg/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accountFrom, accountTo Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: accountFrom.ID,
		ToAccountID: accountTo.ID,
		Amount: utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T)  {
	accountFrom := createRandomAccount(t)
	accountTo := createRandomAccount(t)
	createRandomTransfer(t, accountFrom, accountTo)
}

func TestGetTransfer(t *testing.T)  {
	accountFrom := createRandomAccount(t)
	accountTo := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, accountFrom, accountTo)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T)  {
	accountFrom := createRandomAccount(t)
	accountTo := createRandomAccount(t)

	for i := 0; i < 4; i++ {
		createRandomTransfer(t, accountFrom, accountTo)
	}

	arg := ListTransferParams{
		FromAccountID: accountFrom.ID,
		ToAccountID: accountTo.ID,
		Limit: 2,
		Offset: 2,
	}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 2)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
		require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	}
}
