package main

import (
	"github.com/yylego/protoenum"
	"github.com/yylego/protoenum/protos/protoenumresult"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

// ResultType represents a Go native enum of result
// ResultType 代表结果的 Go 原生枚举
type ResultType string

const (
	ResultTypeUnknown ResultType = "unknown"
	ResultTypePass    ResultType = "pass"
	ResultTypeMiss    ResultType = "miss"
	ResultTypeSkip    ResultType = "skip"
)

var enums = protoenum.NewEnums(
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown, "其它"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_PASS, ResultTypePass, "通过"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_MISS, ResultTypeMiss, "出错"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_SKIP, ResultTypeSkip, "跳过"),
).WithDefault()

func main() {
	if skip, ok := enums.GetByCodeFallbackDefault(int32(protoenumresult.ResultEnum_SKIP)); ok {
		zaplog.LOG.Debug("basic", zap.String("msg", string(skip.Basic())))
		zaplog.LOG.Debug("desc", zap.String("msg", skip.Meta().Desc()))
	}

	if pass, ok := enums.GetByBasicFallbackDefault(ResultTypePass); ok {
		base := protoenumresult.ResultEnum(pass.Code())
		zaplog.LOG.Debug("base", zap.String("msg", base.String()))

		if base == protoenumresult.ResultEnum_PASS {
			zaplog.LOG.Debug("pass")
		}
	}

	if miss, ok := enums.GetByNameFallbackDefault("MISS"); ok {
		zaplog.LOG.Debug("basic", zap.String("msg", string(miss.Basic())))
		zaplog.LOG.Debug("desc", zap.String("msg", miss.Meta().Desc()))
	}

	basics := enums.ListBasics()
	for _, basic := range basics {
		zaplog.LOG.Debug("list", zap.String("basic", string(basic)))
	}
}
