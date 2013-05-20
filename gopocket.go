package gopocket

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

const (
	pocketAddUrl    = "https://getpocket.com/v3/add"
	pocketModifyUrl = "https://getpocket.com/v3/send"
)

type Pocket struct {
	key, token string
}

func Init(consumerKey, accessToken string) *Pocket {
	return &Pocket{key: consumerKey, token: accessToken}
}

func NewBatch() *Batch {
	return &Batch{}
}

func (p *Pocket) Add(url, title string, tags []string) (*AddResponse, *ApiRate, error) {
	data := struct {
		ConsumerKey string   `json:"consumer_key"`
		AccessToken string   `json:"access_token"`
		Url         string   `json:"url"`
		Title       string   `json:"title,omitempty"`
		Tags        []string `json:"tags"`
	}{
		p.key, p.token, url, title, tags,
	}

	var resp *AddResponse
	rate, err := p.post(pocketAddUrl, &data, &resp)
	return resp, rate, err
}

func (p *Pocket) Modify(batch *Batch) (*ModifyResponse, *ApiRate, error) {
	data := struct {
		ConsumerKey string        `json:"consumer_key"`
		AccessToken string        `json:"access_token"`
		Actions     []interface{} `json:"actions"`
	}{
		p.key, p.token, batch.actions,
	}

	var resp *ModifyResponse
	rate, err := p.post(pocketModifyUrl, &data, &resp)
	return resp, rate, err
}

func (p *Pocket) post(url string, requestModel interface{}, result interface{}) (rate *ApiRate, err error) {
	// marshal the request struct to a json object within an io.Reader
	body, err := marshalRequest(requestModel)
	if err != nil {
		return
	}

	// create a POST request to pocket
	req, err := createRequest(url, body)
	if err != nil {
		return
	}

	// execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	// read & unmarshal the response into the expected result, along with the current api rate limits
	defer resp.Body.Close()
	rate = apiRateFromResponse(resp)
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(contents, &result)
	return
}

func marshalRequest(request interface{}) (io.Reader, error) {
	rawBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(rawBody), nil
}

func createRequest(url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest("POST", url, body)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Accept", "application/json")
	return
}

func apiRateFromResponse(response *http.Response) *ApiRate {
	var rate = &ApiRate{}
	t := reflect.ValueOf(rate).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Type().Field(i)
		value := t.Field(i)
		tag := field.Tag.Get("header")
		header := response.Header.Get(tag)

		switch field.Type.Name() {
		case "int":
			num, _ := strconv.ParseInt(header, 10, 32)
			value.SetInt(num)
		case "string":
			value.SetString(header)
		case "int64":
			num, _ := strconv.ParseInt(header, 10, 64)
			value.SetInt(num)
		}
	}
	return rate
}
