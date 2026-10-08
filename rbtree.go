package main

import (
	"fmt"
	"strconv"
	"strings"
)

type RBTColor int

const (
	BLACK RBTColor = iota
	RED
)

type RBTNode struct {
	Data   int
	Color  RBTColor
	Left   *RBTNode
	Right  *RBTNode
	Parent *RBTNode
}

type RBTree struct {
	Root *RBTNode
	Nil  *RBTNode
}

func initTree() *RBTree {
	nilNode := &RBTNode{Data: 0, Color: BLACK, Left: nil, Right: nil, Parent: nil}
	return &RBTree{Root: nilNode, Nil: nilNode}
}

func (t *RBTree) destroy() {
	// В Go сборщик мусора сам освободит память
	t.Root = nil
	t.Nil = nil
}

func (t *RBTree) rotateLeft(x *RBTNode) {
	y := x.Right
	x.Right = y.Left
	if y.Left != t.Nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == t.Nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

func (t *RBTree) rotateRight(x *RBTNode) {
	y := x.Left
	x.Left = y.Right
	if y.Right != t.Nil {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == t.Nil {
		t.Root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}
	y.Right = x
	x.Parent = y
}

func (t *RBTree) fixInsert(k *RBTNode) {
	for k.Parent.Color == RED {
		if k.Parent == k.Parent.Parent.Left {
			u := k.Parent.Parent.Right
			if u.Color == RED {
				k.Parent.Color = BLACK
				u.Color = BLACK
				k.Parent.Parent.Color = RED
				k = k.Parent.Parent
			} else {
				if k == k.Parent.Right {
					k = k.Parent
					t.rotateLeft(k)
				}
				k.Parent.Color = BLACK
				k.Parent.Parent.Color = RED
				t.rotateRight(k.Parent.Parent)
			}
		} else {
			u := k.Parent.Parent.Left
			if u.Color == RED {
				k.Parent.Color = BLACK
				u.Color = BLACK
				k.Parent.Parent.Color = RED
				k = k.Parent.Parent
			} else {
				if k == k.Parent.Left {
					k = k.Parent
					t.rotateRight(k)
				}
				k.Parent.Color = BLACK
				k.Parent.Parent.Color = RED
				t.rotateLeft(k.Parent.Parent)
			}
		}
	}
	t.Root.Color = BLACK
}

func (t *RBTree) insertNode(key int) {
	node := &RBTNode{Data: key, Color: RED, Left: t.Nil, Right: t.Nil, Parent: t.Nil}
	y := t.Nil
	x := t.Root

	for x != t.Nil {
		y = x
		if node.Data == x.Data {
			return // дубликат
		}
		if node.Data < x.Data {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	node.Parent = y
	if y == t.Nil {
		t.Root = node
	} else if node.Data < y.Data {
		y.Left = node
	} else {
		y.Right = node
	}

	if node.Parent == t.Nil {
		node.Color = BLACK
		return
	}
	if node.Parent.Parent == t.Nil {
		return
	}
	t.fixInsert(node)
}

func (t *RBTree) transplant(u, v *RBTNode) {
	if u.Parent == t.Nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Parent = u.Parent
}

func (t *RBTree) minimumNode(node *RBTNode) *RBTNode {
	for node.Left != t.Nil {
		node = node.Left
	}
	return node
}

func (t *RBTree) fixDelete(x *RBTNode) {
	for x != t.Root && x.Color == BLACK {
		if x == x.Parent.Left {
			s := x.Parent.Right
			if s.Color == RED {
				s.Color = BLACK
				x.Parent.Color = RED
				t.rotateLeft(x.Parent)
				s = x.Parent.Right
			}
			if s.Left.Color == BLACK && s.Right.Color == BLACK {
				s.Color = RED
				x = x.Parent
			} else {
				if s.Right.Color == BLACK {
					s.Left.Color = BLACK
					s.Color = RED
					t.rotateRight(s)
					s = x.Parent.Right
				}
				s.Color = x.Parent.Color
				x.Parent.Color = BLACK
				s.Right.Color = BLACK
				t.rotateLeft(x.Parent)
				x = t.Root
			}
		} else {
			s := x.Parent.Left
			if s.Color == RED {
				s.Color = BLACK
				x.Parent.Color = RED
				t.rotateRight(x.Parent)
				s = x.Parent.Left
			}
			if s.Right.Color == BLACK && s.Left.Color == BLACK {
				s.Color = RED
				x = x.Parent
			} else {
				if s.Left.Color == BLACK {
					s.Right.Color = BLACK
					s.Color = RED
					t.rotateLeft(s)
					s = x.Parent.Left
				}
				s.Color = x.Parent.Color
				x.Parent.Color = BLACK
				s.Left.Color = BLACK
				t.rotateRight(x.Parent)
				x = t.Root
			}
		}
	}
	x.Color = BLACK
}

func (t *RBTree) deleteNode(key int) {
	z := t.Root
	for z != t.Nil && z.Data != key {
		if key < z.Data {
			z = z.Left
		} else {
			z = z.Right
		}
	}
	if z == t.Nil {
		return
	}

	y := z
	var x *RBTNode
	yOriginalColor := y.Color

	if z.Left == t.Nil {
		x = z.Right
		t.transplant(z, z.Right)
	} else if z.Right == t.Nil {
		x = z.Left
		t.transplant(z, z.Left)
	} else {
		y = t.minimumNode(z.Right)
		yOriginalColor = y.Color
		x = y.Right
		if y.Parent == z {
			x.Parent = y
		} else {
			t.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		t.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}
	if yOriginalColor == BLACK {
		t.fixDelete(x)
	}
}

func (t *RBTree) searchNode(key int) bool {
	curr := t.Root
	for curr != t.Nil {
		if key == curr.Data {
			return true
		}
		if key < curr.Data {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return false
}

func (t *RBTree) printHelper(root *RBTNode, indent string, last bool) {
	if root == t.Nil {
		return
	}
	fmt.Print(indent)
	if last {
		fmt.Print("R----")
		indent += "     "
	} else {
		fmt.Print("L----")
		indent += "|    "
	}
	color := "(BLK)"
	if root.Color == RED {
		color = "(RED)"
	}
	fmt.Printf("%d %s\n", root.Data, color)
	t.printHelper(root.Left, indent, false)
	t.printHelper(root.Right, indent, true)
}

func (t *RBTree) print() {
	if t.Root == t.Nil {
		fmt.Println("[Пустое дерево]")
		return
	}
	t.printHelper(t.Root, "", true)
}

func (t *RBTree) extractHelper(node *RBTNode, outItems *[]string) {
	if node != t.Nil {
		*outItems = append(*outItems, strconv.Itoa(node.Data))
		t.extractHelper(node.Left, outItems)
		t.extractHelper(node.Right, outItems)
	}
}

func (t *RBTree) extractElements() []string {
	items := []string{}
	t.extractHelper(t.Root, &items)
	return items
}

// вспомогательная — для сравнения слайсов
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// чтобы не было предупреждения о неиспользуемом импорте strings
var _ = strings.TrimSpace