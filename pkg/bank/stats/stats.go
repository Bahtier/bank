package stats

import(
	"github.com/Bahtier/bank/pkg/bank/types"
)

// Avg ф-я вычисляет среднюю сумму платежей
func Avg(payments []types.Payment) types.Money {	
	var total types.Money = 0
	
	for _, payments := range payments {				
			total += payments.Amount					
	}

	total = total / types.Money(len(payments))

	return total
}