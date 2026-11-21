package protoenum

import (
	"strings"

	"github.com/yylego/erero"
)

// CheckComplete checks that each value the proto enum declares has an element in the collection
// Use it once during construction, like WithDefault, so a value added to the proto and not mapped here
// shows up when the program starts, not when a request first carries that value
// Panics naming the values that have no element
//
// CheckComplete 检查 proto 枚举声明的每个值在集合里都有对应元素
// 跟 WithDefault 一样在构造时调用一次：proto 里加了新值、这里忘了登记，程序启动时就能发现，
// 不用等到第一个带着这个值的请求进来
// 有值没登记时 panic，并点名是哪些值
func (c *Enums[P, B, M]) CheckComplete() *Enums[P, B, M] {
	var zero P
	enumDesc := zero.Descriptor()

	var missing []string
	values := enumDesc.Values()
	for idx := range values.Len() {
		value := values.Get(idx)
		if _, ok := c.mapCode2Enum[int32(value.Number())]; !ok {
			missing = append(missing, string(value.Name()))
		}
	}
	if len(missing) > 0 {
		panic(erero.Errorf("ENUM %s HAS VALUES WITHOUT AN ELEMENT: %s", enumDesc.FullName(), strings.Join(missing, ", ")))
	}
	return c
}
