package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	entity "github.com/bmolina1993/mailbox/src/entities"
)

func GetIdsByUserFolder(userFolder string, log bool) ([]entity.HitsData, error) {
	URL_GET := "http://localhost:4080/api/" + userFolder + "/_search"
	var getDocProps entity.GetDocProps //instanciado
	AUX_BODY_GET := entity.GetDocBody{
		Search_type: "matchall",
		From:        0,
		Max_results: 517424, //total archivos
	}

	//parseo struct a json para post api
	BODY_GET, err := json.Marshal(AUX_BODY_GET)
	if err != nil {
		return []entity.HitsData{}, err
	}

	req, err := http.NewRequest("POST", URL_GET, bytes.NewBuffer(BODY_GET))
	if err != nil {
		return []entity.HitsData{}, err
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []entity.HitsData{}, err
	}

	//se cierra despues de terminar todo el proceso
	defer resp.Body.Close()

	//decodifica la respuesta
	json.NewDecoder(resp.Body).Decode(&getDocProps)
	data := getDocProps.Hits.Hits

	//muestra o no el LOG por parametro
	if log {
		fmt.Println("\n--> LOG GET DOCUMENTS API <--")
		fmt.Println("userFolder:", userFolder)
		fmt.Println("StatusCode:", resp.StatusCode)
		fmt.Println("qtyIds:", len(data))
	}

	return data, nil
}
