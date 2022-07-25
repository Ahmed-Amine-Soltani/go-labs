// This is a client for PoetryDb API (https://poetrydb.org).

package poetrydb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// These global variables makes it easy
// to mock these dependencies
// in unit tests.
var (
	jsonUnmarshal = json.Unmarshal
	ioUtilReadAll = ioutil.ReadAll
)

// PoetryResponse holds data about a poetry.
type PoetryResponse struct {
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	Lines     []string `json:"lines"`
	LineCount string   `json:"linecount"`
}

// PoetriesResponse contains a list of poetries.
type PoetriesResponse struct {
	List []PoetryResponse `json:"list"`
}

// PoetryDb interface defines the available endpoints
// for PoetryDb API.
type PoetryDb interface {
	// Random returns a given number of random poetries.
	Random(number int) (PoetriesResponse, error)
}

// HttpClient is an interface that makes it easier to
// mock http.Client in unit tests.
type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

// poetryDb implements PoetryDb interface.
type poetryDb struct {
	baseUrl    string
	httpClient HttpClient
}

func (p *poetryDb) Random(number int) (PoetriesResponse, error) {
	var poetryResponseList PoetriesResponse
	poetryResponse := make([]PoetryResponse, 0)
	endpoint := fmt.Sprintf("%s/random/%d", p.baseUrl, number)
	response, err := p.httpClient.Get(endpoint)
	if err != nil {
		return poetryResponseList, errors.Wrapf(err, `performing request to "%s"`, endpoint)
	}
	defer response.Body.Close()
	statusCode := response.StatusCode
	if !(statusCode >= 200 && statusCode <= 299) {
		return poetryResponseList, errors.New(fmt.Sprintf("got status code %d", statusCode))
	}
	body, err := ioUtilReadAll(response.Body)
	if err != nil {
		return poetryResponseList, errors.Wrap(err, "reading response")
	}
	err = jsonUnmarshal(body, &poetryResponse)
	if err != nil {
		return poetryResponseList, errors.Wrap(err, "parsing response")
	}
	poetryResponseList.List = poetryResponse
	return poetryResponseList, nil
}

// NewPoetryDb creates a new PoetryDb client.
func NewPoetryDb(baseUrl string, timeoutInSeconds int) PoetryDb {
	httpClient := &http.Client{
		Timeout: time.Duration(timeoutInSeconds) * time.Second,
	}
	return &poetryDb{
		baseUrl:    baseUrl,
		httpClient: httpClient,
	}
}
