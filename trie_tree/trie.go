package trie_tree

func (t *Trie) InsertWord(word string) {
	if word == "" {
		return
	}
	curr := t.RootNode
	curr.Pass++
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if curr.Children[index] == nil {
			curr.Children[index] = NewNode()
		}
		curr.Children[index].Pass++
		curr = curr.Children[index]
	}
	curr.End++
}

func (t *Trie) DeleteWord(word string) {
	// 如果之前插入过word
	if t.SearchWord(word) == 0 {
		return
	}
	curr := t.RootNode
	curr.Pass--
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		curr.Children[index].Pass--
		if curr.Children[index].Pass == 0 {
			curr.Children[index] = nil
			return
		}
		curr = curr.Children[index]
	}
	curr.End--
}

// SearchWord 查询word 插入过几次
func (t *Trie) SearchWord(word string) int {
	if word == "" {
		return 0
	}

	curr := t.RootNode
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		// 查询过程中发现 之前没insert过
		if curr.Children[index] == nil {
			return 0
		}
		curr = curr.Children[index]
	}
	return curr.End
}

// PrefixNumber 查询以word 为前缀的单词数量
func (t *Trie) PrefixNumber(word string) int {

	if word == "" {
		return 0
	}

	curr := t.RootNode
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		// 查询过程中发现 之前没insert过
		if curr.Children[index] == nil {
			return 0
		}
		curr = curr.Children[index]
	}
	return curr.Pass
}
