package doubleLinkedList

import "fmt"

type node struct {
	data       int
	prev, next *node
}

func newNode(data int) *node {
	n := new(node)
	n.data = data
	return n
}

type Dll struct {
	head *node
}

func New() *Dll {
	l := new(Dll)
	return l
}

func (l *Dll) addFirst(data int) {
	n := newNode(data)

	if l.head == nil {
		l.head = n
		return
	}

	l.head.prev = n
	n.next = l.head
	l.head = n
	return
}

func (l *Dll) addLast(data int) {
	n := newNode(data)

	if l.head == nil {
		l.head = n
		return
	}

	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}

	tmp.next = n
	n.prev = tmp
	return
}

func (l *Dll) Delete(data int) {

	tmp := l.head
	for tmp.next != nil {
		if tmp.data == data {
			break
		}
		tmp = tmp.next
	}

	if tmp.next == nil {
		return
	}

	nn := tmp.next
	prev := tmp.prev
	prev.next = nn
	nn.prev = prev
	return
}

func (l *Dll) InsertAfter(data, newData int) {
	n := newNode(newData)

	tmp := l.head
	for tmp.next != nil {
		if tmp.data == data {
			break
		}
		tmp = tmp.next
	}

	if tmp.next == nil {
		return
	}

	nn := tmp.next
	tmp.next = n
	n.next = nn
	n.prev = tmp
	nn.prev = n
	return
}

func (l Dll) print() (f, r string) {
	tmp := l.head
	f = ""
	var t2 *node
	for tmp != nil {
		t2 = tmp
		f += fmt.Sprintf("%d ->", tmp.data)
		tmp = tmp.next
	}
	r = ""
	for t2 != nil {
		r += fmt.Sprintf("%d ->", t2.data)
		t2 = t2.prev
	}

	return
}
