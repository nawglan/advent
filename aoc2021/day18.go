package main

import (
	"fmt"
	"math"
	p "advent/util/parse"
)

type SnailFish struct {
	value     int
	depth     int
	magnitude int

	left   *SnailFish
	right  *SnailFish
	parent *SnailFish
}

// returns the number to the left of the stop node
//
//      ,
//     , 5
//    2 ,
//     3 4
// numberLeft(node(5)) returns node(4)
// numberLeft(node(4)) returns node(3)
// numberLeft(node(3)) returns node(2)
// numberLeft(node(2)) returns nil
//
// return the right most node of the left tree

func numberLeft(start, stop *SnailFish) *SnailFish {
	stack := []*SnailFish{}
	curr := start
	var last *SnailFish

	for len(stack) > 0 || curr != nil {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if curr == stop {
				return last
			} else {
				if curr.value != -1 {
					last = curr
				}
			}
			curr = curr.right
		}
	}

	return nil
}

// returns the number to the left of the stop node
//
//      ,
//     , 5
//    2 ,
//     3 4
// numberRight(node(5)) returns nil
// numberRight(node(4)) returns node(5)
// numberRight(node(3)) returns node(4)
// numberRight(node(2)) returns node(3)
//
// return the left most node of the right tree

func numberRight(start, stop *SnailFish) *SnailFish {
	stack := []*SnailFish{}
	curr := start
	var last *SnailFish

	for len(stack) > 0 || curr != nil {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.right
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if curr == stop {
				return last
			} else {
				if curr.value != -1 {
					last = curr
				}
			}
			curr = curr.left
		}
	}

	return nil
}

func (sn *SnailFish) dump() string {
	out := ""

	if sn.value == -1 {
		out = fmt.Sprintf("[%s,%s]", sn.left.dump(), sn.right.dump())
	} else {
		out = fmt.Sprintf("%d", sn.value)
	}

	return out
}

func (sn *SnailFish) updateMagnitude() int {
	if sn.value == -1 {
		sn.magnitude = 3*sn.left.updateMagnitude() + 2*sn.right.updateMagnitude()
	} else {
		sn.magnitude = sn.value
	}

	return sn.magnitude
}

func (sn *SnailFish) updateDepth() {
	if sn != nil {
		if sn.parent != nil {
			sn.depth = sn.parent.depth + 1
		} else {
			sn.depth = 0
		}
		if sn.left != nil {
			sn.left.updateDepth()
		}
		if sn.right != nil {
			sn.right.updateDepth()
		}
	}
}

//
// explodes a , node
//           9 8
func (sn *SnailFish) explode() *SnailFish {
	root := sn
	for root.parent != nil {
		root = root.parent
	}
	if sn.left != nil {
		if numLeft := numberLeft(root, sn.left); numLeft != nil {
			numLeft.value += sn.left.value
		}
		sn.left.parent = nil
		sn.left = nil
	}
	if sn.right != nil {
		if numRight := numberRight(root, sn.right); numRight != nil {
			numRight.value += sn.right.value
		}
		sn.right.parent = nil
		sn.right = nil
	}
	sn.value = 0
	return sn
}

/*
func (sn *SnailFish) checkExplode() bool {
	if sn.value == -1 {
		if sn.left.value != -1 && sn.right.value != -1 {
			if sn.depth > 3 {
				sn = sn.explode()
				return true
			}
		} else {
			if sn.left != nil && sn.left.value == -1 {
				return sn.left.checkExplode()
			}
			if sn.right != nil && sn.right.value == -1 {
				return sn.right.checkExplode()
			}
		}
	}
	return false
}
*/

func (sn *SnailFish) checkExplode() bool {
	stack := []*SnailFish{}
	curr := sn

	for len(stack) > 0 || curr != nil {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if curr.value == -1 && curr.depth >= 4 {
				curr.explode()
				return true
			}
			curr = curr.right
		}
	}

	return false
}

func split(sn *SnailFish) *SnailFish {
	sn.left = &SnailFish{
		value:  int(math.Floor(float64(sn.value) / 2)),
		depth:  sn.depth + 1,
		parent: sn,
	}

	sn.right = &SnailFish{
		value:  int(math.Ceil(float64(sn.value) / 2)),
		depth:  sn.depth + 1,
		parent: sn,
	}

	sn.value = -1

	return sn
}

