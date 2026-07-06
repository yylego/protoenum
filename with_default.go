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

// WithDefault fixes the first element as the default; GetByXxxFallbackDefault returns it on a miss
// Use it once during construction; without it the collection carries no default
// Panics when the collection has no elements / a default is present
//
// WithDefault 把首元素固定为默认值，GetByXxxFallbackDefault 查不到时返回它
// 构造时调用一次即可；不调用则集合没有默认值
// 当集合为空或默认值已存在时 panic
func (c *Enums[P, B, M]) WithDefault() *Enums[P, B, M] {
	must.TRUE(len(c.enumElements) > 0)
	must.Done(c.setDefaultOnce(c.enumElements[0]))
	return c
}
