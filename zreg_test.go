package reg

import (
	"strings"
	"testing"
)

// StringReg
func TestStringReg_IsEmpty(t *testing.T) {
	if !String("").AllowEmpty() {
		t.Error("String(\"\").AllowEmpty() = false, want true")
	}
}
func TestStringReg_NotEmpty(t *testing.T) {
	if String("").NotAllowEmpty() {
		t.Error("String(\"\").NotAllowEmpty() = true, want false")
	}
}

// VersionReg
func TestVersionReg_HighThan(t *testing.T) {
	temps := map[string]bool{
		"1.0.0>0.1.0": true,
		"v1.1>v1.1.2": false,
		"1.2.3>1.1":   true,
		"v2>1.*":      true,
		"1.2.*>1.2":   true,
		"1.2>1.2":     false,
		"1.2>1.2.*":   false,
		"1.2.3>1.2.3": false,
	}
	for k, v := range temps {
		ks := strings.Split(k, ">")
		res := Version(ks[0]).HighThan(ks[1]).B()
		if res != v {
			t.Errorf("%s.HighThan(%s) = %v, want %v", ks[0], ks[1], res, v)
		}
	}
}
func TestVersionReg_LowThan(t *testing.T) {
	temps := map[string]bool{
		"1.0.0<0.1.0": false,
		"v1.1<v1.1.2": true,
		"1.2.3<1.1":   false,
		"v2<1.*":      false,
		"1.2.*<1.2":   false,
		"1.2<1.2":     false,
		"1.2<1.2.*":   true,
		"1.2.3<1.2.3": false,
	}
	for k, v := range temps {
		ks := strings.Split(k, "<")
		res := Version(ks[0]).LowThan(ks[1]).B()
		if res != v {
			t.Errorf("%s.HighThan(%s) = %v, want %v", ks[0], ks[1], res, v)
		}
	}
}
func TestVersionReg_Support(t *testing.T) {
	version := "2.3.4"
	temps := map[string]bool{
		"2.3.4":   true,
		">=2.3":   true,
		"<=2.4.1": true,
		"^2.2.9":  true,
		"~2.3.3":  true,
		"2.3.5":   false,
		">=2.4":   false,
		"<=2.3.1": false,
		"^3.9.9":  false,
		"~2.4.9":  false,
	}
	for k, v := range temps {
		res := Version(version).Support(k).B()
		if res != v {
			t.Errorf("%s.Support(%s) = %v, want %v", version, k, res, v)
		}
	}
}
