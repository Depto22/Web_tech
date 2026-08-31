package main

//go doesn't have inbuilt enums..we have to do it using const
import "fmt"

// enumerated types
// type orderStatus int

// const (
//
//	//iota is an integer type which is auto incrementing
//	Recieved orderStatus = iota
//	Confirmed
//	Prepared
//	Delivered
//
// )
type orderStatus string

const (
	//iota is an integer type which is auto incrementing
	Recieved  orderStatus = "recieved"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("changing order status to ", status)
}

//	func changeOrderStatus(status string) {
//		fmt.Println("changing order status to ", status)
//	}
func main() {
	changeOrderStatus(Recieved)
	changeOrderStatus(Prepared)

}
