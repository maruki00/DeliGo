package factories

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/maruki00/deligo/internal/order/app/enums"
	aggrigate "github.com/maruki00/deligo/internal/order/domain/aggrigate"
	"github.com/maruki00/deligo/internal/order/infra/model"
	product_infra_model "github.com/maruki00/deligo/internal/product/infrastructure/model"
	shared_model "github.com/maruki00/deligo/internal/shared/model"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func NewOrderAggrigate(CostomerId int, prods map[int]int) (*aggrigate.OrderAggrigate, error) {

	ids := ""
	for id := range prods {
		ids += fmt.Sprintf("%d,", id)
	}

	ids = strings.TrimSuffix(ids, ",")

	products, err := getProducts(ids)
	if err != nil {
		return nil, err
	}

	cost := float32(0)
	for _, product := range products {
		pQty := prods[product.Id]
		cost += float32(pQty) * product.Price
	}
	fingerPrint := uuid.New().String()
	Order := model.Order{
		CostumerId:       CostomerId,
		OrderFingerprint: fingerPrint,
		Cost:             cost,
		Status:           enums.ORDER_CREATED,
	}

	return &aggrigate.OrderAggrigate{
		Order: &Order,
		// Items: products,
		Price: cost,
	}, nil

}

func getProducts(ids string) ([]*product_infra_model.Product, error) {

	Client := http.Client{
		Timeout: time.Second * 2,
	}
	data := map[string]string{
		"ids": ids,
	}
	d, _ := json.Marshal(data)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/product/multiple", bytes.NewReader(d))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	dd := shared_model.ResponseModel{}
	err = json.Unmarshal(body, &dd)
	if err != nil {
		return nil, err
	}
	products, ok := dd.Result.(map[string]any)["products"]
	if !ok {
		return nil, errors.New("products are missing in the response")
	}
	product := &product_infra_model.Product{}
	productsResult := []*product_infra_model.Product{}
	for _, p := range products.([]interface{}) {
		mapstructure.Decode(p, product)
		productsResult = append(productsResult, product)
	}
	return productsResult, nil

}
