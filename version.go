package reg

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

const (
	PatternVersion         = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)$`
	PatternVersionSupport  = `^(\^|~|>=|<=)?(0|[1-9]\d*)\.(0|[1-9]\d*)\.?(0|[1-9]\d*)?$`
	PatternVersionThan     = `^[vV]?(\*|0|[1-9]\d*)(\.(\*|0|[1-9]\d*))?(\.(\*|0|[1-9]\d*))?.*$`
	PatternVersionSemantic = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
)

type VersionReg struct {
	StringReg
}

func Version(val string) *VersionReg {
	return &VersionReg{
		StringReg: *String(val),
	}
}
func (r *VersionReg) IsVersion() *VersionReg {
	r.Match(PatternVersion)
	return r
}
func (r *VersionReg) NotVersion() *VersionReg {
	r.NotMatch(PatternVersion)
	return r
}
func (r *VersionReg) IsSemanticVersion() *VersionReg {
	r.Match(PatternVersionSemantic)
	return r
}
func (r *VersionReg) NotSemanticVersion() *VersionReg {
	r.NotMatch(PatternVersionSemantic)
	return r
}
func (r *VersionReg) IsVersionSupport() *VersionReg {
	r.Match(PatternVersionSupport)
	return r
}
func (r *VersionReg) NotVersionSupport() *VersionReg {
	r.NotMatch(PatternVersionSupport)
	return r
}
func (r *VersionReg) HighThan(ver string) *VersionReg {
	arr1 := formatVersionThan(r.String())
	arr2 := formatVersionThan(ver)
	for i := 0; i < 3; i++ {
		v := arr1[i] - arr2[i]
		if v > 0 {
			break
		}
		if v < 0 || i == 2 {
			r.res = false
			break
		}
	}
	return r
}
func (r *VersionReg) LowThan(ver string) *VersionReg {
	arr1 := formatVersionThan(r.String())
	arr2 := formatVersionThan(ver)
	for i := 0; i < 3; i++ {
		v := arr1[i] - arr2[i]
		if v < 0 {
			break
		}
		if v > 0 || i == 2 {
			r.res = false
			break
		}
	}
	return r
}
func (r *VersionReg) Support(ver string) *VersionReg {
	arr1 := formatVersionThan(r.String())
	sp := regexp.MustCompile(PatternVersionSupport).FindStringSubmatch(ver)
	arr2 := formatVersionThan(fmt.Sprintf("%s.%s.%s", sp[2], sp[3], sp[4]))
	version := fmt.Sprintf("%d.%d.%d", arr2[0], arr2[1], arr2[2])
	if fmt.Sprintf("%d.%d.%d", arr1[0], arr1[1], arr1[2]) == version {
		return r
	}
	switch sp[1] {
	case "":
		r.res = false
		return r
	case "^":
		if arr1[0] != arr2[0] {
			r.res = false
			return r
		}
		return r.HighThan(version)
	case "~":
		if arr1[0] != arr2[0] || arr1[1] != arr2[1] {
			r.res = false
			return r
		}
		return r.HighThan(version)
	case ">=":
		return r.HighThan(version)
	case "<=":
		return r.LowThan(version)
	}
	return r
}

func formatVersionThan(version string) (res []int64) {
	arr := regexp.MustCompile(PatternVersionThan).FindStringSubmatch(version)
	for _, v := range arr {
		if v == "*" {
			res = append(res, math.MaxInt64)
			continue
		}
		d, _ := strconv.ParseInt(v, 10, 64)
		res = append(res, d)
	}
	return []int64{res[1], res[3], res[5]}
}
