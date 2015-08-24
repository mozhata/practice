package mp

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {

	var P Player

	switch mtype {
	case "MP3":
		P = &MP3Player{}
	case "WAV":
		P = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
	}

	P.Play(source)

	var p Player
	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source)

}
