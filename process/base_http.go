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
		log.Print(err)
		return respText, false
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return respText, false
	}

	if resp.StatusCode >= 400 {
		log.Print(fmt.Sprintf("%s request not success for %s, status: %s, return code: %v", method, url, resp.Status, resp.StatusCode))
		return respText, false
	}

	if needResult {
		content, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		resp.Body.Close()

		return string(content), true
	}

	return respText, true

}
