package main

import (
	"fmt"
	"geektrust/app/bank"
)

func main__() {
	// loan := bank.InitLoan()
	// loan.BankName = "MDB"
	// loan.BorrowerName = "Dale"
	// loan.PrincipalAmount = 10000
	// loan.NoOfYears = 1
	// loan.RateOfInterest = 2

	// _, _ = loan.BorrowLoan()

	// fetch abalance
	bl := bank.InitBalance()
	bl.BankName = "MDB"
	bl.BorrowerName = "Dale"
	bl.EmiNumber = 4

	response := bl.Balance()

	if response.Status == true {
		fmt.Println(response.BankName, response.BorrowerName, response.AmountPaid, response.NoOfEmiLeft)
	} else {
		fmt.Printf("\033[1;31m%s\033[0m \n", response.Error)
	}

	// payment
	// payment := bank.InitPayment()
	// payment.BankName = "MDB"
	// payment.BorrowerName = "Dale"
	// payment.EmiNumber = 0
	// payment.LumpSumAmount = 3000

	// response_payment := payment.Payment()
	// fmt.Println(response_payment)
}
