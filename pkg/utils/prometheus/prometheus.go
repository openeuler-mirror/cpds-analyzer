package prometheus

import (
	"github.com/prometheus/prometheus/promql/parser"
)

func IsExprValid(expr string) bool {
	if _, err := parser.ParseExpr(expr); err != nil {
		return false
	}
	return true
}
