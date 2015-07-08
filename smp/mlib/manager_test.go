package mlib

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManager failed,not empy!")
	}
	m0 := &MusicEntry{"1", "匆匆那年", "王菲",
		"http://7xityj.com1.z0.glb.clouddn.com/yearsbefore.mp3", "Mp3"}
	mm.Add(m0)
	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed .")
	}
	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find failed.")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist || m.Name != m0.Name ||
		m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager.Find() failed. Found item mismatch.")
	}
	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed. ", err)
	}
	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed.", err)
	}
}
