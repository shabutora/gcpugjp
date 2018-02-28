// generated by qbg -output misc/fixture/e/model_query.go misc/fixture/e; DO NOT EDIT

package e

import (
	"github.com/favclip/qbg/qbgutils"
	"google.golang.org/appengine/datastore"
)

// SampleQueryBuilder build query for Sample.
type SampleQueryBuilder struct {
	q      *datastore.Query
	plugin qbgutils.Plugin
	Foo    *SampleQueryProperty
}

// SampleQueryProperty has property information for SampleQueryBuilder.
type SampleQueryProperty struct {
	bldr *SampleQueryBuilder
	name string
}

// NewSampleQueryBuilder create new SampleQueryBuilder.
func NewSampleQueryBuilder() *SampleQueryBuilder {
	return NewSampleQueryBuilderWithKind("Sample")
}

// NewSampleQueryBuilderWithKind create new SampleQueryBuilder with specific kind.
func NewSampleQueryBuilderWithKind(kind string) *SampleQueryBuilder {
	q := datastore.NewQuery(kind)
	bldr := &SampleQueryBuilder{q: q}
	bldr.Foo = &SampleQueryProperty{
		bldr: bldr,
		name: "Foo",
	}

	if plugger, ok := interface{}(bldr).(qbgutils.Plugger); ok {
		bldr.plugin = plugger.Plugin()
		bldr.plugin.Init("Sample")
	}

	return bldr
}

// Ancestor sets parent key to ancestor query.
func (bldr *SampleQueryBuilder) Ancestor(parentKey *datastore.Key) *SampleQueryBuilder {
	bldr.q = bldr.q.Ancestor(parentKey)
	if bldr.plugin != nil {
		bldr.plugin.Ancestor(parentKey)
	}
	return bldr
}

// KeysOnly sets keys only option to query.
func (bldr *SampleQueryBuilder) KeysOnly() *SampleQueryBuilder {
	bldr.q = bldr.q.KeysOnly()
	if bldr.plugin != nil {
		bldr.plugin.KeysOnly()
	}
	return bldr
}

// Start setup to query.
func (bldr *SampleQueryBuilder) Start(cur datastore.Cursor) *SampleQueryBuilder {
	bldr.q = bldr.q.Start(cur)
	if bldr.plugin != nil {
		bldr.plugin.Start(cur)
	}
	return bldr
}

// Offset setup to query.
func (bldr *SampleQueryBuilder) Offset(offset int) *SampleQueryBuilder {
	bldr.q = bldr.q.Offset(offset)
	if bldr.plugin != nil {
		bldr.plugin.Offset(offset)
	}
	return bldr
}

// Limit setup to query.
func (bldr *SampleQueryBuilder) Limit(limit int) *SampleQueryBuilder {
	bldr.q = bldr.q.Limit(limit)
	if bldr.plugin != nil {
		bldr.plugin.Limit(limit)
	}
	return bldr
}

// Query returns *datastore.Query.
func (bldr *SampleQueryBuilder) Query() *datastore.Query {
	return bldr.q
}

// Filter with op & value.
func (p *SampleQueryProperty) Filter(op string, value interface{}) *SampleQueryBuilder {
	switch op {
	case "<=":
		p.LessThanOrEqual(value)
	case ">=":
		p.GreaterThanOrEqual(value)
	case "<":
		p.LessThan(value)
	case ">":
		p.GreaterThan(value)
	case "=":
		p.Equal(value)
	default:
		p.bldr.q = p.bldr.q.Filter(p.name+" "+op, value) // error raised by native query
	}
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, op, value)
	}
	return p.bldr
}

// LessThanOrEqual filter with value.
func (p *SampleQueryProperty) LessThanOrEqual(value interface{}) *SampleQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<=", value)
	}
	return p.bldr
}

// GreaterThanOrEqual filter with value.
func (p *SampleQueryProperty) GreaterThanOrEqual(value interface{}) *SampleQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >=", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">=", value)
	}
	return p.bldr
}

// LessThan filter with value.
func (p *SampleQueryProperty) LessThan(value interface{}) *SampleQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" <", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "<", value)
	}
	return p.bldr
}

// GreaterThan filter with value.
func (p *SampleQueryProperty) GreaterThan(value interface{}) *SampleQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" >", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, ">", value)
	}
	return p.bldr
}

// Equal filter with value.
func (p *SampleQueryProperty) Equal(value interface{}) *SampleQueryBuilder {
	p.bldr.q = p.bldr.q.Filter(p.name+" =", value)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Filter(p.name, "=", value)
	}
	return p.bldr
}

// Asc order.
func (p *SampleQueryProperty) Asc() *SampleQueryBuilder {
	p.bldr.q = p.bldr.q.Order(p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Asc(p.name)
	}
	return p.bldr
}

// Desc order.
func (p *SampleQueryProperty) Desc() *SampleQueryBuilder {
	p.bldr.q = p.bldr.q.Order("-" + p.name)
	if p.bldr.plugin != nil {
		p.bldr.plugin.Desc(p.name)
	}
	return p.bldr
}