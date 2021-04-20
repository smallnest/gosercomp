package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	libs, tooks, marshaledBytes := readMarshalLog()
	tookFile, _ := os.OpenFile("marshal_took.csv", os.O_CREATE|os.O_RDWR, 0777)
	marshaledBytesFile, _ := os.OpenFile("marshaledBytes.csv", os.O_CREATE|os.O_RDWR, 0777)
	defer tookFile.Close()
	defer marshaledBytesFile.Close()

	for i, name := range libs {
		tookFile.WriteString(fmt.Sprintf("%s,%d\n", name, tooks[i]))
		marshaledBytesFile.WriteString(fmt.Sprintf("%s,%d\n", name, marshaledBytes[i]))
	}

	libs, tooks = readUnmarshalLog()
	tookFile2, _ := os.OpenFile("unmarshal_took.csv", os.O_CREATE|os.O_RDWR, 0777)
	defer tookFile2.Close()

	for i, name := range libs {
		tookFile2.WriteString(fmt.Sprintf("%s,%d\n", name, tooks[i]))
	}
}

func readMarshalLog() ([]string, []int, []int) {
	file, err := os.Open("marshal.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var libs []string
	var tooks []int
	var marshaledBytes []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "BenchmarkMarshalBy") {
			line = strings.TrimPrefix(line, "BenchmarkMarshalBy")
			fields := strings.Fields(line)
			name := fields[0]
			if strings.Index(name, "-") > 0 {
				name = name[:strings.Index(name, "-")]
			}
			libs = append(libs, name)
			tooks = append(tooks, int(s2f(fields[2])))
			marshaledBytes = append(marshaledBytes, int(s2f(fields[4])))
		}
	}

	return libs, tooks, marshaledBytes
}

func readUnmarshalLog() ([]string, []int) {
	file, err := os.Open("unmarshal.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var libs []string
	var tooks []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "BenchmarkUnmarshalBy") {
			line = strings.TrimPrefix(line, "BenchmarkUnmarshalBy")
			fields := strings.Fields(line)
			name := fields[0]
			if strings.Index(name, "-") > 0 {
				name = name[:strings.Index(name, "-")]
			}
			libs = append(libs, name)
			tooks = append(tooks, int(s2f(fields[2])))
		}
	}

	return libs, tooks
}

func s2i(s string) int {
	f, _ := strconv.Atoi(s)
	return f
}

func s2f(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
