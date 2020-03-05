package main

import "./sample010"

func main() {
  game := sample010.NewGame(sample010.DefaultPossibleLetters, 0)
  game.Start()
  game.WaitFinish()
  println("Why give up? Come back again!")
}

