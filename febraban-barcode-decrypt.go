package fbd

import (
	"strconv"
	"strings"
)

const (
	SegmentTown          = "Prefeitura"
	SegmentSanitation    = "Saneamento"
	SegmentEnergyGas     = "Energia eletrica e gas"
	SegmentTelecom       = "Telecomunicacoes"
	SegmentGovernment    = "Orgaos Governamentais"
	SegmentGeneric       = "Carnes e Assemelhados ou demais empresas/orgaos que serao identificadas atraves do CNPJ"
	SegmentTrafficTicket = "Multa de transito"
	SegmentBank          = "Uso exclusivo do banco"
)

const (
	ProductType = "Arrecadacao"
)

const (
	ValueTypeSix   = "Valor a ser cobrado efetivamente em reais"
	ValueTypeSeven = "Quantidade de moeda"
	ValueTypeEight = "Valor a ser cobrado efetivamente em reais"
	valueTypeNine  = "Quantidade de moeda"
)

type ResultStruct struct {
	Product           string
	Segment           string
	ValueType         string
	VerificationDigit string
	Value             int
	CreditAccountInfo string
	Obs               string
}

func Decrypt(barcode string) ResultStruct {
	barcode = replaceAtIndex(barcode, 11, ' ')
	barcode = replaceAtIndex(barcode, 23, ' ')
	barcode = replaceAtIndex(barcode, 35, ' ')
	barcode = replaceAtIndex(barcode, 47, ' ')
	barcode = strings.Replace(barcode, " ", "", -1)

	bcProduct := string(barcode[0])
	bcSegment := string(barcode[1:2])
	bcValueType := string(barcode[2:3])
	bcVerificationDigit := string(barcode[3:4])
	bcValue := string(barcode[4:15])
	bcCreditAccountInfo := string(barcode[15:19])
	bcObs := string(barcode[19:])

	formattedValue, _ := strconv.Atoi(bcValue)

	finalResult := ResultStruct{
		Product:           resolveProduct(bcProduct),
		Segment:           resolveSegment(bcSegment),
		ValueType:         resolveValueType(bcValueType),
		VerificationDigit: bcVerificationDigit,
		Value:             formattedValue,
		CreditAccountInfo: bcCreditAccountInfo,
		Obs:               bcObs,
	}

	return finalResult
}

func resolveProduct(item string) string {
	if item == "8" {
		return ProductType
	}
	return ""
}

func resolveSegment(item string) string {
	switch item {
	case "1":
		return SegmentTown
	case "2":
		return SegmentSanitation
	case "3":
		return SegmentEnergyGas
	case "4":
		return SegmentTelecom
	case "5":
		return SegmentGovernment
	case "6":
		return SegmentGeneric
	case "7":
		return SegmentTrafficTicket
	case "9":
		return SegmentBank
	}
	return ""
}

func resolveValueType(item string) string {
	switch item {
	case "6":
		return ValueTypeSix
	case "7":
		return ValueTypeSeven
	case "8":
		return ValueTypeEight
	case "9":
		return valueTypeNine
	}
	return ""
}

func replaceAtIndex(in string, i int, r rune) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
