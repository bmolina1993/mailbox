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

func PostApiBulkData(data entity.QueryMail) error {
	URL := "http://localhost:4080/api/_bulkv2"

	//parseo struct a json para post api
	dataJson, err := json.Marshal(data)
	if err != nil {
		return err
	}

	//post api
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(dataJson))
	if err != nil {
		return err
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
		return err
	}

	defer resp.Body.Close()

	//en caso de querer ver LOG de respuesta post api [opcional]
	/*
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println("\n--> LOG POST API <--")
		fmt.Println("StatusCode:", resp.StatusCode)
		fmt.Println("body:", string(body))
	*/

	return nil
}
