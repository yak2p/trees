package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	avl_tree "github.com/yak2p/trees/avl-tree"
	"github.com/yak2p/trees/binary-search-tree"
	red_black_tree "github.com/yak2p/trees/red-black-tree"
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

	maxNum := 100000
	list := rand.Perm(maxNum)

	//list := make([]int, maxNum)
	//for i := 0; i < maxNum; i++ {
	//	list[i] = i
	//}

	fmt.Printf("Insert %d nodes elapse:\n ", maxNum)

	startTime := time.Now()
	tree := &binary_search_tree.BinarySearchTree{}
	for _, v := range list {
		tree.Insert(v, 1)
	}
	fmt.Printf("\t BST Non-Recursive : \t\t%s\n", time.Now().Sub(startTime).String())

	startTime = time.Now()
	tree1 := &binary_search_tree.BinarySearchTree{}
	for _, v := range list {
		tree1.InsertRecursive(v, 1)
	}
	fmt.Printf("\t BST Recursive : \t\t%s\n", time.Now().Sub(startTime).String())

	startTime = time.Now()
	tree2 := &avl_tree.Tree{}
	for _, v := range list {
		tree2.Insert(v, 1)
	}
	fmt.Printf("\t AVL tree : \t\t\t%s\n", time.Now().Sub(startTime).String())

	startTime = time.Now()
	tree3 := &red_black_tree.Tree{}
	for _, v := range list {
		tree3.Insert(v, 1)
	}
	fmt.Printf("\t Red-Black tree : \t\t%s\n", time.Now().Sub(startTime).String())

	startTime = time.Now()
	sort.Sort(&List{list})
	fmt.Printf("\t sort.Sort : \t\t\t%s\n", time.Now().Sub(startTime).String())

}
