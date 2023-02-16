package runtime

import "kit.golaxy.org/tiny/util"

// GetInheritor 获取运行时上下文的继承者
func GetInheritor[T any](ctx Context) T {
	return util.Cache2Iface[T](ctx.getOptions().Inheritor.Cache)
}
