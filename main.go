package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    fmt.Println("Vowel/Consonant Shuffler")
    fmt.Println("=========================")
    fmt.Println("Enter a word or a sentence and I will shuffle vowels and consonants separately:")

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("> ")
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)

    shuffled := shuffleVowelsConsonants(text)
    fmt.Printf("Shuffled text:\n%s\n", shuffled)
}

// Function that returns true if a character is a vowel (English)
func isVowel(c byte) bool {
    vowels := "aeiouAEIOU"
    return strings.ContainsRune(vowels, rune(c))
}

// Shuffles vowels among vowels and consonants among consonants in the text
func shuffleVowelsConsonants(text string) string {
    var vowels []byte
    var consonants []byte
    var types []string

    // Separate vowels and consonants, and remember their positions
    for i := 0; i < len(text); i++ {
        c := text[i]
        if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
            if isVowel(c) {
                vowels = append(vowels, c)
                types = append(types, "vowel")
            } else {
                consonants = append(consonants, c)
                types = append(types, "consonant")
            }
        } else {
            types = append(types, "other")
        }
    }

    // Shuffle both slices
    rand.Shuffle(len(vowels), func(i, j int) { vowels[i], vowels[j] = vowels[j], vowels[i] })
    rand.Shuffle(len(consonants), func(i, j int) { consonants[i], consonants[j] = consonants[j], consonants[i] })

    // Reconstruct the text
    var result []byte
    vi, ci := 0, 0
    for i := 0; i < len(text); i++ {
        if types[i] == "vowel" {
            result = append(result, vowels[vi])
            vi++
        } else if types[i] == "consonant" {
            result = append(result, consonants[ci])
            ci++
        } else {
            result = append(result, text[i])
        }
    }
    return string(result)
}
