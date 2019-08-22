package process

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func postRequest(url string, data io.Reader) (string, bool) {
	return httpModRequest(url, http.MethodPost, true, data)
}

func putRequest(url string, data io.Reader) bool {
	_, result := httpModRequest(url, http.MethodPut, false, data)
	return result
}

func httpModRequest(url, method string, needResult bool, data io.Reader) (string, bool) {
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

func delRequest(url string) bool {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if resp.StatusCode >= 400 {
		log.Fatal(fmt.Sprintf("Delete request not success, status: %s, return code: %v", resp.Status, resp.StatusCode))
		return false
	}

	defer resp.Body.Close()

	return true
}
