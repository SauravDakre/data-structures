package trie

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tr := Constructor()
	tr.Insert("abcd")
	tr.Insert("ac")

	v := tr.links["a"]
	ch := make([]string, 0)
	for k := range v.links {
		ch = append(ch, string(k))
	}
	if len(ch) != 2 {
		t.Error("expected length to be 2 but found", len(ch))
	}
}

func TestSearch(t *testing.T) {
	tr := Constructor()
	tr.Insert("apple")
	tr.Insert("app")

	if tr.Search("app") != true {
		t.Error("expected search app to be true but found false")
	}

	if tr.Search("ap") != false {
		t.Error("expected search app to be false but found true")
	}
}

func TestStartsWith(t *testing.T) {
	tr := Constructor()
	tr.Insert("apple")
	tr.Insert("app")

	if tr.StartsWith("ap") != true {
		t.Error("expected StartsWith ap to be true but found false")
	}

	if tr.StartsWith("ab") != false {
		t.Error("expected StartsWith ab to be false but found true")
	}
}
