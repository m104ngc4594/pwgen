// A password generator that generates passwords based on a set of rules.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// The characters used in generated passwords.
const (
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	Numbers   = "0123456789"
	Symbols   = "!@#$%^&*()_+-=[]{}|;':,./<>?"
)

type Generator struct {
	length    int    // The length of the password.
	uppercase bool   // Whether to include uppercase characters.
	lowercase bool   // Whether to include lowercase characters.
	numbers   bool   // Whether to include numbers.
	symbols   bool   // Whether to include symbols.
	exclude   string // A string of characters to exclude from the password.
	nums      int    // The numbers of the password generated.
}

// load the password generator from config.toml if exits.
func loadConfig(g *Generator) {
	viper.SetConfigName("config") // name of config file
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		log.Println(err)
		return
	}
	g.length = viper.GetInt("length")
	g.uppercase = viper.GetBool("uppercase")
	g.lowercase = viper.GetBool("lowercase")
	g.numbers = viper.GetBool("numbers")
	g.symbols = viper.GetBool("symbols")
	g.exclude = viper.GetString("exclude")
	g.nums = viper.GetInt("nums")
}

// NewGenerator creates a new password generator.
func NewGenerator() *Generator {
	g := Generator{20, true, true, true, true, "", 60}
	loadConfig(&g)
	return &g
}

// Generate generates nums of the password based on the rules defined in generator.
func (g *Generator) Generate() []string {
	var passwords []string
	for i := 0; i < g.nums; i++ {
		passwords = append(passwords, g.generate())
	}
	return passwords
}

// generate generates a password based on the rules defined in generator.
func (g *Generator) generate() string {
	var password string
	for i := 0; i < g.length; i++ {
		password += g.getChar()
	}
	return password
}

// getChar returns a random character based on the rules defined in generator.
func (g *Generator) getChar() string {
	var chars string
	if g.uppercase {
		chars += Uppercase
	}
	if g.lowercase {
		chars += Lowercase
	}
	if g.numbers {
		chars += Numbers
	}
	if g.symbols {
		chars += Symbols
	}
	return getRandomChar(chars, g.exclude)
}

// getRandomChar returns a random character exclude from the given string.
func getRandomChar(chars, exclude string) string {
	for {
		rand.Seed(time.Now().UnixNano())
		char := string(chars[rand.Intn(len(chars))])
		if !strings.Contains(exclude, char) {
			return char
		}
	}
}

func main() {
	// Create a password generator.
	generator := NewGenerator()

	// Generate the nums of passwords.
	passwords := generator.Generate()

	// Print the passwords.
	for _, password := range passwords {
		fmt.Println(password)
	}
}
