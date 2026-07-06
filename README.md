[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/protoenum/release.yml?branch=main&label=BUILD)](https://github.com/yylego/protoenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/protoenum)](https://pkg.go.dev/github.com/yylego/protoenum)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/protoenum/main.svg)](https://coveralls.io/github/yylego/protoenum?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
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
🛡️ **Strict Design**: panics fast on a nil / duplicate; the default is an explicit set-once opt-in
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
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
).WithDefault()

func main() {
	if item, ok := enums.GetByCodeFallbackDefault(int32(protoenumstatus.StatusEnum_SUCCESS)); ok {
		zaplog.LOG.Debug("basic", zap.String("msg", string(item.Basic())))
	}

	if enum, ok := enums.GetByNameFallbackDefault("SUCCESS"); ok {
		base := protoenumstatus.StatusEnum(enum.Code())
		zaplog.LOG.Debug("base", zap.String("msg", base.String()))

		if base == protoenumstatus.StatusEnum_SUCCESS {
			zaplog.LOG.Debug("done")
		}
	}

	if defaultEnum, ok := enums.GetDefault(); ok {
		zaplog.LOG.Debug("default", zap.String("msg", string(defaultEnum.Basic())))
	}
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
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

## API Reference

### Single Enum Operations

| Method                                        | Description                               | Returns                  |
| --------------------------------------------- | ----------------------------------------- | ------------------------ |
| `NewEnum(protoEnum, basicEnum)`               | Create enum instance without metadata     | `*Enum[P, B, *MetaNone]` |
| `NewEnumWithDesc(protoEnum, basicEnum, desc)` | Create enum instance with description     | `*Enum[P, B, *MetaDesc]` |
| `NewEnumWithMeta(protoEnum, basicEnum, meta)` | Create enum instance with custom metadata | `*Enum[P, B, M]`         |
| `enum.Proto()`                                | Get underlying protobuf enum              | `P`                      |
| `enum.Code()`                                 | Get numeric code                          | `int32`                  |
| `enum.Name()`                                 | Get enum name                             | `string`                 |
| `enum.Basic()`                                | Get Go native enum value                  | `B`                      |
| `enum.Meta()`                                 | Get custom metadata                       | `M`                      |

### Collection Creation

| Method               | Description                                                                                              | Returns           |
| -------------------- | -------------------------------------------------------------------------------------------------------- | ----------------- |
| `NewEnums(items...)` | Create collection; panics on a nil / duplicate (no default — chain WithDefault to fix the first element) | `*Enums[P, B, M]` |

### Existence Check (Get)

| Method                    | Description                             | Returns                  |
| ------------------------- | --------------------------------------- | ------------------------ |
| `enums.GetByProto(proto)` | Get via protobuf enum, check existence  | `(*Enum[P, B, M], bool)` |
| `enums.GetByCode(code)`   | Get via code, check existence           | `(*Enum[P, B, M], bool)` |
| `enums.GetByName(name)`   | Get via name, check existence           | `(*Enum[P, B, M], bool)` |
| `enums.GetByBasic(basic)` | Get via Go native enum, check existence | `(*Enum[P, B, M], bool)` |

### Fallback Access

| Method                                   | Description                                                                               | Returns                  |
| ---------------------------------------- | ----------------------------------------------------------------------------------------- | ------------------------ |
| `enums.GetByProtoFallbackDefault(proto)` | Get via protobuf enum (returns (default, true) if not found, (nil, false) if no default)  | `(*Enum[P, B, M], bool)` |
| `enums.GetByCodeFallbackDefault(code)`   | Get via code (returns (default, true) if not found, (nil, false) if no default)           | `(*Enum[P, B, M], bool)` |
| `enums.GetByNameFallbackDefault(name)`   | Get via name (returns (default, true) if not found, (nil, false) if no default)           | `(*Enum[P, B, M], bool)` |
| `enums.GetByBasicFallbackDefault(basic)` | Get via Go native enum (returns (default, true) if not found, (nil, false) if no default) | `(*Enum[P, B, M], bool)` |

### Enumeration (List)

| Method                         | Description                                | Returns            |
| ------------------------------ | ------------------------------------------ | ------------------ |
| `enums.ListProtos()`           | Returns a slice of each protoEnum value    | `[]P`              |
| `enums.ListBasics()`           | Returns a slice of each basicEnum value    | `[]B`              |
| `enums.ListNonDefaultProtos()` | Returns protoEnum values excluding default | `[]P`              |
| `enums.ListNonDefaultBasics()` | Returns basicEnum values excluding default | `[]B`              |
| `enums.ListEnums()`            | Returns a slice of each Enum instance      | `[]*Enum[P, B, M]` |
| `enums.ListNonDefaultEnums()`  | Returns Enum instances excluding default   | `[]*Enum[P, B, M]` |

### Default Value Management

The default is an explicit opt-in: `NewEnums` sets none. Chain `WithDefault` once to fix the first element; the default stays put and has no unset.

| Method                | Description                                              | Returns                  |
| --------------------- | -------------------------------------------------------- | ------------------------ |
| `enums.GetDefault()`  | Get the default, returns (default, true) or (nil, false) | `(*Enum[P, B, M], bool)` |
| `enums.WithDefault()` | Chain: fix the first element as the default, just once   | `*Enums[P, B, M]`        |

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

statusEnums := protoenum.NewEnums(
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, "未知状态"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "成功"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure, "失败"),
)
// Fix the first element (UNKNOWN) as the fallback default: GetByXxxFallbackDefault returns it on a miss, ListNonDefaultXxx skips it
statusEnums.WithDefault()
```

**Multiple lookup methods:**

```go
// Using numeric code - returns (default, true) if not found, (nil, false) if no default
enum := statusEnums.GetByCodeFallbackDefault(1)
fmt.Printf("Found: %s\n", enum.Meta().Desc())

