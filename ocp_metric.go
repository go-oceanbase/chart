package main

import (
	"encoding/json"
	"sort"
	"time"
)

type TimeData struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

func parseTimeData(data ...[]TimeData) interface{} {
	ret := make([][]interface{}, 0, len(data))
	for _, it := range data {
		sort.Slice(it, func(i, j int) bool {
			return it[i].Timestamp < it[j].Timestamp
		})
		for _, vs := range it {
			t := time.Unix(vs.Timestamp, 0)
			ts := t.Format("15:04:05")
			ret = append(ret, []interface{}{ts, vs.Value})
		}
		break
	}
	return ret
}

func decodeTimeDatas(input []byte) ([]TimeData, error) {
	ret := []TimeData{}
	err := json.Unmarshal(input, &ret)
	return ret, err
}