package sample010

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//DefaultPossibleLetters is a default keys usage in game
var DefaultPossibleLetters = map[string]string{
  "W": mountColoredText(Pentagon + " ", Yellow), 
  "S": mountColoredText(Elipse + " ", Purple), 
  "A": mountColoredText(Circle + " ", Blue), 
  "D": mountColoredText(Hexagon + " ", Red),
}

// Game represents a game struct main control
type Game struct {
  Score, TopScore int
  possibleLetters map[string]string
  possibleLettersKeys []string
  TerminalCtrl TerminalControl
}

// NewGame Instantiate a new Game with default values
func NewGame(possibleLetters map[string]string, topScore int) Game {
	return Game{
		possibleLetters: possibleLetters,
		TopScore: topScore,
		possibleLettersKeys: extractPossibleLettersKeys(possibleLetters),
		TerminalCtrl: NewTerminalControl(possibleLetters),
		Score: 0,
	}
}

//Start initialize the game rotines
func (g *Game) Start() {
  g.TerminalCtrl.StartKeyCapture()
  go g.gameLoop()
}

//WaitFinish Wait all neccessary things finish to return
func (g *Game) WaitFinish() {
  g.TerminalCtrl.WaitExitKeyPress()
}

func (g *Game) gameLoop(){
  g.printIstructions()
  var letterSequence []string
  var stage, score, topScore int
  for {
    stage++
    g.printSumary(stage, score, topScore)
    letterSequence = g.drawLetterSequence(stage)
    g.printSequence(letterSequence)
    points, correct := g.verifySequence(letterSequence)
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

func (g *Game) printIstructions() {
  println("Hi, wellcome to our Simon game.\nPlease press the same sequence order:")
  for char, symbol := range g.possibleLetters {
    println("Press ", char, " to ", symbol)
  }
  println("Or press 'Esc' to go out!")
}

func (g *Game) printSumary(stage, score, topScore int) {
  fmt.Printf("\nOk, score=%d | stage=%d | top score=%d\n", score, stage, topScore)
  println("Be prepered to the symbol sequence!")
  g.printBackCounter(3, 500)
}

func (g *Game) printBackCounter(from, ms int) {
  for i := from; i > 0; i-- {
    fmt.Printf("%s%d", strings.Repeat(" ", from-i), i)
    time.Sleep(time.Duration(ms) * time.Millisecond)
    fmt.Printf("\r%s\r", strings.Repeat(" ", (from-i+1)*2))
  }
}
func (g *Game) drawLetterSequence(stage int) []string {
  sequence := make([]string, stage)
  for i := range sequence {
    sequence[i] = g.drawLetter()
  }
  return sequence
}

func (g *Game) drawLetter() string {
  rand.Seed(time.Now().UnixNano())
  size := len(g.possibleLettersKeys)
  return g.possibleLettersKeys[rand.Intn(size)]
}

func (g *Game) printSequence(letterSequence []string) {
  size := len(letterSequence)
  waitTime := time.Duration(1000 - size*10)
  for i, letter := range letterSequence {
    fmt.Printf("%s", strings.Repeat(" ", i*2)+g.possibleLetters[letter])
    time.Sleep(waitTime * time.Millisecond)
    fmt.Printf("\r%s\r", strings.Repeat(" ", (i+1)*2))
  }
}
func (g *Game) verifySequence(letterSequence []string) (int, bool) {
  defer g.TerminalCtrl.DisableKeyPress()
  g.TerminalCtrl.EnableKeyPress()
  points := 0
  for _, letter := range letterSequence {
    if letter != g.TerminalCtrl.ReadKeyPressed() {
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