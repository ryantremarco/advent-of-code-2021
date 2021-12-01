package util

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readInputFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	str := (string)(bytes)
	return strings.Split(str, "\n"), nil
}

func ReadInputStrings(path string) ([]string, error) {
	return readInputFile(path)
}

func ReadInputInts(path string) ([]int, error) {
	strs, err := readInputFile(path)
	if err != nil {
		return nil, err
	}

	ints := make([]int, len(strs))
	for i, str := range strs {
		if str == "" {
			continue
		}

		conv, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		ints[i] = conv
	}

	return ints, nil
}
