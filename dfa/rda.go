package dfa

import (
	"fmt"

	"github.com/t14raptor/go-fast/ast"
)

// Data Flow Analysis

type rdaContext struct {
	scopeDepth         int
	functionScopeDepth int
	scopeMaxDepth      int
	scopeStack         []*Scope
	Debug              bool
	UseDefs            []*UseDef
}

type ScopeDefs map[string][]*ScopeDef

func (s ScopeDefs) AppendScopeDefs(src ScopeDefs) {
	for key, exprs := range src {
		s[key] = append(s[key], exprs...)
	}
}

var DefCount int64

type ScopeDefType int

const (
	BlockScope ScopeDefType = iota
	FunctionScope
	Assignment
)

type ScopeDef struct {
	Val   *ast.Expression
	Depth int
	Typ   ScopeDefType
	Count int64
}

type Scope struct {
	Conditional   bool
	FunctionScope bool
	Definitions   ScopeDefs
}

func (s *Scope) AddValue(id string, v *ast.Expression, overwrite bool, typ ScopeDefType, depth int) {
	val := &ScopeDef{
		Val:   v,
		Typ:   typ,
		Depth: depth,
		Count: DefCount,
	}

	DefCount++

	if overwrite {
		s.Definitions[id] = []*ScopeDef{val}
		return
	}

	if v, ok := s.Definitions[id]; ok {
		s.Definitions[id] = append(v, val)
		return
	}

	s.Definitions[id] = []*ScopeDef{val}
}

func (s *Scope) Get(id string) ([]*ScopeDef, bool) {
	res, ok := s.Definitions[id]
	return res, ok
}

func (s *Scope) HasDef(id string, def *ScopeDef) bool {
	for _, d := range s.Definitions[id] {
		if d == def {
			return true
		}
	}

	return false
}

// NewScope creates a new scope with the following params:
// - cond: is the scope a conditional scope.
// - funcscope: is the scope a function scope.
func NewScope(cond bool, funcscope bool) *Scope {
	return &Scope{
		Conditional:   cond,
		FunctionScope: funcscope,
		Definitions:   make(ScopeDefs),
	}
}

func CreateContextRDA(maxScopeDepth int) *rdaContext {
	stk := make([]*Scope, maxScopeDepth)
	stk[0] = NewScope(false, true)
	return &rdaContext{
		scopeStack:    stk,
		scopeMaxDepth: maxScopeDepth,
	}
}

func (r *rdaContext) Start(a *ast.Program) {
	dfaVisitor := DfaVisitor{
		Ctx: r,
	}

	DefCount = 0
	a.VisitWith(&dfaVisitor)
	if r.Debug {
		fmt.Println("Definitions:", r.scopeStack[0].Definitions)
	}
}

func (r *rdaContext) pushScope(scope *Scope) {
	if r.Debug {
		fmt.Printf("Push Scope: %d->%d\n", r.scopeDepth, r.scopeDepth+1)
	}
	if r.scopeDepth >= r.scopeMaxDepth {
		panic("exceeded max scope depth")
	}
	r.scopeDepth++

	if scope.FunctionScope {
		r.functionScopeDepth = r.scopeDepth
	}

	scope.Definitions.AppendScopeDefs(r.scopeStack[r.scopeDepth-1].Definitions)

	r.scopeStack[r.scopeDepth] = scope
}

func (r *rdaContext) popScope() *Scope {
	if r.Debug {
		fmt.Printf("Pop  Scope: %d->%d\n", r.scopeDepth, r.scopeDepth-1)
	}
	if r.scopeDepth <= 0 {
		panic("can't pop further down than 0")
	}

	x := r.scopeStack[r.scopeDepth]

	r.scopeDepth--
	if x.FunctionScope {
		r.functionScopeDepth = 0

		for i := r.scopeDepth; i >= 0; i-- {
			if r.scopeStack[i].FunctionScope {
				r.functionScopeDepth = i
			}
		}
	}

	if r.Debug {
		fmt.Println("Scope Defs:", x.Definitions)
	}
	return x
}

// mergeDown will merge defintions from the scope "a" into the current scope.
func (r *rdaContext) mergeDown(scopeDepth int, a *Scope) {
	for id, vals := range a.Definitions {
		currentVals := r.scopeStack[r.scopeDepth].Definitions[id]
		carryVals := []*ScopeDef{}
		for _, val := range vals {
			switch val.Typ {
			case BlockScope:
				continue
			case FunctionScope:
				if a.FunctionScope {
					continue
				}

				if !r.scopeStack[r.scopeDepth].HasDef(id, val) {
					carryVals = append(carryVals, val)
				}
			case Assignment:
				// Assignment depths are found when an assignment is hit.
				if val.Depth < scopeDepth && !r.scopeStack[r.scopeDepth].HasDef(id, val) {
					// merge down if it's a new assignment and the val expiration depth is lower or equal to current scope depth.
					carryVals = append(carryVals, val)
				}
			}
		}

		// Ensure the slice is initialized before appending
		if currentVals == nil {
			// Add default undefined value because no values existed before.
			// TODO: Make undefined block
			r.scopeStack[r.scopeDepth].Definitions[id] = append(carryVals, &ScopeDef{
				Val: nil,
				Typ: FunctionScope,
			})
			continue
		}

		r.scopeStack[r.scopeDepth].Definitions[id] = append(
			currentVals,
			carryVals...,
		)
	}

}

// mergeSameDepth merges defintions from scope A and scope B, storing the definitions in scope A.
func (r *rdaContext) mergeSameDepth(a, b *Scope) {
	for id, defs := range b.Definitions {
		orig, needsCheck := a.Definitions[id]
		if needsCheck {
			original := []*ScopeDef{}

			for _, def := range defs {
				if !a.HasDef(id, def) {
					original = append(original, def)
				}
			}

			a.Definitions[id] = append(orig, original...)
		} else {
			a.Definitions[id] = defs
		}
	}
}

func (r *rdaContext) cutExpiring(defs []*ScopeDef, currentDepth int) []*ScopeDef {
	result := []*ScopeDef{}
	for _, v := range defs {
		// If the depth is
		if v.Typ != BlockScope && currentDepth >= v.Depth {
			result = append(result, v)
		}
	}

	return result
}
