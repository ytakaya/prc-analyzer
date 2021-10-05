package visitor

import (
	"github.com/ytakaya/prc-parser/parser"
)

type PrcListener struct {
	*parser.BaseTSqlParserListener
	Rules []int
}

// func (p *PrcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	p.Rules = append(p.Rules, ctx.GetRuleIndex())
// }

func (p *PrcListener) EnterMerge_statement(ctx *parser.Merge_statementContext) {
	p.Rules = append(p.Rules, ctx.GetRuleIndex())
}

func (p *PrcListener) EnterDelete_statement(ctx *parser.Delete_statementContext) {
	p.Rules = append(p.Rules, ctx.GetRuleIndex())
}

func (p *PrcListener) EnterInsert_statement(ctx *parser.Insert_statementContext) {
	p.Rules = append(p.Rules, ctx.GetRuleIndex())
}

func (p *PrcListener) EnterSelect_statement(ctx *parser.Select_statementContext) {
	p.Rules = append(p.Rules, ctx.GetRuleIndex())
}

func (p *PrcListener) EnterUpdate_statement(ctx *parser.Update_statementContext) {
	p.Rules = append(p.Rules, ctx.GetRuleIndex())
}

func (p *PrcListener) EnterIf_statement(ctx *parser.If_statementContext) {
	p.Rules = append(p.Rules, ctx.GetRuleIndex())
}

func (p *PrcListener) EnterWhile_statement(ctx *parser.While_statementContext) {
	p.Rules = append(p.Rules, ctx.GetRuleIndex())
}

func (p *PrcListener) EnterReturn_statement(ctx *parser.Return_statementContext) {
	p.Rules = append(p.Rules, ctx.GetRuleIndex())
}
