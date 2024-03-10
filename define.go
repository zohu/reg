package reg

import (
	"reflect"
	"regexp"
)

type Reg struct {
	val interface{}
	res bool
}

func New(val interface{}) *Reg {
	return &Reg{
		val: val,
		res: true,
	}
}

func (r *Reg) B() bool {
	return r.res
}
func (r *Reg) NotB() bool {
	return !r.res
}
func (r *Reg) AllowZero() *Reg {
	r.res = r.res || reflect.ValueOf(r.val).IsZero()
	return r
}
func (r *Reg) NotAllowZero() *Reg {
	r.res = r.res && !reflect.ValueOf(r.val).IsZero()
	return r
}
func (r *Reg) And(dest *Reg) *Reg {
	r.res = r.res && dest.res
	return r
}
func (r *Reg) Or(dest *Reg) *Reg {
	r.res = r.res || dest.res
	return r
}
func (r *Reg) Kind() reflect.Kind {
	return reflect.TypeOf(r.val).Kind()
}
func (r *Reg) String() string {
	return reflect.ValueOf(r.val).String()
}
func (r *Reg) Match(pattern string) *Reg {
	ok, _ := regexp.MatchString(pattern, r.String())
	r.res = r.res && ok
	return r
}
func (r *Reg) NotMatch(pattern string) *Reg {
	ok, _ := regexp.MatchString(pattern, r.String())
	r.res = r.res && !ok
	return r
}
