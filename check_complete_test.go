package protoenum_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/protoenum"
	"github.com/yylego/protoenum/protos/protoenumstatus"
)

func TestEnums_CheckComplete(t *testing.T) {
	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "unknown"),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "success"),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "failure"),
	).CheckComplete()
	require.Len(t, enums.ListEnums(), 3)
}

func TestEnums_CheckComplete_Missing(t *testing.T) {
	// FAILURE is declared in the proto, yet not mapped here, so the panic names it
	// FAILURE 在 proto 里有声明，这里却没登记，panic 里点名它
	require.PanicsWithError(t, "ENUM protoenumstatus.StatusEnum HAS VALUES WITHOUT AN ELEMENT: FAILURE", func() {
		protoenum.NewEnums(
			protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "unknown"),
			protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "success"),
		).CheckComplete()
	})
	// Each missing value is named, in the sequence the proto declares them
	// 缺的值逐个点名，次序与 proto 里的声明一致
	require.PanicsWithError(t, "ENUM protoenumstatus.StatusEnum HAS VALUES WITHOUT AN ELEMENT: UNKNOWN, FAILURE", func() {
		protoenum.NewEnums(
			protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "success"),
		).CheckComplete()
	})
}
