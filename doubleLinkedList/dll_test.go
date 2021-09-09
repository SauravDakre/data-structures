package doubleLinkedList

import (
	"testing"
)

func TestDll(t *testing.T) {
	l := New()
	l.addLast(1)
	l.addFirst(7)
	l.addLast(2)
	l.addFirst(9)
	l.addFirst(8)
	l.addLast(3)
	l.InsertAfter(2, 5)
	l.Delete(2)
	l.addFirst(4)

	f, r := l.print()

	if f != "4 ->8 ->9 ->7 ->1 ->5 ->3 ->" {
		t.Error("wrong double linked list")
	}
	if r != "3 ->5 ->1 ->7 ->9 ->8 ->4 ->" {
		t.Error("wrong double linked list")
	}
}

func TestDelete(t *testing.T) {
	l := New()
	l.addLast(1)
	l.addLast(2)
	l.addLast(3)
	l.InsertAfter(2, 5)
	l.Delete(2)
	f, r := l.print()

	if f != "1 ->5 ->3 ->" {
		t.Error("wrong double linked list")
	}
	if r != "3 ->5 ->1 ->" {
		t.Error("wrong double linked list")
	}
}

func TestInsertAfter(t *testing.T) {
	l := New()
	l.addLast(1)
	l.addLast(2)
	l.addLast(3)
	l.InsertAfter(2, 5)
	f, r := l.print()

	if f != "1 ->2 ->5 ->3 ->" {
		t.Error("wrong double linked list")
	}
	if r != "3 ->5 ->2 ->1 ->" {
		t.Error("wrong double linked list")
	}
}

func TestAdddLast(t *testing.T) {
	l := New()
	l.addLast(1)
	l.addLast(2)
	l.addLast(3)
	f, r := l.print()

	if f != "1 ->2 ->3 ->" {
		t.Error("wrong double linked list")
	}
	if r != "3 ->2 ->1 ->" {
		t.Error("wrong double linked list")
	}
}

func TestAddFirst(t *testing.T) {
	l := New()
	l.addFirst(1)
	l.addFirst(2)
	l.addFirst(3)
	f, r := l.print()

	if f != "3 ->2 ->1 ->" {
		t.Error("wrong double linked list")
	}
	if r != "1 ->2 ->3 ->" {
		t.Error("wrong double linked list")
	}

}
