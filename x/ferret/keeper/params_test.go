package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "Ferret/testutil/keeper"
	"Ferret/x/ferret/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.FerretKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
