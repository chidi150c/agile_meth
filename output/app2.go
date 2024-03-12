// Allow users to provide detailed information about their pets (e.g., breed, age, special needs).
// Add Pet struct to User for pet information
type Pet struct {
    Name       string
    Breed      string
    Age        int
    SpecialNeeds string
}

// Update User struct to include pets slice
type User struct {
    ID       int
    Name     string
    Email    string
    Phone    string
    Address  string
    Pets     []Pet
}

func main() {
    // Instantiate a new user with pets
    user := User{
        ID:      1,
        Name:    "Jane Smith",
        Email:   "jane.smith@example.com",
        Phone:   "987-654-3210",
        Address: "456 Oak Avenue",
        Pets: []Pet{
            {Name: "Charlie", Breed: "Labrador", Age: 5, SpecialNeeds: "None"},
            {Name: "Bella", Breed: "Maine Coon", Age: 3, SpecialNeeds: "Allergies"},
        },
    }

    // Print user and pet details
    fmt.Println("User Details:")
    fmt.Println("ID:", user.ID)
    fmt.Println("Name:", user.Name)
    fmt.Println("Email:", user.Email)
    fmt.Println("Phone:", user.Phone)
    fmt.Println("Address:", user.Address)

    fmt.Println("\nPets:")
    for i, pet := range user.Pets {
        fmt.Printf("Pet %d\n", i+1)
        fmt.Println("Name:", pet.Name)
        fmt.Println("Breed:", pet.Breed)
        fmt.Println("Age:", pet.Age)
        fmt.Println("Special Needs:", pet.SpecialNeeds)
        fmt.Println()
    }
}
