package reg

import (
	"reflect"
	"regexp"
	"unicode/utf8"
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
func (r *Reg) AllowZero() bool {
	return r.res && reflect.ValueOf(r.val).IsZero()
}
func (r *Reg) NotAllowZero() bool {
	return r.res && !reflect.ValueOf(r.val).IsZero()
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
func (r *Reg) MaxLen(length int) *Reg {
	switch r.Kind() {
	case reflect.String:
		r.res = r.res && utf8.RuneCountInString(r.String()) <= length
	case reflect.Slice, reflect.Array, reflect.Chan, reflect.Map:
		r.res = r.res && reflect.ValueOf(r.val).Len() <= length
	case reflect.Ptr:
		r.val = reflect.ValueOf(r.val).Elem()
		return r.MaxLen(length)
	default:
		r.res = false
	}
	return r
}
func (r *Reg) MinLen(length int) *Reg {
	switch r.Kind() {
	case reflect.String:
		r.res = r.res && utf8.RuneCountInString(r.String()) >= length
	case reflect.Slice, reflect.Array, reflect.Chan, reflect.Map:
		r.res = r.res && reflect.ValueOf(r.val).Len() >= length
	case reflect.Ptr:
		r.val = reflect.ValueOf(r.val).Elem()
		return r.MinLen(length)
	default:
		r.res = false
	}
	return r
}
