package main

import (
	"fmt"
	"geektrust/app/bank"
	"geektrust/app/helpers/color"
	"os"
	"strconv"
)

func main() {

	err := ProcessCommands()
	if err != nil {
		fmt.Printf("\033[1;31m[error]\033[0m %s \n", err)
	}
}

/**
 * -------------------------------------------------------------------
 * Routing the functions based on the input comands
 * -------------------------------------------------------------------
 */
func ProcessCommands() error {

	allowedCommands := map[string]string{
		"LOAN":    "LOAN",
		"PAYMENT": "PAYMENT",
		"BALANCE": "BALANCE",
	}

	if len(os.Args) > 1 {

		//check comnand exist allowed commands
		if _, ok := allowedCommands[os.Args[1]]; !ok {
			return fmt.Errorf("Invalid command")
		}

		switch os.Args[1] {
		case "LOAN":
			return _process_loan()
		case "PAYMENT":
			return _process_payment()

		case "BALANCE":
			return _process_balance()
		}
	} else {
		return fmt.Errorf("Should provide commands")
	}

	return nil
}

/**
 * -------------------------------------------------------------------
 * Process LOAN
 * -------------------------------------------------------------------
 */
func _process_loan() error {
	if len(os.Args) < 7 {
		return fmt.Errorf("insufficient parameters")
	}
	principal_converted_val, _ := strconv.ParseFloat(os.Args[4], 64)
	no_of_years_converted_val, _ := strconv.Atoi(os.Args[5])
	rate_of_converted_val, _ := strconv.ParseFloat(os.Args[6], 64)

	loan := &bank.Loan{
		BankName:        os.Args[2],
		BorrowerName:    os.Args[3],
		PrincipalAmount: principal_converted_val,
		NoOfYears:       no_of_years_converted_val,
		RateOfInterest:  rate_of_converted_val,
	}
	err, loan_data := loan.BorrowLoan()

	if err != nil {
		color.Error(fmt.Sprintf("%v", err))
	} else {
		//
		// Render the content
		//
		fmt.Println("| ----------------------------------------------------")
		fmt.Printf("| \t\t \033[1;32m%s\033[0m \n", " Your Loan has been approved")
		fmt.Printf("| \033[1;31m%s\033[0m \n", "Loan Details")
		fmt.Println("| ----------------------------------------------------")
		fmt.Println("| Bank Name : ", loan_data.BankName)
		fmt.Println("| Borrower Name : ", loan_data.BorrowerName)
		fmt.Println("| Principal Amount : ", loan_data.PrincipalAmount)
		fmt.Println("| No Of Years : ", loan_data.NoOfYears)
		fmt.Println("| Rate Of Interest : ", loan_data.RateOfInterest)

		fmt.Println("| ----------------------------------------------------")
		fmt.Printf("| \033[1;31m%s\033[0m \n", "EMI Details")
		fmt.Println("| ----------------------------------------------------")
		fmt.Println("| Total Repay : ", loan_data.TotalRepay)
		fmt.Println("| Total Emis : ", loan_data.NoOfEmis)
		fmt.Println("| Monthly Payment : ", loan_data.MonthlyEmi)
		fmt.Println("| ----------------------------------------------------")
	}
	return nil
}

/**
 * -------------------------------------------------------------------
 * Process Payments
 * -------------------------------------------------------------------
 */
func _process_payment() error {
	if len(os.Args) < 6 {
		return fmt.Errorf("insufficient parameters")
	}
	lupmsump_converted_val, _ := strconv.ParseFloat(os.Args[4], 64)
	emi_converted_val, _ := strconv.Atoi(os.Args[5])

	payment := &bank.PaymentRequest{
		BankName:      os.Args[2],
		BorrowerName:  os.Args[3],
		LumpSumAmount: lupmsump_converted_val,
		EmiNumber:     emi_converted_val,
	}
	response_payment := payment.Payment()

	if response_payment.Status == true {
		color.Success("Your Loan payment been transfered successfully.")
	} else {
		color.Error(fmt.Sprintf("%v", response_payment.Error))
	}
	return nil
}

/**
 * -------------------------------------------------------------------
 * Process Balance
 * -------------------------------------------------------------------
 */
func _process_balance() error {
	if len(os.Args) < 5 {
		return fmt.Errorf("insufficient parameters")
	}
	converted_emi_val, _ := strconv.Atoi(os.Args[4])

	balanceObj := &bank.BalanceRequest{
		BankName:     os.Args[2],
		BorrowerName: os.Args[3],
		EmiNumber:    converted_emi_val,
	}
	response := balanceObj.Balance()

	if response.Status == true {
		fmt.Println(response.BankName, response.BorrowerName, response.AmountPaid, response.NoOfEmiLeft)
	} else {
		color.Error(fmt.Sprintf("%v", response.Error))
	}

	return nil
}

// go run main.go LOAN IDIDI bank 1000 1 2
// go run main.go LOAN IDIDI Dale 10000 1 1.2
// go run main.go PAYMENT IDIDI bank 170 1
// go run main.go BALANCE IDIDI bank 0
