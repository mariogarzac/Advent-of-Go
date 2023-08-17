
package tree

import (
    "fmt"
)

type Node struct {
    value int
    right *Node
    left *Node
}

type BinaryTree struct {
    Root *Node
}

func (bt *BinaryTree) Insert(value int){
    newNode := &Node{value: value}
    if bt.Root == nil {
        bt.Root = newNode
        return
    }

    curr := bt.Root

    for {
        if value > curr.value {
            if curr.right == nil{
                curr.right = newNode
                return
            }
            curr = curr.right
        }else{
            if curr.left == nil{
                curr.left = newNode
                return
            }
            curr = curr.left
        }
    }
}

func (bt *BinaryTree) PreOrder(node *Node){
    if (node != nil){
        fmt.Printf("%d ", node.value)
        bt.PreOrder(node.left)
        bt.PreOrder(node.right)
    }
}

func (bt *BinaryTree) InOrder(node *Node, values *[]int){
    if (node != nil){
        bt.InOrder(node.left, values)
        *values = append(*values, node.value) 
        bt.InOrder(node.right, values)
    }else{
        return
    }
}

func (bt *BinaryTree) PostOrder(node *Node){
    if (node != nil){
        bt.PostOrder(node.left)
        bt.PostOrder(node.right)
        fmt.Printf("%d ", node.value)
    }
}

func (bt *BinaryTree) PrettyPrint(node *Node, level int) {
    if node == nil {
        return
    }

    bt.PrettyPrint(node.right, level+1)

    for i := 0; i < level; i++ {
        fmt.Print("    ")
    }
    fmt.Println(node.value)

    bt.PrettyPrint(node.left, level+1)
}

func (bt *BinaryTree) Find(node *Node, value int) bool {

    curr := node

    if node == nil {
        return false
    }

    if value == curr.value {
        return true
    }else{
        if value > curr.value {
            return bt.Find(node.right, value)
        }else{
            return bt.Find(node.left, value)
        }
    }
}

func (bt *BinaryTree) FindTopThree(node *Node) int {

    var values []int

    var total int
    total = 0

    bt.InOrder(node, &values)

    for i := len(values) - 3; i < len(values); i++ {
        total += values[i]
    }
    return total
}
