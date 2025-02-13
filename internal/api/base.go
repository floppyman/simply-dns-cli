package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/umbrella-sh/um-common/logging/ulog"
)

func getRequest(path string) ([]byte, error) {
	req, err := createRequest(http.MethodGet, path, nil)
	if err != nil {
		ulog.Console.Debug().Msg("1")
		return nil, err
	}

	return doRequest(req)
}

func postRequest(path string, record *SimplyDnsRecord) ([]byte, error) {
	bodyBytes, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}

	req, err := createRequest(http.MethodPost, path, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	return doRequest(req)
}

func putRequest(path string, record *SimplyDnsRecord) ([]byte, error) {
	bodyBytes, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}

	req, err := createRequest(http.MethodPut, path, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	return doRequest(req)
}

func deleteRequest(path string) ([]byte, error) {
	req, err := createRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return doRequest(req)
}

func createRequest(method string, path string, reader io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, createFullUrl(path), reader)
	if err != nil {
		ulog.Console.Debug().Msg("1")
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", currentConfig.AccountNumber, currentConfig.AccountApiKey)))))

	return req, nil
}

func doRequest(req *http.Request) ([]byte, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	bodyStr, err := readBody(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("request failed with status code %d | %s | %s", res.StatusCode, res.Status, bodyStr))
	}

	return bodyStr, nil
}

func readBody(body io.ReadCloser) ([]byte, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func createFullUrl(path string) string {
	res, _ := url.JoinPath(currentConfig.Url, path)
	return res
}
