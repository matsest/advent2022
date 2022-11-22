package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadInts(path string) ([]int, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	contents := string(bytes)
	s := strings.Split(contents, ",")
	var out []int

	for v := range s {
		i, err := strconv.Atoi(s[v])
		if err != nil {
			return nil, err
		}
		out = append(out, i)
	}

	return out, nil
}

func SliceAtoi(sa []string) ([]int, error) {
    si := make([]int, 0, len(sa))
    for _, a := range sa {
        i, err := strconv.Atoi(a)
        if err != nil {
            return si, err
        }
        si = append(si, i)
    }
    return si, nil
}