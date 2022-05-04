package keeper_test

import (
	"testing"

	testkeeper "github.com/Karan-3108/Fortress/testutil/keeper"
	"github.com/Karan-3108/Fortress/x/fortress/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FortressKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
