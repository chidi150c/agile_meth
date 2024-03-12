// Determine the required fields for each profile type (e.g., bio, experience, skills, photos).
package main

import "fmt"

type Profile struct {
    Bio       string
    Experience int
    Skills    []string
    Photos    []string
}

func main() {
    ownerProfile := Profile{
        Bio:       "Passionate pet owner looking for reliable pet sitter",
        Experience: 5,
        Skills:    []string{"Dog walking", "Basic grooming"},
        Photos:    []string{"photo1.jpg", "photo2.jpg"},
    }

    sitterProfile := Profile{
        Bio:       "Experienced pet sitter available for dog and cat care",
        Experience: 3,
        Skills:    []string{"Pet training", "First-aid certified"},
        Photos:    []string{"photo3.jpg", "photo4.jpg"},
    }

    fmt.Printf("Owner Profile: %+v\n", ownerProfile)
    fmt.Printf("Sitter Profile: %+v\n", sitterProfile)
}
