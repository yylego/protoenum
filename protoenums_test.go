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

func TestEnums_ListNonDefaultProtos(t *testing.T) {
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
	enums.WithDefault()

	protoEnums := enums.ListNonDefaultProtos()
	t.Log(protoEnums)
	require.Len(t, protoEnums, 3)
	require.Equal(t, protoenumresult.ResultEnum_PASS, protoEnums[0])
	require.Equal(t, protoenumresult.ResultEnum_MISS, protoEnums[1])
	require.Equal(t, protoenumresult.ResultEnum_SKIP, protoEnums[2])

	// NoDefault: no default to exclude, so ListNonDefault returns each element
	enumsNoDefault, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)
	require.NoError(t, err)
	allEnums := enumsNoDefault.ListNonDefaultProtos()
	t.Log(allEnums)
	require.Len(t, allEnums, 4)
}

func TestEnums_ListNonDefaultBasics(t *testing.T) {
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
	enums.WithDefault()

	validBasics := enums.ListNonDefaultBasics()
	t.Log(validBasics)
	require.Len(t, validBasics, 3)
	require.Equal(t, ResultTypePass, validBasics[0])
	require.Equal(t, ResultTypeMiss, validBasics[1])
	require.Equal(t, ResultTypeSkip, validBasics[2])

	// NoDefault: no default to exclude, so ListNonDefault returns each element
	enumsNoDefault, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)
	require.NoError(t, err)
	allBasics := enumsNoDefault.ListNonDefaultBasics()
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

// TestEnums_LookupByCode checks lookup via numeric code — true on a hit, false on a miss
//
// 验证按数字代码查找：命中返回 true，未命中返回 false
func TestEnums_LookupByCode(t *testing.T) {
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

	found, ok := enums.LookupByCode(int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.True(t, ok)
	t.Log(found.Basic())
	require.Equal(t, StatusTypeSuccess, found.Basic())

	_, ok = enums.LookupByCode(999)
	require.False(t, ok)
}

// TestEnums_LookupByName checks lookup via name — true on a hit, false on a miss
//
// 验证按名称查找：命中返回 true，未命中返回 false
func TestEnums_LookupByName(t *testing.T) {
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

	found, ok := enums.LookupByName(protoenumstatus.StatusEnum_SUCCESS.String())
	require.True(t, ok)
	t.Log(found.Basic())
	require.Equal(t, StatusTypeSuccess, found.Basic())

	_, ok = enums.LookupByName("NOT_EXISTS")
	require.False(t, ok)
}

// TestNewEnums_DuplicateBasic checks that two distinct protos sharing one basic value cause an error
//
// 验证两个不同 proto 共用同一 basic 值时返回错误
func TestNewEnums_DuplicateBasic(t *testing.T) {
	type StatusType string
	_, err := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusType("same")),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusType("same")),
	)
	require.Error(t, err)
	t.Log("expected:", err)
}

// TestNewEnums_NilElement checks that a nil Enum element gives an error
//
// 验证传入 nil Enum 元素时返回错误
func TestNewEnums_NilElement(t *testing.T) {
	type StatusType string
	_, err := protoenum.NewEnums[protoenumstatus.StatusEnum, StatusType, *protoenum.MetaNone](
		nil,
	)
	require.Error(t, err)
	t.Log("expected:", err)
}
