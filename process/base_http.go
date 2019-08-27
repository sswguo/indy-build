package process

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func getRequest(url string) (string, bool) {
	return httpRequest(url, http.MethodGet, true, nil)
}

func postRequest(url string, data io.Reader) (string, bool) {
	return httpRequest(url, http.MethodPost, true, data)
}

func putRequest(url string, data io.Reader) bool {
	_, result := httpRequest(url, http.MethodPut, false, data)
	return result
}

func delRequest(url string) bool {
	_, result := httpRequest(url, http.MethodDelete, false, nil)
	return result
}

func httpRequest(url, method string, needResult bool, data io.Reader) (string, bool) {
	client := &http.Client{}
	respText := ""
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		log.Fatal(err)
		return respText, false
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return respText, false
	}

	if resp.StatusCode >= 400 {
		log.Fatal(fmt.Sprintf("%s request not success, status: %s, return code: %v", method, resp.Status, resp.StatusCode))
		return respText, false
	}

	if needResult {
		content, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		return string(content), true
	}

	defer resp.Body.Close()

	return respText, true

}
