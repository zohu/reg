package reg

const (
	PatternAlphanumeric               = `^[a-zA-Z0-9]+$`
	PatternAlphanumericUnderline      = `^[a-zA-Z0-9_]+$`
	PatternTruthAlphanumericUnderline = `^[a-zA-Z]+[a-zA-Z0-9_]*[a-zA-Z0-9]+$`

	PatternUrl   = `^http(s)?:\/\/([\w-]+\.)+[\w-]+(\/[\w- ./?%&=]*)?$`
	PatternPhone = `^((\+|00)86)?1[3-9]\d{9}$`
	PatternEmail = `^\w+([-\.]\w+)*@(\w+\.)+[A-Za-z]{2,6}$`
	PatternIpv4  = `^((25[0-5]|(2[0-4]|1\d|[1-9])?\d)\.){3}(25[0-5]|(2[0-4]|1\d|[1-9])?\d)$`
)

type StringReg struct {
	Reg
}

func String(val string) *StringReg {
	return &StringReg{
		Reg: Reg{
			val: val,
			res: true,
		},
	}
}
func (r *StringReg) AllowEmpty() *StringReg {
	r.res = r.res || r.val == ""
	return r
}
func (r *StringReg) NotAllowEmpty() *StringReg {
	r.res = r.res && r.val != ""
	return r
}
func (r *StringReg) IsAlphanumeric() *StringReg {
	r.Match(PatternAlphanumeric)
	return r
}
func (r *StringReg) NotAlphanumeric() *StringReg {
	r.NotMatch(PatternAlphanumeric)
	return r
}
func (r *StringReg) IsAlphanumericUnderline() *StringReg {
	r.Match(PatternAlphanumericUnderline)
	return r
}
func (r *StringReg) NotAlphanumericUnderline() *StringReg {
	r.NotMatch(PatternAlphanumericUnderline)
	return r
}
func (r *StringReg) IsTruthAlphanumericUnderline() *StringReg {
	r.Match(PatternTruthAlphanumericUnderline)
	return r
}
func (r *StringReg) NotTruthAlphanumericUnderline() *StringReg {
	r.NotMatch(PatternTruthAlphanumericUnderline)
	return r
}
func (r *StringReg) IsUrl() *StringReg {
	r.Match(PatternUrl)
	return r
}
func (r *StringReg) NotUrl() *StringReg {
	r.NotMatch(PatternUrl)
	return r
}
func (r *StringReg) IsPhone() *StringReg {
	r.Match(PatternPhone)
	return r
}
func (r *StringReg) NotPhone() *StringReg {
	r.NotMatch(PatternPhone)
	return r
}
func (r *StringReg) IsEmail() *StringReg {
	r.Match(PatternEmail)
	return r
}
func (r *StringReg) NotEmail() *StringReg {
	r.NotMatch(PatternEmail)
	return r
}
func (r *StringReg) IsIpv4() *StringReg {
	r.Match(PatternIpv4)
	return r
}
func (r *StringReg) NotIpv4() *StringReg {
	r.NotMatch(PatternIpv4)
	return r
}
