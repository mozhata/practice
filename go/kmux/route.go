package kmux

import (
	"net/http"
	"strings"
)

type Handle func(http.ResponseWriter, *http.Request, Params)

type route struct {
	allRegistered map[string]bool
	statics       map[string]Handle
	// :id
	dynamics []*routePattern
	// // *file
	// matchAlls []routePattern
}

type routePattern struct {
	parts        []string
	partsCount   int
	pathVariable map[int]Param
	handle       Handle
}

type Param struct {
	Key   string
	Value string
}

func (r *route) add(pattern string, handle Handle) {
	if r.allRegistered == nil {
		r.allRegistered = make(map[string]bool)
	}
	if r.allRegistered[pattern] {
		panic("pattern" + pattern + "have registered")
	}
	dynamic := buildRoutePattern(pattern)
	if dynamic == nil {
		r.statics[pattern] = handle
	} else {
		r.dynamics = append(r.dynamics, dynamic)
	}
	r.allRegistered[pattern] = true
}

// buildRoutePattern retrun routePattern, isStatic
func buildRoutePattern(pattern string) *routePattern {
	p := &routePattern{}
	p.parts = strings.Split(pattern, "/")
	p.pathVariable = make([int]Param)

	for i, segment := range p.parts {
		if len(segment) < 1 {
			panic("segment of pattern should at least own one character")
		}
		if segment[0] == ':' {
			if len(segment) < 2 {
				panic("invalid segment: " + segment)
			}
			p.pathVariable[i] = Param{Key: segment[1:]}
		}
	}
	if len(p.pathVariable) == 0 {
		return nil
	}
	p.partsCount = len(p.parts)
	return p
}

func (rp *routePattern) match(pattern string) bool {
	segments := strings.Split(pattern, ",")
	if len(segments) != rp.partsCount {
		return false
	}
	for i, seg := range segments {
		if rp.parts[i] == seg {
			continue
		}
		param, found := rp.pathVariable[i]
		if !found {
			return false
		}
		param.Value = seg
		rp.pathVariable[i] = param
	}
	return false
}
