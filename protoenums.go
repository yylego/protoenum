// Package protoenum: Collection management to handle Protocol Buffer enum metadata
// Provides indexed collections of enum descriptors with multiple lookup methods
// Looks up via proto, code, name, and basic value
//
// protoenum: Protocol Buffer 枚举元数据集合管理
// 提供带有多种查找方法的枚举描述符索引集合
// 支持按代码、名称或 basic 枚举值快速检索，实现高效枚举处理
package protoenum

import (
	"slices"

	"github.com/yylego/must"
)

// Enums manages a collection of Enum instances with indexed lookups
// Keeps a map for proto, code, name, and basic so lookup stays O(1)
// Provides O(1) lookup when searching proto, code, name, and basic value
// Includes a configurable default value returned when lookups miss
//
// Enums 管理 Enum 实例集合并提供索引查找
// 维护四个映射表以通过不同标识符高效检索
// 为 proto、代码、名称和 basic 枚举值搜索提供 O(1) 查找性能
// 支持在查找失败时返回可选的默认值
type Enums[P ProtoEnum, B comparable, M any] struct {
	enumElements []*Enum[P, B, M]          // Holds complete Enum instances in defined sequence // 存放所有 Enum 实例，并维持其定义的次序
	mapProtoEnum map[P]*Enum[P, B, M]      // Map from proto enum to Enum // 从 proto 枚举到 Enum 的映射
	mapCode2Enum map[int32]*Enum[P, B, M]  // Map from numeric code to Enum // 从数字代码到 Enum 的映射
	mapName2Enum map[string]*Enum[P, B, M] // Map from name string to Enum // 从名称字符串到 Enum 的映射
	mapBasicEnum map[B]*Enum[P, B, M]      // Map from basic enum to Enum // 从 basic 枚举到 Enum 的映射
	defaultValue *Enum[P, B, M]            // Fallback returned by GetByXxxFallbackDefault on a miss; nil under NoDefault // GetByXxxFallbackDefault 查不到时回落的兜底值；NoDefault 时为 nil
}

// NewEnums creates a new Enums collection from the given Enum instances
// Builds a lookup map for each of proto, code, name, and basic
// No default is set here; the default is an explicit opt-in — chain WithDefault once to fix the first element
// Panics on a nil element / a duplicate proto / code / name / basic value
//
// 从给定的 Enum 实例创建新的 Enums 集合
// 构建索引映射以通过 proto、代码、名称和 basic 枚举值高效查找
// 此处不设默认值；默认值是显式可选项——如需兜底，链式调用一次 WithDefault 把首元素固定为默认
// 遇到 nil 成员或重复的 proto、代码、名称、basic 值时 panic
func NewEnums[P ProtoEnum, B comparable, M any](params ...*Enum[P, B, M]) *Enums[P, B, M] {
	res := &Enums[P, B, M]{
		enumElements: slices.Clone(params), // Clone the slice to preserve the defined sequence // 克隆切片以保持枚举元素的定义次序
		mapProtoEnum: make(map[P]*Enum[P, B, M], len(params)),
		mapCode2Enum: make(map[int32]*Enum[P, B, M], len(params)),
		mapName2Enum: make(map[string]*Enum[P, B, M], len(params)),
		mapBasicEnum: make(map[B]*Enum[P, B, M], len(params)),
	}
	for _, enum := range params {
		must.Full(enum) // reject a nil element // 拒绝 nil 成员

		// Reject a duplicate proto / code / name / basic key // 拒绝重复的 proto / code / name / basic 键
		must.Null(res.mapProtoEnum[enum.Proto()])
		res.mapProtoEnum[enum.Proto()] = enum
		must.Null(res.mapCode2Enum[enum.Code()])
		res.mapCode2Enum[enum.Code()] = enum
		must.Null(res.mapName2Enum[enum.Name()])
		res.mapName2Enum[enum.Name()] = enum
		must.Null(res.mapBasicEnum[enum.Basic()])
		res.mapBasicEnum[enum.Basic()] = enum
	}
	return res
}

