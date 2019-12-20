# Febraban Barcode Decrypt

Decrypt any barcode using Febraban layout.

## Description

Febraban Barcode Decrypt is a lib to decode barcode and provide a struct with all data found.

## Usage

Functions
- `func Decrypt(barcode string) ResultStruct`

Samples
- `import (
	"github.com/loadfms/fbd"
)`
- `result := fbd.Decrypt("836700000042808100470000000000001032979211190058")`

Result Structure
```
type ResultStruct struct {
	Product           string
	Segment           string
	ValueType         string
	VerificationDigit string
	Value             int
	CreditAccountInfo string
	Obs               string
}
```
