package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}

	endWordExists := false

	combos := map[string][]string{}
	for _, w := range wordList {
		if w == endWord {
			endWordExists = true
		}
		for i := range w {
			c := w[:i]+"*"+w[i+1:]
			if _, ok := combos[c]; !ok {
				combos[c] = []string{w}
			} else {
				combos[c] = append(combos[c], w)
			}
		}
	}

	if !endWordExists {
		return 0
	}

	type wordStep struct {
		w string
		s int
	}
	queue := []wordStep{
		{w: beginWord, s: 1},
	}

	for q := 0; q < len(queue); q++ {
		for i := range queue[q].w {
			c := queue[q].w[:i]+"*"+queue[q].w[i+1:]
			if _, ok := combos[c]; ok {
				for _, w := range combos[c] {
					if w == endWord {
						return queue[q].s+1
					}
					queue = append(queue, wordStep{w: w, s: queue[q].s+1})
					delete(combos, c)
				}
			}
		}
	}

	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot","dot","dog","lot","log","cog"}))
}