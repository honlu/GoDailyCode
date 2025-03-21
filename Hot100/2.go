package hot100

import "sort"

func groupAnagrams(strs []string) [][]string {
	// key: 相同字母的string, value：结果列表
	sameStrsMap := map[string][]string{}
	for i, str := range strs {
		temp := []byte(str)
		sort.Slice(temp, func(i, j int) bool {
			if temp[i] > temp[j] {
				return true
			}
			return false
		})
		sameStrsMap[string(temp)] = append(sameStrsMap[string(temp)], strs[i])
	}
	var res [][]string
	for _, value := range sameStrsMap {
		res = append(res, value)
	}
	return res
}
