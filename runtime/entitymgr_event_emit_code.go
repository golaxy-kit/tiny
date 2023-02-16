// Code generated by eventcode --decl_file=entitymgr_event.go gen_emit --package=runtime; DO NOT EDIT.

package runtime

import (
	localevent "kit.golaxy.org/tiny/localevent"
	"kit.golaxy.org/tiny/ec"
	"kit.golaxy.org/tiny/util"
)

func emitEventEntityMgrAddEntity(event localevent.IEvent, entityMgr IEntityMgr, entity ec.Entity) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[EventEntityMgrAddEntity](delegate).OnEntityMgrAddEntity(entityMgr, entity)
		return true
	})
}

func emitEventEntityMgrRemovingEntity(event localevent.IEvent, entityMgr IEntityMgr, entity ec.Entity) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[EventEntityMgrRemovingEntity](delegate).OnEntityMgrRemovingEntity(entityMgr, entity)
		return true
	})
}

func emitEventEntityMgrRemoveEntity(event localevent.IEvent, entityMgr IEntityMgr, entity ec.Entity) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[EventEntityMgrRemoveEntity](delegate).OnEntityMgrRemoveEntity(entityMgr, entity)
		return true
	})
}

func emitEventEntityMgrEntityAddComponents(event localevent.IEvent, entityMgr IEntityMgr, entity ec.Entity, components []ec.Component) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[EventEntityMgrEntityAddComponents](delegate).OnEntityMgrEntityAddComponents(entityMgr, entity, components)
		return true
	})
}

func emitEventEntityMgrEntityRemoveComponent(event localevent.IEvent, entityMgr IEntityMgr, entity ec.Entity, component ec.Component) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[EventEntityMgrEntityRemoveComponent](delegate).OnEntityMgrEntityRemoveComponent(entityMgr, entity, component)
		return true
	})
}
