[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/protoenum/release.yml?branch=main&label=BUILD)](https://github.com/yylego/protoenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/protoenum)](https://pkg.go.dev/github.com/yylego/protoenum)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/protoenum/main.svg)](https://coveralls.io/github/yylego/protoenum?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.23--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/protoenum.svg)](https://github.com/yylego/protoenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/protoenum)](https://goreportcard.com/report/github.com/yylego/protoenum)

# protoenum

`protoenum` provides utilities to manage Protobuf enum metadata in Go. It bridges Protobuf enums with Go native enums (`type StatusType string`) via the `Basic()` method, and offers enum collections with simple code, name, and basic-value lookups.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Core Features

🎯 **Smart Enum Management**: Wrap Protobuf enums with Go native enums and custom metadata
🔗 **Go Native Enum Bridge**: Seamless conversion via `Basic()` method to Go native enum types
⚡ **Multi-Lookup Support**: Fast code, name, and basic-value lookups
🔄 **Type-Safe Operations**: Triple generics maintain strict type checks across protobuf, Go native enums, and metadata
🛡️ **Strict Design**: Single usage pattern prevents misuse with required defaults
🌍 **Production Grade**: Battle-tested enum handling in enterprise applications

## Installation

```bash
go get github.com/yylego/protoenum
```

## Quick Start

### Define Proto Enum

The project includes example proto files:
- `protoenumstatus.proto` - Basic status enum
- `protoenumresult.proto` - Test result enum

### Basic Collection Usage

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
	// Get Go native enum from protobuf enum (returns default when not found)
	item := enums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS))
	zaplog.LOG.Debug("basic", zap.String("msg", string(item.Basic())))

	// Convert between protoenum and native enum (safe with default fallback)
	enum := enums.GetByName("SUCCESS")
	base := protoenumstatus.StatusEnum(enum.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	if base == protoenumstatus.StatusEnum_SUCCESS {
		zaplog.LOG.Debug("done")
	}

	// Get default basic enum value (first item becomes default)
	defaultBasic, err := enums.GetDefaultBasic()
	if err != nil {
		panic(err)
	}
	zaplog.LOG.Debug("default", zap.String("msg", string(defaultBasic)))
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Advanced Lookup Methods

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

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)


## API Reference

### Single Enum Operations

| Method | Description | Returns |
|--------|-------------|--------|
| `NewEnum(protoEnum, basicEnum)` | Create enum instance without metadata | `*Enum[P, B, *MetaNone]` |
| `NewEnumWithDesc(protoEnum, basicEnum, desc)` | Create enum instance with description | `*Enum[P, B, *MetaDesc]` |
| `NewEnumWithMeta(protoEnum, basicEnum, meta)` | Create enum instance with custom metadata | `*Enum[P, B, M]` |
| `enum.Proto()` | Get underlying protobuf enum | `P` |
| `enum.Code()` | Get numeric code | `int32` |
| `enum.Name()` | Get enum name | `string` |
| `enum.Basic()` | Get Go native enum value | `B` |
| `enum.Meta()` | Get custom metadata | `M` |

### Collection Creation

| Method | Description | Returns |
|--------|-------------|--------|
| `NewEnums(items...)` | Create collection with validation (first item becomes default) | `(*Enums[P, B, M], error)` |

### Existence Check (Lookup)

| Method | Description | Returns |
|--------|-------------|--------|
| `enums.LookupByProto(proto)` | Lookup by protobuf enum, check existence | `(*Enum[P, B, M], bool)` |
| `enums.LookupByCode(code)` | Lookup by code, check existence | `(*Enum[P, B, M], bool)` |
| `enums.LookupByName(name)` | Lookup by name, check existence | `(*Enum[P, B, M], bool)` |
| `enums.LookupByBasic(basic)` | Lookup by Go native enum, check existence | `(*Enum[P, B, M], bool)` |

### Safe Access (Get)

| Method | Description | Returns |
|--------|-------------|--------|
| `enums.GetByProto(proto)` | Get by protobuf enum (returns default if not found, nil if no default) | `*Enum[P, B, M]` |
| `enums.GetByCode(code)` | Get by code (returns default if not found, nil if no default) | `*Enum[P, B, M]` |
| `enums.GetByName(name)` | Get by name (returns default if not found, nil if no default) | `*Enum[P, B, M]` |
| `enums.GetByBasic(basic)` | Get by Go native enum (returns default if not found, nil if no default) | `*Enum[P, B, M]` |

### Enumeration (List)

| Method | Description | Returns |
|--------|-------------|--------|
| `enums.ListProtos()` | Returns a slice of each protoEnum value | `[]P` |
| `enums.ListBasics()` | Returns a slice of each basicEnum value | `[]B` |
| `enums.ListValidProtos()` | Returns protoEnum values excluding default | `[]P` |
| `enums.ListValidBasics()` | Returns basicEnum values excluding default | `[]B` |

### Default Value Management

