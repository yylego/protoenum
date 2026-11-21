package main

import (
	"github.com/yylego/protoenum"
	"github.com/yylego/protoenum/protos/protoenumstatus"
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

// MetaI18n represents a custom metadata type with English/Chinese descriptions
// MetaI18n 代表带有中英文描述的自定义元数据类型
type MetaI18n struct {
	zh string // Chinese description // 中文描述
	en string // English description // 英文描述
}

func (c *MetaI18n) Chinese() string { return c.zh }
func (c *MetaI18n) English() string { return c.en }

// Build enum collection with custom English/Chinese metadata
// 构建带有自定义中英文元数据的枚举集合
var enums = protoenum.NewEnums(
	protoenum.NewEnumWithMeta(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, &MetaI18n{zh: "未知", en: "Unknown"}),
	protoenum.NewEnumWithMeta(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, &MetaI18n{zh: "成功", en: "Success"}),
	protoenum.NewEnumWithMeta(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure, &MetaI18n{zh: "失败", en: "Failure"}),
).WithDefault()

func main() {
	if success, ok := enums.GetByBasicFallbackDefault(StatusTypeSuccess); ok {
		zaplog.LOG.Debug("basic", zap.String("msg", string(success.Basic())))
		zaplog.LOG.Debug("zh", zap.String("msg", success.Meta().Chinese()))
		zaplog.LOG.Debug("en", zap.String("msg", success.Meta().English()))

		if success.Basic() == StatusTypeSuccess {
			zaplog.LOG.Debug("done")
		}
	}

	if failure, ok := enums.GetByCodeFallbackDefault(int32(protoenumstatus.StatusEnum_FAILURE)); ok {
		zaplog.LOG.Debug("basic", zap.String("msg", string(failure.Basic())))
		zaplog.LOG.Debug("zh", zap.String("msg", failure.Meta().Chinese()))
		zaplog.LOG.Debug("en", zap.String("msg", failure.Meta().English()))
	}

	validBasics := enums.ListNonDefaultBasics()
	for _, basic := range validBasics {
		zaplog.LOG.Debug("valid", zap.String("basic", string(basic)))
	}
}
