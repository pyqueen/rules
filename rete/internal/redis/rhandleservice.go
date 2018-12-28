package redis

import (
	"github.com/project-flogo/rules/common/model"
	"github.com/project-flogo/rules/rete/internal/types"
)

type handleServiceImpl struct {
	allHandles map[string]types.ReteHandle
}

func NewHandleCollection(config map[string]interface{}) types.HandleService {
	hc := handleServiceImpl{}
	hc.allHandles = make(map[string]types.ReteHandle)
	return &hc
}

func (hc *handleServiceImpl) Init() {

}

func (hc *handleServiceImpl) AddHandle(hdl types.ReteHandle) {
	hc.allHandles[hdl.GetTupleKey().String()] = hdl
}

func (hc *handleServiceImpl) RemoveHandle(tuple model.Tuple) types.ReteHandle {
	rh, found := hc.allHandles[tuple.GetKey().String()]
	if found {
		delete(hc.allHandles, tuple.GetKey().String())
		return rh
	}
	return nil
}

func (hc *handleServiceImpl) GetHandle(tuple model.Tuple) types.ReteHandle {
	return hc.allHandles[tuple.GetKey().String()]
}

func (hc *handleServiceImpl) GetHandleByKey(key model.TupleKey) types.ReteHandle {
	return hc.allHandles[key.String()]
}

func (hc *handleServiceImpl) GetOrCreateHandle(nw types.Network, tuple model.Tuple) types.ReteHandle {
	h, found := hc.allHandles[tuple.GetKey().String()]
	if !found {
		h = newReteHandleImpl(nw, tuple)
		hc.allHandles[tuple.GetKey().String()] = h
	}

	return h
}
