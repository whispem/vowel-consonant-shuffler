# Vowel/Consonant Shuffler (Go mini project)

This mini-project takes a user input (a word or a sentence)
and shuffles vowels among vowels, and consonants among consonants, while
keeping each character type in its original position.

For example:
- `Hello world!` might become `Holle werld!`
- `Programming is fun` might become `Prigrammong os fan`

## How to use

1. Open this folder in your terminal.
2. Run:
   ```sh
   go run main.go
   ```
3. Type/paste your text, press Enter.
4. The shuffled text will be displayed.

---

### How it works

- Only shuffles alphabetic (A-Z, a-z) vowels and consonants, other characters (spaces, punctuation) are not changed.
- Uses a simple list of English vowels: A, E, I, O, U (both uppercase and lowercase).
- Not case-sensitive for vowel/consonant detection.
- Random shuffle: each run gives a different result.
