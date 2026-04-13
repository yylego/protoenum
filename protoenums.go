// Package protoenum: Collection management to handle Protocol Buffer enum metadata
// Provides indexed collections of enum descriptors with multiple lookup methods
// Enables fast lookup using code, name, and basic value with efficient enum handling
//
// protoenum: Protocol Buffer 枚举元数据集合管理
// 提供带有多种查找方法的枚举描述符索引集合
// 支持按代码、名称或 basic 枚举值快速检索，实现高效枚举处理
package protoenum

import (
	"slices"

	"github.com/yylego/erero"
	"github.com/yylego/protoenum/internal/utils"
	"github.com/yylego/tern/slicetern"
)

// Enums manages a collection of Enum instances with indexed lookups
// Maintains multiple maps enabling efficient lookup using different identifiers
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
	defaultValue *Enum[P, B, M]            // Configurable default value when lookup misses // 查找失败时的可选默认值
	defaultValid *bool                     // When true, default is treated as valid in ListValidXxx // 为 true 时，ListValidXxx 将默认值视为有效
}

// NewEnums creates a new Enums collection from the given Enum instances
// Builds indexed maps enabling efficient lookup using proto, code, name, and basic value
// The first item becomes the default value if provided
// Returns an error if duplicate proto, code, name, or basic values are detected
//
// 从给定的 Enum 实例创建新的 Enums 集合
// 构建索引映射以通过 proto、代码、名称和 basic 枚举值高效查找
// 如果提供了参数，第一个项成为默认值
// 如果检测到重复的 proto、代码、名称或 basic 值则返回错误
func NewEnums[P ProtoEnum, B comparable, M any](params ...*Enum[P, B, M]) (*Enums[P, B, M], error) {
	res := &Enums[P, B, M]{
		enumElements: slices.Clone(params), // Clone the slice to preserve the defined sequence // 克隆切片以保持枚举元素的定义次序
		mapProtoEnum: make(map[P]*Enum[P, B, M], len(params)),
		mapCode2Enum: make(map[int32]*Enum[P, B, M], len(params)),
		mapName2Enum: make(map[string]*Enum[P, B, M], len(params)),
		mapBasicEnum: make(map[B]*Enum[P, B, M], len(params)),
		defaultValue: slicetern.V0(params), // Set first item as default if available // 如果有参数，将第一个设置为默认值
		defaultValid: nil,
	}
	for _, enum := range params {
		if enum == nil {
			return nil, erero.New("ENUM ELEMENT IS MISSING")
		}
		if _, ok := res.mapProtoEnum[enum.Proto()]; ok {
			return nil, erero.Errorf("DUPLICATE PROTO ENUM: %v", enum.Proto())
		}
		res.mapProtoEnum[enum.Proto()] = enum
		if _, ok := res.mapCode2Enum[enum.Code()]; ok {
			return nil, erero.Errorf("DUPLICATE ENUM CODE: %v", enum.Code())
		}
		res.mapCode2Enum[enum.Code()] = enum
		if _, ok := res.mapName2Enum[enum.Name()]; ok {
			return nil, erero.Errorf("DUPLICATE ENUM NAME: %v", enum.Name())
		}
		res.mapName2Enum[enum.Name()] = enum
		if _, ok := res.mapBasicEnum[enum.Basic()]; ok {
			return nil, erero.Errorf("DUPLICATE BASIC ENUM: %v", enum.Basic())
		}
		res.mapBasicEnum[enum.Basic()] = enum
	}
	return res, nil
}

// LookupByProto finds an Enum using its Protocol Buffer enum value
// Returns the Enum and true if found, nil and false otherwise
//
// 通过 Protocol Buffer 枚举值查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
func (c *Enums[P, B, M]) LookupByProto(proto P) (*Enum[P, B, M], bool) {
	res, ok := c.mapProtoEnum[proto]
	return res, ok
}

// GetByProto finds an Enum using its Protocol Buffer enum value
// Returns default value if the enum is not found in the collection
// Returns nil if no default value has been configured
//
// 通过 Protocol Buffer 枚举值检索 Enum
// 如果在集合中找不到枚举则返回默认值
// 如果未配置默认值则返回 nil
func (c *Enums[P, B, M]) GetByProto(proto P) *Enum[P, B, M] {
	if res, ok := c.mapProtoEnum[proto]; ok {
		return res
	}
	return c.defaultValue
}