// GetByProto finds an Enum using its Protocol Buffer enum value
// Returns the Enum and true if found, nil and false otherwise
//
// 通过 Protocol Buffer 枚举值查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
func (c *Enums[P, B, M]) GetByProto(proto P) (*Enum[P, B, M], bool) {
	res, ok := c.mapProtoEnum[proto]
	return res, ok
}

// GetByProtoFallbackDefault finds an Enum using its Protocol Buffer enum value
// Returns default value if the enum is not found in the collection
// Returns (result, true) when usable; (nil, false) when it misses and no default is set
//
// 通过 Protocol Buffer 枚举值检索 Enum
// 如果在集合中找不到枚举则返回默认值
// 命中或有默认兜底时返回 (enum, true)；都没有时返回 (nil, false)
func (c *Enums[P, B, M]) GetByProtoFallbackDefault(proto P) (*Enum[P, B, M], bool) {
	if res, ok := c.mapProtoEnum[proto]; ok {
		return res, true
	}
	return c.defaultValue, c.defaultValue != nil
}

// GetByCode finds an Enum using its numeric code
// Returns the Enum and true if found, nil and false otherwise
//
// 通过数字代码查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
func (c *Enums[P, B, M]) GetByCode(code int32) (*Enum[P, B, M], bool) {
	res, ok := c.mapCode2Enum[code]
	return res, ok
}

// GetByCodeFallbackDefault finds an Enum using its numeric code
// Returns default value if no enum with the given code exists
// Returns (result, true) when usable; (nil, false) when it misses and no default is set
//
// 通过数字代码检索 Enum
// 如果不存在具有给定代码的枚举则返回默认值
// 命中或有默认兜底时返回 (enum, true)；都没有时返回 (nil, false)
func (c *Enums[P, B, M]) GetByCodeFallbackDefault(code int32) (*Enum[P, B, M], bool) {
	if res, ok := c.mapCode2Enum[code]; ok {
		return res, true
	}
	return c.defaultValue, c.defaultValue != nil
}

// GetByName finds an Enum using its string name
// Returns the Enum and true if found, nil and false otherwise
//
// 通过字符串名称查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
func (c *Enums[P, B, M]) GetByName(name string) (*Enum[P, B, M], bool) {
	res, ok := c.mapName2Enum[name]
	return res, ok
}

// GetByNameFallbackDefault finds an Enum using its string name
// Returns default value if no enum with the given name exists
// Returns (result, true) when usable; (nil, false) when it misses and no default is set
//
// 通过字符串名称检索 Enum
// 如果不存在具有给定名称的枚举则返回默认值
// 命中或有默认兜底时返回 (enum, true)；都没有时返回 (nil, false)
func (c *Enums[P, B, M]) GetByNameFallbackDefault(name string) (*Enum[P, B, M], bool) {
	if res, ok := c.mapName2Enum[name]; ok {
		return res, true
	}
	return c.defaultValue, c.defaultValue != nil
}

// GetByBasic finds an Enum using its Go native enum value
// Returns the Enum and true if found, nil and false otherwise
//
// 通过 Go 原生枚举值查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
func (c *Enums[P, B, M]) GetByBasic(basic B) (*Enum[P, B, M], bool) {
	res, ok := c.mapBasicEnum[basic]
	return res, ok
}

// GetByBasicFallbackDefault finds an Enum using its Go native enum value
// Returns default value if no enum with the given basic enum exists
// Returns (result, true) when usable; (nil, false) when it misses and no default is set
//
// 通过 Go 原生枚举值检索 Enum
// 如果不存在具有给定 basic 枚举的枚举则返回默认值
// 命中或有默认兜底时返回 (enum, true)；都没有时返回 (nil, false)
func (c *Enums[P, B, M]) GetByBasicFallbackDefault(basic B) (*Enum[P, B, M], bool) {
	if res, ok := c.mapBasicEnum[basic]; ok {
		return res, true
	}
	return c.defaultValue, c.defaultValue != nil
}

