package linkedlist

import . "review-go/node"

func Contains[T comparable](ll *LinkedList[T], cmpFn func(T, T) bool, data T) bool {
	ll.RLock()
	defer ll.RUnlock()

	for v := range ll.Values() {
		if cmpFn(v, data) {
			return true
		}
	}

	return false
}

func Find[T comparable](ll *LinkedList[T], cmpFn func(T, T) bool, data T) *Node[T] {
	for n := ll.Head(); n != nil; n = n.Next() {
		if cmpFn(n.Value(), data) {
			return n
		}
	}
	return nil
}
