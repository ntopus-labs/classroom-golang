package main

/**
  ANSI escape is what can help us make visual changes
  https://en.wikipedia.org/wiki/ANSI_escape_code
*/
func main() {
  var message string
  var colorCode string = "33"
  var CSIColorClear = "\033[0m" 
  CSIColorStart := "\033[1;"+ colorCode +"m"
  message = "Texto colorido"
	println(CSIColorStart+message+CSIColorClear)
}