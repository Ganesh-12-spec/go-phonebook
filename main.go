package main

import (
	"fmt"
	"go-phonebook/phonebook"
)

func main() {
    pb := phonebook.NewPhoneBook()

    // Add contacts
    pb.Add(&phonebook.Contact{
        Name:  "Rushi",
        Phone: "9999911111",
        Email: "rushi@email.com",
    })
    pb.Add(&phonebook.Contact{
        Name:  "Priya",
        Phone: "8888822222",
        Email: "priya@email.com",
    })
    pb.Add(&phonebook.Contact{
        Name:  "Arjun",
        Phone: "7777733333",
        Email: "arjun@email.com",
    })

    fmt.Println("=== All Contacts ===")
    pb.List()

    // Get aur modify — pointer se
    if c, ok := pb.Get("Rushi"); ok {
        c.UpdatePhone("9000000001")   // pointer se original change
        c.AddGroup("Friends")
        c.AddGroup("Work")
    }

    // Search
    fmt.Println("\n=== Search: 'ri' ===")
    results := pb.Search("ri")
    for _, c := range results {
        fmt.Println(c.String())
    }

    // Updated contact dekho
    fmt.Println("\n=== Rushi after update ===")
    if c, ok := pb.Get("Rushi"); ok {
        fmt.Println(c.String())
        fmt.Println("Groups:", c.Groups)
    }

    // Delete
    pb.Delete("Arjun")
    fmt.Printf("\nAfter delete — Total contacts: %d\n", pb.Size())

    fmt.Println("\n=== Final List ===")
    pb.List()
}