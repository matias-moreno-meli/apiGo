package domains

import (
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Site struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	CountryID          string   `json:"country_id"`
	SaleFeesMode       string   `json:"sale_fees_mode"`
	MercadopagoVersion int32    `json:"mercadopago_version"`
	DefaultCurrencyId  string   `json:"default_currency_id"`
	ImmediatePayment   string   `json:"immediate_payment"`
	PaymentMethodIds   []string `json:"payment_method_ids"`
	Settings           struct {
		IdentificationTypes      []string      `json:"identification_types"`
		TaxpayerTypes            []interface{} `json:"taxpayer_types"`
		IdentificationTypesRules interface{}   `json:"identification_types_rules"`
	} `json:"settings"`
	Currencies []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Categories []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"categories"`
}

func (site *Site) Get() *utils.ApiError {
	if len(site.ID) == 0 {
		return &utils.ApiError{
			Message: "El id de usuario es 0",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)

	response, err := http.Get(url)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &site); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil

}
