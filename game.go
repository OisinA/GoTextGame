package main

import (
    "bufio"
    "fmt"
    "os"
)

type CharClass struct {
    name string
    dmgmodifier float64
}

type Player struct {
  name string
  class CharClass
  health int
}

type Enemy struct {
  name string
  difficulty int
  health int
}

func ClearLine() {
  for i := 0; i < 34; i++ {
    fmt.Print("-")
  }
  fmt.Print("\n")
}

func DisplayName(p Player) {
  fmt.Print("You are playing as ", p.name, " the ", p.class.name, ".\n")
}

var terminate = false

func main() {

  reader := bufio.NewReader(os.Stdin)
  ClearLine()
  fmt.Print("Please enter your name: ")

  text, _ := reader.ReadString('\n')
  text = text[:len(text) - 2]
  ClearLine()
  p := Player{text, CharClass{"Wizard", 0.25}, 20}

  DisplayName(p)

  enemy := Enemy{"Zombie", 1, 20}

  ClearLine()

  fmt.Print("You stumble across a ", enemy.name, ".\n")

  for {
    if terminate {
      break
    }

    fmt.Print("What do you do?\n> ")

    action, _ := reader.ReadString('\n')
    fmt.Print(action)

  }

  ClearLine()

}
