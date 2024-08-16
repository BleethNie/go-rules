// Code generated from E:/02_Resource/01_Code/jc_go_manager/go-rules/parser/Rules.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Rules

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by RulesParser.
type RulesVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RulesParser#rules.
	VisitRules(ctx *RulesContext) interface{}

	// Visit a parse tree produced by RulesParser#argumentExpressionList.
	VisitArgumentExpressionList(ctx *ArgumentExpressionListContext) interface{}

	// Visit a parse tree produced by RulesParser#numberList.
	VisitNumberList(ctx *NumberListContext) interface{}

	// Visit a parse tree produced by RulesParser#stringList.
	VisitStringList(ctx *StringListContext) interface{}

	// Visit a parse tree produced by RulesParser#valueList.
	VisitValueList(ctx *ValueListContext) interface{}

	// Visit a parse tree produced by RulesParser#MultiplicativeExpr.
	VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{}

	// Visit a parse tree produced by RulesParser#Ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by RulesParser#ContainsCond.
	VisitContainsCond(ctx *ContainsCondContext) interface{}

	// Visit a parse tree produced by RulesParser#AdditiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by RulesParser#NumberLit.
	VisitNumberLit(ctx *NumberLitContext) interface{}

	// Visit a parse tree produced by RulesParser#InCond.
	VisitInCond(ctx *InCondContext) interface{}

	// Visit a parse tree produced by RulesParser#DurationLit.
	VisitDurationLit(ctx *DurationLitContext) interface{}

	// Visit a parse tree produced by RulesParser#ParentExpr.
	VisitParentExpr(ctx *ParentExprContext) interface{}

	// Visit a parse tree produced by RulesParser#FuncExpr.
	VisitFuncExpr(ctx *FuncExprContext) interface{}

	// Visit a parse tree produced by RulesParser#BoolLit.
	VisitBoolLit(ctx *BoolLitContext) interface{}

	// Visit a parse tree produced by RulesParser#StringLit.
	VisitStringLit(ctx *StringLitContext) interface{}

	// Visit a parse tree produced by RulesParser#EqualityExpr.
	VisitEqualityExpr(ctx *EqualityExprContext) interface{}

	// Visit a parse tree produced by RulesParser#PrimaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by RulesParser#RelationalExpr.
	VisitRelationalExpr(ctx *RelationalExprContext) interface{}

	// Visit a parse tree produced by RulesParser#LogicalExpr.
	VisitLogicalExpr(ctx *LogicalExprContext) interface{}

	// Visit a parse tree produced by RulesParser#BooleanExpr.
	VisitBooleanExpr(ctx *BooleanExprContext) interface{}
}