// LookupByCode finds an Enum using its numeric code
// Returns the Enum and true if found, nil and false otherwise
//
// 通过数字代码查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
func (c *Enums[P, B, M]) LookupByCode(code int32) (*Enum[P, B, M], bool) {
	res, ok := c.mapCode2Enum[code]
	return res, ok
}

// GetByCode finds an Enum using its numeric code
// Returns default value if no enum with the given code exists
// Returns nil if no default value has been configured
//
// 通过数字代码检索 Enum
// 如果不存在具有给定代码的枚举则返回默认值
// 如果未配置默认值则返回 nil
func (c *Enums[P, B, M]) GetByCode(code int32) *Enum[P, B, M] {
	if res, ok := c.mapCode2Enum[code]; ok {
		return res
	}
	return c.defaultValue
}

// LookupByName finds an Enum using its string name
// Returns the Enum and true if found, nil and false otherwise
//
// 通过字符串名称查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
func (c *Enums[P, B, M]) LookupByName(name string) (*Enum[P, B, M], bool) {
	res, ok := c.mapName2Enum[name]
	return res, ok
}

// GetByName finds an Enum using its string name
// Returns default value if no enum with the given name exists
// Returns nil if no default value has been configured
//
// 通过字符串名称检索 Enum
// 如果不存在具有给定名称的枚举则返回默认值
// 如果未配置默认值则返回 nil
func (c *Enums[P, B, M]) GetByName(name string) *Enum[P, B, M] {
	if res, ok := c.mapName2Enum[name]; ok {
		return res
	}
	return c.defaultValue
}

// LookupByBasic finds an Enum using its Go native enum value
// Returns the Enum and true if found, nil and false otherwise
//
// 通过 Go 原生枚举值查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
func (c *Enums[P, B, M]) LookupByBasic(basic B) (*Enum[P, B, M], bool) {
	res, ok := c.mapBasicEnum[basic]
	return res, ok
}

// GetByBasic finds an Enum using its Go native enum value
// Returns default value if no enum with the given basic enum exists
// Returns nil if no default value has been configured
//
// 通过 Go 原生枚举值检索 Enum
// 如果不存在具有给定 basic 枚举的枚举则返回默认值
// 如果未配置默认值则返回 nil
func (c *Enums[P, B, M]) GetByBasic(basic B) *Enum[P, B, M] {
	if res, ok := c.mapBasicEnum[basic]; ok {
		return res
	}
	return c.defaultValue
}

// GetDefault returns the current default Enum value
// Returns error if no default value has been configured
//
// 返回当前的默认 Enum 值
// 如果未配置默认值则返回错误
func (c *Enums[P, B, M]) GetDefault() (*Enum[P, B, M], error) {
	if c.defaultValue == nil {
		return nil, erero.New("NO DEFAULT VALUE CONFIGURED")
	}
	return c.defaultValue, nil
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

// ListValidProtos returns a slice excluding the default protoEnum value.
// If no default value is configured, returns each protoEnum value.
//
// 返回一个切片，排除默认 protoEnum 值，其余按定义次序排列。
// 如果未配置默认值，则返回所有 protoEnum 值。
func (c *Enums[P, B, M]) ListValidProtos() []P {
	if c.defaultValue != nil && !utils.GetPointerValue(c.defaultValid) {
		var results []P
		for _, item := range c.enumElements {
			if item.Code() != c.defaultValue.Code() {
				results = append(results, item.Proto())
			}
		}
		return results
	}
	return c.ListProtos()
}

// ListValidBasics returns a slice excluding the default basicEnum value.
// If no default value is configured, returns each basicEnum value.
//
// 返回一个切片，排除默认 basicEnum 值，其余按定义次序排列。
// 如果未配置默认值，则返回所有 basicEnum 值。
func (c *Enums[P, B, M]) ListValidBasics() []B {
	if c.defaultValue != nil && !utils.GetPointerValue(c.defaultValid) {
		var results []B
		for _, item := range c.enumElements {
			if item.Basic() != c.defaultValue.Basic() {
				results = append(results, item.Basic())
			}
		}
		return results
	}
	return c.ListBasics()
}
