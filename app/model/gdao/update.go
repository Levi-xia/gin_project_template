package gdao

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"project/global"
	"reflect"
	"strings"
)

func (s UpdateEndPoint[T]) Update() (int64, error) {
	var (
		query          string
		rowsArgs       []any
		conditionsArgs []any
		args           []any
		result         sql.Result
		affectRows     int64
		err            error
	)
	if query, rowsArgs, conditionsArgs, err = s.point2Sql(); err != nil {
		return 0, err
	}
	
	fmt.Println(query)

	args = append(rowsArgs, conditionsArgs...)

	if result, err = global.App.DB.Exec(query, args...); err != nil {
		return 0, err
	}
	if affectRows, err = result.RowsAffected(); err != nil {
		return 0, err
	}
	return affectRows, nil
}

func (s UpdateEndPoint[T]) point2Sql() (string, []any, []any, error) {
	var (
		query           string
		tableQuery      string
		rowsQuery       string
		rowsArgs        []any
		conditionsQuery string
		conditionsArgs  []any
		appendsQuery    string
		err             error
	)
	if tableQuery, err = s.table2string(); err != nil {
		return query, rowsArgs, conditionsArgs, errors.New("table transfer failed")
	}
	if rowsQuery, rowsArgs, err = s.rows2string(); err != nil {
		return query, rowsArgs, conditionsArgs, errors.New("rows transfer failed")
	}
	if conditionsQuery, conditionsArgs, err = s.conditions2string(); err != nil {
		return query, rowsArgs, conditionsArgs, errors.New("condition transfer failed")
	}
	if appendsQuery, err = s.appends2string(); err != nil {
		return query, rowsArgs, conditionsArgs, errors.New("appends transfer failed")
	}
	query = fmt.Sprintf("UPDATE %v SET %v %v %v", tableQuery, rowsQuery, conditionsQuery, appendsQuery)

	return query, rowsArgs, conditionsArgs, err
}

func (s UpdateEndPoint[T]) table2string() (string, error) {
	if s.Table == "" {
		return s.Table, errors.New("empty table")
	}
	return s.Table, nil
}

func (s UpdateEndPoint[T]) rows2string() (string, []any, error) {
	if len(s.Rows) == 0 {
		return "", nil, errors.New("empty rows")
	}
	var (
		query       string
		prepareRows []string
		args        []any
	)
	for k, v := range s.Rows {
		prepareRows = append(prepareRows, fmt.Sprintf("%v = ?", k))
		args = append(args, v)
	}
	query = strings.Join(prepareRows, ",")
	return query, args, nil
}

func (s UpdateEndPoint[T]) conditions2string() (string, []any, error) {
	var (
		query             string
		prepareConditions []string
		args              []any
		err               error
	)
	if len(s.Conditions) == 0 {
		return query, args, err
	}
	for k, v := range s.Conditions {
		if reflect.ValueOf(v).Kind() == reflect.Slice {
			// 如果是切片类型，需要格式化为IN查询
			var (
				inQuery string
				inArgs  []any
			)
			inQuery, inArgs, _ = sqlx.In(" IN (?)", v)

			k = fmt.Sprintf("(%v %v)", k, inQuery)
			args = append(args, inArgs...)
		} else {
			// 非切片类型格式化为普通等值或比较查询
			k = fmt.Sprintf("(%v %v)", k, "?")
			args = append(args, v)
		}
		prepareConditions = append(prepareConditions, k)
	}
	query = fmt.Sprintf("WHERE %v %v", query, strings.Join(prepareConditions, " AND "))
	return query, args, nil
}

func (s UpdateEndPoint[T]) appends2string() (string, error) {
	var query string

	if len(s.Appends) > 0 {
		query = strings.Join(s.Appends, " ")
	}
	return query, nil
}
