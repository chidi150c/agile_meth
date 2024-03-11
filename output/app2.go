//   - Include details on the duration, activities involved, and any special accommodations provided.
package main

import "fmt"

type Service struct {
	Name         string
	Description  string
	Price        float64
	Duration     string
	Activities   string
	Accommodations string
}

func main() {
	dogWalking := Service{"Dog Walking", "Professional dog walking service for your beloved pet.", 15.00, "30 minutes", "Walking and exercise for your dog", "N/A"}
	petSitting := Service{"Pet Sitting", "In-home pet sitting service to make sure your pet is comfortable while you're away.", 25.00, "Per day", "Feeding, playtime, and companionship", "Food and water provided"}
	boarding := Service{"Boarding", "Safe and secure boarding facility for overnight stays.", 35.00, "Per night", "Playtime, feeding, and socialization", "Individual sleeping area"}

	printServiceInfo(dogWalking)
	printServiceInfo(petSitting)
	printServiceInfo(boarding)
}

func printServiceInfo(service Service) {
	fmt.Println("Service: ", service.Name)
	fmt.Println("Description: ", service.Description)
	fmt.Println("Price: $", service.Price)
	fmt.Println("Duration: ", service.Duration)
	fmt.Println("Activities: ", service.Activities)
	fmt.Println("Accommodations: ", service.Accommodations)
	fmt.Println()
}
