package cores

import (
	"fmt"
)

type Model struct {
	TableName       string
	LimitStr        int
	OffsetStr       int
	WhereStr        string
	ParamStr        []interface{}
	OrderStr        string
	ColumnStr       string
	PrimaryKey      string
	JoinStr         string
	GroupByStr      string
	HavingStr       string
	QuoteIdentifier string
	ParamIdentifier string
	ParamIteration  int
}

func NewSqlBuild(options ...interface{}) (m *Model) {
	if len(options) == 0 {
		m = &Model{ColumnStr: "*", PrimaryKey: "Id", QuoteIdentifier: "`", ParamIdentifier: "?", ParamIteration: 1}
	} else if options[0] == "pg" {
		m = &Model{ColumnStr: "id", PrimaryKey: "Id", QuoteIdentifier: "\"", ParamIdentifier: options[0].(string), ParamIteration: 1}
	} else if options[0] == "mssql" {
		m = &Model{ColumnStr: "id", PrimaryKey: "id", QuoteIdentifier: "", ParamIdentifier: options[0].(string), ParamIteration: 1}
	}
	return
}

func (orm *Model) Table(tbname string) *Model {
	orm.TableName = tbname
	return orm
}

func (orm *Model) PK(pk string) *Model {
	orm.PrimaryKey = pk
	return orm
}

func (orm *Model) Where(queryString interface{}, args ...interface{}) *Model {
	switch queryString := queryString.(type) {
	case string:
		orm.WhereStr = queryString
		args = append(args, queryString)
	}
	orm.ParamStr = args
	return orm
}

func (orm *Model) Limit(start int, size ...int) *Model {
	orm.LimitStr = start
	if len(size) > 0 {
		orm.OffsetStr = size[0]
	}
	return orm
}

func (orm *Model) Offset(offset int) *Model {
	orm.OffsetStr = offset
	return orm
}

func (orm *Model) OrderBy(order string) *Model {
	orm.OrderStr = order
	return orm
}

func (orm *Model) Select(colums string) *Model {
	orm.ColumnStr = colums
	return orm
}

//The join_operator should be one of INNER, LEFT OUTER, CROSS etc - this will be prepended to JOIN
func (orm *Model) Join(join_operator, tablename, condition string) *Model {
	if orm.JoinStr != "" {
		orm.JoinStr = orm.JoinStr + fmt.Sprintf(" %v JOIN %v ON %v", join_operator, tablename, condition)
	} else {
		orm.JoinStr = fmt.Sprintf("%v JOIN %v ON %v", join_operator, tablename, condition)
	}

	return orm
}

func (orm *Model) GroupBy(keys string) *Model {
	orm.GroupByStr = fmt.Sprintf("GROUP BY %v", keys)
	return orm
}

func (orm *Model) Having(conditions string) *Model {
	orm.HavingStr = fmt.Sprintf("HAVING %v", conditions)
	return orm
}

func (orm *Model) GenerateSelectSql() (a string) {
	if orm.ParamIdentifier == "mssql" {
		if orm.OffsetStr > 0 {
			a = fmt.Sprintf("select ROW_NUMBER() OVER(order by %v )as rownum,%v from %v",
				orm.PrimaryKey,
				orm.ColumnStr,
				orm.TableName)
			if orm.WhereStr != "" {
				a = fmt.Sprintf("%v WHERE %v", a, orm.WhereStr)
			}
			a = fmt.Sprintf("select * from (%v) "+
				"as a where rownum between %v and %v",
				a,
				orm.OffsetStr,
				orm.LimitStr)
		} else if orm.LimitStr > 0 {
			a = fmt.Sprintf("SELECT top %v %v FROM %v", orm.LimitStr, orm.ColumnStr, orm.TableName)
			if orm.WhereStr != "" {
				a = fmt.Sprintf("%v WHERE %v", a, orm.WhereStr)
			}
			if orm.GroupByStr != "" {
				a = fmt.Sprintf("%v %v", a, orm.GroupByStr)
			}
			if orm.HavingStr != "" {
				a = fmt.Sprintf("%v %v", a, orm.HavingStr)
			}
			if orm.OrderStr != "" {
				a = fmt.Sprintf("%v ORDER BY %v", a, orm.OrderStr)
			}
		} else {
			a = fmt.Sprintf("SELECT %v FROM %v", orm.ColumnStr, orm.TableName)
			if orm.WhereStr != "" {
				a = fmt.Sprintf("%v WHERE %v", a, orm.WhereStr)
			}
			if orm.GroupByStr != "" {
				a = fmt.Sprintf("%v %v", a, orm.GroupByStr)
			}
			if orm.HavingStr != "" {
				a = fmt.Sprintf("%v %v", a, orm.HavingStr)
			}
			if orm.OrderStr != "" {
				a = fmt.Sprintf("%v ORDER BY %v", a, orm.OrderStr)
			}
		}
	} else {
		a = fmt.Sprintf("SELECT %v FROM %v", orm.ColumnStr, orm.TableName)
		if orm.JoinStr != "" {
			a = fmt.Sprintf("%v %v", a, orm.JoinStr)
		}
		if orm.WhereStr != "" {
			a = fmt.Sprintf("%v WHERE %v", a, orm.WhereStr)
		}
		if orm.GroupByStr != "" {
			a = fmt.Sprintf("%v %v", a, orm.GroupByStr)
		}
		if orm.HavingStr != "" {
			a = fmt.Sprintf("%v %v", a, orm.HavingStr)
		}
		if orm.OrderStr != "" {
			a = fmt.Sprintf("%v ORDER BY %v", a, orm.OrderStr)
		}
		if orm.OffsetStr > 0 {
			a = fmt.Sprintf("%v LIMIT %v OFFSET %v", a, orm.LimitStr, orm.OffsetStr)
		} else if orm.LimitStr > 0 {
			a = fmt.Sprintf("%v LIMIT %v", a, orm.LimitStr)
		}
	}
	return
}
