package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/Yesphet/trees/binary-search-tree"
)

type List struct {
	list []int
}

func (l *List) Len() int {
	return len(l.list)
}

func (l *List) Less(i, j int) bool {
	return l.list[i] < l.list[j]
}

func (l *List) Swap(i, j int) {
	tmp := l.list[i]
	l.list[i] = l.list[j]
	l.list[j] = tmp
}

func main() {

	maxNum := 10000000
	list := rand.Perm(maxNum)

	startTime := time.Now()
	tree := &binary_search_tree.BinarySearchTree{}
	for _, v := range list {
		tree.Insert(v, 1)
	}
	fmt.Printf("Non-Recursive Duration: %s\n", time.Now().Sub(startTime).String())

	startTime = time.Now()
	tree1 := &binary_search_tree.BinarySearchTree{}
	for _, v := range list {
		tree1.InsertRecursive(v, 1)
	}
	fmt.Printf("Recursive Duration: %s\n", time.Now().Sub(startTime).String())

	startTime = time.Now()
	sort.Sort(&List{list})
	fmt.Printf("Sort Duration: %s\n", time.Now().Sub(startTime).String())

}
