package handler

import (
	"net/http"
	"time"

	"github.com/antonholmquist/jason"
)

const remoteSourceData = "https://www.cbr-xml-daily.ru/daily_json.js"

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetRemoteDataValute() (*jason.Object, error) {
	r, err := myClient.Get(remoteSourceData)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	res, _ := jason.NewObjectFromReader(r.Body)

	return res, nil
}
