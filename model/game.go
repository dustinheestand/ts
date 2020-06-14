package model

import (
	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model
	US        User    `gorm:"many2many:user_games;"`
	USSR      User    `gorm:"many2many:user_games;"`
	State     int64   // [8]space[7]score[3]defcon[3]turn[4]round[1]phasing-player[2]turn-state[2]actions-remaining[2]winner[2]early/mid/late
	Countries []int32 // 4 bits per country per player, so 1 8-bit int for each country
	Cards     []int64 // states: deck/not-yet-in-deck(based on round), discard, in effect, US hand, USSR hand, removed, in play -- think I can get by with one byte per card, so I need an int for every eight cards; I can add one more state
	Summary   []int64 // probably total points rec'd, played for events/influence/coup/realign, total scoring cards received, points spaced, average roll in coups/wars/realigns
}
