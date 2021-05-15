import (
	"fmt"
	"http"
)

type Coin struct {
	Name string `json:"name"`
}

var coins []Coin 

func getCoinsHandler(w http.responseWriter , r *http.Request) {
	coinListBytes, err := json.Marshal(coins)

	if err != nil {
		fmt.Println(fmt.Errorf("Error in getting coins %v", err))
		w.writeHeader(http.StatusInternalServerError)
		return
	}
	w.Write(coinListBytes)
}

func createCoinHandler(w http.responseWriter, r *http.Request) {
	coin := Coin{}
}