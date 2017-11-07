package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

//This struct stores the information about the character's class.
type CharClass struct {
    name string
    dmgmodifier float64
}

//This struct stores the name, class and health of the player.
type Player struct {
    name string
    class CharClass
    health int
}

//This struct stores the name, difficulty and health of the enemy.
type Enemy struct {
    name string
    difficulty int
    health int
}

var player Player
var enemy Enemy

//Player method to damage the player.
func (p *Player) Damage(damage int) {
    p.health = p.health - damage
}

//Enemy method to damage the enemy.
func (e *Enemy) Damage(damage int) {
    e.health = e.health - damage
}

//Player method allowing the player to attack the enemy for a certain amount of damage.
func (p Player) Attack(enemy *Enemy, damage int) {
    enemy.Damage(damage)
    fmt.Print("You damaged the enemy for ", damage, " damage.\n")
}

//Parse the user's inputted action and convert it to the appropriate method.
func (p *Player) ParseAction(action string) {
    switch(action) {
        case "attack" {
            p.attack(&enemy, 3)
        }
    }
}

//Print a line of the - character to break up the text.
func ClearLine() {
    for i := 0; i < 34; i++ {
      fmt.Print("-")
    }
    fmt.Print("\n")
}

//Display the user's name and class to the user.
func (p Player) DisplayName() {
    fmt.Print("You are playing as ", p.name, " the ", p.class.name, ".\n")
}

//This function returns whether either the player or enemy are dead.
func CheckHealth() bool {
    return player.health <= 0 || enemy.health <= 0
}

//This function is called when the game is over. It prints out the ending messages.
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

//Main function of the game.
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

    //Main for loop of the game. Continues until the 'terminate' variable is true.
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
