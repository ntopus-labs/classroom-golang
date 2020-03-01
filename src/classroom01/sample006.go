package main

import (
	"strconv"
  "os"
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

/**
  Args Usage
  https://gobyexample.com/command-line-arguments
*/
func main() {
  colorList := []string{
    yellow,
    blue,
    purple,
    gray,
  }
  textList := []string{
    pentagon,
    elipse,
    hexagon,
    circle,
  }
  defer println("")
  argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		loopTimes, err := strconv.Atoi(argsWithoutProg[0])
		if err != nil {
      println("Fail to undestand Loop Times Value:")
      panic(err)
    }
    i:= 0;
    for  i < loopTimes {
      printAllColoredTextInList(textList, colorList); i++
    } 
    return
	}
  for {
    printAllColoredTextInList(textList, colorList)
  }
}

func printAllColoredTextInList(textList, colorList [] string) {
  for index := 0; index < len(textList); index++ {
    printExpecificColoredTextForEachColorInList(textList[index], colorList)
  }
}

func printExpecificColoredTextForEachColorInList(text string, colorList []string) {
	for _, color := range colorList {
    if text == pentagon {
      if color != yellow { continue }  
      printColoredText(text+" ", color)
    } else if color == purple {
      if text == elipse { continue }  
      printColoredText(text+" ", color)
    } else {
      printColoredText(text+" ", color)
      if text == circle { return }
    }
	}
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