// Using enum name
enum = statusEnums.GetByNameFallbackDefault("SUCCESS")
fmt.Printf("Status: %s\n", enum.Meta().Desc())

// Using Go native enum value - type-safe lookup
enum = statusEnums.GetByBasicFallbackDefault(StatusTypeSuccess)
fmt.Printf("Basic: %s\n", enum.Basic())

// Existence check - returns (enum, bool)
if found, ok := statusEnums.GetByCode(1); ok {
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

// Get non-default values (excluding the default)
validProtos := statusEnums.ListNonDefaultProtos()
// > [SUCCESS, FAILURE] (UNKNOWN is default, excluded)

validBasics := statusEnums.ListNonDefaultBasics()
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
if enum, ok := enums.GetByCodeFallbackDefault(1); ok {
    basicValue := enum.Basic()  // StatusType("success")

    // Use in business logic with Go native enum
    switch basicValue {
    case StatusTypeSuccess:
        fmt.Println("Operation succeeded")
    case StatusTypeUnknown:
        fmt.Println("Unknown status")
    }
}

// Look up using Go native enum value
if found, ok := enums.GetByBasicFallbackDefault(StatusTypeSuccess); ok {
    fmt.Printf("Code: %d, Name: %s\n", found.Code(), found.Name())
}
```

**Type conversion patterns:**

```go
// Convert from enum instance to native protobuf enum
if statusEnum, ok := enums.GetByNameFallbackDefault("SUCCESS"); ok {
    native := protoenumstatus.StatusEnum(statusEnum.Code())
    fmt.Println(native) // use native enum in protobuf operations
}
```

**Lookup patterns:**

```go
// GetByXxxFallbackDefault returns (default, true) on unknown values; (nil, false) if no default
if result, ok := enums.GetByCodeFallbackDefault(999); ok {  // uses the default (UNKNOWN)
    fmt.Printf("Fallback: %s\n", result.Name())
}

// GetByXxx returns (enum, bool)
if found, ok := enums.GetByCode(1); ok {
    fmt.Printf("Found: %s\n", found.Name())
}
```

### Default Values and Chain Configuration

**Set the default up front, once:**

```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
)

// NewEnums sets no default; chain WithDefault once to fix the first element (UNKNOWN) as the fallback
enums := protoenum.NewEnums(
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, "未知"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "成功"),
)
enums.WithDefault()

// GetDefault returns (enum, bool)
if defaultEnum, ok := enums.GetDefault(); ok {
    fmt.Println(defaultEnum.Basic())
}
```

**Fallback path:**

```go
// With a default set, GetByXxxFallbackDefault returns it on a miss
if notFound, ok := enums.GetByCodeFallbackDefault(999); ok {  // uses UNKNOWN (the default)
    fmt.Printf("Fallback: %s\n", notFound.Meta().Desc())
}

// The default is fixed once and has no unset. A second WithDefault panics.
// Without a default, GetByXxxFallbackDefault returns nil on a miss (use GetByXxx instead).
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
