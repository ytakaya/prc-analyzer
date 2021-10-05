package visitor

import (
	"github.com/ytakaya/prc-parser/parser"
)

type PRCVisitor struct {
	*parser.BaseTSqlParserVisitor

	Count  int
	PrcCtx []parser.ICreate_or_alter_procedureContext
}

func (p *PRCVisitor) VisitTsql_file(ctx *parser.Tsql_fileContext) interface{} {
	for _, child := range ctx.GetChildren() {
		switch childCtx := child.(type) {
		case *parser.BatchContext:
			childCtx.Accept(p)
		}
	}
	return nil
}

func (p *PRCVisitor) VisitBatch(ctx *parser.BatchContext) interface{} {
	for _, child := range ctx.GetChildren() {
		switch childCtx := child.(type) {
		case *parser.Batch_level_statementContext:
			childCtx.Accept(p)
		}
	}
	return nil
}

func (p *PRCVisitor) VisitBatch_level_statement(ctx *parser.Batch_level_statementContext) interface{} {
	if prcCtx := ctx.Create_or_alter_procedure(); prcCtx != nil {
		p.Count += 1
		p.PrcCtx = append(p.PrcCtx, prcCtx)
	}
	return nil
}
