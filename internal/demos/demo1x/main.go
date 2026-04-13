package main

import (
	"github.com/yylego/protoenum"
	"github.com/yylego/protoenum/protos/protoenumstatus"
	"github.com/yylego/rese"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

// StatusType represents a Go native enum of status
// StatusType 代表状态的 Go 原生枚举
type StatusType string

const (
	StatusTypeUnknown StatusType = "unknown"
	StatusTypeSuccess StatusType = "success"
	StatusTypeFailure StatusType = "failure"
)

// Build status enum collection
// 构建状态枚举集合
var enums = rese.P1(protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
))

func main() {
	item := enums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS))
	zaplog.LOG.Debug("basic", zap.String("msg", string(item.Basic())))

	enum := enums.GetByName("SUCCESS")
	base := protoenumstatus.StatusEnum(enum.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	if base == protoenumstatus.StatusEnum_SUCCESS {
		zaplog.LOG.Debug("done")
	}

	defaultBasic, err := enums.GetDefaultBasic()
	if err != nil {
		panic(err)
	}
	zaplog.LOG.Debug("default", zap.String("msg", string(defaultBasic)))
}
