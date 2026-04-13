[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/protoenum/release.yml?branch=main&label=BUILD)](https://github.com/yylego/protoenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/protoenum)](https://pkg.go.dev/github.com/yylego/protoenum)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/protoenum/main.svg)](https://coveralls.io/github/yylego/protoenum?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.23--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/protoenum.svg)](https://github.com/yylego/protoenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/protoenum)](https://goreportcard.com/report/github.com/yylego/protoenum)

# protoenum

`protoenum` 是一个 Go 语言包，提供管理 Protobuf 枚举元数据的工具。它通过 `Basic()` 方法桥接 Protobuf 枚举和 Go 原生枚举（`type StatusType string`），并提供枚举集合支持简单的代码、名称和 Basic 值查找。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🎯 **智能枚举管理**：将 Protobuf 枚举与 Go 原生枚举和自定义元数据包装
🔗 **Go 原生枚举桥接**：通过 `Basic()` 方法无缝转换到 Go 原生枚举类型
⚡ **多方式查找**：支持代码、名称和 Basic 值快速查找
🔄 **类型安全操作**：三泛型保持 protobuf、Go 原生枚举和元数据的类型安全
🛡️ **严格设计**：单一使用模式防止误用，强制要求默认值
🌍 **生产级别**：经过实战检验的企业级枚举处理方案

## 安装

```bash
go get github.com/yylego/protoenum
```

## 快速开始

### 定义 Proto 枚举

项目包含示例 proto 文件：
- `protoenumstatus.proto` - 基础状态枚举
- `protoenumresult.proto` - 测试结果枚举

### 基础集合使用

```go
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

// Build status enum collection
// 构建状态枚举集合
var enums = rese.P1(protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
))

func main() {
	// 从 protobuf 枚举获取 Go 原生枚举（找不到时返回默认值）
	item := enums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS))
	zaplog.LOG.Debug("basic", zap.String("msg", string(item.Basic())))

	// 在 protoenum 和原生枚举之间转换（安全且有默认值回退）
	enum := enums.GetByName("SUCCESS")
	base := protoenumstatus.StatusEnum(enum.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	if base == protoenumstatus.StatusEnum_SUCCESS {
		zaplog.LOG.Debug("done")
	}

	// 获取默认 basic 枚举值（第一个元素成为默认值）
	defaultBasic, err := enums.GetDefaultBasic()
	if err != nil {
		panic(err)
	}
	zaplog.LOG.Debug("default", zap.String("msg", string(defaultBasic)))
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 高级查找方法

```go
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

// Build enum collection with description
// 构建带描述的枚举集合
var enums = rese.P1(protoenum.NewEnums(
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown, "其它"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_PASS, ResultTypePass, "通过"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_MISS, ResultTypeMiss, "出错"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_SKIP, ResultTypeSkip, "跳过"),
))

