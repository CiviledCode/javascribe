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

var Undefined = &ScopeDef{
	Val:   nil,
	Typ:   FunctionScope,
	Count: -1,
}

type ScopeDefType int

const (
	BlockScope ScopeDefType = iota
	FunctionScope
	GlobalScope
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

// NewScope creates a new scope.
// cond depicts if the scope is a conditional scope.
// funcscope depicts if the scope is a function scope.
func NewScope(cond bool, funcscope bool) *Scope {
	return &Scope{
		Conditional:   cond,
		FunctionScope: funcscope,
		Definitions:   make(ScopeDefs),
	}
}

// AddValue adds a definition to the scope.
// id denotes the identifier.
// v denotes the expression.
// overwrite denotes if the value overwrites previous declarations in this scope.
// typ denotes the type of definition (block, function, global)
// depth is the depth that the declaration expires at.
func (s *Scope) AddValue(id string, v *ast.Expression, overwrite bool, typ ScopeDefType, depth int) {
	val := &ScopeDef{
		Val:   v,
		Typ:   typ,
		Depth: depth,
		Count: DefCount,
	}

	if v == nil {
		val = Undefined
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

// Get retrieves a list of definitions for an identifier in that scope.
func (s *Scope) Get(id string) ([]*ScopeDef, bool) {
	res, ok := s.Definitions[id]
	return res, ok
}

// HasDef determines if a ScopeDef exists in this scope or not.
func (s *Scope) HasDef(id string, def *ScopeDef) bool {
	for _, d := range s.Definitions[id] {
		if d == def {
			return true
		}
	}

	return false
}

// RemoveParentDefs removes all ScopeDefs that exist in the parent scope.
// This prevents propogating definitions that already exist into the parent scope.
func (s *Scope) RemoveParentDefs(parentScope *Scope) {
	for id, defs := range s.Definitions {
		for idx, def := range defs {
			if parentScope.HasDef(id, def) {
				s.Definitions[id][idx] = nil
			}
		}
	}
}

// MergeSameDepth merges defintions from scope A and scope B, storing the definitions in scope A.
func (s *Scope) MergeSameDepth(b *Scope) {
	for id, defs := range b.Definitions {
		s.MergeDefs(defs, id)
	}
}

// MergeDefs merges definitions for a given identifier into this scope.
// This ensures values are not duplicated when merged.
func (s *Scope) MergeDefs(defs []*ScopeDef, id string) {
	orig, needsCheck := s.Definitions[id]
	if needsCheck {
		original := []*ScopeDef{}

		for _, def := range defs {
			if def == nil {
				continue
			}

			if !s.HasDef(id, def) {
				original = append(original, def)
			}
		}

		s.Definitions[id] = append(orig, original...)
	} else {
		s.Definitions[id] = defs
	}
}

// AddUndefined adds Undefined objects to all declarations within this scope.
// parentScope is used to determine which values were propogated from the parent scope.
func (s *Scope) AddUndefined(parentScope *Scope) {
	for id, x := range s.Definitions {
		// If no value exists in the parent scope, then add undefined.
		if _, found := parentScope.Get(id); !found && !s.HasDef(id, Undefined) {
			s.Definitions[id] = append(x, Undefined)
		}
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
func (r *rdaContext) mergeDown(scopeDepth int, a *Scope, conditional bool) {
	parentScope := r.scopeStack[r.scopeDepth]
outer:
	for id, vals := range a.Definitions {

		currentVals := parentScope.Definitions[id]
		carryVals := []*ScopeDef{}

		for _, val := range vals {
			if val == nil {
				continue
			}

			switch val.Typ {
			case BlockScope:
				// The value hasn't expired scope.
				if r.scopeDepth < val.Depth {
					continue
				}

				// Append definition to scope if it's
				if !parentScope.HasDef(id, val) || !conditional {
					carryVals = append(carryVals, val)
				}
			case FunctionScope:
				if a.FunctionScope && r.scopeDepth < val.Depth {
					continue
				}

				if !parentScope.HasDef(id, val) || !conditional {
					carryVals = append(carryVals, val)
				}
			case GlobalScope:
				if !parentScope.HasDef(id, val) || !conditional {
					carryVals = append(carryVals, val)
				}
			}
		}

		// Ensure the slice is initialized before appending
		if currentVals == nil {
			for _, def := range carryVals {
				if def == Undefined {
					parentScope.Definitions[id] = carryVals
					continue outer
				}
			}

			parentScope.Definitions[id] = carryVals
			continue
		}

		if !conditional {
			parentScope.Definitions[id] = carryVals
		} else {
			parentScope.Definitions[id] = append(
				currentVals,
				carryVals...,
			)
		}
	}

}

func (r *rdaContext) findNotExpiring(s *Scope, id string, blockParents bool) []*ScopeDef {
	result := []*ScopeDef{}
	for _, val := range s.Definitions[id] {
		if val.Typ == GlobalScope {
			result = append(result, val)
			continue
		}

		if blockParents && r.scopeStack[r.scopeDepth].HasDef(id, val) {
			continue
		}

		if r.scopeDepth < val.Depth {
			continue
		}

		result = append(result, val)
	}

	return result
}
