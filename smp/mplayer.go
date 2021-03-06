package main

import (
	"bufio"
	"fmt"
	"os"
	"practice/smp/mlib"
	"practice/smp/mp"
	"strconv"
	"strings"
)

var lib *mlib.MusicManager
var id int = 1
var ctrl, siginal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&mlib.MusicEntry{strconv.Itoa(id), tokens[2],
				tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {

			name, _ := strconv.Atoi(tokens[2])
			lib.Remove(name)

			lib.Remove(tokens[2])

		} else {
			fmt.Println("USAGE: lib remove<id>")
		}
	default:
		fmt.Println("Unrecognized command:", tokens[1])
	}
}
func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE:lib paly<name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	mp.Play(e.Source, e.Type)
}

func main() {

func mian() {

	fmt.Println(`
		Enter following commands to control the player:
		lib list -- View the existing musicx lib
		lib add <name><artist><source><type> -- Add a music to the music lib
		lib remove <name> --Remove the specified music from the lib
		play <name> -- Play the specified music
		`)


	lib = mlib.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command -> ")


	lib = mlib.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command -> ")

		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized commad: ", tokens[0])
		}
	}
}



// type ReadWriter interface {
// 	Read(buf []byte) (n int, err error)
// 	Write(buf []byte) (n int, err error)
// }
// type Deleter interface {
// 	Delete(soucre string) (err, error)
// }
// type Booker interface {
// 	ReadWriter
// 	Deleter
// }

