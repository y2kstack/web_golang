package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"webservices/cors"
)

const productsPath = "products"

func handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productList, err := getProductList()
		j, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPost:
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		productID, err := insertProduct(product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"productId":%d}`, productID)))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func handleProduct(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", productsPath))

	fmt.Println(r.URL.Path)
	fmt.Println(urlPathSegments, "1-------")
	fmt.Println(len(urlPathSegments[1:]), "2------")

	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {

	case http.MethodGet:
		product, err := getProduct(productID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if product == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodDelete:
		err := removeProduct(productID)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if *product.ProductID != productID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateProduct(product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	productsHandler := http.HandlerFunc(handleProducts)
	productHandler := http.HandlerFunc(handleProduct)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsPath), cors.Middleware(productsHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsPath), cors.Middleware(productHandler))
}
