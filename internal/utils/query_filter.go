package utils

import (
	"fmt"
	"strings"
)

func CreateQueryFilters(filters map[string]interface{}, exactMatch bool) string {
	sql := ""
	limitQuery := ""

	if len(filters) > 0 {

		limit, limitExists := filters["per_page"].(int)
		if limitExists {
			limitQuery = fmt.Sprintf(" LIMIT %d", limit)
			delete(filters, "per_page")
		}

		sql += ` WHERE `
		for key, value := range filters {
			switch v := value.(type) {
			case int, int8, int16, int32, int64:
				sql += fmt.Sprintf(`%s = %d AND `, key, v)
			case float32, float64:
				sql += fmt.Sprintf(`%s = %.2f AND `, key, v)
			case string:
				if exactMatch {
					sql += fmt.Sprintf(`%s = '%s' AND `, key, v)
				} else {
					sql += fmt.Sprintf(`%s LIKE '%%%s%%' AND `, key, v)
				}
			case []interface{}:
				values := make([]string, len(v))
				for i, val := range v {
					values[i] = fmt.Sprintf("'%v'", val)
				}
				sql += fmt.Sprintf(`%s IN (%s) AND `, key, strings.Join(values, ", "))
			}
		}
		sql = sql[:len(sql)-5]
		sql += limitQuery
	}

	return sql
}
