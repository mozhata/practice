package mmux

import (
	"fmt"
	"net/http"
	"strings"
)

type Handle func(http.ResponseWriter, *http.Request, PathVars)

type route struct {
	*trie
}

type trie struct {
	segment  string
	handle   Handle
	branches []*trie
}

type PathVar struct {
	Key   string
	Value string
}

type PathVars []PathVar

// right
// return Handle, score, matched
func (t *trie) match(segments []string, vars PathVars) (Handle, PathVars) {
	if len(segments) == 1 {
		if segments[0] == t.segment {
			return t.handle, vars
		}
		if t.segment[0] == ':' {
			vars = append(vars, PathVar{t.segment[1:], segments[0]})
			return t.handle, vars
		}
		return nil, nil
	}
	// len(segments) > 1
	if t.segment[0] == ':' {
		vars = append(vars, PathVar{t.segment[1:], segments[0]})
	} else if t.segment != segments[0] {
		return nil, nil
	}
	var (
		handlFunc Handle
	)
	for _, tree := range t.branches {
		handlFunc, vars = tree.match(segments[1:], vars)
		if handlFunc != nil {
			return handlFunc, vars
		}
	}
	// not match
	// vars = vars[:len(vars)-1]
	return nil, nil
}

// len(segments)>0
func (t *trie) add(segments []string, handle Handle) {
	for i := 0; i < len(segments); i++ {
		for _, branch := range t.branches {
			if segments[i] == branch.segment {
				if i == len(segments)-1 {
					branch.handle = handle
					return
				}
				branch.add(segments[i+1:], handle)
			}
		}
		newBranch := trie{segment: segments[i]}
		if t.branches == nil {
			t.branches = make([]*trie, 0, 1)
		}
		t.branches = append(t.branches, &newBranch)
		if i == len(segments)-1 {
			newBranch.handle = handle
			return
		}
		newBranch.add(segments[i+1:], handle)
	}
}

func newRoute() *route {
	r := route{}
	r.trie = &trie{} // root segment is ""
	return &r
}

// right
func (r *route) match(pattern string) (Handle, PathVars) {
	if pattern == "/" {
		return r.trie.handle, nil
	}
	segments := strings.Split(pattern, "/")
	segments = segments[1 : len(segments)-1]
	vars := make(PathVars, 0, 1)
	var h Handle
	for _, branch := range r.branches {
		if h, vars = branch.match(segments, vars); h != nil {
			return h, vars
		}
	}
	return nil, nil
}

// TODO check every panic to make sure logic
func (r *route) add(pattern string, handle Handle) {
	// check validation
	if pattern == "" {
		panic("emtpty pattern !")
	}
	if pattern[0] != '/' {
		panic("path must begin with '/'. pattern: '" + pattern + "'")
	}
	if handle == nil {
		panic("handler should not empty. register pattern: " + pattern)
	}
	if pattern[len(pattern)-1] != '/' {
		pattern += "/"
	}
	// check segments validation
	segments := strings.Split(pattern, "/")
	for _, seg := range segments[1 : len(segments)-1] {
		if seg == "" || seg == ":" {
			panic("pattern \"" + pattern + "\" is not valid")
		}
	}
	h, pathVariables := r.match(pattern)
	if h != nil {
		fmt.Printf("pathVariables: %#v\n\n", pathVariables)
		// already exist
		shadows := make([]string, len(segments))
		var regxIndex int
		for i, seg := range segments[1 : len(segments)-1] {

			if seg[0] == ':' {
				param := pathVariables[regxIndex]
				shadows[i+1] = ":" + param.Key
				regxIndex = regxIndex + 1
				continue
			}
			shadows[i+1] = seg
		}
		panic(fmt.Sprintf("pattern %s conflict with %s", pattern, strings.Join(shadows, "/")))
	}
	if pattern == "/" {
		r.trie.handle = handle
		return
	}

	r.trie.add(segments[1:len(segments)-1], handle)
}

func (r *route) getHandle(pattern string) (Handle, PathVars) {
	var (
		pathVariables = make(PathVars, 0, 1)
		parts         = strings.Split(pattern, "/")
		h             Handle
	)
	for _, branch := range r.trie.branches {
		h, pathVariables = branch.match(parts[1:len(parts)-1], pathVariables)
		if h != nil {
			return h, pathVariables
		}
	}
	return nil, nil
}
