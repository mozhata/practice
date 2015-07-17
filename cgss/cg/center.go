package cg

import (
	"encoding/json"
	"errors"
	"practice/cgss/ipc"
	"sync"
)

// var _ ipc.Server = &CenterServer{}

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}
type CenterServer struct {
	servers map[string]ipc.Server
	palers  []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)
	return &CenterServer{servers: servers, players: players}
}
func (server *CenterServer) addPlayer(params string) error {
	player := NewPlayer()
	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.palers = append(server.palers, player)
}
func (server *CenterServer) removePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	for i, v := range server.palers {
		if v.Name == params {
			if len(server.palers) == 1 {
				server.palers = make([]*Player, 0)
			} else if i == len(server.palers)-1 {
				server.palers = server.palers[:i]
			} else if i == 0 {
				server.palers = server.palers[1:]
			} else {
				server.palers = append(server.palers[:i], server.palers[i+1:])
			}
			return nil
		}
	}
	return errors.New("Player not found")
}
func (server *CenterServer) listPlayer(params string) (palyers string, err errors) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	if len(server.palers) > 0 {
		b, _ := json.Marshal(server.palers)
		palyers = string(b)
	} else {
		err = errors.New("No players online")
	}
	return
}
func (server *CenterServer) broadcast(params string) error {
	var message Message
	err := json.Unmarshal([]byte(params), &message)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()
	if len(server.palers) > 0 {
		for _, palyer := range server.palers {
			palyer.mq <- &message
		}
	} else {
		err = errors.New("No player online")
	}
	return err
}
func (server *CenterServer) Handle(method, params string) *ipc.Response {
	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "removeplayer":
		err := server.removePlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "listplayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{"200", players}
	case "broadcast":
		err := server.broadcast(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{Code: "200"}
	default:
		return &ipc.Response{Code: "404", Body: method + ":" + params}
	}
	return &ipc.Response{Code: "200"}
}
func (server *CenterServer) Name() string {
	return "CenterServer"
}
