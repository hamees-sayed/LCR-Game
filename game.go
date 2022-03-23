package main
import ("fmt"
		"math/rand"
		"time"
)

func main() {
	fmt.Println("Welcome to LCR dice game :)")
	
	g := new()
	
	fmt.Println("Please enter how many players will play the game: ")
	fmt.Println("NOTE: enter number more than 2.")
	
}


type game struct {
	players []*player
}

// join add new player to the game
func(g *game) Join (playerName string) *player {
	// create new player
	p := player {
		name :	 playerName,
		tokens : 3,
	}
	playerCount := len(g.player)
		if playerCount > 0 {
			lastPlayer := g.players[playerCount-1]
			p.left = lastPlayer
			p.right = g.players[0]
			lastplayer.right = &p
			g.players[0].left = &p
		}
		g.players = append(g.players, &p)
		return
}

// fininshed check do we have only one player with Tokens
//this should be done after every turn
func (g *game) fininshed() (p *players) {
	playersWithTokens := 0
	for _, value : range g.players {
		if value.tokens > 0 {
			playersWithTokens++
				if playersWithTokens > 1 {
					return nil
  			}
			p = value
		}
	}
	return
}		



// new int game

func new() *game {
	return &game{}
}

// player
type player struct {
	name string
	tokens int
	right *player
	left *player
}

// roll uses rand to return diceface
func (p *player) rollDice() (result []string) {
	// find out how many dices we should roll
	// if user have more than 3 tokens he can only roll 3 dices
	// or he roll exact number of tokens as dices.
	
	dices := p.tokens
		if p.tokens > 2 {
			dices = 3
		}
		
		// roll the dices and update the value on the players
		for index := 0; index < dices :index++ {
			d := dice{}
			diceResult = d.roll()
			result = append(result, diceResult)
			
			
			switch diceResult {
			case "right" :
				p.tokens--
				p.right.tokens++
			case "left" :
				p.tokens--
				p.left.tokens++
			case "center" :
				p.tokens--
			}
		}
		return
}

// dice
type dice struct {
}

func (d *dice) roll() string {
	rand.Seed(time.Now().UnixNano)
	r := rand.Intn(6)
	
	switch r {
		case 0 : 
			return "right"
		case 1 :
			return "left"
		case 2 :
			return "center"
		default :
			return "DoNothing" 
	}
}