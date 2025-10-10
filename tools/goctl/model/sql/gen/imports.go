package gen

import (
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/model/sql/template"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func genImports(table Table, withCache, timeImport bool) (string, error) {
	var thirdImports []string
	var m = map[string]struct{}{}
	for _, c := range table.Fields {
		if len(c.DataType.Pkg) > 0 {
			if _, ok := m[c.DataType.Pkg]; ok {
				continue
			}
			m[c.DataType.Pkg] = struct{}{}
			thirdImports = append(thirdImports, fmt.Sprintf("%q", c.DataType.Pkg))
		}
	}

	if withCache {
		text, err := pathx.LoadTemplate(category, importsTemplateFile, template.Imports)
		if err != nil {
			return "", err
		}

		buffer, err := util.With("import").Parse(text).Execute(map[string]any{
			"time":       timeImport,
			"containsPQ": table.ContainsPQ,
			"data":       table,
			"third":      strings.Join(thirdImports, "\n"),
		})
		if err != nil {
			return "", err
		}

		return buffer.String(), nil
	}

	text, err := pathx.LoadTemplate(category, importsWithNoCacheTemplateFile, template.ImportsNoCache)
	if err != nil {
		return "", err
	}

	buffer, err := util.With("import").Parse(text).Execute(map[string]any{
		"time":       timeImport,
		"containsPQ": table.ContainsPQ,
		"data":       table,
		"third":      strings.Join(thirdImports, "\n"),
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
