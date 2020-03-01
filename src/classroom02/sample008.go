package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
  "strings"
  "math/rand"
	"sync"
)

/**
  Unicode / UTF8 Characters
  https://en.wikipedia.org/wiki/UTF-8
  https://en.wikipedia.org/wiki/Unicode#UTF
  https://en.wikibooks.org/wiki/Unicode/Character_reference/2000-2FFF
*/
const pentagon = "\u2B1F"
const hexagon = "\u2B22"
const circle = "\u2B24"
const elipse = "\u2B2E"
const black = "30"
const red = "31"
const green = "32"
const yellow = "33"
const blue = "34"
const purple = "35"
const lightblue = "36"
const gray = "37"

var wg = sync.WaitGroup{}
var keyPressed = make(chan string)
var possibleLetters = []string{"W", "S", "A", "D"}

func main() {
  defer println("")
  defer wg.Wait()
  go gameLoop()
  wg.Add(1)
  go keyCaptureLoop()
}

func gameLoop(){
  printIstructions()
  var letter string
  var correct, wrong int
  for {
    printSumary(correct, wrong)
    letter = drawLetter()
    if letter == <-keyPressed {
      correct++
      println("\nCongratulations you guessed!")
    } else {
      wrong++
      println("\nAs expected you missed!")
    }
  }
}

func printIstructions() {
  println("Hi, wellcome to our game.\nPlease guess which symbol I thought:")
  println("Press W to " + mountColoredText(pentagon, yellow))
  println("Press S to " + mountColoredText(elipse, purple))
  println("Press A to " + mountColoredText(circle, blue))
  println("Press D to " + mountColoredText(hexagon, red))
  println("Or press 'Esc' to go out!")
}

func printSumary(correct, wrong int) {
  fmt.Printf("\nOk, me=%d | you=%d\n", wrong, correct)
  println("What symbol did I think?")
}

func drawLetter() string {
  return possibleLetters[rand.Intn(len(possibleLetters))]
}

func keyCaptureLoop() {
  defer wg.Done()
  err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()
	for {
    char, key, err := keyboard.GetKey()
		if (err != nil) {
			panic(err)
		} 
    if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC {
      return
    }
    handleKeyPressed(char)
	}
}

func handleKeyPressed(char rune) {
  charUpcase := strings.ToUpper(string(char)) 
  switch charUpcase {
    case "W":
      printColoredText(pentagon+" ", yellow)
    case "S":
      printColoredText(elipse+" ", purple)
    case "A":
      printColoredText(circle+" ", blue)
    case "D":
      printColoredText(hexagon+" ", red)
    default:
      return
  }
  keyPressed <- charUpcase
}
func printColoredText(text string, colorCode string) {
	print(mountColoredText(text, colorCode))
}

/**
  ANSI escape is what can help us make visual changes
  https://en.wikipedia.org/wiki/ANSI_escape_code
*/
func mountColoredText(text, colorCode string) string {
	CSIColorClear := "\033[0m"
	CSIColorStart := "\033[1;" + colorCode + "m"
	return CSIColorStart + text + CSIColorClear
}
