package dfa

import (
	"fmt"

	"github.com/t14raptor/go-fast/ast"
)

type DfaVisitor struct {
	Ctx *rdaContext
}

func (lv *DfaVisitor) VisitArrayLiteral(n *ast.ArrayLiteral) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitArrayPattern(n *ast.ArrayPattern) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitArrowFunctionLiteral(n *ast.ArrowFunctionLiteral) {
	n.VisitChildrenWith(lv)
}

func (lv *DfaVisitor) VisitAssignExpression(n *ast.AssignExpression) {
	currentScope := lv.Ctx.scopeStack[lv.Ctx.scopeDepth]
	id := n.Left.Expr.(*ast.Identifier).Name
	foundDepth := 0
	conditional := currentScope.Conditional
	typ := GlobalScope
outer:
	for i := lv.Ctx.scopeDepth; i >= 0; i-- {
		if lv.Ctx.scopeStack[i].Conditional {
			conditional = true
		}

		if f, ok := lv.Ctx.scopeStack[i].Definitions[id]; ok {
			for _, x := range f {
				typ = x.Typ
				foundDepth = x.Depth
				break outer
			}
		}
	}

	currentScope.AddValue(id, n.Right, !conditional, typ, foundDepth)
}

func (lv *DfaVisitor) VisitAwaitExpression(n *ast.AwaitExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitBadStatement(n *ast.BadStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitBinaryExpression(n *ast.BinaryExpression) {
	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitBindingTarget(n *ast.BindingTarget) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitBlockStatement(n *ast.BlockStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitBooleanLiteral(n *ast.BooleanLiteral) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitBreakStatement(n *ast.BreakStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitCallExpression(n *ast.CallExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitCaseStatement(n *ast.CaseStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitCaseStatements(n *ast.CaseStatements) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitCatchStatement(n *ast.CatchStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitClassDeclaration(n *ast.ClassDeclaration) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitClassElement(n *ast.ClassElement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitClassElements(n *ast.ClassElements) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitClassLiteral(n *ast.ClassLiteral) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitClassStaticBlock(n *ast.ClassStaticBlock) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitComputedProperty(n *ast.ComputedProperty) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitConciseBody(n *ast.ConciseBody) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitConditionalExpression(n *ast.ConditionalExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitContinueStatement(n *ast.ContinueStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitDebuggerStatement(n *ast.DebuggerStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitDoWhileStatement(n *ast.DoWhileStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitEmptyStatement(n *ast.EmptyStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitExpression(n *ast.Expression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitExpressionStatement(n *ast.ExpressionStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitExpressions(n *ast.Expressions) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitFieldDefinition(n *ast.FieldDefinition) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitForInStatement(n *ast.ForInStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitForInto(n *ast.ForInto) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitForLoopInitializer(n *ast.ForLoopInitializer) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitForOfStatement(n *ast.ForOfStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitForStatement(n *ast.ForStatement) {
	forScope := NewScope(true, false)
	lv.Ctx.pushScope(forScope)
	n.VisitChildrenWith(lv)
	lv.Ctx.popScope()

	lv.Ctx.mergeDown(lv.Ctx.scopeDepth+1, forScope)
}

func (lv *DfaVisitor) VisitFunctionDeclaration(n *ast.FunctionDeclaration) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitFunctionLiteral(n *ast.FunctionLiteral) {
	n.VisitChildrenWith(lv)
}

func (lv *DfaVisitor) VisitIdentifier(n *ast.Identifier) {
	if n.Name != "log" {
		defs := lv.Ctx.scopeStack[lv.Ctx.scopeDepth].Definitions[n.Name]

		ud := &UseDef{
			Usage:       n,
			Definitions: defs,
		}

		lv.Ctx.UseDefs = append(lv.Ctx.UseDefs, ud)
	}
	n.VisitChildrenWith(lv)
}

func (lv *DfaVisitor) VisitIfStatement(n *ast.IfStatement) {
	ifScope := NewScope(true, false)

	lv.Ctx.pushScope(ifScope)
	lv.VisitExpression(n.Test)
	lv.VisitStatement(n.Consequent)
	lv.Ctx.popScope()

	//currDepth := lv.Ctx.scopeDepth + 1
	if n.Alternate != nil {
		var elseScope *Scope
		var elifScopes []*Scope
		x := n.Alternate

		// Visit all the blocks.
		for x != nil {
			if elif, ok := x.Stmt.(*ast.IfStatement); ok {
				elifScope := NewScope(true, false)
				elifScopes = append(elifScopes, elifScope)

				lv.Ctx.pushScope(elifScope)
				lv.VisitExpression(elif.Test)
				lv.VisitStatement(elif.Consequent)
				lv.Ctx.popScope()

				x = elif.Alternate
			} else {
				elseScope = NewScope(true, false)

				lv.Ctx.pushScope(elseScope)
				lv.VisitStatement(x)
				lv.Ctx.popScope()

				break
			}
		}

		if elseScope != nil {
			// Find variables that exist in all scopes.
		} else {
			// Because no else, merge else if defs with if defs and merge down.
			for _, elif := range elifScopes {
				lv.Ctx.mergeSameDepth(ifScope, elif)
			}

			lv.Ctx.mergeDown(lv.Ctx.scopeDepth+1, ifScope)
		}
	} else {
		lv.Ctx.mergeDown(lv.Ctx.scopeDepth+1, ifScope)
	}
}

func (lv *DfaVisitor) VisitInvalidExpression(n *ast.InvalidExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitLabelledStatement(n *ast.LabelledStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitMemberExpression(n *ast.MemberExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitMemberProperty(n *ast.MemberProperty) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitMetaProperty(n *ast.MetaProperty) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitMethodDefinition(n *ast.MethodDefinition) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitNewExpression(n *ast.NewExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitNullLiteral(n *ast.NullLiteral) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitNumberLiteral(n *ast.NumberLiteral) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitObjectLiteral(n *ast.ObjectLiteral) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitObjectPattern(n *ast.ObjectPattern) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitOptional(n *ast.Optional) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitOptionalChain(n *ast.OptionalChain) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitParameterList(n *ast.ParameterList) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitPrivateDotExpression(n *ast.PrivateDotExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitPrivateIdentifier(n *ast.PrivateIdentifier) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitProgram(n *ast.Program) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitProperties(n *ast.Properties) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitProperty(n *ast.Property) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitPropertyKeyed(n *ast.PropertyKeyed) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitPropertyShort(n *ast.PropertyShort) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitRegExpLiteral(n *ast.RegExpLiteral) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitReturnStatement(n *ast.ReturnStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitSequenceExpression(n *ast.SequenceExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitSpreadElement(n *ast.SpreadElement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitStatement(n *ast.Statement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitStatements(n *ast.Statements) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitStringLiteral(n *ast.StringLiteral) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitSuperExpression(n *ast.SuperExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitSwitchStatement(n *ast.SwitchStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitTemplateElement(n *ast.TemplateElement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitTemplateElements(n *ast.TemplateElements) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitTemplateLiteral(n *ast.TemplateLiteral) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitThisExpression(n *ast.ThisExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitThrowStatement(n *ast.ThrowStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitTryStatement(n *ast.TryStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitUnaryExpression(n *ast.UnaryExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitUpdateExpression(n *ast.UpdateExpression) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitVariableDeclaration(n *ast.VariableDeclaration) {
	// TODO: Declarations with multiple declarations
	switch n.Token.String() {
	case "var":
		if i, ok := n.List[0].Target.Target.(*ast.Identifier); ok {
			lv.Ctx.scopeStack[lv.Ctx.scopeDepth].AddValue(i.Name, n.List[0].Initializer, true, FunctionScope, lv.Ctx.functionScopeDepth)
			lv.VisitExpression(n.List[0].Initializer)
		}
	case "let":
		if i, ok := n.List[0].Target.Target.(*ast.Identifier); ok {
			lv.Ctx.scopeStack[lv.Ctx.scopeDepth].AddValue(i.Name, n.List[0].Initializer, true, BlockScope, lv.Ctx.scopeDepth)
			lv.VisitExpression(n.List[0].Initializer)
		}
	case "const":
		if i, ok := n.List[0].Target.Target.(*ast.Identifier); ok {
			lv.Ctx.scopeStack[lv.Ctx.scopeDepth].AddValue(i.Name, n.List[0].Initializer, true, BlockScope, lv.Ctx.scopeDepth)
			lv.VisitExpression(n.List[0].Initializer)
		}
	default:
		fmt.Println("Didn't find a keyboard")
	}
}

func (lv *DfaVisitor) VisitVariableDeclarator(n *ast.VariableDeclarator) {
	// Skip the initializer.
	lv.VisitBindingTarget(n.Target)
}
func (lv *DfaVisitor) VisitVariableDeclarators(n *ast.VariableDeclarators) {
	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitWhileStatement(n *ast.WhileStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitWithStatement(n *ast.WithStatement) {

	n.VisitChildrenWith(lv)
}
func (lv *DfaVisitor) VisitYieldExpression(n *ast.YieldExpression) {

	n.VisitChildrenWith(lv)
}
