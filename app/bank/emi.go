package bank

import (
	"fmt"
	"geektrust/app"
	"geektrust/app/models"
)

type BalanceRequest struct {
	BankName     string
	BorrowerName string
	EmiNumber    int
}

type BalanceResponse struct {
	BankName     string
	BorrowerName string
	EmiNumber    int
	AmountPaid   float64
	NoOfEmiLeft  int
	Status       bool
	Error        error
}

/**
 * ----------------------------------------------------------------------------
 * Initiate Object for Balance
 * ----------------------------------------------------------------------------
 */
func InitBalance() *BalanceRequest {
	return &BalanceRequest{}
}

/**
 * ----------------------------------------------------------------------------
 * @function : Balance
 *
 * Find out Balance amount after a particular emi number
 * Steps
 *		Find the loan details exists
 *		Get the unique emi id string and fetch total number of emis
 *		Fetch the emis upto given emi number
 *		Find the total paid sum of fetched emis
 *		Calculate remaining emis
 *
 * @return
 *		BalanceResponse object
 * ----------------------------------------------------------------------------
 */
func (bl *BalanceRequest) Balance() BalanceResponse {
	result := map[string]interface{}{}
	var emi__paid float64
	var record__length int
	var total_emis int
	var total_paid_emis int
	var BalanceRespoObjec BalanceResponse

	app := app.InitApp()

	//get the record of borrower detail
	objectInstance := app.DB
	objectInstance = objectInstance.Model(&models.LoanDetailsLedger{})
	objectInstance = objectInstance.Where("bank_name = ? AND borrower_name = ?", bl.BankName, bl.BorrowerName)
	objectInstance = objectInstance.Find(&result)

	if len(result) > 0 {
		emil__id := result["emi_id"]

		// find total no of emis
		app.DB.Table("emi_payment_detail_ledgers").Select("count(id)").Where("emi_id = ? ", emil__id).Row().Scan(&total_emis)

		// find total no of emis
		app.DB.Table("emi_payment_detail_ledgers").Select("count(id)").Where("emi_id = ? AND payment_status = ? ", emil__id, 1).Row().Scan(&total_paid_emis)
		emis_left := total_emis - total_paid_emis - bl.EmiNumber

		// if inpu is greater than no of emis
		if emis_left < 0 {
			BalanceRespoObjec.Status = false
			BalanceRespoObjec.Error = fmt.Errorf("Invalid emi number")
			return BalanceRespoObjec
		}

		var records []models.EmiPaymentDetailLedger

		emiobjectInstance := app.DB
		emiobjectInstance = emiobjectInstance.Where("emi_id = ? ", emil__id)
		emiobjectInstance = emiobjectInstance.Limit(bl.EmiNumber + 1)
		emiobjectInstance = emiobjectInstance.Order("id ASC")
		emiobjectInstance.Find(&records)

		record__length = len(records)

		for i := 0; i < record__length; i++ {
			_item := records[i]
			// emi__paid += _item.EmiAmount
			emi__paid += _item.EmiPaidAmount
		}

		BalanceRespoObjec.Status = true
		BalanceRespoObjec.BankName = bl.BankName
		BalanceRespoObjec.BorrowerName = bl.BorrowerName
		BalanceRespoObjec.EmiNumber = bl.EmiNumber
		BalanceRespoObjec.AmountPaid = emi__paid
		BalanceRespoObjec.NoOfEmiLeft = emis_left

	} else {
		BalanceRespoObjec.Status = false
		BalanceRespoObjec.Error = fmt.Errorf("Unable to find loan details of user with bank name")
	}
	return BalanceRespoObjec
}