| Method | Description | Returns |
|--------|-------------|--------|
| `enums.GetDefault()` | Get current default value (nil if unset) | `*Enum[P, B, M]` |
| `enums.GetDefaultProto()` | Get default protoEnum value | `(P, error)` |
| `enums.GetDefaultBasic()` | Get default basicEnum value | `(B, error)` |
| `enums.SetDefault(enum)` | Set default (returns error if existing default) | `error` |
| `enums.UnsetDefault()` | Remove default (returns error if unset) | `error` |
| `enums.WithDefault(enum)` | Chain: set default by enum instance | `*Enums[P, B, M]` |
| `enums.WithDefaultCode(code)` | Chain: set default by code (panics if not found) | `*Enums[P, B, M]` |
| `enums.WithDefaultName(name)` | Chain: set default by name (panics if not found) | `*Enums[P, B, M]` |
| `enums.WithUnsetDefault()` | Chain: remove default value | `*Enums[P, B, M]` |

## Examples

### Working with Single Enums

**Creating enhanced enum instance:**
```go
type StatusType string
const StatusTypeSuccess StatusType = "success"

statusEnum := protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "操作成功")
fmt.Printf("Code: %d, Name: %s, Basic: %s, Description: %s\n",
    statusEnum.Code(), statusEnum.Name(), statusEnum.Basic(), statusEnum.Meta().Desc())
```

**Accessing underlying protobuf enum:**
```go
originalEnum := statusEnum.Proto()
if originalEnum == protoenumstatus.StatusEnum_SUCCESS {
    fmt.Println("Success status detected")
}
```

### Collection Operations

**Building enum collections:**
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

**Multiple lookup methods:**
```go
// Using numeric code - returns default if not found, nil if no default
enum := statusEnums.GetByCode(1)
fmt.Printf("Found: %s\n", enum.Meta().Desc())

// Using enum name
enum = statusEnums.GetByName("SUCCESS")
fmt.Printf("Status: %s\n", enum.Meta().Desc())

// Using Go native enum value - type-safe lookup
enum = statusEnums.GetByBasic(StatusTypeSuccess)
fmt.Printf("Basic: %s\n", enum.Basic())

// Existence check - returns (enum, bool)
if found, ok := statusEnums.LookupByCode(1); ok {
    fmt.Printf("Found: %s\n", found.Meta().Desc())
}
```

**Listing values:**
```go
// Get a slice of each registered proto enum
protoEnums := statusEnums.ListProtos()
// > [UNKNOWN, SUCCESS, FAILURE]

// Get a slice of each registered basic Go enum
basicEnums := statusEnums.ListBasics()
// > ["unknown", "success", "failure"]

// Get valid values (excluding default)
validProtos := statusEnums.ListValidProtos()
// > [SUCCESS, FAILURE] (UNKNOWN is default, excluded)

validBasics := statusEnums.ListValidBasics()
// > ["success", "failure"]
```

### Advanced Usage

**Go native enum bridge via Basic():**
```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
)

// Bridge protobuf enum to Go native enum
enum := enums.GetByCode(1)
basicValue := enum.Basic()  // Returns StatusType("success")

// Use in business logic with Go native enum
switch basicValue {
case StatusTypeSuccess:
    fmt.Println("Operation succeeded")
case StatusTypeUnknown:
    fmt.Println("Unknown status")
}

// Lookup using Go native enum value
found := enums.GetByBasic(StatusTypeSuccess)
fmt.Printf("Code: %d, Name: %s\n", found.Code(), found.Name())
```

**Type conversion patterns:**
```go
// Convert from enum instance to native protobuf enum
// Always returns valid enum (with default fallback)
statusEnum := enums.GetByName("SUCCESS")
native := protoenumstatus.StatusEnum(statusEnum.Code())
// Use native enum in protobuf operations with safe access
```

**Lookup patterns:**
```go
// GetByXxx returns default on unknown values, nil if no default
result := enums.GetByCode(999)  // Returns default (UNKNOWN)
if result != nil {
    fmt.Printf("Fallback: %s\n", result.Name())
}

// LookupByXxx returns (enum, bool) - use with rese to panic if not found
found, ok := enums.LookupByCode(1)
// Or use rese.P1(enums.LookupByCode(1)) to panic if not found
```

### Default Values and Chain Configuration

**Automatic default value (first item):**
```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
)

enums, err := protoenum.NewEnums(
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, "未知"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "成功"),
)
// First item (UNKNOWN) becomes default on creation
defaultEnum := enums.GetDefault()  // Returns nil if no default
```

**Default management:**
```go
// Lookup returns default if not found, nil if no default
notFound := enums.GetByCode(999)  // Returns UNKNOWN (default)
if notFound != nil {
    fmt.Printf("Fallback: %s\n", notFound.Meta().Desc())
}

// Change default: unset first, then set new one
enums.UnsetDefault()
successEnum, ok := enums.LookupByCode(1)
if ok {
    enums.SetDefault(successEnum)
}

// Once UnsetDefault is invoked, GetByXxx returns nil if not found
```

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/protoenum.svg?variant=adaptive)](https://starchart.cc/yylego/protoenum)
