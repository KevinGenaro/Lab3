package tp5

import (
	"encoding/json"
	"log"
	"net/http"
)

type Item struct {
	Name       string `json:"name"`
	CountryId  string `json:"country_id"`
	Currencies []struct {
		Id     string
		Symbol string
	} `json:"currencies"`
}

func ConsumoApi() {

	resp, err := http.Get("https://api.mercadolibre.com/sites/MLA")
	if err != nil {
		log.Fatal(err)
	}
	var resultado Item
	err = json.NewDecoder(resp.Body).Decode(&resultado)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resultado)
}
