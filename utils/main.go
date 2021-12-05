package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func ReadInputFile(path string) string {
	dir := "/home/nate/www/adventofcode/2021"
	filePath := fmt.Sprintf("%s/%s", dir, path)
	log.Println("Reading Data from: ", filePath)
	contents, _ := ioutil.ReadFile(filePath)
	return string(contents)
}

func SplitByLine(content string) []string {
	return strings.Split(content, "\n")
}

func StringsToInts(content []string) []int64 {
	var vals []int64
	for _, val := range content {
		trimmed := strings.TrimSpace(val)
		if trimmed != "" {
			intVal, _ := strconv.Atoi(trimmed)
			vals = append(vals, int64(intVal))
		}
	}
	return vals
}

func StringsToIntsNormal(content []string) []int {
	var vals []int
	for _, val := range content {
		trimmed := strings.TrimSpace(val)
		if trimmed != "" {
			intVal, _ := strconv.Atoi(trimmed)
			vals = append(vals, intVal)
		}
	}
	return vals
}

func SetField(v interface{}, name string, value string) error {
	// v must be a pointer to a struct
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return errors.New("v must be pointer to struct")
	}

	// Dereference pointer
	rv = rv.Elem()

	// Lookup field by name
	fv := rv.FieldByName(name)
	if !fv.IsValid() {
		return fmt.Errorf("not a field name: %s", name)
	}

	// Field must be exported
	if !fv.CanSet() {
		return fmt.Errorf("cannot set field %s", name)
	}

	// We expect a string field
	if fv.Kind() != reflect.String {
		return fmt.Errorf("%s is not a string field", name)
	}

	// Set the value
	fv.SetString(value)
	return nil
}

func FindMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
