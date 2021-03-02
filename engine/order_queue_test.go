package engine

import (
	"matching-engine/enum"
	"testing"
)

func TestOrderQueue_addOrder(t *testing.T) {
	q := new(orderQueue)
	q.init(enum.DESC)
	q.addOrder(&Order{})
}
