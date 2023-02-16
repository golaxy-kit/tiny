package ec

import (
	"kit.golaxy.org/tiny/localevent"
	"kit.golaxy.org/tiny/util"
	"kit.golaxy.org/tiny/util/container"
)

func UnsafeComponent(comp Component) _UnsafeComponent {
	return _UnsafeComponent{
		Component: comp,
	}
}

type _UnsafeComponent struct {
	Component
}

func (uc _UnsafeComponent) Init(name string, entity Entity, inheritor Component, hookCache container.Cache[localevent.Hook]) {
	uc.init(name, entity, inheritor, hookCache)
}

func (uc _UnsafeComponent) SetID(id ID) {
	uc.setID(id)
}

func (uc _UnsafeComponent) GetContext() util.IfaceCache {
	return uc.getContext()
}

func (uc _UnsafeComponent) SetState(state ComponentState) {
	uc.setState(state)
}

func (uc _UnsafeComponent) GetInheritor() Component {
	return uc.getInheritor()
}

func (uc _UnsafeComponent) EventComponentDestroySelf() localevent.IEvent {
	return uc.eventComponentDestroySelf()
}

func (uc _UnsafeComponent) GetInnerGC() container.GC {
	return uc.getInnerGC()
}

func (uc _UnsafeComponent) GetInnerGCCollector() container.GCCollector {
	return uc.getInnerGCCollector()
}
