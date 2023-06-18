package facad

import (
	"errors"
	"log"
)

type details struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
}

// NewPersonAccount ...
func NewPersonAccount(accountId string, code int) *details {
	return &details{
		account:      newAccount(accountId),
		wallet:       initAmount(),
		securityCode: newSC(code),
	}
}

// CreditAmount ...
func (d *details) CreditAmount(accountId string, code, amount int) error {
	// Verify account details
	accountVerified := d.account.verifyAccount(accountId)
	if !accountVerified {
		return errors.New("invalid account id")
	}

	// Checks security code
	codeVerified := d.securityCode.checkSecurity(code)
	if !codeVerified {
		return errors.New("invalid security code")
	}

	// Credit balance
	d.wallet.addAmount(amount)
	return nil
}

// DebitAmount ...
func (d *details) DebitAmount() {
	// Verify account details

	// Checks security code

	// Debit balance
}

// Status ...
func (d *details) Status() {
	log.Println(d.account.id, d.wallet.amount)
}
