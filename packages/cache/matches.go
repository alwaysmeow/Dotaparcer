package cache

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CacheMatches(numbers []int) error {
	file, err := os.Create("./cache/matches.txt")
	if err != nil {
		return fmt.Errorf("can't create file: %w", err)
	}
	defer file.Close()

	ids := unique(numbers)

	for _, number := range ids {
		_, err := fmt.Fprintf(file, "%d\n", number)
		if err != nil {
			return fmt.Errorf("can't write into file: %w", err)
		}
	}

	return nil
}

func LoadCachedMatches() ([]int, error) {
	file, err := os.Open("./cache/matches.txt")
	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("не удалось преобразовать строку в число: %w", err)
		}
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	return numbers, nil
}

func unique(arr []int) []int {
	seen := make(map[int]struct{})
	uniqueArr := []int{}

	for _, val := range arr {
		_, ok := seen[val]
		if !ok {
			seen[val] = struct{}{}
			uniqueArr = append(uniqueArr, val)
		}
	}

	return uniqueArr
}

func UniqueMatches() {
	ids, _ := LoadCachedMatches()
	ids = unique(ids)
	CacheMatches(ids)
}
