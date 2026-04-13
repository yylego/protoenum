package protoenum_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/protoenum"
	"github.com/yylego/protoenum/protos/protoenumresult"
	"github.com/yylego/protoenum/protos/protoenumstatus"
)

func TestEnums_GetByProto(t *testing.T) {
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

	enum := enums.GetByProto(protoenumstatus.StatusEnum_SUCCESS)
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Basic(), StatusTypeSuccess)
}

func TestEnums_GetByCode(t *testing.T) {
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

	enum := enums.GetByCode(int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_FAILURE.String())
	require.Equal(t, enum.Basic(), StatusTypeFailure)
}

func TestEnums_GetByName(t *testing.T) {
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

	enum := enums.GetByName(protoenumstatus.StatusEnum_UNKNOWN.String())
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_UNKNOWN.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_UNKNOWN.String())
	require.Equal(t, enum.Basic(), StatusTypeUnknown)
}

func TestEnums_GetByBasic(t *testing.T) {
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

	enum := enums.GetByBasic(StatusTypeSuccess)
	require.NotNil(t, enum)
	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Basic(), StatusTypeSuccess)

	// Use LookupByBasic to check existence
	enumFailure, ok := enums.LookupByBasic(StatusTypeFailure)
	require.True(t, ok)
	require.Equal(t, enumFailure.Code(), int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.Equal(t, enumFailure.Basic(), StatusTypeFailure)
}

func TestEnums_LookupByProto_NotFound(t *testing.T) {
	type StatusType string
	const (
		StatusTypeSuccess StatusType = "success"
	)

	enums, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	)
	require.NoError(t, err)

	_, ok := enums.LookupByProto(protoenumstatus.StatusEnum_FAILURE)
	require.False(t, ok)
}

func TestEnums_ListProtos(t *testing.T) {
	type ResultType string

	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	enums, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)
	require.NoError(t, err)

	protoEnums := enums.ListProtos()
	t.Log(protoEnums)
	require.Len(t, protoEnums, 4)
	require.Equal(t, protoenumresult.ResultEnum_UNKNOWN, protoEnums[0])
	require.Equal(t, protoenumresult.ResultEnum_PASS, protoEnums[1])
	require.Equal(t, protoenumresult.ResultEnum_MISS, protoEnums[2])
	require.Equal(t, protoenumresult.ResultEnum_SKIP, protoEnums[3])
}

func TestEnums_ListBasics(t *testing.T) {
	type ResultType string

	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	enums, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)
	require.NoError(t, err)

	basicEnums := enums.ListBasics()
	t.Log(basicEnums)
	require.Len(t, basicEnums, 4)
	require.Equal(t, ResultTypeUnknown, basicEnums[0])
	require.Equal(t, ResultTypePass, basicEnums[1])
	require.Equal(t, ResultTypeMiss, basicEnums[2])
	require.Equal(t, ResultTypeSkip, basicEnums[3])
}

func TestEnums_ListValidProtos(t *testing.T) {
	type ResultType string

	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	enums, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)
	require.NoError(t, err)

	protoEnums := enums.ListValidProtos()
	t.Log(protoEnums)
	require.Len(t, protoEnums, 3)
	require.Equal(t, protoenumresult.ResultEnum_PASS, protoEnums[0])
	require.Equal(t, protoenumresult.ResultEnum_MISS, protoEnums[1])
	require.Equal(t, protoenumresult.ResultEnum_SKIP, protoEnums[2])

	require.NoError(t, enums.UnsetDefault())
	allEnums := enums.ListValidProtos()
	t.Log(allEnums)
	require.Len(t, allEnums, 4)
}

func TestEnums_ListValidBasics(t *testing.T) {
	type ResultType string

	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	enums, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)
	require.NoError(t, err)

	validBasics := enums.ListValidBasics()
	t.Log(validBasics)
	require.Len(t, validBasics, 3)
	require.Equal(t, ResultTypePass, validBasics[0])
	require.Equal(t, ResultTypeMiss, validBasics[1])
	require.Equal(t, ResultTypeSkip, validBasics[2])

	require.NoError(t, enums.UnsetDefault())
	allBasics := enums.ListValidBasics()
	t.Log(allBasics)
	require.Len(t, allBasics, 4)
}

func TestNewEnums_DuplicateProto(t *testing.T) {
	type StatusType string
	_, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusType("a")),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusType("b")),
	)
	require.Error(t, err)
	t.Log("expected:", err)
}
