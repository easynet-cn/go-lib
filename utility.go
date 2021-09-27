package golib

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}

	return m
}

func StringSliceToMap(strs []string) map[string]string {
	m := make(map[string]string)

	for _, v := range strs {
		m[v] = v
	}

	return m
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func JsonToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)

	err := json.Unmarshal([]byte(jsonStr), &m)

	return m, err
}

func MapToJson(m map[string]string) (string, error) {
	if bytes, err := json.Marshal(m); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func FormatFloat(f float64, scale int) float64 {
	result, _ := strconv.ParseFloat(strconv.FormatFloat(f, 'f', scale+1, 64), 64)

	pow := math.Pow(10, float64(scale))

	return math.Round(result*pow) / pow
}

func SlicePage(pageIndex int, pageSize int, total int64) (start, end int64) {
	if pageIndex < 0 {
		pageIndex = 1
	}

	if pageSize < 0 {
		pageSize = 20
	}

	if int64(pageSize) > total {
		return 0, total
	}

	// total pages
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	if pageIndex > totalPages {
		return 0, 0
	}

	start = int64((pageIndex - 1) * pageSize)
	end = start + int64(pageSize)

	if end > total {
		end = total
	}

	return start, end
}
