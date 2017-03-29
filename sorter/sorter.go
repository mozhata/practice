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

var (
	infile    *string = flag.String("i", "infile", "File contains values for sorting")
	outfile   *string = flag.String("o", "outfile", "File to receive sorted values")
	algorithm *string = flag.String("a", "qsort", "Sort algorithm")
)

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile= ", *infile, "outfile= ", outfile, "algorithm= ", algorithm)
	}
	values, err := readValues(*infile)
	if err == nil {

		// 返回当前时间 eg: 2015-08-15 20:54:41.970192591 +0800 CST

		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unkown or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
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
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}