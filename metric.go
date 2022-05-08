package main

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/sirupsen/logrus"
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
	var raw []map[string]float64
	err := json.Unmarshal(input, &raw)
	if err != nil {
		return nil, err
	}
	key := ""
	for k, _ := range raw[0] {
		if k != "timestamp" {
			key = k
		}
	}
	logrus.Infof("metric: %s", key)
	for _, it := range raw {
		ret = append(ret, TimeData{
			Timestamp: int64(it["timestamp"]),
			Value:     it[key],
		})
	}
	return ret, err
}