func (sn *SnailFish) checkSplit() bool {
	stack := []*SnailFish{}
	curr := sn

	for len(stack) > 0 || curr != nil {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if curr.value >= 10 {
				curr = split(curr)
				curr.updateDepth()
				return true
			}
			curr = curr.right
		}
	}

	return false
}

func (sn *SnailFish) reduce() bool {
	result := false

	// check for explode
	if result = sn.checkExplode(); result {
	} else {
		sn.updateDepth()
		if result = sn.checkSplit(); result {
		}
		sn.updateDepth()
	}

	return result
}

func add(a, b *SnailFish) *SnailFish {
	result := &SnailFish{
		value: -1,
		depth: 0,
		left:  a,
		right: b,
	}

	result.left.parent = result
	result.right.parent = result

	result.left.updateDepth()
	result.right.updateDepth()

	for result.reduce() {
	}

	return result
}

func (sn *SnailFish) updateParents() {
	if sn != nil {
		if sn.left != nil {
			sn.left.parent = sn
			sn.left.updateParents()
		}
		if sn.right != nil {
			sn.right.parent = sn
			sn.right.updateParents()
		}
	}
}

func clone(sn *SnailFish) *SnailFish {
	ret := &SnailFish{
		value:     sn.value,
		depth:     sn.depth,
		magnitude: sn.magnitude,
	}

	if sn.left != nil {
		ret.left = clone(sn.left)
	}
	if sn.right != nil {
		ret.right = clone(sn.right)
	}

	ret.updateParents()

	return ret
}

func day18(puzzle_data []string) {
	numbers := []*SnailFish{}

	for _, dataVal := range puzzle_data {
		output := []string{}
		operators := []string{}

		for i := 0; i < len(dataVal); i++ {
			token := dataVal[i]
			switch token {
			case '[':
				operators = append(operators, string(token))
			case ']':
				if len(operators) > 0 {
					for operators[len(operators)-1] != "[" {
						output = append(output, operators[len(operators)-1])
						operators = operators[0 : len(operators)-1]
						if len(operators) == 0 {
							break
						}
					}
					if len(operators) == 0 || operators[len(operators)-1] != "[" {
						panic("there should be a left bracket here")
					} else {
						// discard left bracket
						operators = operators[0 : len(operators)-1]
					}
				}
			case ',': // make number
				if len(operators) > 0 {
					for operators[len(operators)-1] != "[" {
						output = append(output, operators[len(operators)-1])
						operators = operators[0 : len(operators)-1]
					}
					operators = append(operators, string(token))
				}
			default: // number
				output = append(output, string(token))
			}
		}

		stack := []*SnailFish{}
		for _, token := range output {
			if token == "," {
				left := stack[len(stack)-2]
				right := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				node := &SnailFish{value: -1}
				right.parent = node
				left.parent = node
				node.left = left
				node.right = right

				stack = append(stack, node)
			} else {
				node := &SnailFish{
					value: p.MakeInt(token),
				}
				node.updateDepth()
				stack = append(stack, node)
			}
		}

		if len(stack) > 0 {
			numbers = append(numbers, stack[0])
		}
	}

	sum := add(clone(numbers[0]), clone(numbers[1]))
	for i := 2; i < len(numbers); i++ {
		sum = add(clone(sum), clone(numbers[i]))
	}
	fmt.Printf("Day 18 (part 1) The magnitude of the result is %d\n", sum.updateMagnitude())

	magnitude := -1
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j {
				result1 := add(clone(numbers[i]), clone(numbers[j]))
				result1.updateMagnitude()
				result2 := add(clone(numbers[j]), clone(numbers[i]))
				result2.updateMagnitude()
				if result1.magnitude > magnitude {
					magnitude = result1.magnitude
				}
				if result2.magnitude > magnitude {
					magnitude = result2.magnitude
				}
			}
		}
	}

	fmt.Printf("Day 18 (part 2) The maximum magnitude is %d\n", magnitude)
}
