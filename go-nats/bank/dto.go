package bank

import (
	"encoding/json"
	"log"
)

type RequestWithdraw struct {
	Id     uint
	Amount uint
}

func (rw RequestWithdraw) String() string {
	out, err := json.MarshalIndent(rw, "", "   ")
	if err != nil {
		log.Printf("Could not marshal request --- error: %s", err.Error())
		return ""
	}
	return string(out)
}
