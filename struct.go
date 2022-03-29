package main

// developing only uncompleted...
//Insert struct declaration here
type jediKnight struct {
	FName            string  `json:"fname"`
	LName            string  `json:"lName"`
	Age              int     `json:"age"`
	Subscriber       bool    `json:"subscriber"`
	Phone            int     `json:"phone"`
	CreditAvailable  string  `json:"creditAvailable"`
	CurrentCartCost  float32 `json:"currentCartCost"`
	CurrentOrderCost float32 `json:"currentOrderCost"`
}

var jedis = []jediKnight{}

func main() {
	//Insert code here
}
