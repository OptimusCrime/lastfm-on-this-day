package lastfm

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/optimuscrime/lastfm-on-this-day/pgk/config"
)

type requester struct {
	config *config.Config
}

var apiUrl = "http://ws.audioscrobbler.com/2.0/"

func (r *requester) makeLastFmRequest(operation string, params map[string]string) ([]byte, error) {
	requestUrl := r.supplyDefaultParamsAndBuildRequestUrl(operation, params)

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "LastFm On This Day (github.com/optimuscrime/lastfm-on-this-day)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = resp.Body.Close(); err != nil {
		return nil, err
	}

	if resp.StatusCode == 500 {
		return nil, errors.New("server responded with 500")
	}

	return bodyBytes, nil
}

func (r *requester) supplyDefaultParamsAndBuildRequestUrl(operation string, params map[string]string) string {
	params["api_key"] = r.config.ApiKey
	params["method"] = operation

	params["api_sig"] = r.createSignedChecksum(params)

	// Note: The format parameter should not be a part of the signed checksum calculation for some (undocumented) reason
	params["format"] = "json"

	return buildRequestUrl(params)
}

func buildRequestUrl(params map[string]string) string {
	queryParams := make([]string, 0, len(params))
	for k, v := range params {
		queryParams = append(queryParams, k+"="+v)
	}

	joinedQueryParams := strings.Join(queryParams, "&")

	return apiUrl + "?" + joinedQueryParams
}

func (r *requester) createSignedChecksum(params map[string]string) string {
	mapKeys := make([]string, 0, len(params))
	for k := range params {
		mapKeys = append(mapKeys, k)
	}

	sort.Strings(mapKeys)

	joinedKeyValuePairs := make([]string, 0, len(params))
	for _, k := range mapKeys {
		joinedKeyValuePairs = append(joinedKeyValuePairs, fmt.Sprintf("%s%s", k, params[k]))
	}

	joinedParamsString := strings.Join(joinedKeyValuePairs, "")

	md5Hash := md5.Sum([]byte(joinedParamsString + r.config.SharedSecret))
	return fmt.Sprintf("%x", md5Hash)
}
