package services

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"strconv"
	"strings"
)

const (
	YES string = "ТАК"
	NO  string = "НІ"
)

type GoogleSheetFetcher struct {
}

func (GoogleSheetFetcher) Fetch(category string) ViewData {
	var products = make([]Product, 0, 100)
	log.Printf("fetch %s products from google sheet ", category)
	ctx := context.Background()
	sheetsService, err := sheets.NewService(ctx, option.WithAPIKey("AIzaSyDtw-RYGAz9G1iv-0QUnHcqRJZLGo3EgiI"))

	if err != nil {
		log.Fatal("failed to fetch categories")
	}

	if sheetsService != nil {
		log.Println("Connected to Google Sheets")
	}

	call := sheetsService.Spreadsheets.Values.Get("1eIoMdlPLoNqb3qJrLt-DhXmndYa7ZlqnMgVO5tFmBeY", fmt.Sprintf("%v!A1:G1000", category))
	if sheetValues, err := call.Do(); err == nil {
		log.Println("processing rows")
		for _, row := range sheetValues.Values[1:] {
			if fmt.Sprintf("%v", row[1]) == "" {
				log.Printf("skip the item if name %v is empty \n", fmt.Sprintf("%v", row[1]))
				continue
			}

			if strings.Compare(fmt.Sprintf("%v", row[6]), YES) != 0 {
				continue
			}

			s, err := strconv.ParseFloat(fmt.Sprintf("%v", row[4]), 32)
			if err != nil {
				s = 0.00
			}

			imageUrl := fmt.Sprintf("%v", row[5])
			if imageUrl == "" {
				imageUrl = "https://via.placeholder.com/300?Text=Немає"
			}

			products = append(products, Product{
				true,
				fmt.Sprintf("%s", row[1]),
				fmt.Sprintf("%s", row[2]),
				ParseQuantity(fmt.Sprintf("%s", row[3])),
				s,
				imageUrl,
			})
		}

		log.WithField("products", products).Println("Products")

		return ViewData{
			Name:  category,
			Items: products,
		}

	} else {
		log.Error(err.Error())
	}

	return ViewData{
		Name:  category,
		Items: make([]Product, 0),
	}
}
