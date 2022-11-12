package services

import (
	"bytes"
	"fmt"
	"net/http"
)

func DeleteDocuments(userFolder string) error {
	var statusCode int
	URL_DELETE := "http://localhost:4080/api/" + userFolder + "/_doc/"
	listIds, _ := GetIdsByUserFolder(userFolder, false)

	fmt.Println("\n\n--> LOG DELETE DOCUMENTS API <--")
	fmt.Print("En progreso...")
	for _, idDocument := range listIds {
		Id := idDocument.Id
		URL_DELETE_ID := URL_DELETE + Id

		req, err := http.NewRequest("DELETE", URL_DELETE_ID, bytes.NewBuffer([]byte{}))
		if err != nil {
			return err
		}
		req.SetBasicAuth("admin", "Complexpass#123")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		//se cierra despues de terminar todo el proceso
		defer resp.Body.Close()

		statusCode = resp.StatusCode
		/*
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}


			fmt.Println("\nuserFolder:", userFolder)
			fmt.Println("StatusCode:", resp.StatusCode)
			fmt.Println("body:", string(body))
		*/
	}
	fmt.Println("\nStatusCode:", statusCode)

	return nil
}
