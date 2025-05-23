package main

import (
	"sort"
	"strings"
)

func HasPrefix[s ~string](str s, prefix string) bool {
	return strings.HasPrefix(string(str), prefix)
}

func TrimPrefix[s ~string](str s, prefix string) s {
	return s(strings.TrimPrefix(string(str), prefix))
}

func HasSuffix[s ~string](str s, suffix string) bool {
	return strings.HasSuffix(string(str), suffix)
}

func TrimSuffix[s ~string](str s, suffix string) s {
	return s(strings.TrimSuffix(string(str), suffix))
}

func ReplaceAll[s, t, u ~string](str s, old t, new u) s {
	return s(strings.ReplaceAll(string(str), string(old), string(new)))
}

func Contains[t ~string, s ~string](str s, substr t) bool {
	return strings.Contains(string(str), string(substr))
}

func Split[s ~string](str s, sep string) []s {
	var ret []s
	for _, x := range strings.Split(string(str), sep) {
		ret = append(ret, s(x))
	}
	return ret
}

func Replace[s, ot, nt ~string](str s, old ot, new nt, n int) s {
	return s(strings.Replace(string(str), string(old), string(new), n))
}

func Join[s ~string](strs []s, sep string) s {
	return s(strings.Join(func() (ret []string) {
		for _, x := range strs {
			ret = append(ret, string(x))
		}
		return
	}(), sep))
}

func ContainsAny[s ~string](str s, chars string) bool {
	return strings.ContainsAny(string(str), chars)
}

func Index[s ~string](str s, substr string) int {
	return strings.Index(string(str), substr)
}

func Capitalize[s ~string](str s) s {
	return s(strings.ToUpper(string(str[0]))) + str[1:]
}

func SortStrings[s ~string](str []s) {
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
}

func SliceToMap[T comparable](s []T) map[T]bool {
	m := make(map[T]bool, len(s))
	for _, v := range s {
		m[v] = true
	}
	return m
}

func MergeMaps[T comparable](maps ...map[T]bool) map[T]bool {
	result := make(map[T]bool)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func RemoveMapValues[T comparable, S any](m map[T]S) map[T]bool {
	result := make(map[T]bool)
	for k := range m {
		result[k] = true
	}

	return result
}

func MapContainsAny[T ~string, S ~string](key T, dict []S) bool {
	for _, v := range dict {
		if Contains(key, v) {
			return true
		}
	}

	return false
}
