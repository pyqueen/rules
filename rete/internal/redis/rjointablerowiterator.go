package redis

import (
	"container/list"
	"github.com/project-flogo/rules/rete/internal/types"
)

type rowIteratorImpl struct {
	table map[int]types.JoinTableRow
	kList list.List
	curr  *list.Element
}

func newRowIterator(jTable map[int]types.JoinTableRow) types.RowIterator {
	ri := rowIteratorImpl{}
	ri.table = jTable
	ri.kList = list.List{}
	for k, _ := range jTable {
		ri.kList.PushBack(k)
	}
	ri.curr = ri.kList.Front()
	return &ri
}

func (ri *rowIteratorImpl) HasNext() bool {
	return ri.curr != nil
}

func (ri *rowIteratorImpl) Next() types.JoinTableRow {
	id := ri.curr.Value.(int)
	val := ri.table[id]
	ri.curr = ri.curr.Next()
	return val
}
