package common

import "sort"

type kv struct {
	Key   string
	Value int
}

func ReverseSortByValue(tosort map[string]int) ([]kv, error) {
	ss, err := SortByValue(tosort)
	if ss != nil {
		ss = sort.Reverse(ss)
	}
	return ss, err
}

func SortByValue(tosort map[string]int) ([]kv, error) {
	var ss []kv
	for k, v := range tosort {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	return ss, nil
}