func main() {
	// Lookup using enum code (returns default when not found)
	// 按枚举代码查找（找不到时返回默认值）
	skip := enums.GetByCode(int32(protoenumresult.ResultEnum_SKIP))
	zaplog.LOG.Debug("basic", zap.String("msg", string(skip.Basic())))
	zaplog.LOG.Debug("desc", zap.String("msg", skip.Meta().Desc()))

	// Lookup using Go native enum value (type-safe)
	// 按 Go 原生枚举值查找（类型安全查找）
	pass := enums.GetByBasic(ResultTypePass)
	base := protoenumresult.ResultEnum(pass.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	// Business logic with native enum
	// 使用原生枚举的业务逻辑
	if base == protoenumresult.ResultEnum_PASS {
		zaplog.LOG.Debug("pass")
	}

	// Lookup using enum name (safe with default fallback)
	// 按枚举名称查找（安全且有默认值回退）
	miss := enums.GetByName("MISS")
	zaplog.LOG.Debug("basic", zap.String("msg", string(miss.Basic())))
	zaplog.LOG.Debug("desc", zap.String("msg", miss.Meta().Desc()))

	// List each basic enum value in defined sequence
	// 按定义次序列出各 basic 枚举值
	basics := enums.ListBasics()
	for _, basic := range basics {
		zaplog.LOG.Debug("list", zap.String("basic", string(basic)))
	}
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)


## API 参考

### 单个枚举操作

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `NewEnum(protoEnum, basicEnum)` | 创建枚举实例（无元数据） | `*Enum[P, B, *MetaNone]` |
| `NewEnumWithDesc(protoEnum, basicEnum, desc)` | 创建枚举实例（带描述） | `*Enum[P, B, *MetaDesc]` |
| `NewEnumWithMeta(protoEnum, basicEnum, meta)` | 创建枚举实例（带自定义元数据） | `*Enum[P, B, M]` |
| `enum.Proto()` | 获取底层 protobuf 枚举 | `P` |
| `enum.Code()` | 获取数值代码 | `int32` |
| `enum.Name()` | 获取枚举名称 | `string` |
| `enum.Basic()` | 获取 Go 原生枚举值 | `B` |
| `enum.Meta()` | 获取自定义元数据 | `M` |

### 创建集合

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `NewEnums(items...)` | 创建集合并验证（第一项成为默认值） | `(*Enums[P, B, M], error)` |

### 存在性检查 (Lookup)

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `enums.LookupByProto(proto)` | 按 protobuf 枚举查找，检查是否存在 | `(*Enum[P, B, M], bool)` |
| `enums.LookupByCode(code)` | 按代码查找，检查是否存在 | `(*Enum[P, B, M], bool)` |
| `enums.LookupByName(name)` | 按名称查找，检查是否存在 | `(*Enum[P, B, M], bool)` |
| `enums.LookupByBasic(basic)` | 按 Go 原生枚举查找，检查是否存在 | `(*Enum[P, B, M], bool)` |

### 安全访问 (Get)

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `enums.GetByProto(proto)` | 按 protobuf 枚举获取（找不到返回默认值，无默认值返回 nil） | `*Enum[P, B, M]` |
| `enums.GetByCode(code)` | 按代码获取（找不到返回默认值，无默认值返回 nil） | `*Enum[P, B, M]` |
| `enums.GetByName(name)` | 按名称获取（找不到返回默认值，无默认值返回 nil） | `*Enum[P, B, M]` |
| `enums.GetByBasic(basic)` | 按 Go 原生枚举获取（找不到返回默认值，无默认值返回 nil） | `*Enum[P, B, M]` |

### 枚举列表 (List)

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `enums.ListProtos()` | 返回各 protoEnum 值的切片 | `[]P` |
| `enums.ListBasics()` | 返回各 basicEnum 值的切片 | `[]B` |
| `enums.ListValidProtos()` | 返回排除默认值的 protoEnum 切片 | `[]P` |
| `enums.ListValidBasics()` | 返回排除默认值的 basicEnum 切片 | `[]B` |

### 默认值管理

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `enums.GetDefault()` | 获取当前默认值（未设置返回 nil） | `*Enum[P, B, M]` |
| `enums.GetDefaultProto()` | 获取默认 protoEnum 值 | `(P, error)` |
| `enums.GetDefaultBasic()` | 获取默认 basicEnum 值 | `(B, error)` |
| `enums.SetDefault(enum)` | 设置默认值（已有默认值返回 error） | `error` |
| `enums.UnsetDefault()` | 移除默认值（未设置返回 error） | `error` |
| `enums.WithDefault(enum)` | 链式：通过枚举实例设置默认值 | `*Enums[P, B, M]` |
| `enums.WithDefaultCode(code)` | 链式：通过代码设置默认值（找不到则 panic） | `*Enums[P, B, M]` |
| `enums.WithDefaultName(name)` | 链式：通过名称设置默认值（找不到则 panic） | `*Enums[P, B, M]` |
| `enums.WithUnsetDefault()` | 链式：移除默认值 | `*Enums[P, B, M]` |

## 使用示例

### 单个枚举操作

**创建增强枚举包装器：**
```go
type StatusType string
const StatusTypeSuccess StatusType = "success"

statusEnum := protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "操作成功")
fmt.Printf("代码: %d, 名称: %s, Basic: %s, 描述: %s\n",
    statusEnum.Code(), statusEnum.Name(), statusEnum.Basic(), statusEnum.Meta().Desc())
```

**访问底层 protobuf 枚举：**
```go
originalEnum := statusEnum.Proto()
if originalEnum == protoenumstatus.StatusEnum_SUCCESS {
    fmt.Println("检测到成功状态")
}
```

### 集合操作

**构建枚举集合：**
```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
    StatusTypeFailure StatusType = "failure"
)

statusEnums, err := protoenum.NewEnums(
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, "未知状态"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "成功"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure, "失败"),
)
```

**多种查找方式：**
```go
// 按数字代码查找（找不到返回默认值，无默认值返回 nil）
enum := statusEnums.GetByCode(1)
fmt.Printf("找到: %s\n", enum.Meta().Desc())

// 按枚举名称查找
enum = statusEnums.GetByName("SUCCESS")
fmt.Printf("状态: %s\n", enum.Meta().Desc())

// 按 Go 原生枚举值查找 - 类型安全查找
enum = statusEnums.GetByBasic(StatusTypeSuccess)
fmt.Printf("Basic: %s\n", enum.Basic())

// 存在性检查 - 返回 (enum, bool)
if found, ok := statusEnums.LookupByCode(1); ok {
    fmt.Printf("找到: %s\n", found.Meta().Desc())
}
```

**列出枚举值:**
```go
// 获取各已注册 proto 枚举的切片
protoEnums := statusEnums.ListProtos()
// > [UNKNOWN, SUCCESS, FAILURE]

// 获取各已注册 basic Go 原生枚举的切片
basicEnums := statusEnums.ListBasics()
// > ["unknown", "success", "failure"]

// 获取有效值（排除默认值）
validProtos := statusEnums.ListValidProtos()
// > [SUCCESS, FAILURE]（UNKNOWN 是默认值，被排除）

validBasics := statusEnums.ListValidBasics()
// > ["success", "failure"]
```

### 高级用法

**通过 Basic() 桥接 Go 原生枚举：**
```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
)

// 桥接 protobuf 枚举到 Go 原生枚举
enum := enums.GetByCode(1)
basicValue := enum.Basic()  // 返回 StatusType("success")

// 在业务逻辑中使用 Go 原生枚举
switch basicValue {
case StatusTypeSuccess:
    fmt.Println("操作成功")
case StatusTypeUnknown:
    fmt.Println("未知状态")
}

// 通过 Go 原生枚举值查找
found := enums.GetByBasic(StatusTypeSuccess)
fmt.Printf("代码: %d, 名称: %s\n", found.Code(), found.Name())
```

**类型转换模式：**
```go
// 从枚举包装器转换为原生 protobuf 枚举
// 始终返回有效枚举（带默认值回退）
statusEnum := enums.GetByName("SUCCESS")
native := protoenumstatus.StatusEnum(statusEnum.Code())
// 在 protobuf 操作中安全使用原生枚举
```

**查找模式：**
```go
// GetByXxx 对未知值返回默认值，无默认值返回 nil
result := enums.GetByCode(999)  // 返回默认值（UNKNOWN）
if result != nil {
    fmt.Printf("回退: %s\n", result.Name())
}

// LookupByXxx 返回 (enum, bool)，配合 rese 使用可在找不到时 panic
found, ok := enums.LookupByCode(1)
// 或使用 rese.P1(enums.LookupByCode(1)) 在找不到时 panic
```

### 默认值和链式配置

**自动默认值（第一项）：**
```go
enums, err := protoenum.NewEnums(
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, "未知"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "成功"),
)
// 第一项（UNKNOWN）自动成为默认值
defaultEnum, err := enums.GetDefault()  // 未设置时返回 error
```

**默认值管理：**
```go
// 查找失败返回默认值，无默认值返回 nil
notFound := enums.GetByCode(999)  // 返回 UNKNOWN（默认值）
if notFound != nil {
    fmt.Printf("回退值: %s\n", notFound.Meta().Desc())
}

// 更改默认值：先取消设置，再设置新值
enums.UnsetDefault()
successEnum, ok := enums.LookupByCode(1)
if ok {
    enums.SetDefault(successEnum)
}

// UnsetDefault 后，GetByXxx 找不到时返回 nil
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yylego/protoenum.svg?variant=adaptive)](https://starchart.cc/yylego/protoenum)
