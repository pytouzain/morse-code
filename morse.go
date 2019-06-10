package main

import (
	"fmt"
	"os"
	"time"
	"unicode"

	"github.com/stianeikeland/go-rpio"
)

var morseAlphabet = map[rune]string{
	'A': ".-",
	'B': "-...",
	'C': "-.-.",
	'D': "-..",
	'E': ".",
	'F': "..-",
	'G': "--.",
	'H': "....",
	'I': "..",
	'J': ".---",
	'K': "-.-",
	'L': ".-..",
	'M': "--",
	'N': "-.",
	'O': "---",
	'P': ".--.",
	'Q': "--.-",
	'R': ".-.",
	'S': "...",
	'T': "-",
	'U': "..-",
	'V': "...-",
	'W': ".--",
	'X': "-..-",
	'Y': "-.--",
	'Z': "--..",
	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--....",
	'8': "---..",
	'9': "----.",
	'0': "-----",
}

func main() {
	fmt.Println("Opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("Unable to open gpio", err.Error()))
	}
	defer rpio.Close()

	message := getMessage()

	pin := rpio.Pin(18)
	pin.Output()

	for _, char := range message {
		if char == ' ' {
			wordCut(pin)
			continue
		}
		morseTranslation := morseAlphabet[unicode.ToUpper(char)]
		for _, char := range morseTranslation {
			if char == '.' {
				shortSignal(pin)
			} else if char == '-' {
				longSignal(pin)
			}
			time.Sleep(time.Millisecond * 100)
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func getMessage() string {
	if len(os.Args) < 2 {
		fmt.Println("You need to provide a message !")
		os.Exit(1)
	}
	return string(os.Args[1])
}

func shortSignal(pin rpio.Pin) {
	fmt.Println("Short Signal")
	pin.High()
	time.Sleep(time.Millisecond * 200)
	pin.Low()
}

func longSignal(pin rpio.Pin) {
	fmt.Println("Long Signal")
	pin.High()
	time.Sleep(time.Millisecond * 400)
	pin.Low()
}

func wordCut(pin rpio.Pin) {
	fmt.Println("Word Cut")
	time.Sleep(time.Millisecond * 600)
}
