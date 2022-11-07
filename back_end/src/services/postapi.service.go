package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	entity "github.com/bmolina1993/mailbox/src/entities"
)

func PostApiBulkData(data entity.QueryMail) {
	URL := "http://localhost:4080/api/_bulkv2"

	//parseo struct a json para post api
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v", string(dataJson))

	//post api
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(dataJson))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println("-->LOG POST API<--")
	log.Println("resp.StatusCode:", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
