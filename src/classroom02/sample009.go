package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/eiannone/keyboard"
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
var possibleLetters = map[string]string{
  "W": mountColoredText(pentagon + " ", yellow), 
  "S": mountColoredText(elipse + " ", purple), 
  "A": mountColoredText(circle + " ", blue), 
  "D": mountColoredText(hexagon + " ", red),
}
var possibleLettersKeys = extractPossibleLettersKeys(possibleLetters)
var canPress = false

func main() {
  defer func () {
    println("Why give up? Come back again!")
  }()
  defer wg.Wait()
  go gameLoop()
  wg.Add(1)
  go keyCaptureLoop()
}

func gameLoop(){
  printIstructions()
  var letterSequence []string
  var stage, score, topScore int
  for {
    stage++
    printSumary(stage, score, topScore)
    letterSequence = drawLetterSequence(stage)
    printSequence(letterSequence)
    points, correct := verifySequence(letterSequence)
    score += points
    if score > topScore {
      topScore = score
    }
    if !correct {
      score=0
      stage=0
      println("\nAs expected you missed!")
      continue
    }
    println("\nCongratulations you guessed! Good Memory!")
  }
}

func printIstructions() {
  println("Hi, wellcome to our Simon game.\nPlease press the same sequence order:")
  for char, symbol := range possibleLetters {
    println("Press ", char, " to ", symbol)
  }
  println("Or press 'Esc' to go out!")
}

func printSumary(stage, score, topScore int) {
  fmt.Printf("\nOk, score=%d | stage=%d | top score=%d\n", score, stage, topScore)
  println("Be prepered to the symbol sequence!")
  printBackCounter(3, 500)
}

func printBackCounter(from, ms int) {
  for i := from; i > 0; i-- {
    fmt.Printf("%s%d", strings.Repeat(" ", from-i), i)
    time.Sleep(time.Duration(ms) * time.Millisecond)
    fmt.Printf("\r%s\r", strings.Repeat(" ", (from-i+1)*2))
  }
}
func drawLetterSequence(stage int) []string {
  sequence := make([]string, stage)
  for i := range sequence {
    sequence[i] = drawLetter()
  }
  return sequence
}

func drawLetter() string {
  rand.Seed(time.Now().UnixNano())
  size := len(possibleLettersKeys)
  return possibleLettersKeys[rand.Intn(size)]
}

func printSequence(letterSequence []string) {
  size := len(letterSequence)
  waitTime := time.Duration(1000 - size*10)
  for i, letter := range letterSequence {
    fmt.Printf("%s", strings.Repeat(" ", i*2)+possibleLetters[letter])
    time.Sleep(waitTime * time.Millisecond)
    fmt.Printf("\r%s\r", strings.Repeat(" ", (i+1)*2))
  }
}
func verifySequence(letterSequence []string) (int, bool) {
  defer func() { canPress = false }()
  canPress = true
  points := 0
  for _, letter := range letterSequence {
    if letter != <-keyPressed {
      return points, false
    }
    points += 10
  }
  return points, true
}

func extractPossibleLettersKeys(possibleLettersMap map[string]string) []string {
  keys := make([]string, 0, len(possibleLettersMap))
  for k := range possibleLettersMap {
    keys = append(keys, k)
  }
  return keys
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
    if !canPress {
      continue
    }
    handleKeyPressed(char)
	}
}

func handleKeyPressed(char rune) {
  charUpcase := strings.ToUpper(string(char)) 
  if _, ok := possibleLetters[charUpcase]; ok == false {
    return
  }
  print(possibleLetters[charUpcase])
  keyPressed <- charUpcase
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
