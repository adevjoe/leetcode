// https://leetcode.com/problems/word-ladder

package leetcode

import "reflect"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool)
	for i := range wordList {
		wordMap[wordList[i]] = false
	}
	queue := NewQueue()
	queue.Push(beginWord)
	ladder := 1
	for queue.Len() != 0 {
		l := queue.Len()
		for j := 0; j < l; j++ {
			word := queue.Touch()
			queue.Pop()
			if word.Equal(endWord) {
				return ladder
			}
			delete(wordMap, string(word))

			for k := range word {
				for i := 0; i < 26; i++ {
					c := word[k]
					word[k] = 'a' + byte(i)
					if _, ok := wordMap[string(word)]; ok && !reflect.DeepEqual(word, queue.End()) {
						queue.Push(string(word))
					}
					word[k] = c
				}
			}
		}
		ladder++
	}
	return 0
}

type Queue struct {
	values []SS
}

type SS []byte

func (ss *SS) Equal(s string) bool {
	return string(*ss) == s
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Len() int {
	return len(q.values)
}

func (q *Queue) Empty() bool {
	return q.Len() == 0
}

func (q *Queue) Push(i string) {
	q.values = append(q.values, SS(i))
}

func (q *Queue) Pop() SS {
	if q.Empty() {
		return nil
	}
	v := q.values[0]
	q.values = q.values[1:]
	return v
}

func (q *Queue) Touch() SS {
	if q.Empty() {
		return nil
	}
	return q.values[0]
}

func (q *Queue) End() SS {
	if q.Empty() {
		return nil
	}
	return q.values[q.Len()-1]
}