// GetDefault returns the default Enum
// Returns (default, true) when fixed; (nil, false) when no default has been fixed
//
// 返回默认 Enum
// 已固定时返回 (默认值, true)，未固定时返回 (nil, false)
func (c *Enums[P, B, M]) GetDefault() (*Enum[P, B, M], bool) {
	return c.defaultValue, c.defaultValue != nil
}

// ListProtos returns a slice containing each protoEnum value in the defined sequence
// Maintains the same sequence as enum values were registered
//
// 返回一个包含各 protoEnum 值的切片，次序与定义时一致
// 保持枚举值注册时的顺序
func (c *Enums[P, B, M]) ListProtos() []P {
	var results = make([]P, 0, len(c.enumElements))
	for _, item := range c.enumElements {
		results = append(results, item.Proto())
	}
	return results
}

// ListBasics returns a slice containing each basicEnum value in the defined sequence
// Maintains the same sequence as enum values were registered
//
// 返回一个包含各 basicEnum 值的切片，次序与定义时一致
// 保持枚举值注册时的顺序
func (c *Enums[P, B, M]) ListBasics() []B {
	var results = make([]B, 0, len(c.enumElements))
	for _, item := range c.enumElements {
		results = append(results, item.Basic())
	}
	return results
}

// ListNonDefaultProtos returns a slice excluding the default protoEnum value.
// If no default value is configured, returns each protoEnum value.
//
// 返回一个切片，排除默认 protoEnum 值，其余按定义次序排列。
// 如果未配置默认值，则返回所有 protoEnum 值。
func (c *Enums[P, B, M]) ListNonDefaultProtos() []P {
	if c.defaultValue == nil {
		return c.ListProtos()
	}
	results := make([]P, 0, len(c.enumElements))
	for _, item := range c.enumElements {
		if item.Proto() != c.defaultValue.Proto() {
			results = append(results, item.Proto())
		}
	}
	return results
}

// ListNonDefaultBasics returns a slice excluding the default basicEnum value.
// If no default value is configured, returns each basicEnum value.
//
// 返回一个切片，排除默认 basicEnum 值，其余按定义次序排列。
// 如果未配置默认值，则返回所有 basicEnum 值。
func (c *Enums[P, B, M]) ListNonDefaultBasics() []B {
	if c.defaultValue == nil {
		return c.ListBasics()
	}
	results := make([]B, 0, len(c.enumElements))
	for _, item := range c.enumElements {
		if item.Basic() != c.defaultValue.Basic() {
			results = append(results, item.Basic())
		}
	}
	return results
}

// ListEnums returns a slice of each Enum instance in the defined sequence
// Maintains the same sequence as enum values were registered
//
// ListEnums 返回一个包含各 Enum 实例的切片，次序与定义时一致
// 保持枚举值注册时的顺序
func (c *Enums[P, B, M]) ListEnums() []*Enum[P, B, M] {
	return slices.Clone(c.enumElements)
}

// ListNonDefaultEnums returns a slice of Enum instances excluding the default
// If no default value is configured, returns each Enum instance
//
// ListNonDefaultEnums 返回排除默认值的 Enum 实例切片，其余按定义次序排列
// 如果未配置默认值，则返回所有 Enum 实例
func (c *Enums[P, B, M]) ListNonDefaultEnums() []*Enum[P, B, M] {
	if c.defaultValue == nil {
		return c.ListEnums()
	}
	results := make([]*Enum[P, B, M], 0, len(c.enumElements))
	for _, item := range c.enumElements {
		if item.Proto() != c.defaultValue.Proto() {
			results = append(results, item)
		}
	}
	return results
}
