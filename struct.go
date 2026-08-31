package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

// pass the object by reference
func (o *order) changeStatus(s string) {
	o.status = s
}
func (o order) getAmount() float32 {
	return o.amount
}

// go doesnt have constructor system...so alternative
func const_order(id string, amount float32, status string, createdAt time.Time) *order {
	newOrder := order{
		id:        "1",
		amount:    50.00,
		status:    "Recieved",
		createdAt: time.Now(),
	}
	return &newOrder
}
func main() {
	myOrder1 := order{
		id:        "1",
		amount:    50.00,
		status:    "Recieved",
		createdAt: time.Now(),
	}
	fmt.Println(myOrder1)
	myOrder1.amount = 100
	fmt.Println(myOrder1)
	myOrder1.changeStatus("confirmed")
	fmt.Println(myOrder1)
	fmt.Println(myOrder1.getAmount())
	myOrder2 := const_order("2", 90.00, "delivered", time.Now())
	fmt.Println(myOrder2)
	myOrder3 := order{
		id:        "3",
		amount:    929.00,
		status:    "Recieved",
		createdAt: time.Now(),
		customer: customer{
			name:  "Raju",
			phone: "01722207398",
		},
	}
	fmt.Println(myOrder3)

}
