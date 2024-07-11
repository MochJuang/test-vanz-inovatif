package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type FruitType string

const (
	FruitTypeImport FruitType = "IMPORT"
	FruitTypeLocal  FruitType = "LOCAL"
)

type Fruits struct {
	ID    int       `json:"fruitId"`
	Name  string    `json:"fruitName"`
	Type  FruitType `json:"fruitType"`
	Stock int       `json:"Stock"`
}

func case_1() {
	fruits, err := loadFruitsFromJSON("fruits.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("Fruits:", fruits)

	fruitNames := getUniqueFruitNames(fruits)
	fmt.Println("1). Fruits owned by Andi:", fruitNames)
	fmt.Println()

	fruitByType, stockByType := separateFruitsByType(fruits)
	fmt.Println("2). Number of containers needed:", len(fruitByType))
	fmt.Println()

	fmt.Println("3). Containers:")
	for t, names := range fruitByType {
		fmt.Printf("Container for type %s contains fruits: %v, with total stock: %d\n", t, names, stockByType[t])
	}
	fmt.Println()

	fmt.Println("Comments:")
	fmt.Println("1. Some fruits have the same name with different uppercase/lowercase letters, they should be standardized.")
	fmt.Println("2. There are the same IDs for several fruits, each fruit should have a unique ID.")
	fmt.Println("3. Some JSON formats mistakenly use 'colon', they should use \"colon\".")
}

func loadFruitsFromJSON(filePath string) ([]Fruits, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var fruits []Fruits
	err = json.Unmarshal(bytes, &fruits)
	if err != nil {
		return nil, err
	}
	return fruits, nil
}

func getUniqueFruitNames(fruits []Fruits) []string {
	nameSet := make(map[string]struct{})
	var uniqueNames []string
	for _, fruit := range fruits {
		normalized := strings.ToLower(fruit.Name)
		if _, exists := nameSet[normalized]; !exists {
			nameSet[normalized] = struct{}{}
			uniqueNames = append(uniqueNames, fruit.Name)
		}
	}
	return uniqueNames
}

func separateFruitsByType(fruits []Fruits) (map[FruitType][]string, map[FruitType]int) {
	fruitByType := make(map[FruitType][]string)
	stockByType := make(map[FruitType]int)
	nameSet := make(map[string]struct{})
	for _, fruit := range fruits {
		normalized := strings.ToLower(fruit.Name)
		if _, exists := nameSet[normalized]; !exists {
			nameSet[normalized] = struct{}{}
			fruitByType[fruit.Type] = append(fruitByType[fruit.Type], fruit.Name)
			stockByType[fruit.Type] += fruit.Stock
		} else {
			stockByType[fruit.Type] += fruit.Stock
		}
	}

	return fruitByType, stockByType
}
