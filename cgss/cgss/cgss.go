package main

import (
	"fmt"
	"practice/cgss/cg"
	"practice/cgss/ipc"
	"strconv"
)

var centerClient *cg.CenterClient

func startCenterServer() error {
	server := ipc.NewIpcServer(&cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{client}
	return nil
}
func Help(args []string) int {
	fmt.Println(`
		Commands:
			login <username><level><exp>
			logout <username>
			listplayer
			quite(q)
			help(h)
		`)
	return 0
}
func Quit(args []string) int {
	return 1
}
func Logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("USAGE: logout <username>")
		return 0
	}
	centerClient.RemovePlayer(args[1])
}
func Login(args []string) int {
	if len(args) != 4 {
		fmt.Println("USAGE: login <username><level><exp>")
		return 0
	}
	level, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid Parameter: <exp> should be an interger.")
		return 0
	}
	palyer := cg.NewPlayer()
	palyer.Name = args[1]
	palyer.Lever = level
	palyer.Exp = exp
	err = centerClient.AddPlayer(palyer)
	if err != nil {
		fmt.Println("Failed adding player", err)
	}
	return 0
}
