package dimets

type DimeTransactionFee struct {
	BaseDimeTransactionLog
}
type DimeFeeTransaction interface {
	ToJson() (*DimeTransactionFee, error)
}

func NewDimeTransactionFee(text string) DimeFeeTransaction {
	return DimeTafTransaction{Text: text}
}
