package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GetDocBody struct {
	Search_type string
	From        int
	Size        int
}

type HitsData struct {
	//Index string `json:"_index"`
	Id string `json:"_id"`
}

type HitsProps struct {
	Hits []HitsData
}

type GetDocProps struct {
	//Timed_out bool
	Hits HitsProps
}

func GetIdsDocuments() GetDocProps {
	URL_GET := "http://localhost:4080/es/_search"
	AUX_BODY_GET := GetDocBody{
		Search_type: "matchall",
		From:        0,
		Size:        517424,
	}
	//parseo struct a json para post api
	BODY_GET, err := json.Marshal(AUX_BODY_GET)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", URL_GET, bytes.NewBuffer(BODY_GET))
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

	fmt.Println("--> LOG GET DOCUMENTS API <--")
	fmt.Println("resp.StatusCode:", resp.StatusCode)

	//instanciado
	var getDocProps GetDocProps

	json.NewDecoder(resp.Body).Decode(&getDocProps)

	//fmt.Printf("%+v \n", getDocProps)
	return getDocProps
}

func DeleteDocuments(userFolder string) {
	URL_DELETE := "http://localhost:4080/api/" + userFolder + "/_doc/"
	idDocumetns := GetIdsDocuments()
	listIds := idDocumetns.Hits.Hits

	fmt.Println("\n\n--> LOG DELETE DOCUMENTS API <--")
	for _, idDocument := range listIds {
		Id := idDocument.Id
		URL_DELETE_ID := URL_DELETE + Id

		req, err := http.NewRequest("DELETE", URL_DELETE_ID, bytes.NewBuffer([]byte{}))
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
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nuserFolder:", userFolder)
		fmt.Println("StatusCode:", resp.StatusCode)
		fmt.Println("body:", string(body))
	}
}
