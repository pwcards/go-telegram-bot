package models

const (
	TimeKeySend08 = "start_08"
	TimeKeySend09 = "start_09"
	TimeKeySend10 = "start_10"
)

type TimeSend struct {
	TimeValue string `json:"time_value"`
}

var TimeMap = map[string]TimeSend{
	TimeKeySend08: {
		TimeValue: "08:00",
	},
	TimeKeySend09: {
		TimeValue: "09:00",
	},
	TimeKeySend10: {
		TimeValue: "10:00",
	},
}

func GetTimeMapValue(key string) string {
	return TimeMap[key].TimeValue
}
