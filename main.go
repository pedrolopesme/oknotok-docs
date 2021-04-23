package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pedrolopesme/oknotok"
)

func test1() {
}

func main() {
	fmt.Println("basic sample testing on OkNotOK")
	settings := oknotok.Settings{
		Name:              "test-ci",
		MaxHalfOkRequests: 2,
		Interval:          time.Duration(2 * time.Second),
		Timeout:           time.Duration(5 * time.Second),
	}

	oknok := oknotok.NewOkNotOk(settings)

	body, err := oknok.Call(func() (interface{}, error) {
		resp, err := http.Get("https://httpstat.us/200")
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})

	if err != nil {
		fmt.Println("something bad happened", err)
	}

	fmt.Println("request sent", body)
}
