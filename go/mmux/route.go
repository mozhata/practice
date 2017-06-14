package kmux

import (
	"net/http"
	"strings"
)

type Handle func(http.ResponseWriter, *http.Request, Params)

type route struct {
	statics  map[string]Handle
	dynamics *trie
}

type trie struct {
	segment  string
	handle   Handle
	branches []*trie
}

func (t *trie) match(segments []string, pathVariables map[string]string) (Handle, bool) {
	// seg 相同, 比较下一级
	// seg 不同, 1, segement是regx, 比较下一级
	// seg 不同, 2, segement是static, return false
	// TODO: not neccessary
	if len(segments) < 1 {
		return nil, false
	}
	if len(segments) == 1 {
		if segments[0] == t.segment {
			return t.handle, true
		}
		if t.segment[0] == ':' {
			pathVariables[t.segment[1:]] = segments[0]
			return t.handle, true
		}
		return nil, false
	}
	// len(segements) > 1
	if t.segment[0] == ':' {
		pathVariables[t.segment[1:]] = segments[0]
	} else {
		if t.segment != segments[0] {
			return nil, false
		}
	}
	for _, tree := range t.branches {
		if handler, matched := tree.match(segments[1:], pathVariables); matched {
			return handler, true
		}
	}
	delete(pathVariables, t.segment[1:])
	return nil, false
}

func newRoute() *route {
	r := route{}
	r.allRegistered = make(map[string]bool)
	r.statics = make(map[string]Handle)
	return &r
}

type routePattern struct {
	parts        []string
	partsCount   int
	pathVariable Params
	handle       Handle
}

type Params map[int]Param

type Param struct {
	Key   string
	Value string
}

func (r *route) add(pattern string, handle Handle) {
	if handle == nil {
		panic("handler should not empty")
	}
	// TODO: change this

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
		dynamic.handle = handle
		r.dynamics = append(r.dynamics, dynamic)
	}
	r.allRegistered[pattern] = true
}

/*func (r *route) add(pattern string, handle Handle) {

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
		dynamic.handle = handle
		r.dynamics = append(r.dynamics, dynamic)
	}
	r.allRegistered[pattern] = true
}
*/
// buildRoutePattern retrun routePattern, isStatic
func buildRoutePattern(pattern string) *routePattern {
	// static
	if pattern == "/" {
		return nil
	}

	p := &routePattern{}
	parts := strings.Split(pattern, "/")
	p.parts = parts[1 : len(parts)-1]
	p.partsCount = len(p.parts)
	p.pathVariable = make(map[int]Param)
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
	return p
}

func (rp *routePattern) match(pattern string) bool {
	segments := strings.Split(pattern, "/")
	segments = segments[1 : len(segments)-1]
	if len(segments) != rp.partsCount {
		return false
	}
	params := make(map[int]Param, len(rp.pathVariable))
	for i, seg := range segments {
		if rp.parts[i] == seg {
			continue
		}
		param, found := rp.pathVariable[i]
		if !found {
			return false
		}
		param.Value = seg
		params[i] = param
	}
	rp.pathVariable = params
	return true
}

func (r *route) getHandle(pattern string) (Handle, Params) {
	h, ok := r.statics[pattern]
	if ok {
		return h, nil
	}
	for _, rp := range r.dynamics {
		if rp.match(pattern) {
			return rp.handle, rp.pathVariable
		}
	}
	return nil, nil
}
