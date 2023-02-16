//go:generate go run kit.golaxy.org/tiny/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE
package runtime

import "kit.golaxy.org/tiny/ec"

// EventECTreeAddChild [EmitUnExport] 事件：EC树中子实体加入父实体
type EventECTreeAddChild interface {
	OnAddChild(ecTree IECTree, parent, child ec.Entity)
}

// EventECTreeRemoveChild [EmitUnExport] 事件：EC树中子实体离开父实体
type EventECTreeRemoveChild interface {
	OnRemoveChild(ecTree IECTree, parent, child ec.Entity)
}
