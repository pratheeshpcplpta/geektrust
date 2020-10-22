package bank

import (
	"fmt"
	"geektrust/app"
	"geektrust/app/models"
	"math"
	"math/rand"
	"time"
)

type Loan struct {
	BankName        string
	BorrowerName    string
	PrincipalAmount float64
	NoOfYears       int
	RateOfInterest  float64
	EmiId           string
	TotalRepay      float64
	NoOfEmis        int
	MonthlyEmi      float64
}

/**
 * ----------------------------------------------------------------------------
 * Initiate Object for Loan
 * ----------------------------------------------------------------------------
 */
func InitLoan() *Loan {
	return &Loan{}
}

/**
 * ----------------------------------------------------------------------------
 * Initiate borrow loan functionality
 *
 * Steps :
 *
 *	Generate a random EMI ID
 *	Check the loan details already exists with the bank name and borrower name
 * 	Add new entry to ledger
 *	Calculate the Monthly EMI
 *
 * ----------------------------------------------------------------------------
 */
func (loan *Loan) BorrowLoan() (error, *Loan) {
	app := app.InitApp()

	_emiid := loan.GenerateEMIID()

	//check the entry already exists
	var result []models.LoanDetailsLedger
	objectInstance := app.DB
	objectInstance = objectInstance.Model(&models.LoanDetailsLedger{})
	objectInstance = objectInstance.Where("bank_name = ? AND borrower_name = ?", loan.BankName, loan.BorrowerName)
	objectInstance = objectInstance.Find(&result)

	if len(result) > 0 {
		return fmt.Errorf("You already have an approved loan from this bank."), loan
	}
	// Create
	app.DB.Create(&models.LoanDetailsLedger{
		BankName:        loan.BankName,
		BorrowerName:    loan.BorrowerName,
		PrincipalAmount: loan.PrincipalAmount,
		NoOfYears:       loan.NoOfYears,
		RateOfInterest:  loan.RateOfInterest,
		EmiId:           _emiid,
	})

	// calculate EMI and insert into table
	loan.CalculateEMI(app)
	return nil, loan
}

/**
 * ----------------------------------------------------------------------------
 * @Function : CalculateEMI
 * Calculate the monthly EMI for the loan
 *
 * Calculate the Interest
 *			i = P * N * R
 *			P - Principal Amont , N - No of Years, R - Rate of Interest
 *			The total amount to repay will be A = P + I
 *
 * Calculate Monthly payment by
 *			Total EMis = no of years * 12
 *			Monthly Payment = A / Total EMis
 *
 * ----------------------------------------------------------------------------
 */
func (loan *Loan) CalculateEMI(app *app.App) {
	_princ_amount := loan.PrincipalAmount
	_no_of_years := loan.NoOfYears
	_rate_of_interest := loan.RateOfInterest

	//calculate the interest
	// i = P*N*R
	_intrest := _princ_amount * float64(_no_of_years) * float64(_rate_of_interest/100)
	_total_return := _intrest + _princ_amount

	_emi_total_months := (_no_of_years * 12) //find the no of months to pay

	// find the monthly payment
	_monthly_payment := _total_return / float64(_emi_total_months)
	_monthly_payment = math.Ceil(_monthly_payment)

	year, month, day := time.Now().Date()

	emiDetails := []models.EmiPaymentDetailLedger{}

	emiMonth := int(month)
	emiYear := int(year)

	for i := 0; i < _emi_total_months; i++ {
		emiDetails = append(emiDetails, models.EmiPaymentDetailLedger{
			EmiId:         loan.EmiId,
			EmiAmount:     _monthly_payment,
			EmiYear:       emiYear,
			EmiMonth:      emiMonth,
			EmiDay:        day,
			PaymentStatus: 0,
			AddedLumpsum:  0,
		})

		emiMonth = emiMonth + 1
		if emiMonth > 12 {
			emiMonth = 1
			emiYear = emiYear + 1
		}

	}

	//inser the emi detail to table
	result := app.DB.Create(&emiDetails)

	if result.Error != nil {
		panic(result.Error)
	}

	loan.TotalRepay = _total_return
	loan.NoOfEmis = _emi_total_months
	loan.MonthlyEmi = _monthly_payment
}

/**
 * ----------------------------------------------------------------------------
 * Generate a randoom string
 * ----------------------------------------------------------------------------
 */
func (loan *Loan) GenerateEMIID() string {
	_str := RandStringRunes(5)
	loan.EmiId = _str

	return _str
}

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
