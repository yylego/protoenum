package protoenum

import (
	"github.com/yylego/erero"
	"github.com/yylego/must"
)

// GetDefaultProto returns the protoEnum value of the default Enum
// Returns zero value and error if no default value has been configured
//
// 返回默认 Enum 的 protoEnum 值
// 如果未配置默认值则返回零值和错误
func (c *Enums[P, B, M]) GetDefaultProto() (P, error) {
	def, err := c.GetDefault()
	if err != nil {
		var zero P
		return zero, err
	}
	return def.Proto(), nil
}

// GetDefaultBasic returns the basicEnum value of the default Enum
// Returns zero value and error if no default value has been configured
//
// 返回默认 Enum 的 basicEnum 值
// 如果未配置默认值则返回零值和错误
func (c *Enums[P, B, M]) GetDefaultBasic() (B, error) {
	def, err := c.GetDefault()
	if err != nil {
		var zero B
		return zero, err
	}
	return def.Basic(), nil
}

// SetDefault sets the default Enum value to return when lookups miss
// Returns error if defaultEnum is nil or if a default has been set
//
// 设置查找失败时返回的默认 Enum 值
// 如果 defaultEnum 为 nil 或已设置过默认值则返回错误
func (c *Enums[P, B, M]) SetDefault(enum *Enum[P, B, M]) error {
	if c.defaultValue != nil {
		return erero.New("DEFAULT VALUE EXISTED WHEN SETTING")
	}
	if enum == nil {
		return erero.New("DEFAULT ENUM IS MISSING")
	}
	c.defaultValue = enum
	c.defaultValid = nil
	return nil
}

// SetDefaultProto sets the default using a Protocol Buffer enum value
// Returns error if the specified proto enum is not found in the collection
//
// 使用 Protocol Buffer 枚举值设置默认值
// 如果指定的 proto 枚举不存在则返回错误
func (c *Enums[P, B, M]) SetDefaultProto(proto P) error {
	res, ok := c.LookupByProto(proto)
	if !ok {
		return erero.Errorf("PROTO ENUM NOT FOUND: %v", proto)
	}
	return c.SetDefault(res)
}

// SetDefaultBasic sets the default using a Go native enum value
// Returns error if the specified basic enum is not found in the collection
//
// 使用 Go 原生枚举值设置默认值
// 如果指定的 basic 枚举不存在则返回错误
func (c *Enums[P, B, M]) SetDefaultBasic(basic B) error {
	res, ok := c.LookupByBasic(basic)
	if !ok {
		return erero.Errorf("BASIC ENUM NOT FOUND: %v", basic)
	}
	return c.SetDefault(res)
}

// SetDefaultValid marks the default value as active when true
// When active, ListValidProtos and ListValidBasics include the default
// Returns error if no default value exists or if defaultValid has been set
//
// 标记默认值是否应被视为有效
// 当 valid 为 true 时，ListValidProtos 和 ListValidBasics 包含默认值
// 如果无默认值或 defaultValid 已设置则返回错误
func (c *Enums[P, B, M]) SetDefaultValid(valid bool) error {
	if c.defaultValue == nil {
		return erero.New("NO DEFAULT VALUE CONFIGURED")
	}
	if c.defaultValid != nil {
		return erero.New("DEFAULT VALID EXISTED WHEN SETTING")
	}
	c.defaultValid = &valid
	return nil
}

// UnsetDefault unsets the default Enum value
// Once invoked, GetByXxx lookups return nil if not found
// Returns error if no default value exists at the moment
//
// 取消设置默认 Enum 值
// 调用后，GetByXxx 查找失败时返回 nil
// 如果当前无默认值则返回错误
func (c *Enums[P, B, M]) UnsetDefault() error {
	if c.defaultValue == nil {
		return erero.New("NO DEFAULT VALUE TO UNSET")
	}
	c.defaultValue = nil
	c.defaultValid = nil
	return nil
}

// WithDefault sets the default Enum value and returns the Enums instance
// Enables fluent chain-style configuration during initialization
//
// 设置默认 Enum 值并返回 Enums 实例
// 支持初始化时的流式链式配置
func (c *Enums[P, B, M]) WithDefault(enum *Enum[P, B, M]) *Enums[P, B, M] {
	must.Done(c.SetDefault(enum))
	return c
}

// WithDefaultProto sets the default using a proto enum and returns the Enums instance
//
// 使用 proto 枚举设置默认值并返回 Enums 实例
func (c *Enums[P, B, M]) WithDefaultProto(proto P) *Enums[P, B, M] {
	must.Done(c.SetDefaultProto(proto))
	return c
}

// WithDefaultBasic sets the default using a basic enum and returns the Enums instance
//
// 使用 basic 枚举设置默认值并返回 Enums 实例
func (c *Enums[P, B, M]) WithDefaultBasic(basic B) *Enums[P, B, M] {
	must.Done(c.SetDefaultBasic(basic))
	return c
}

// WithDefaultCode sets the default using a numeric code and returns the Enums instance
//
// 使用数字代码设置默认值并返回 Enums 实例
func (c *Enums[P, B, M]) WithDefaultCode(code int32) *Enums[P, B, M] {
	res, ok := c.LookupByCode(code)
	must.TRUE(ok)
	return c.WithDefault(res)
}

// WithDefaultName sets the default using an enum name and returns the Enums instance
//
// 使用枚举名称设置默认值并返回 Enums 实例
func (c *Enums[P, B, M]) WithDefaultName(name string) *Enums[P, B, M] {
	res, ok := c.LookupByName(name)
	must.TRUE(ok)
	return c.WithDefault(res)
}

// WithDefaultValid marks the default as active and returns the Enums instance
//
// 标记默认值为有效并返回 Enums 实例
func (c *Enums[P, B, M]) WithDefaultValid(valid bool) *Enums[P, B, M] {
	must.Done(c.SetDefaultValid(valid))
	return c
}

// WithUnsetDefault unsets the default Enum value and returns the Enums instance
//
// 取消设置默认 Enum 值并返回 Enums 实例
func (c *Enums[P, B, M]) WithUnsetDefault() *Enums[P, B, M] {
	must.Done(c.UnsetDefault())
	return c
}
