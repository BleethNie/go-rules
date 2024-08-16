// Code generated from E:/02_Resource/01_Code/jc_go_manager/go-rules/parser/Rules.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Rules

import "github.com/antlr4-go/antlr/v4"

type BaseRulesVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRulesVisitor) VisitRules(ctx *RulesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitArgumentExpressionList(ctx *ArgumentExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitNumberList(ctx *NumberListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitStringList(ctx *StringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitValueList(ctx *ValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitContainsCond(ctx *ContainsCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitAdditiveExpr(ctx *AdditiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitNumberLit(ctx *NumberLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitInCond(ctx *InCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitDurationLit(ctx *DurationLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitParentExpr(ctx *ParentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitFuncExpr(ctx *FuncExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitBoolLit(ctx *BoolLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitStringLit(ctx *StringLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitRelationalExpr(ctx *RelationalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitLogicalExpr(ctx *LogicalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRulesVisitor) VisitBooleanExpr(ctx *BooleanExprContext) interface{} {
	return v.VisitChildren(ctx)
}
