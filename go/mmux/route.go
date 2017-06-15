package mmux

import (
	"fmt"
	"net/http"
	"strings"
)

type Handle func(http.ResponseWriter, *http.Request, map[string]string)

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
	if t.segment == "" && len(t.branches) == 0 {
		return nil, false
	}
	// if segments[0]
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
	fmt.Printf("t.segment: %q\n", t.segment)
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

func (t *trie) add(segments []string, handle Handle) {
	// the root segment is ""
	for i := 0; i < len(segments); i++ {
		for _, branch := range t.branches {
			// i == len(segments) can not happen
			if segments[i] == branch.segment {
				branch.add(segments[i+1:], handle)
			}
		}
		newBranch := trie{segment: segments[i]}
		if t.branches == nil {
			t.branches = make([]*trie, 0, 1)
		}
		if i == len(segments)-1 {
			newBranch.handle = handle
		}
		t.branches = append(t.branches, &newBranch)
		newBranch.add(segments[i+1:], handle)
	}
}

func newRoute() *route {
	r := route{}
	r.statics = make(map[string]Handle)
	r.dynamics = &trie{} // root segment is ""
	return &r
}

// TODO check every panic to make sure logic
func (r *route) add(pattern string, handle Handle) {
	if handle == nil {
		panic("handler should not empty. register pattern: " + pattern)
	}
	// check if the pattern have registered to statics
	if _, ok := r.statics[pattern]; ok {
		panic("pattern " + pattern + "already registered.")
	}
	// check whether the pattern contains regx segment
	if !strings.Contains(pattern, "/:") {
		r.statics[pattern] = handle
		return
	}
	// pattern contains regx segments.
	// check whether the pattern registered to dynamics
	segments := strings.Split(pattern, "/")
	// ignore the first and last emtyp string
	fmt.Printf("\n\nroute.add, segments: %#v, segments[1:len(segments) -1]: %#v\n\n", segments, segments[1:len(segments)-1])
	segments = segments[1 : len(segments)-1]
	// check segment is valid
	for _, seg := range segments {
		if seg == "" || seg == ":" {
			panic("pattern " + pattern + "is not vaid")
		}
	}
	pathVariables := make(map[string]string)
	if _, ok := r.dynamics.match(segments, pathVariables); ok {
		// build exist pattern
		shadows := make([]string, 0, len(segments))
		for _, seg := range segments {
			if seg[0] == ':' {
				regxSeg, _ := findKeyByValue(pathVariables, seg)
				shadows = append(shadows, ":"+regxSeg)
				continue
			}
			shadows = append(shadows, seg)
		}
		existPattern := "/" + strings.Join(shadows, "/")
		panic(fmt.Sprintf("pattern %s conflict with %s", pattern, existPattern))
	}
	r.dynamics.add(segments, handle)
}

func findKeyByValue(dict map[string]string, val string) (string, bool) {
	for k, v := range dict {
		if v == val {
			return k, true
		}
	}
	return "", false
}

func (r *route) getHandle(pattern string) (Handle, map[string]string) {
	h, ok := r.statics[pattern]
	if ok {
		return h, nil
	}
	pathVariables := make(map[string]string)
	parts := strings.Split(pattern, "/")
	for _, branch := range r.dynamics.branches {
		h, ok = branch.match(parts[1:len(parts)-1], pathVariables)
		if ok {
			return h, pathVariables
		}
	}
	return nil, nil
}
