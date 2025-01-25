package main  
  
import (  
	"fmt"  
)  
  
// Fungsi untuk membalikkan string  
func reverseString(input string) string {  
	// Konversi string ke slice of runes (untuk menangani Unicode)  
	runes := []rune(input)  
	// Balikkan slice  
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {  
		runes[i], runes[j] = runes[j], runes[i]  
	}  
	// Kembalikan sebagai string  
	return string(runes)  
}  
  
func main() {  
	// Contoh penggunaan  
	original := "Luky Mulana"  
	reversed := reverseString(original)  
	fmt.Printf("Original: %s\n", original)  
	fmt.Printf("Reversed: %s\n", reversed)  
}  