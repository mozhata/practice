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
	key   string
	value string
}

type PathVars []PathVar

// TODO: pathVariables use slice, not map
// return Handle, score, matched
func (t *trie) match(segments []string, vars PathVars) Handle {
	if len(segments) == 1 {
		if segments[0] == t.segment {
			return t.handle
		}
		if t.segment[0] == ':' {
			vars = append(vars, PathVar{t.segment[1:], segments[0]})
			return t.handle
		}
		return nil
	}
	// len(segments) > 1
	if t.segment[0] == ':' {
		vars = append(vars, PathVar{t.segment[1:], segments[0]})
	} else if t.segment != segments[0] {
		return nil
	}
	var (
		handlFunc Handle
	)
	for _, tree := range t.branches {
		handlFunc = tree.match(segments[1:], vars)
		if handlFunc != nil {
			return handlFunc
		}
	}
	// not match
	vars = vars[:len(vars)-1]
	return nil
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
		if h = branch.match(segments, vars); h != nil {
			return h, vars
		}
	}
	return nil, nil
}

// TODO check every panic to make sure logic
func (r *route) add(pattern string, handle Handle) {
	if handle == nil {
		panic("handler should not empty. register pattern: " + pattern)
	}
	// TODO: should use slice, not map
	h, pathVariables := r.match(pattern)
	segments := strings.Split(pattern, "/")
	if h != nil {
		// already exist
		shadows := make([]string, 0, len(segments))
		var regxIndex int
		for _, seg := range segments {
			if seg[0] == ':' {
				param := pathVariables[regxIndex]
				shadows = append(shadows, ":"+param.key)
				continue
			}
			shadows = append(shadows, seg)
		}
		panic(fmt.Sprintf("pattern %s conflict with %s", pattern, strings.Join(shadows, "/")))
	}

	if pattern == "/" {
		r.trie.handle = handle
		return
	}
	r.trie.add(segments[1:len(segments)-1], handle)
}

// TODO: seems not to use
func findKeyByValue(dict map[string]string, val string) (string, bool) {
	for k, v := range dict {
		if v == val {
			return k, true
		}
	}
	return "", false
}

func (r *route) getHandle(pattern string) (Handle, PathVars) {
	pathVariables := make(PathVars, 0, 1)
	parts := strings.Split(pattern, "/")
	for _, branch := range r.trie.branches {
		h := branch.match(parts[1:len(parts)-1], pathVariables)
		if h != nil {
			return h, pathVariables
		}
	}
	return nil, nil
}
