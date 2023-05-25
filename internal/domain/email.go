package domain

type Email struct {
	To   []string
	Body BodyMail
}

type BodyMail struct {
	Headers string
	Message string
}

type BodyEmailData struct {
	TotalBalance      float64
	MonthTransactions map[string]int
	AverageDebit      float64
	AverageCredit     float64
}

func UserTransactionToEmail(totalCredit float64, totalDebit float64, cantCredit int, cantDebit int, transCountByMonth map[string]int) *BodyEmailData {
	return &BodyEmailData{
		TotalBalance:      totalCredit + totalDebit,
		MonthTransactions: transCountByMonth,
		AverageDebit:      totalDebit / float64(cantDebit),
		AverageCredit:     totalCredit / float64(cantCredit),
	}
}
