package util

import (
	"strconv"
	"github.com/astaxie/beego/context"
)

func Datatables(tableName string, aColumns []string, Input *context.BeegoInput, where interface{}) (string, []interface{}, string) {

	/*
	 * Ordering
	 * 排序请求
	 */
	orderStr := ""
	isFirst := true
	if iSortCol_0, _ := strconv.Atoi(Input.Query("iSortCol_0")); iSortCol_0 > -1 {
		ranges, _ := strconv.Atoi(Input.Query("iSortingCols"))
		for i := 0; i < ranges; i++ {
			istring := strconv.Itoa(i)
			if iSortcol := Input.Query("bSortable_" + Input.Query("iSortCol_" + istring)); iSortcol == "true" {
				sordir := Input.Query("sSortDir_" + istring)
				thisSortCol, _ := strconv.Atoi(Input.Query("iSortCol_" + istring))
				if sordir == "asc" {
					if isFirst {
						orderStr += tableName + "." + aColumns[thisSortCol] + " asc"
						isFirst = false
					} else {
						orderStr += "," + orderStr + tableName + "." + aColumns[thisSortCol] + " asc"
					}
				} else {
					if isFirst {
						orderStr += tableName + "." + aColumns[thisSortCol] + " desc"
						isFirst = false
					} else {
						orderStr += "," + orderStr + tableName + "." + aColumns[thisSortCol] + " desc"
					}
				}
			}
		}
	}


	/*
	 * where条件
	 */
	whereStr := ""
	isFirst = true
	wheres, ok := where.(map[string]interface{})
	args := make([]interface{}, 0)

	/*
	 * Filtering
	 * 快速过滤器
	 */
	if ok {
		for k, v := range wheres {
			if isFirst {
				whereStr += k + " ? "
				isFirst = false
			} else {
				whereStr += " And " + k + " ? "
			}
			args = append(args, v)
		}
	}

	isFilter := false
	isFilterFirst := true
	filterStr := ""
	for i := 0; i < len(aColumns); i++ {
		if Input.Query("bSearchable_" + strconv.Itoa(i)) == "true" && len(Input.Query("sSearch")) > 0 {
			isFilter = true
			if isFilterFirst {
				filterStr += aColumns[i] + " like ? "
				isFilterFirst = false
			} else {
				filterStr += " OR " + aColumns[i] + " like ? "
			}
			args = append(args, "%" + Input.Query("sSearch") + "%")
		}
	}

	if isFilter {
		whereStr +=  " AND ( " + filterStr + ") "
	}

	return orderStr, args, whereStr
}
