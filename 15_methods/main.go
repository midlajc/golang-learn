package main

import "fmt"

func main() {
	fmt.Println("Structs in Go-lang")
	// no inheritance in golang; No super or parent

	midlaj := User{Name: "MIDLAJ C", Email: "mail.midlajc@gmail.com", Age: 21, Status: true}

	fmt.Println(midlaj)
	fmt.Printf("Details For %v is: %+v\n", midlaj.Name, midlaj)
	fmt.Println("Age:", midlaj.Age)

	userSlice := make([]User, 0)
	userSlice = append(userSlice, User{Name: "MIDLAJ C", Email: "mail.midlajc@gmail.com", Age: 21, Status: true})
	fmt.Println(userSlice)

	userMap := make(map[string]User)
	userMap["midlaj"] = User{Name: "MIDLAJ C", Email: "mail.midlajc@gmail.com", Age: 21, Status: true}
	fmt.Println(userMap)
	fmt.Printf("%+v\n", userMap)
	fmt.Printf("%+v\n", userMap["midlaj"])
	midlaj.UpdateEmailPBV("test@gmail.com")
	fmt.Println(midlaj.Email)
	midlaj.UpdateEmailPBR("test@gmail.com")
	fmt.Println(midlaj.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() bool {
	return u.Status
}
func (u User) UpdateEmailPBV(email string) {
	u.Email = email
}
func (u *User) UpdateEmailPBR(email string) {
	u.Email = email
}
