package models

import (
	"fmt"
	"math/rand"
)

type Player struct {
	name  string
	isBot bool
}

var (
	randomNames = []string{"Dmitry", "Daniil", "Alexander", "Oleg", "Andrew"}
)

func NewPlayer(name string) *Player {
	if len(name) == 0 {
		return &Player{
			name:  fmt.Sprintf("bot_%s", randomNames[rand.Intn(len(randomNames))]),
			isBot: true,
		}
	}
	return &Player{
		name:  name,
		isBot: false,
	}
}

func (p *Player) IsBot() bool {
	return p.isBot
}

func (p *Player) Name() string {
	return p.name
}
