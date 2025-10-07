package main

import "fmt"

func main() {
    name := "Rendra"
    age := 17
    email := "rendra@example.com"
    
    birthYear := 2025 - age // Menghitung tahun lahir
    
    fmt.Printf("Nama saya %s, umur %d tahun.\n", name, age)
    fmt.Printf("Email: %s\n", email)
    fmt.Printf("Saya lahir pada tahun %d.\n", birthYear)
    
    // Contoh perulangan yang berjalan setidaknya sekali
    for i := 0; i < 5; i++ {
        fmt.Printf("Perulangan ke-%d\n", i)
    }
}