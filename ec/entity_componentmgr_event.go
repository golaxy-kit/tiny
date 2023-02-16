//go:generate go run kit.golaxy.org/tiny/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE
package ec

// EventCompMgrAddComponents [EmitUnExport] 事件：实体的组件管理器加入一些组件
type EventCompMgrAddComponents interface {
	OnCompMgrAddComponents(entity Entity, components []Component)
}

// EventCompMgrRemoveComponent [EmitUnExport] 事件：实体的组件管理器删除组件
type EventCompMgrRemoveComponent interface {
	OnCompMgrRemoveComponent(entity Entity, component Component)
}
