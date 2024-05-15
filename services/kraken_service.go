package services

import (
	"fmt"
	"io"
	"main/helpers"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)


type KrakenApi struct {
	Key    string
	Secret string
	Client *http.Client
}

const baseUrl = "https://api.kraken.com"

func (kraken *KrakenApi) GetBalance() ([]byte, error) {
	uri := "/0/private/Balance"
	method := "POST"

	params := url.Values{}
	params.Add("nonce", fmt.Sprintf("%d", time.Now().UnixMilli()))
	
	req, err := http.NewRequest(method,
		baseUrl+uri,
		strings.NewReader(params.Encode()))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	signature := helpers.GetSignature(uri, params)

	req.Header.Set("API-Key", os.Getenv("KRAKEN_API_KEY"))
	req.Header.Set("API-Sign", signature)

	fmt.Println(req)
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(string(body))
	return body, nil
}

func (kraken *KrakenApi) GetBalanceExtend() ([]byte, error) {
	uri := "/0/private/BalanceEx"
	method := "POST"

	params := url.Values{}
	params.Add("nonce", fmt.Sprintf("%d", time.Now().UnixMilli()))
	
	req, err := http.NewRequest(method,
		baseUrl+uri,
		strings.NewReader(params.Encode()))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	signature := helpers.GetSignature(uri, params)

	req.Header.Set("API-Key", os.Getenv("KRAKEN_API_KEY"))
	req.Header.Set("API-Sign", signature)

	fmt.Println(req)
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(string(body))
	return body, nil
}