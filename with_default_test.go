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

	enums, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	)
	require.NoError(t, err)
	enums.WithDefault()

	defProto, err := enums.GetDefaultProto()
	require.NoError(t, err)
	require.Equal(t, protoenumstatus.StatusEnum_UNKNOWN, defProto)

	defBasic, err := enums.GetDefaultBasic()
	require.NoError(t, err)
	require.Equal(t, StatusTypeUnknown, defBasic)

	require.Equal(t, StatusTypeUnknown, enums.GetByCode(999).Basic())                // miss falls back to default
	require.Equal(t, StatusTypeSuccess, enums.GetByBasic(StatusTypeSuccess).Basic()) // hit returns the element
}

// Without WithDefault, GetDefault reports an error and a miss returns nil.
func TestEnums_NoDefault(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
	)

	enums, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	)
	require.NoError(t, err)

	_, err = enums.GetDefault()
	require.Error(t, err)
	_, err = enums.GetDefaultProto()
	require.Error(t, err)
	_, err = enums.GetDefaultBasic()
	require.Error(t, err)
	require.Nil(t, enums.GetByCode(999))                                             // miss returns nil
	require.Equal(t, StatusTypeSuccess, enums.GetByBasic(StatusTypeSuccess).Basic()) // hit returns the element
}

// The default is fixed just once; a second WithDefault panics.
func TestEnums_Default_SetOnce(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
	)

	enums, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	)
	require.NoError(t, err)

	enums.WithDefault()

	// A second WithDefault panics; there is no unset
	require.Panics(t, func() {
		enums.WithDefault()
	})
}
