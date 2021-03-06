// Gopocket provides a Go wrapper around the getpocket API (http://www.getpocket.com)
package gopocket

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

const (
	pocketAddUrl      = "https://getpocket.com/v3/add"
	pocketModifyUrl   = "https://getpocket.com/v3/send"
	pocketRetrieveUrl = "https://getpocket.com/v3/get"
)

// Pocket type performs all API calls, created by Init function.
type Pocket struct {
	key, token string
	post       func(url string, requestModel interface{}, result interface{}) (rate *ApiRate, err error)
}

// Init creates a new Pocket type with the provided app consumer key and user access token.
func Init(consumerKey, accessToken string) *Pocket {
	return &Pocket{key: consumerKey, token: accessToken, post: post}
}

// NewBatch creates a new Batch type for requesting multiple operations in a single request.
func NewBatch() *Batch {
	return &Batch{}
}

// NewOptions creates a new Options type for retrieving user articles.
func NewOptions() *Options {
	return &Options{dict: make(map[string]interface{})}
}

// Add submits a url to a users pocket queue. Title will be used as a fallback if the API couldn't accurately parse the
// title of the article, sending an empty string is acceptable. If you have multiple urls to add at once,
// create a Batch & use the Modify method.
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

// Modify submits an array of operations in a single request to the API. See the Batch methods for more details.
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

// Retrieve submits a request to return a filtered array of articles. See the Options type for more details.
func (p *Pocket) Retrieve(opts *Options) (*RetrieveResponse, *ApiRate, error) {
	opts.dict["consumer_key"] = p.key
	opts.dict["access_token"] = p.token
	defer func() {
		delete(opts.dict, "consumer_key")
		delete(opts.dict, "access_token")
	}()

	var resp *RetrieveResponse
	rate, err := p.post(pocketRetrieveUrl, &opts.dict, &resp)
	return resp, rate, err
}

func post(url string, requestModel interface{}, result interface{}) (rate *ApiRate, err error) {
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
	defer resp.Body.Close()

	// grab the error information along with the current api rate limits.
	rate = apiRateFromResponse(resp)
	if rate.Error != "" {
		err = errors.New(rate.Error)
		return
	}

	// read & unmarshal the response into the expected result
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

func apiRateFromResponse(response *http.Response) (rate *ApiRate) {
	rate = new(ApiRate)
	t := reflect.ValueOf(rate).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Type().Field(i)
		value := t.Field(i)
		tag := field.Tag.Get("header")
		header := response.Header.Get(tag)

		if header != "" {
			switch field.Type.Name() {
			case "int":
				num, _ := strconv.ParseInt(header, 10, 32)
				value.SetInt(num)
			case "string":
				value.SetString(header)
			}
		}
	}

	return
}
