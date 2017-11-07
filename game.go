package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
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

var player Player
var enemy Enemy

func (p *Player) Damage(damage int) {
    p.health = p.health - damage
}

func (e *Enemy) Damage(damage int) {
    e.health = e.health - damage
}

func (p Player) Attack(enemy *Enemy, damage int) {
    enemy.Damage(damage)
    fmt.Print("You damaged the enemy for ", damage, " damage.\n")
}

func (p *Player) ParseAction(action string) {
    if action == "attack" {
        p.Attack(&enemy, 3)
    }
}

func ClearLine() {
    for i := 0; i < 34; i++ {
      fmt.Print("-")
    }
    fmt.Print("\n")
}

func (p Player) DisplayName() {
    fmt.Print("You are playing as ", p.name, " the ", p.class.name, ".\n")
}

func CheckHealth() bool {
    return player.health <= 0 || enemy.health <= 0
}

func EndGame() {
    phealth := player.health
    ehealth := enemy.health

    fmt.Print("The game is over.\n")

    if(phealth <= 0) {
        fmt.Print("You have died! Unfortunate.\n")
        return
    }

    if(ehealth <= 0) {
        fmt.Print("You have slain the ", enemy.name, "!\n")
        return
    }
}

var terminate = false

func main() {

    reader := bufio.NewReader(os.Stdin)
    ClearLine()
    fmt.Print("Please enter your name: ")

    text, _ := reader.ReadString('\n')
    text = text[:len(text) - 2]
    ClearLine()
    player = Player{text, CharClass{"Wizard", 0.25}, 20}

    player.DisplayName()

    enemy = Enemy{"Zombie", 1, 20}

    ClearLine()

    fmt.Print("You stumble across a ", enemy.name, ".\n")

    for {
        if terminate {
            EndGame()
            break
        }

        fmt.Print("What do you do?\n> ")

        action, _ := reader.ReadString('\n')
        action = action[:len(action) - 2]
        action = strings.ToLower(action)
        ClearLine()

        validActions := map[string]bool {
            "attack": true,
            "flee": true,
        }

        if validActions[action] {
              player.ParseAction(action)
        } else {
              continue
        }

        fmt.Print("You have ", player.health, " health remaining.\n")
        fmt.Print("The enemy has ", enemy.health, " health remaining.\n")
        ClearLine()

        terminate = CheckHealth()
    }

    ClearLine()

}
