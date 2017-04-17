package main

import (
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
)

type Weibo struct {
        Id           uint64
        Timestamp    uint64
        UserName     string
        RepostsCount uint64
        Text         string
}

var searcher engine.Engine
searcher.Init(types.EngineInitOptions{
	SegmenterDictionaries: "../../data/dictionary.txt",
	StopTokenFile:         "../../data/stop_tokens.txt",
	IndexerInitOptions: &types.IndexerInitOptions{
		IndexType: types.LocationsIndex,
	},
})
func main() {
	
}