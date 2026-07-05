package protoenum

import (
	"github.com/yylego/erero"
	"github.com/yylego/must"
)

// setDefaultOnce fixes the default to the given element just once
// Returns an error when a default is present, so the default stays fixed
//
// setDefaultOnce 把默认值一次性固定到给定成员
// 若默认值已存在则报错，因此默认值永不改变
func (c *Enums[P, B, M]) setDefaultOnce(enum *Enum[P, B, M]) error {
	if c.defaultValue != nil {
		return erero.New("DEFAULT VALUE EXISTED WHEN SETTING")
	}
	c.defaultValue = enum
	return nil
}

// WithDefault fixes the first element as the default; GetByXxx returns it on a miss
// Use it once during construction; without it the collection carries no default
// Panics when the collection has no elements / a default is present
//
// WithDefault 把首元素固定为默认值，GetByXxx 查不到时返回它
// 构造时调用一次即可；不调用则集合没有默认值
// 当集合为空或默认值已存在时 panic
func (c *Enums[P, B, M]) WithDefault() *Enums[P, B, M] {
	must.TRUE(len(c.enumElements) > 0)
	must.Done(c.setDefaultOnce(c.enumElements[0]))
	return c
}

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
