package main

/**
  Unicode / UTF8 Characters
  https://en.wikipedia.org/wiki/UTF-8
  https://en.wikipedia.org/wiki/Unicode#UTF
  https://en.wikibooks.org/wiki/Unicode/Character_reference/2000-2FFF
*/
func main() {
  const black string = "30" 
  const red = "31" 
  const green = "32" 
  const yellow = "33" 
  const blue = "34" 
  const purple = "35" 
  const lightBlue = "36" 
  const gray = "37" 
  const pentagon = "\u2B1F"
  const hexagon = "\u2B22"
  const circle = "\u2B24"
  const elipse = "\u2B2E"
	printColoredText(pentagon, yellow)
	printColoredText(hexagon, blue)
	printColoredText(circle, purple)
	printColoredText(elipse, lightBlue)
}

func printColoredText(text string, colorCode string) {
	println(mountColoredText(text, colorCode))
}
/**
  ANSI escape is what can help us make visual changes
  https://en.wikipedia.org/wiki/ANSI_escape_code
*/
func mountColoredText(text, colorCode string) string {
  CSIColorClear := "\033[0m" 
  CSIColorStart := "\033[1;"+ colorCode +"m"
  return CSIColorStart+text+CSIColorClear
}
