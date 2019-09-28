package main


type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func NewCustomer (Id int, Name string, Gender string, Age int, Phone string, Email string) Customer {
	return Customer {
		Id: Id,
		Name: Name,
		Gender: Gender,
		Age: Age,
		Phone: Phone,
		Email: Email,
	}
}



func main () {

}
