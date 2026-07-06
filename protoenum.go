// Package protoenum: manages Protocol Buffer enum metadata
// Wraps a protobuf enum with a Go native enum and custom metadata
// Uses three generics: protoEnum, basic, and metaType
//
// protoenum: Protocol Buffer 枚举元数据管理
// 用 Go 原生枚举和自定义元数据包装 protobuf 枚举
// 使用三个泛型：protoEnum、basic、metaType
package protoenum

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ProtoEnum constrains a generic type param to a protobuf enum
// A protobuf enum has String() + Number() and is comparable, so it works as a map index
//
// ProtoEnum 约束泛型参数为 protobuf 枚举
// protobuf 枚举有 String() + Number() 且可比较，因此能当 map 键
type ProtoEnum interface {
	// String gives the enum name from the protobuf schema
	// String 返回 protobuf 模式里的枚举名称
	String() string
	// Number gives the numeric wire value
	// Number 返回数字线格式值
	Number() protoreflect.EnumNumber

	// comparable lets the enum act as a map index
	// comparable 让枚举能当 map 键
	comparable
}

// Enum wraps a protobuf enum with a Go native enum and custom metadata
// Basic() gives the Go native enum; Meta() gives the metadata
// Uses three generics to keep type checks across the three parts
//
// Enum 用 Go 原生枚举和自定义元数据包装 protobuf 枚举
// Basic() 取 Go 原生枚举值，Meta() 取元数据
// 使用三个泛型在三部分之间保持类型检查
type Enum[protoEnum ProtoEnum, basicEnum comparable, metaType any] struct {
	proto protoEnum // Source Protocol Buffer enum value // 源 Protocol Buffer 枚举值
	basic basicEnum // Go native enum value (e.g. type StatusType string) // Go 原生枚举值（如 type StatusType string）
	meta  metaType  // Custom metadata of the enum // 枚举的自定义元数据
}

// NewEnum creates a new Enum instance binding protobuf enum with Go native enum
// Use this when you just need enum mapping without description
// The basic param accepts Go native enum type (e.g. type StatusType string)
// Returns a reference to the created Enum instance, supporting chained invocation
//
// 创建新的 Enum 实例，绑定 protobuf 枚举与 Go 原生枚举
// 当只需要枚举映射而不需要描述时使用此函数
// basic 参数接受 Go 原生枚举类型（如 type StatusType string）
// 返回创建的 Enum 实例指针以便链式调用
func NewEnum[protoEnum ProtoEnum, basicEnum comparable](proto protoEnum, basic basicEnum) *Enum[protoEnum, basicEnum, *MetaNone] {
	return &Enum[protoEnum, basicEnum, *MetaNone]{
		proto: proto,
		basic: basic,
		meta:  &MetaNone{},
	}
}

// NewEnumWithDesc creates a new Enum instance with protobuf enum, Go native enum, and description
// Use this when you need both enum mapping and human-readable description
// The basic param accepts Go native enum type (e.g. type StatusType string)
// The description param provides custom description used in docs and UI rendering
//
// 创建带有 protobuf 枚举、Go 原生枚举和描述的新 Enum 实例
// 当需要枚举映射和人类可读描述时使用此函数
// basic 参数接受 Go 原生枚举类型（如 type StatusType string）
// description 参数提供用于文档和显示的自定义描述
func NewEnumWithDesc[protoEnum ProtoEnum, basicEnum comparable](proto protoEnum, basic basicEnum, description string) *Enum[protoEnum, basicEnum, *MetaDesc] {
	return &Enum[protoEnum, basicEnum, *MetaDesc]{
		proto: proto,
		basic: basic,
		meta:  &MetaDesc{description: description},
	}
}

// NewEnumWithMeta creates a new Enum instance with protobuf enum, Go native enum, and custom metadata
// Use this when you need customized metadata types beyond simple string description
// The basic param accepts Go native enum type (e.g. type StatusType string)
// The meta param accepts custom metadata types (e.g. i18n descriptions with multiple languages)
//
// 创建带有 protobuf 枚举、Go 原生枚举和自定义元数据的新 Enum 实例
// 当需要超越简单字符串描述的灵活元数据类型时使用此函数
// basic 参数接受 Go 原生枚举类型（如 type StatusType string）
// meta 参数接受任意自定义元数据类型（如双语描述）
func NewEnumWithMeta[protoEnum ProtoEnum, basicEnum comparable, metaType any](proto protoEnum, basic basicEnum, meta metaType) *Enum[protoEnum, basicEnum, metaType] {
	return &Enum[protoEnum, basicEnum, metaType]{
		proto: proto,
		basic: basic,
		meta:  meta,
	}
}

// Proto returns the underlying Protocol Buffer enum value
//
// 返回底层的 Protocol Buffer 枚举值
func (c *Enum[protoEnum, basicEnum, metaType]) Proto() protoEnum {
	return c.proto
}

// Basic returns the Go native enum value bound to this enum
// Use it to move from a protobuf enum to a Go native enum (e.g. type StatusType string)
//
// 返回与此枚举绑定的 Go 原生枚举值
// 用它从 protobuf 枚举转到 Go 原生枚举（如 type StatusType string）
func (c *Enum[protoEnum, basicEnum, metaType]) Basic() basicEnum {
	return c.basic
}

// Code returns the numeric code of the enum as int32
// Converts the Protocol Buffer enum value to a standard int32 type
//
// 返回枚举的数字代码为 int32
// 将 Protocol Buffer 枚举数字转换成标准 int32 类型
func (c *Enum[protoEnum, basicEnum, metaType]) Code() int32 {
	return int32(c.proto.Number())
}

// Name returns the string name of the enum value
// Gets the Protocol Buffer enum's string representation
//
// 返回枚举值的字符串名称
// 获取 Protocol Buffer 枚举的字符串表示
func (c *Enum[protoEnum, basicEnum, metaType]) Name() string {
	return c.proto.String()
}

// Meta returns the metadata associated with this enum
// Provides access to custom metadata like description via MetaDesc
// Use this when you need to access extended enum metadata
//
// 返回与此枚举关联的元数据
// 提供对自定义元数据（如通过 MetaDesc 获取描述）的访问
// 在需要访问额外的枚举元数据时使用此方法
func (c *Enum[protoEnum, basicEnum, metaType]) Meta() metaType {
	return c.meta
}
