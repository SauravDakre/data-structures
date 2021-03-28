package trie

import "fmt"

type Trie struct {
	links map[string]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{}
	t.links = make(map[string]*Trie)
	t.isEnd = false
	return t
}

func (t *Trie) read() {
	for k, v := range t.links {
		fmt.Println("k:", k, "v:", *v)
		t = v
		t.read()
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	tmp := t
	for _, c := range word {
		if _, val := tmp.links[string(c)]; !val {
			nt := Constructor()
			tmp.links[string(c)] = &nt
		}
		tmp = tmp.links[string(c)]
	}
	tmp.isEnd = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	tmp := t
	for _, c := range word {
		if l, valid := tmp.links[string(c)]; valid {
			tmp = l
		} else {
			return false
		}
	}
	return tmp.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	tmp := t
	for _, c := range prefix {
		if l, valid := tmp.links[string(c)]; valid {
			tmp = l
		} else {
			return false
		}
	}
	return true
}
