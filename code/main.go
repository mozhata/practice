package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	// tryShit()
	// tryPrintNilHashSum()
	// tryHashWriteSum()
	PrintMD5("../chap1/chap1.go")
}

func tryShit() {
	src := []int{0}
	shiftZero(src)
}

// []int{1,0,0,4,0}, 非零元素移到到左边

func shiftZero(arr []int) {
	if len(arr) < 1 {
		return
	}
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] != 0 {
			left++
		}
		if left >= right {
			break
		}
		if arr[right] == 0 {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	fmt.Println(arr)
}

func tryPrintNilHashSum() {
	printHash()
	printHash()
	printHash()
}

func printHash() {
	hash := sha1.New()
	hb := hash.Sum(nil)
	fmt.Println(hex.EncodeToString(hb))
}

func tryHashWriteSum() {
	str := "slkgsdhjl;kjghhkjlkjhjg"
	h := md5.New()
	h.Write([]byte(str))
	fmt.Printf("md5 for write: %s\n", hex.EncodeToString(h.Sum(nil)))
	bs := md5.Sum([]byte(str))
	fmt.Printf("md5.Sum:  %s\n", hex.EncodeToString(bs[:]))
	fmt.Printf("md5 sum for prefix: %s\n", hex.EncodeToString(h.Sum([]byte{1, 2})))
}

func PrintMD5(fileName string) {
	fh, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	st, err := fh.Stat()
	if err != nil {
		panic(err)
	}
	h := md5.New()
	size, err := io.Copy(h, fh)
	if err != nil {
		panic(err)
	}
	// ioutil.ReadFile(fileName)
	sum := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("%s\t %s\n", fileName, sum)
	fmt.Printf("size by copy: %d, size by fh: %d\n", size, st.Size())
}
