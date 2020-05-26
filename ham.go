package ham

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
)

func WithContext(ctx context.Context) context.Context {
	return ctx
}

var defaultExecutor = newExecutor()

func getExecutor(ctx context.Context) *executor {
	return defaultExecutor
}

func newExecutor() *executor {
	return &executor{}
}

func (m *executor) registerNumber(n *Number) {
	varName := indexToVarName(m.lastAutoVar)
	m.lastAutoVar++
	n.varName = varName
	m.variables = append(m.variables, n)
}

func (m *executor) registerEdge(n *edge) {
	m.edges = append(m.edges, n)
}

type executor struct {
	variables []*Number
	edges []*edge

	lastAutoVar int
}

type Number struct {
	b decimal.Decimal
	varName string
}

type edge struct {
	left *Number
	right *Number
	result *Number
	op string
}

func Int64(ctx context.Context, x int64) *Number {

	n := &Number{
		b: decimal.NewFromInt(x),
	}
	getExecutor(ctx).registerNumber(n)
	return n
}

func (m *Number) Pow(ctx context.Context, exp *Number) *Number {
	n := &Number{b:m.b.Pow(exp.b)}
	getExecutor(ctx).registerEdge(&edge{
		left: m,
		right: exp,
		result: n,
		op:"^",
	})
	getExecutor(ctx).registerNumber(n)
	return n
}

const varLetters = "abcdefghijklmnopqrstuvwxyz"
func Print(ctx context.Context) {
	e := getExecutor(ctx)
	hitNumbers := map[*Number]bool{}
	for _, edge := range e.edges {
		hitNumbers[edge.result] = true
	}
	for _, num := range e.variables {
		if !hitNumbers[num] {

			fmt.Printf("%s = %s\n", num.varName, num.b.String())
		}
	}
	for _, num := range e.variables {
		if !hitNumbers[num] {
			continue
		}
		var thisEdge *edge
		for _, edge := range e.edges {
			if edge.result == num {
				thisEdge = edge
				break
			}
		}
		edgeOp := fmt.Sprintf("(%s %s %s)", thisEdge.left.varName, thisEdge.op, thisEdge.right.varName)
		fmt.Printf("%s = %s\n", num.varName, edgeOp)
	}
}

func PrintSingleLine(ctx context.Context) {
	e := getExecutor(ctx)
	finalNumbers := map[*Number]bool{}
	for _, edge := range e.edges {
		finalNumbers[edge.left] = true
		finalNumbers[edge.right] = true
	}
	for _, n := range e.variables {
		if finalNumbers[n] {
			continue
		}
		fmt.Printf("%s = %s\n",n.varName, e.edgePrint(n))
	}
}

func (m *executor) edgePrint(n *Number) string {
	var thisEdge *edge
	for _, edge := range m.edges {
		if edge.result == n {
			thisEdge = edge
			break
		}
	}
	var leftEdge *edge
	var rightEdge *edge
	for _, e := range m.edges {
		if e.result == thisEdge.left {
			leftEdge = e
			continue
		}
		if e.result == thisEdge.right {
			rightEdge = e
			continue
		}
	}
	leftStr := ""
	rightStr := ""
	if leftEdge == nil {
		leftStr = thisEdge.left.b.String()
	} else {
		leftStr = m.edgePrint(thisEdge.left)
	}
	if rightEdge == nil {
		rightStr = thisEdge.right.b.String()
	} else {
		rightStr = m.edgePrint(thisEdge.right)
	}
	return fmt.Sprintf("(%s %s %s)", leftStr, thisEdge.op, rightStr)
}

func indexToVarName(i int) string {
	b := ""
	for {
		b += string(varLetters[i % len(varLetters)])
		if i > len(varLetters) {
			i = i - len(varLetters)
		} else {
			break
		}
	}
	return b
}