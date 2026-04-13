package protoenum_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/protoenum"
	"github.com/yylego/protoenum/protos/protoenumstatus"
)

func TestEnums_DefaultValue(t *testing.T) {
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

	// Test that first item becomes default
	defaultEnum, err := enums.GetDefault()
	require.NoError(t, err)
	require.Equal(t, StatusTypeUnknown, defaultEnum.Basic())

	// Test lookup returns default when not found
	notFound := enums.GetByCode(999)
	require.NotNil(t, notFound)
	require.Equal(t, defaultEnum, notFound)

	// Test GetByName returns default when not found
	notFoundByName := enums.GetByName("NOT_EXISTS")
	require.NotNil(t, notFoundByName)
	require.Equal(t, defaultEnum, notFoundByName)

	// Test GetByBasic returns default when not found
	notFoundByBasic := enums.GetByBasic(StatusType("not_exists"))
	require.NotNil(t, notFoundByBasic)
	require.Equal(t, defaultEnum, notFoundByBasic)
}

func TestEnums_GetDefaultProto(t *testing.T) {
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

	defaultProto, err := enums.GetDefaultProto()
	require.NoError(t, err)
	require.Equal(t, protoenumstatus.StatusEnum_UNKNOWN, defaultProto)

	// Test returns error when no default is set
	require.NoError(t, enums.UnsetDefault())
	_, err = enums.GetDefaultProto()
	require.Error(t, err)
}

func TestEnums_GetDefaultBasic(t *testing.T) {
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

	defaultBasic, err := enums.GetDefaultBasic()
	require.NoError(t, err)
	require.Equal(t, StatusTypeUnknown, defaultBasic)

	// Test returns error when no default is set
	require.NoError(t, enums.UnsetDefault())
	_, err = enums.GetDefaultBasic()
	require.Error(t, err)
}

func TestEnums_SetDefault(t *testing.T) {
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

	// Must unset default first before setting new one
	require.NoError(t, enums.UnsetDefault())

	// Now set default to SUCCESS via LookupByBasic
	successEnum, ok := enums.LookupByBasic(StatusTypeSuccess)
	require.True(t, ok)
	require.NoError(t, enums.SetDefault(successEnum))

	// Check new default
	newDefault, err := enums.GetDefault()
	require.NoError(t, err)
	require.Equal(t, StatusTypeSuccess, newDefault.Basic())
}

func TestEnums_SetDefault_ReturnsErrorOnDuplicate(t *testing.T) {
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

	// SetDefault should return error because default is present (first item)
	successEnum, ok := enums.LookupByBasic(StatusTypeSuccess)
	require.True(t, ok)
	err = enums.SetDefault(successEnum)
	require.Error(t, err)
	t.Log("expected:", err)
}

func TestEnums_SetDefault_ReturnsErrorOnNil(t *testing.T) {
	type StatusType string

	enums, err := protoenum.NewEnums[protoenumstatus.StatusEnum, StatusType, *protoenum.MetaNone]()
	require.NoError(t, err)

	err = enums.SetDefault(nil)
	require.Error(t, err)
	t.Log("expected:", err)
}

func TestEnums_UnsetDefault(t *testing.T) {
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

	_, err = enums.GetDefault()
	require.NoError(t, err)

	require.NoError(t, enums.UnsetDefault())

	// GetDefault returns error once unset
	_, err = enums.GetDefault()
	require.Error(t, err)

	// UnsetDefault again returns error
	err = enums.UnsetDefault()
	require.Error(t, err)
}

func TestEnums_WithUnsetDefault(t *testing.T) {
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

	enums = enums.WithUnsetDefault()
	_, err = enums.GetDefault()
	require.Error(t, err)
}

func TestEnums_ChainMethods(t *testing.T) {
	t.Run("with-default-enum", func(t *testing.T) {
		type StatusType string
		const (
			StatusTypeSuccess StatusType = "success"
		)

		enums, err := protoenum.NewEnums[protoenumstatus.StatusEnum, StatusType, *protoenum.MetaNone]()
		require.NoError(t, err)

		enums = enums.WithDefault(protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess))
		def, err := enums.GetDefault()
		require.NoError(t, err)
		require.Equal(t, StatusTypeSuccess, def.Basic())
	})

	t.Run("with-default-panics-on-existing", func(t *testing.T) {
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

		require.Panics(t, func() {
			enums.WithDefaultCode(int32(protoenumstatus.StatusEnum_SUCCESS))
		})
	})

	t.Run("unset-then-set-default", func(t *testing.T) {
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

		require.NoError(t, enums.UnsetDefault())
		successEnum, ok := enums.LookupByBasic(StatusTypeSuccess)
		require.True(t, ok)
		require.NoError(t, enums.SetDefault(successEnum))
		def, err := enums.GetDefault()
		require.NoError(t, err)
		require.Equal(t, StatusTypeSuccess, def.Basic())
	})

	t.Run("with-invalid-code-panics", func(t *testing.T) {
		type StatusType string

		enums, err := protoenum.NewEnums[protoenumstatus.StatusEnum, StatusType, *protoenum.MetaNone]()
		require.NoError(t, err)

		require.Panics(t, func() {
			enums.WithDefaultCode(999)
		})
	})

	t.Run("with-invalid-name-panics", func(t *testing.T) {
		type StatusType string

		enums, err := protoenum.NewEnums[protoenumstatus.StatusEnum, StatusType, *protoenum.MetaNone]()
		require.NoError(t, err)

		require.Panics(t, func() {
			enums.WithDefaultName("NOT_EXISTS")
		})
	})
}
