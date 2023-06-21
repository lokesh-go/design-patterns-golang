package observer

import "fmt"

type UserObserver struct {
	UserId string
}

func (u *UserObserver) send(itemName string) {
	fmt.Printf("%s is available for user %s\n", itemName, u.UserId)
}
