package protoenum_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/protoenum"
	"github.com/yylego/protoenum/protos/protoenumstatus"
)

// WithDefault fixes the first element (UNKNOWN) as the default;
// GetDefault* then return it and a miss falls back to it.
func TestEnums_WithDefault(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	)
	enums.WithDefault()

	def, ok := enums.GetDefault()
	require.True(t, ok)
	require.Equal(t, protoenumstatus.StatusEnum_UNKNOWN, def.Proto())
	require.Equal(t, StatusTypeUnknown, def.Basic())

	// miss falls back to the default
	miss, ok := enums.GetByCodeFallbackDefault(999)
	require.True(t, ok)
	require.Equal(t, StatusTypeUnknown, miss.Basic())
	// hit returns the element
	hit, ok := enums.GetByBasicFallbackDefault(StatusTypeSuccess)
	require.True(t, ok)
	require.Equal(t, StatusTypeSuccess, hit.Basic())
}

// Without WithDefault, GetDefault returns ok=false and a miss returns nil.
func TestEnums_NoDefault(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	)

	def, ok := enums.GetDefault()
	require.False(t, ok)
	require.Nil(t, def)
	// miss + no default → (nil, false)
	miss, ok := enums.GetByCodeFallbackDefault(999)
	require.False(t, ok)
	require.Nil(t, miss)
	// hit returns the element
	hit, ok := enums.GetByBasicFallbackDefault(StatusTypeSuccess)
	require.True(t, ok)
	require.Equal(t, StatusTypeSuccess, hit.Basic())
}

// The default is fixed just once; a second WithDefault panics.
func TestEnums_Default_SetOnce(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	)

	enums.WithDefault()

	// A second WithDefault panics; there is no unset
	require.Panics(t, func() {
		enums.WithDefault()
	})
}
