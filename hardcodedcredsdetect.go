package hardcodedcredsdetect


import (
    "go/ast"
    "go/token"
    "regexp"
    "strconv"

    "golang.org/x/tools/go/analysis"
    "golang.org/x/tools/go/analysis/passes/inspect"
    "golang.org/x/tools/go/ast/inspector"
)

var (
    passwordRegexp    = regexp.MustCompile(`(?i)password|passwd|pwd`)
    credentialRegexp  = regexp.MustCompile(`(?i)credential|cred|auth.*token|api.*key`)
)

var Analyzer = &analysis.Analyzer{
    Name: "hardcodedcredsdetect",
    Doc:  "detects hardcoded credentials in code",
    Run:  run,
    Requires: []*analysis.Analyzer{
        inspect.Analyzer,
    },
}

func run(pass *analysis.Pass) (interface{}, error) {
    var err error
    astInsp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

    astInsp.Preorder(nil, func(n ast.Node) {
        switch n := n.(type) {
        case *ast.AssignStmt:
            for i, expr := range n.Rhs {
                if lit, ok := expr.(*ast.BasicLit); ok && lit.Kind == token.STRING {
                    if isSensitiveVariableName(n.Lhs[i].(*ast.Ident).Name) {
                        if value, err := strconv.Unquote(lit.Value); err == nil {
                            pass.Reportf(lit.Pos(), "hardcoded credential found: %s", value)
                        }
                    }
                }
            }
        }
    })

    return nil, err
}

// judge whether the variable name is sensitive or not
func isSensitiveVariableName(name string) bool {
    return passwordRegexp.MatchString(name) || credentialRegexp.MatchString(name)
}
