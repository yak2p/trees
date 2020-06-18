## 动态查找树实现

该项目实现了以下几个动态查找树：

* [二叉搜索树BST](./binary-search-tree)
* [平衡二叉树AVL](./avl-tree)
* [红黑树](./red-black-tree)

代码实现主要从易于理解角度出发，因此还存在一些优化空间。

实现的思路可以参考文章： [动态查找树](https://yesphet.github.io/posts/平衡树/)



下面给出一些插入数据测试数据，测试方法在[test目录](./test)下

```markdown
* 1-100000，按随机顺序插入：

Insert 100000 nodes elapse:
         BST Non-Recursive :            35.790168ms
         BST Recursive :                42.092132ms
         AVL tree :                     53.018442ms
         Red-Black tree :               41.5808ms
         sort.Sort :                    20.590467ms


* 1-100000，按从小到大有序的插入：

Insert 100000 nodes elapse:
         BST Non-Recursive :            13.520115515s
         BST Recursive :                31.70352614s
         AVL tree :                     23.767831ms
         Red-Black tree :               29.698382ms
         sort.Sort :                    8.535786ms
```

可以发现在随机插入的情况下，性能排序为 二叉搜索树非递归实现 > 红黑树 >二叉搜索树递归实现 > AVL树。

在顺序插入的情况下，性能排序为 AVL树 > 红黑树 > 二叉搜索树非递归实现 > 二叉搜索树递归实现。

因此综合来看，红黑树更适合于用来建立内存索引。