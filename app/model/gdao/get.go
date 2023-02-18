package gdao

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"project/global"
	"reflect"
	"strings"
)

func (s GetEndPoint[T]) Get() error {
	var (
		query string
		args  []any
		err   error
	)
	if query, args, err = s.point2Sql(); err != nil {
		return err
	}

	fmt.Println(query)

	err = global.App.DB.Get(s.Model, query, args...)
	return err
}

func (s GetEndPoint[T]) point2Sql() (string, []any, error) {

	var (
		query           string
		fieldsQuery     string
		tableQuery      string
		conditionsQuery string
		conditionsArgs  []any
		appendsQuery    string
		err             error
	)
	if fieldsQuery, err = s.fields2string(); err != nil {
		return query, conditionsArgs, errors.New("fields transfer failed")
	}
	if tableQuery, err = s.table2string(); err != nil {
		return query, conditionsArgs, errors.New("table transfer failed")
	}
	if conditionsQuery, conditionsArgs, err = s.conditions2string(); err != nil {
		return query, conditionsArgs, errors.New("condition transfer failed")
	}
	if appendsQuery, err = s.appends2string(); err != nil {
		return query, conditionsArgs, errors.New("appends transfer failed")
	}
	query = fmt.Sprintf("SELECT %v FROM %v %v %v", fieldsQuery, tableQuery, conditionsQuery, appendsQuery)

	return query, conditionsArgs, err
}

func (s GetEndPoint[T]) table2string() (string, error) {
	if s.Table == "" {
		return "", errors.New("empty table")
	}
	return s.Table, nil
}

func (s GetEndPoint[T]) fields2string() (string, error) {
	var query string

	if len(s.Fields) == 0 {
		return query, errors.New("empty fields")
	}
	if s.Fields[0] == "*" {
		query = "*"
	} else {
		query = strings.Join(s.Fields, ",")
	}
	return query, nil
}

func (s GetEndPoint[T]) conditions2string() (string, []any, error) {
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

func (s GetEndPoint[T]) appends2string() (string, error) {
	var query string

	if len(s.Appends) > 0 {
		query = strings.Join(s.Appends, " ")
	}
	return query, nil
}
