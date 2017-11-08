package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "math/rand"
    "time"
)

//This struct stores the information about the character's class.
type CharClass struct {
    name string
    maxdmg int
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

var (
    player Player
    enemy Enemy
)

var charclasses = [2]CharClass {
    {"Barbarian", 5, 0.5},
    {"Wizard", 8, 1.0},
}

//Player method to damage the player.
func (p *Player) Damage(damage int) {
    p.health = p.health - damage
}

//Enemy method to damage the enemy.
func (e *Enemy) Damage(damage int) {
    e.health = e.health - damage
}

func (e Enemy) Attack(player *Player, damage int) {
    damage = damage * e.difficulty
    damageModified := float64(damage) * player.class.dmgmodifier
    if damageModified != 1.0 {
        fmt.Println("You are taking ", (player.class.dmgmodifier * 100), "% damage.")
        fmt.Println("You would've taken ", damage, " damage.")
    }
    player.Damage(int(damageModified))
    fmt.Println("The zombie damaged you for ", int(damageModified), " damage.")
}

//Player method allowing the player to attack the enemy for a certain amount of damage.
func (p Player) Attack(enemy *Enemy, damage int) {
    enemy.Damage(damage)
    fmt.Println("You damaged the enemy for ", damage, " damage.")
}

//Parse the user's inputted action and convert it to the appropriate method.
func (p *Player) ParseAction(action string) {
    switch(action) {
        case "attack":
            p.Attack(&enemy, GenRandomNumber(p.class.maxdmg))
    }
}

//Print a line of the - character to break up the text.
func ClearLine() {
    for i := 0; i < 34; i++ {
      fmt.Print("-")
    }
    fmt.Println("")
}

func GenRandomNumber(max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max) + 1
}

//Display the user's name and class to the user.
func (p Player) DisplayName() {
    fmt.Println("You are playing as ", p.name, " the ", p.class.name, ".")
}

//This function returns whether either the player or enemy are dead.
func CheckHealth() bool {
    return player.health <= 0 || enemy.health <= 0
}

//This function is called when the game is over. It prints out the ending messages.
func EndGame() {
    phealth := player.health
    ehealth := enemy.health

    fmt.Println("The game is over.")

    if phealth <= 0 {
        fmt.Println("You have died! Unfortunate.")
        return
    }

    if ehealth <= 0 {
        fmt.Println("You have slain the ", enemy.name, "!")
        return
    }
}

//Main for loop of the game. Continues until the 'terminate' variable is true.
func GameLoop() {
    reader := bufio.NewReader(os.Stdin)
    validActions := map[string]bool {
        "attack": true,
        "flee": true,
    }
    turn := 1
    for {
        if terminate {
            EndGame()
            break
        }

        fmt.Println("Turn #", turn)
        fmt.Print("What do you do?\n> ")

        action, _ := reader.ReadString('\n')
        action = action[:len(action) - 2]
        action = strings.ToLower(action)
        ClearLine()

        if validActions[action] {
            player.ParseAction(action)
        } else {
            continue
        }

        fmt.Println("")
        enemy.Attack(&player, GenRandomNumber(enemy.difficulty * 4))
        fmt.Println("")

        fmt.Println("You have ", player.health, " health remaining.")
        fmt.Println("The enemy has ", enemy.health, " health remaining.")
        ClearLine()

        terminate = CheckHealth()
        turn = turn + 1
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

    fmt.Println("Available classes: ")
    for index, i := range charclasses {
        fmt.Print("#", (index + 1), ": ", i.name)
        if index != len(charclasses) - 1 {
            fmt.Println("")
        }
    }
    fmt.Println("")
    ClearLine()

    classPicked := false
    var charClass CharClass
    for {
        if classPicked {
            break
        }

        fmt.Print("Which class would you like to play?\n> ")

        class, _ := reader.ReadString('\n')
        class = class[:len(class) - 2]
        class = strings.ToLower(class)
        for _, i := range charclasses {
            if class == strings.ToLower(i.name) {
                charClass = i
                classPicked = true
            }
        }
    }

    player = Player{text, charClass, 20}

    player.DisplayName()

    enemy = Enemy{"Zombie", 1, 20}

    ClearLine()

    fmt.Println("You stumble across a ", enemy.name, ".")

    GameLoop()

    ClearLine()

}
