package stats
import(
	"github.com/Bahtier/bank/pkg/bank/types"
	"fmt"
)
func ExampleAvg()  {
	payments := []types.Payment{
		types.Payment{
			Amount: 5_000,			
		},
		types.Payment{
			Amount: 15_000,			
		},		
	}

	fmt.Println(Avg(payments))

	//Output: 10000
}