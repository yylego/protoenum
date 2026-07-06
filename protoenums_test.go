package protoenum_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/protoenum"
	"github.com/yylego/protoenum/protos/protoenumresult"
	"github.com/yylego/protoenum/protos/protoenumstatus"
)

func TestEnums_GetByProtoFallbackDefault(t *testing.T) {
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

	enum, ok := enums.GetByProtoFallbackDefault(protoenumstatus.StatusEnum_SUCCESS)
	require.True(t, ok)
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Basic(), StatusTypeSuccess)
}

func TestEnums_GetByCodeFallbackDefault(t *testing.T) {
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

	enum, ok := enums.GetByCodeFallbackDefault(int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.True(t, ok)
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_FAILURE.String())
	require.Equal(t, enum.Basic(), StatusTypeFailure)
}

func TestEnums_GetByNameFallbackDefault(t *testing.T) {
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

	enum, ok := enums.GetByNameFallbackDefault(protoenumstatus.StatusEnum_UNKNOWN.String())
	require.True(t, ok)
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_UNKNOWN.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_UNKNOWN.String())
	require.Equal(t, enum.Basic(), StatusTypeUnknown)
}

func TestEnums_GetByBasicFallbackDefault(t *testing.T) {
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

	enum, ok := enums.GetByBasicFallbackDefault(StatusTypeSuccess)
	require.True(t, ok)
	require.NotNil(t, enum)
	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Basic(), StatusTypeSuccess)

	// Use GetByBasic to check existence
	enumFailure, ok := enums.GetByBasic(StatusTypeFailure)
	require.True(t, ok)
	require.Equal(t, enumFailure.Code(), int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.Equal(t, enumFailure.Basic(), StatusTypeFailure)
}

// TestEnums_FallbackDefault_Miss covers each fallback branch
// TestEnums_FallbackDefault_Miss 覆盖每个键的回落分支
func TestEnums_FallbackDefault_Miss(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	).WithDefault() // UNKNOWN becomes the default

	// Each miss falls back to the default (UNKNOWN)
	byProto, ok := enums.GetByProtoFallbackDefault(protoenumstatus.StatusEnum(9999))
	require.True(t, ok)
	require.Equal(t, StatusTypeUnknown, byProto.Basic())

	byName, ok := enums.GetByNameFallbackDefault("NOT_EXISTS")
	require.True(t, ok)
	require.Equal(t, StatusTypeUnknown, byName.Basic())

	byBasic, ok := enums.GetByBasicFallbackDefault(StatusType("not_exists"))
	require.True(t, ok)
	require.Equal(t, StatusTypeUnknown, byBasic.Basic())
}

func TestEnums_GetByProto_NotFound(t *testing.T) {
	type StatusType string
	const (
		StatusTypeSuccess StatusType = "success"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	)

	res, ok := enums.GetByProto(protoenumstatus.StatusEnum_FAILURE)
	require.False(t, ok)
	require.Nil(t, res)
}

func TestEnums_ListProtos(t *testing.T) {
	type ResultType string

	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)

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

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)

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

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)
	enums.WithDefault()

	protoEnums := enums.ListNonDefaultProtos()
	t.Log(protoEnums)
	require.Len(t, protoEnums, 3)
	require.Equal(t, protoenumresult.ResultEnum_PASS, protoEnums[0])
	require.Equal(t, protoenumresult.ResultEnum_MISS, protoEnums[1])
	require.Equal(t, protoenumresult.ResultEnum_SKIP, protoEnums[2])

	// NoDefault: no default to exclude, so ListNonDefault returns each element
	enumsNoDefault := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)
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

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)
	enums.WithDefault()

	validBasics := enums.ListNonDefaultBasics()
	t.Log(validBasics)
	require.Len(t, validBasics, 3)
	require.Equal(t, ResultTypePass, validBasics[0])
	require.Equal(t, ResultTypeMiss, validBasics[1])
	require.Equal(t, ResultTypeSkip, validBasics[2])

	// NoDefault: no default to exclude, so ListNonDefault returns each element
	enumsNoDefault := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)
	allBasics := enumsNoDefault.ListNonDefaultBasics()
	t.Log(allBasics)
	require.Len(t, allBasics, 4)
}

func TestEnums_ListEnums(t *testing.T) {
	type ResultType string
	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)

	items := enums.ListEnums()
	require.Len(t, items, 4)
	require.Equal(t, ResultTypeUnknown, items[0].Basic())
	require.Equal(t, ResultTypePass, items[1].Basic())
	require.Equal(t, ResultTypeMiss, items[2].Basic())
	require.Equal(t, ResultTypeSkip, items[3].Basic())
}

func TestEnums_ListNonDefaultEnums(t *testing.T) {
	type ResultType string
	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)
	enums.WithDefault() // UNKNOWN becomes the default

	items := enums.ListNonDefaultEnums()
	require.Len(t, items, 3)
	require.Equal(t, ResultTypePass, items[0].Basic())
	require.Equal(t, ResultTypeMiss, items[1].Basic())
	require.Equal(t, ResultTypeSkip, items[2].Basic())

	// Without a default, nothing is excluded
	enumsNoDefault := protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
	)
	require.Len(t, enumsNoDefault.ListNonDefaultEnums(), 2)
}

func TestNewEnums_DuplicateProto(t *testing.T) {
	type StatusType string
	require.Panics(t, func() {
		protoenum.NewEnums(
			protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusType("a")),
			protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusType("b")),
		)
	})
}

// TestEnums_GetByCode checks lookup via numeric code — true on a hit, false on a miss
//
// 验证按数字代码查找：命中返回 true，未命中返回 false
func TestEnums_GetByCode(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	)

	found, ok := enums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.True(t, ok)
	t.Log(found.Basic())
	require.Equal(t, StatusTypeSuccess, found.Basic())

	found, ok = enums.GetByCode(999)
	require.False(t, ok)
	require.Nil(t, found)
}

// TestEnums_GetByName checks lookup via name — true on a hit, false on a miss
//
// 验证按名称查找：命中返回 true，未命中返回 false
func TestEnums_GetByName(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	)

	found, ok := enums.GetByName(protoenumstatus.StatusEnum_SUCCESS.String())
	require.True(t, ok)
	t.Log(found.Basic())
	require.Equal(t, StatusTypeSuccess, found.Basic())

	found, ok = enums.GetByName("NOT_EXISTS")
	require.False(t, ok)
	require.Nil(t, found)
}

// TestNewEnums_DuplicateBasic checks that two distinct protos sharing one basic value cause a panic
//
// 验证两个不同 proto 共用同一 basic 值时 panic
func TestNewEnums_DuplicateBasic(t *testing.T) {
	type StatusType string
	require.Panics(t, func() {
		protoenum.NewEnums(
			protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusType("same")),
			protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusType("same")),
		)
	})
}

// TestNewEnums_NilElement checks that a nil Enum element causes a panic
//
// 验证传入 nil Enum 元素时 panic
func TestNewEnums_NilElement(t *testing.T) {
	type StatusType string
	require.Panics(t, func() {
		protoenum.NewEnums[protoenumstatus.StatusEnum, StatusType, *protoenum.MetaNone](
			nil,
		)
	})
}
