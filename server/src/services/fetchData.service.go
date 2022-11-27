package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	entity "github.com/bmolina1993/mailbox/src/entities"
	"github.com/joho/godotenv"
)

func FetchData() (string, error) {
	URL := "http://localhost:4080/es/_search"

	//instancia de entidad
	var getDocProps entity.GetDocProps
	AUX_BODY_GET := entity.GetDocBody{
		Search_type: "matchall",
		From:        0,
		Size:        517424, //total archivos
	}

	//parseo struct a json para request
	BODY_GET, err := json.Marshal(AUX_BODY_GET)
	if err != nil {
		return "", err
	}

	//se hace request para obtener data recolectada en zinc
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(BODY_GET))
	if err != nil {
		return "", err
	}

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error cargando archivo .env")
	}
	user := os.Getenv("ZINC_USER")
	password := os.Getenv("ZINC_PASS")

	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	//se cierra despues de terminar todo el proceso
	defer resp.Body.Close()

	//decodifica la respuesta para acceder a datos de respuesta api
	json.NewDecoder(resp.Body).Decode(&getDocProps)
	fetchData := getDocProps.Hits.Hits

	//parseo struct a json
	auxDataJson, err := json.Marshal(fetchData)
	if err != nil {
		return "", err
	}

	dataJson := string(auxDataJson)

	return dataJson, nil
}
