# THE_LEDGER_CO


## The challenge

[Click Here to view the challenge](https://www.geektrust.in/coding-problem/backend/ledger-co)

You work at a startup called The Ledger Co., a marketplace for banks to lend money to borrowers and receive payments for the loans.

The interest for the loan is calculated by I = P*N*R where P is the principal amount, N is the number of years and R is the rate of interest. The total amount to repay will be A = P + I

The amount should be paid back monthly in the form of EMIs. The borrowers can also pay a lump sum (that is, an amount more than their monthly EMI). In such a case, the lump sum will be deducted from the total amount (A) which can reduce the number of EMIs. This doesn’t affect the EMI amount unless the remaining total amount is less than the EMI. All transactions happen through The Ledger Co.

You need to design a system to find out how much amount a user has paid the bank and how many EMIs are remaining .


### Database configuration
Using a simple database to store the information about loan and emi details

You can change db connection details in .env file

Domine : valid config value mysql, sqlite

## EMI
Emi has been calculated by simple interest rate
**SI = P * R * T/100**

* **SI** Simple Interest
* **R**	 Rate of interest
* **T**	 Time
* **P**	 Principal Amount


The formula for calculating EMI using interest rate is:

**EMI = (Principal + Interest)/Period in Months**



## Assumptions and conditions added with Commands

### Balance
- It prints the total amount paid by the borrower, including all the Lump Sum amounts paid including that EMI number, and the no of EMIs remaining.
- Only print Active emi payments
  
  for example : 
    
    If IDIDI have **12** month emi payment with **85** as EMI amount. He have 12 active EMIs. On 1st payment he pays **170**. He payed **85** more as lumpsum. 
    Thus the extra paid **85** has been deducted from last emi and that EMI will be marked as fully paid.

    So that the user have 10 active emi's remain to pay


### Payment
- EMIS can paid only one by one
- Can't pay 4th emi without 3rd emi payment
- Checking the emi already paid for a particular month
- Payment Amount should be greaterthan or equals to the EMI monthly amount
- If pay more than loan amount, remaining balance amount should be credit back to the user


### Loan
- Can't request multiple loan with the same bank and borrower name without completing all the emi payments.




## Execution

./geektrust LOAN IDIDI bank 1000 1 2

./geektrust PAYMENT IDIDI bank 1000 5

./geektrust BALANCE IDIDI bank 0

