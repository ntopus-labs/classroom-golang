package sample010

import (
	"strings"
	"sync"

	"github.com/eiannone/keyboard"
)

//TerminalControl handle all terminal iteractions
type TerminalControl struct {
	wg sync.WaitGroup
	canPress bool
	possibleLetters map[string]string
	keyPressed chan string
}

// NewTerminalControl Instatiate a new TerminalControl with default values
func NewTerminalControl(possibleLetters map[string]string) TerminalControl {
	return TerminalControl{
		sync.WaitGroup{},
		false,
		possibleLetters,
		make(chan string),
	}
}

//StartKeyCapture Start an async loop of key capture
func (t *TerminalControl) StartKeyCapture() {
  t.wg.Add(1)
  go t.keyCaptureLoop()
}

//WaitExitKeyPress wait a exit key be pressed
func (t *TerminalControl) WaitExitKeyPress() {
  t.wg.Wait()
} 

//EnableKeyPress can press possible keys
func (t *TerminalControl) EnableKeyPress() {
	t.canPress = true
}

//DisableKeyPress can press possible keys
func (t *TerminalControl) DisableKeyPress() {
	t.canPress = false
}

//ReadKeyPressed return a pressed key from channel
func (t *TerminalControl) ReadKeyPressed() string {
	return <-t.keyPressed
}

func (t *TerminalControl) keyCaptureLoop() {
  defer t.wg.Done()
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
    if !t.canPress {
      continue
    }
    t.handleKeyPressed(char)
	}
}

func (t *TerminalControl) handleKeyPressed(char rune) {
  charUpcase := strings.ToUpper(string(char)) 
  if _, ok := t.possibleLetters[charUpcase]; ok == false {
    return
  }
  print(t.possibleLetters[charUpcase])
  t.keyPressed <- charUpcase
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