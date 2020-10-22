package main

import (
	"fmt"
	"geektrust/app"
	"geektrust/app/models"
)

func main_() {
	var record models.EmiPaymentDetailLedger

	app := app.InitApp()
	emil__id := "4V32G"
	emiobjectInstance := app.DB
	emiobjectInstance = emiobjectInstance.Where("emi_id = ?", emil__id)
	emiobjectInstance = emiobjectInstance.Limit(1)
	emiobjectInstance = emiobjectInstance.Offset(1)
	emiobjectInstance = emiobjectInstance.Order("id ASC")
	emiobjectInstance.Find(&record)

	fmt.Println(record)
	fmt.Println(emil__id)
}
