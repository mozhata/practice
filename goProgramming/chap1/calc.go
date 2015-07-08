package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"practice/pacpackge/sorter/bubblesort"
	"practice/pacpackge/sorter/qsort"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted value")
var algorthm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open input file ", infile)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}
func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to creat the output file ", outfile)
		return err
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm = ", *algorthm)
	}
	values, err := readValues(*infile)
	if err == nil {
		// fmt.Println("Read values: ", values)
		t1 := time.Now()
		switch *algorthm {
		case "qsort":
			for i := 0; i < 1000; i++ {
				qsort.QuickSort(values)
			}
		case "bubblesort":
			for i := 0; i < 1000; i++ {
				bubblesort.BubbleSort(values)
			}
		default:
			fmt.Println("Sorting algorithm", *algorthm, "is either unkown or unsupported")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs,", t2.Sub(t1), "to complete.")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
	fmt.Println(Quick2Sort([]int{11, 2, 33, 4, 55, 6, 77, 7}))

}
