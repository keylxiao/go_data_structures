package lib

import (
    "errors"
)

// BinaryTree 二叉树
type BinaryTree struct {
    value       ElemType    // 节点储存值
    left, right *BinaryTree // 左右子树
}

// InitBinaryTree 初始化二叉树
func InitBinaryTree(value ElemType) *BinaryTree {
    tree := new(BinaryTree)
    tree.value = value
    return tree
    //return new(BinaryTree)
}

// Insert 插入节点
func (tree *BinaryTree) Insert(value ElemType) {
    // 空树
    if tree.value == 0 {
        tree.value = value
    } else {
        // 非空树
        if value > tree.value {
            if tree.right == nil {
                tree.right = new(BinaryTree)
            }
            tree.right.Insert(value)
        } else {
            if tree.left == nil {
                tree.left = new(BinaryTree)
            }
            tree.left.Insert(value)
        }
    }
}

// Value 返回节点值
func (tree *BinaryTree) Value() ElemType {
    return tree.value
}

// Height 返回树的高度
func (tree *BinaryTree) Height() int {
    if tree == nil {
        return 0
    }
    leftHeight := tree.left.Height()
    rightHeight := tree.right.Height()
    if leftHeight > rightHeight {
        return leftHeight + 1
    } else {
        return rightHeight + 1
    }
}

// Left 返回左子树
func (tree *BinaryTree) Left() *BinaryTree {
    return tree.left
}

// Right 返回右子树
func (tree *BinaryTree) Right() *BinaryTree {
    return tree.right
}

// FindNode 寻找节点
func (tree *BinaryTree) FindNode(value ElemType) (*BinaryTree, error) {
    if tree == nil {
        return nil, errors.New("不存在该节点")
    }
    if value == tree.value {
        return tree, nil
    } else {
        if value > tree.value {
            return tree.right.FindNode(value)
        } else {
            return tree.left.FindNode(value)
        }
    }
}

// PreOrder 前序遍历
func (tree *BinaryTree) PreOrder(array *[]ElemType) *[]ElemType {
    if tree != nil {
        *array = append(*array, tree.value)
        tree.left.PreOrder(array)
        tree.right.PreOrder(array)
    }
    return array
}

// InOrder 中序遍历
func (tree *BinaryTree) InOrder(array *[]ElemType) *[]ElemType {
    if tree != nil {
        tree.left.PreOrder(array)
        *array = append(*array, tree.value)
        tree.right.PreOrder(array)
    }
    return array
}

// PostOrder 后序遍历
func (tree *BinaryTree) PostOrder(array *[]ElemType) *[]ElemType {
    if tree != nil {
        tree.left.PreOrder(array)
        tree.right.PreOrder(array)
        *array = append(*array, tree.value)
    }
    return array
}

// GrandOrder 广度优先遍历(层先遍历)
func (tree *BinaryTree) GrandOrder(array *[][]ElemType) {
    if tree == nil {
        return
    }
    GrandBfs(tree, 0, array)
}

func GrandBfs(tree *BinaryTree, level int, array *[][]ElemType) {
    if tree == nil {
        return
    }
    // 补充切片
    if len(*array) < level+1 {
        *array = append(*array, make([]ElemType, 0))
    }
    (*array)[level] = append((*array)[level], tree.value)
    GrandBfs(tree.left, level+1, array)
    GrandBfs(tree.right, level+1, array)
}