package repository

import (
	"fmt"

	"github.com/uptrace/bun"
)

type Param struct {
	tx      *bun.Tx
	attr    []string
	include []string
	page    int
	limit   int
	sort    []SortField
	filter  []FieldFilter
}

func (p *Param) InjectReturning(selector *bun.InsertQuery) {
	if len(p.attr) > 0 {
		for _, v := range p.attr {
			selector.Returning(v)
		}
	} else {
		selector.Returning("id")
	}
}

func (p *Param) InsertQuery(db *bun.DB, model any) *bun.InsertQuery {
	var qb *bun.InsertQuery

	if p.tx != nil {
		qb = p.tx.NewInsert().Model(model)
	} else {
		qb = db.NewInsert().Model(model)
	}

	return qb
}

func (p *Param) InjectFilter(selector *bun.SelectQuery) {
	for _, v := range p.filter {
		query, err := v.ToQuery()
		if err != nil {
			panic(err)
		}

		selector.Where(query, v.Value)
	}
}

func (p *Param) InjectSort(selector *bun.SelectQuery) {
	for _, v := range p.sort {
		selector.OrderExpr(v.Column + " " + string(v.Dir))
	}
}

func (p *Param) InjectPagination(selector *bun.SelectQuery) {
	page := p.page
	limit := p.limit

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	selector.Limit(limit).Offset(offset)
}

func (p *Param) InjectInclude(selector *bun.SelectQuery) {
	for _, v := range p.include {
		selector.Relation(v)
	}
}

type SortField struct {
	Column string
	Dir    SortDirection
}

type FieldFilter struct {
	Column string
	Value  any
	Op     Op
}

func (s *FieldFilter) ToQuery() (string, error) {
	switch expression := s.Op; expression {
	case EQ:
		return fmt.Sprintf("%s = ?", s.Column), nil
	case NEQ:
		return fmt.Sprintf("%s != ?", s.Column), nil
	case GT:
		return fmt.Sprintf("%s > ?", s.Column), nil
	case LT:
		return fmt.Sprintf("%s < ?", s.Column), nil
	case LIKE:
		return s.Column + " LIKE %?%", nil
	case LIKELF:
		return s.Column + " LIKE %?", nil
	case LIKERG:
		return s.Column + " LIKE ?", nil
	}

	return "", fmt.Errorf("unknown operator: %s", s.Op)
}

type SortDirection string

const (
	ASC  SortDirection = "asc"
	DESC SortDirection = "desc"
)

type Op string

const (
	EQ     Op = "="
	NEQ    Op = "!="
	GT     Op = ">"
	LT     Op = "<"
	LIKE   Op = "%"
	LIKELF Op = "%<"
	LIKERG Op = "%>"
)

type Option func(*Param)

func WithAttr(attr ...string) Option {
	return func(param *Param) {
		param.attr = attr
	}
}

func WithTx(tx *bun.Tx) Option {
	return func(param *Param) {
		param.tx = tx
	}
}

func WithPaginate(page int, limit int) Option {
	return func(param *Param) {
		param.page = page
		param.limit = limit
	}
}

func WithSort(column string, dir SortDirection) Option {
	return func(param *Param) {
		param.sort = append(param.sort, SortField{Column: column, Dir: dir})
	}
}

func parseOptions(options ...Option) *Param {
	param := &Param{}
	for _, option := range options {
		option(param)
	}
	return param
}

func WithFilter(column string, value any, op Op) Option {
	return func(param *Param) {
		param.filter = append(param.filter, FieldFilter{Column: column, Value: value, Op: op})
	}
}

func WithEqFilter(column string, value any) Option {
	return func(param *Param) {
		param.filter = append(param.filter, FieldFilter{Column: column, Value: value, Op: EQ})
	}
}

func WithInclude(include ...string) Option {
	return func(param *Param) {
		param.include = append(param.include, include...)
	}
}
