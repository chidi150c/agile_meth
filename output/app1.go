//   - Develop clear and concise descriptions of each service offered (e.g., dog walking, pet sitting, boarding).
package main

import "fmt"

type Service struct {
	Name        string
	Description string
	Price       float64
}

func main() {
	dogWalking := Service{
		Name:        "Dog Walking",
		Description: "Professional dog walking service for your beloved pet.",
		Price:       15.00,
	}

	petSitting := Service{
		Name:        "Pet Sitting",
		Description: "In-home pet sitting service to make sure your pet is comfortable while you're away.",
		Price:       25.00,
	}

	boarding := Service{
		Name:        "Boarding",
		Description: "Safe and secure boarding facility for overnight stays.",
		Price:       35.00,
	}

	fmt.Println("Service: ", dogWalking.Name)
	fmt.Println("Description: ", dogWalking.Description)
	fmt.Println("Price: $", dogWalking.Price)

	fmt.Println("Service: ", petSitting.Name)
	fmt.Println("Description: ", petSitting.Description)
	fmt.Println("Price: $", petSitting.Price)

	fmt.Println("Service: ", boarding.Name)
	fmt.Println("Description: ", boarding.Description)
	fmt.Println("Price: $", boarding.Price)
}
