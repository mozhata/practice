package mmux

import (
	"fmt"
	"net/http"
	"strings"
)

type Handle func(http.ResponseWriter, *http.Request, map[string]string)

type route struct {
	*trie
}

type trie struct {
	segment  string
	handle   Handle
	branches []*trie
}

// return Handle, score, matched
func (t *trie) match(segments []string, pathVariables map[string]string) Handle {
	if len(segments) == 0 || t.segment == "" && len(t.branches) == 0 {
		return nil
	}
	if len(segments) == 1 {
		if segments[0] == t.segment {
			return t.handle
		}
		if t.segment[0] == ':' {
			pathVariables[t.segment[1:]] = segments[0]
			return t.handle
		}
		return nil
	}
	// len(segments) > 1
	if t.segment[0] == ':' {
		pathVariables[t.segment[1:]] = segments[0]
	} else if t.segment != segments[0] {
		return nil
	}
	var (
		handlFunc Handle
	)
	for _, tree := range t.branches {
		handlFunc = tree.match(segments[1:], pathVariables)
		if handlFunc != nil {
			return handlFunc
		}
	}
	// not match
	delete(pathVariables, t.segment[1:])
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

// TODO check every panic to make sure logic
func (r *route) add(segments []string, handle Handle) {
	if handle == nil {
		panic("handler should not empty. register pattern: " + strings.Join(segments, "/"))
	}
	pathVariables := make(map[string]string)
	for _, branch := range r.branches {
		if h := branch.match(segments, pathVariables); h != nil {
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
			existPattern := "/" + strings.Join(shadows, "/") + "/"
			panic(fmt.Sprintf("pattern %s conflict with %s", strings.Join(segments, "/"), existPattern))
		}
	}
	r.trie.add(segments, handle)
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
	pathVariables := make(map[string]string)
	parts := strings.Split(pattern, "/")
	for _, branch := range r.trie.branches {
		h := branch.match(parts[1:len(parts)-1], pathVariables)
		if h != nil {
			return h, pathVariables
		}
	}
	return nil, nil
}
