package main

import (
	"fmt"
	"math"
	"strconv"
)

type Op int

const (
	OpNon Op = iota
	OpAdd    // 加
	OpSub    // 减
	OpMul    // 乘
	OpDiv    // 除法
)

var opFuncMap = map[Op]func(a, b float64) float64{
	OpAdd: func(a, b float64) float64 {
		return a + b
	},
	OpSub: func(a, b float64) float64 {
		return a - b
	},
	OpMul: func(a, b float64) float64 {
		return a * b
	},
	OpDiv: func(a, b float64) float64 {
		return a / b
	},
}

var opStrMap = map[Op]string{
	OpAdd: "+",
	OpSub: "-",
	OpMul: "*",
	OpDiv: "/",
}

type BinOp struct {
	Left  Expr
	Right Expr
	Op    Op
}

func (e BinOp) Eval(ec *EvalContext) float64 {
	left := e.Left.Eval(ec)
	right := e.Right.Eval(ec)
	f := opFuncMap[e.Op]
	if f == nil {
		return math.NaN()
	}
	return f(left, right)
}

func (e BinOp) ExprString() string {
	left := e.Left.ExprString()
	right := e.Right.ExprString()
	return fmt.Sprintf("(%s %s %s)", left, opStrMap[e.Op], right)
}

type EvalContext struct {
	unknownMap map[string]float64
}

func newEvalContent() *EvalContext {
	ec := &EvalContext{}
	ec.unknownMap = make(map[string]float64)
	return ec
}

func (ec *EvalContext) setUnknown(name string, val float64) {
	ec.unknownMap[name] = val
}

type Expr interface {
	Eval(ec *EvalContext) float64
	ExprString() string
}

type Unknown struct {
	name string
}

func (v *Unknown) GetName() string {
	return v.name
}

func (v *Unknown) ExprString() string {
	return v.name
}

func (v *Unknown) Eval(ec *EvalContext) float64 {
	result, ok := ec.unknownMap[v.name]
	if ok {
		return result
	}
	return math.NaN()
}

type AbsNum int

func (n AbsNum) Eval(ec *EvalContext) float64 {
	return float64(n)
}

func (n AbsNum) ExprString() string {
	return strconv.Itoa(int(n))
}

func add(a, b Expr) Expr {
	return BinOp{
		Left:  a,
		Op:    OpAdd,
		Right: b,
	}
}

func sub(a, b Expr) Expr {
	return BinOp{
		Left:  a,
		Op:    OpSub,
		Right: b,
	}
}

func mul(a, b Expr) Expr {
	return BinOp{
		Left:  a,
		Op:    OpMul,
		Right: b,
	}
}

func div(a, b Expr) Expr {
	return BinOp{
		Left:  a,
		Op:    OpDiv,
		Right: b,
	}
}
