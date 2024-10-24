package server

import (
	"fmt"
	"testing"

	"github.com/sirgwain/craig-stars/ai"
	"github.com/sirgwain/craig-stars/config"
	"github.com/sirgwain/craig-stars/cs"
	"github.com/sirgwain/craig-stars/db"
	"github.com/stretchr/testify/assert"
)

func createTestGameRunner() GameRunner {
	dbConn := db.NewConn()
	cfg := &config.Config{}
	// cfg.Database.Filename = "../data/sqlx.db"
	cfg.Database.Filename = ":memory:"
	cfg.Database.UsersFilename = ":memory:"
	cfg.Database.Recreate = true
	cfg.Database.DebugLogging = true
	if err := dbConn.Connect(cfg); err != nil {
		panic(fmt.Errorf("connect to test database, %w", err))
	}

	return &gameRunner{
		dbConn: dbConn,
		client: cs.NewGamer(),
	}
}

func Test_gameRunner_HostGame(t *testing.T) {

	gr := createTestGameRunner()

	fullGame, err := gr.HostGame(1, cs.NewGameSettings().WithHost(cs.Humanoids()).WithAIPlayer(cs.AIDifficultyNormal, 0))

	if err != nil {
		t.Errorf("host game %v", err)
		return
	}

	// make sure we generate some universes
	assert.Greater(t, len(fullGame.Planets), 0)
	assert.Greater(t, len(fullGame.Players), 0)
}

func Test_gameRunner_GenerateTurns(t *testing.T) {

	dbConn := db.NewConn()
	cfg := &config.Config{}
	cfg.Database.Filename = ":memory:"
	cfg.Database.UsersFilename = ":memory:"
	if err := dbConn.Connect(cfg); err != nil {
		panic(fmt.Errorf("connect to test database, %w", err))
	}

	// create a race per PRT
	for _, prt := range cs.PRTs {
		race := cs.NewRace()
		race.PRT = prt
		race.Name = fmt.Sprintf("%v", prt)
		race.PluralName = fmt.Sprintf("%vs", prt)
	}

	gr := gameRunner{
		dbConn: dbConn,
		client: cs.NewGamer(),
	}

	// create a game with AI players for each PRT
	fullGame, err := gr.HostGame(1, cs.NewGameSettings().
		WithName("All Races Test").
		WithSize(cs.SizeMedium).
		WithAIPlayerRace(ai.Races[0], cs.AIDifficultyNormal, 0).
		WithAIPlayerRace(ai.Races[1], cs.AIDifficultyNormal, 1).
		WithAIPlayerRace(ai.Races[2], cs.AIDifficultyNormal, 2).
		WithAIPlayerRace(ai.Races[3], cs.AIDifficultyNormal, 3).
		WithAIPlayerRace(ai.Races[4], cs.AIDifficultyNormal, 0).
		WithAIPlayerRace(ai.Races[5], cs.AIDifficultyNormal, 1).
		WithAIPlayerRace(ai.Races[6], cs.AIDifficultyNormal, 2).
		WithAIPlayerRace(ai.Races[7], cs.AIDifficultyNormal, 3).
		WithAIPlayerRace(ai.Races[8], cs.AIDifficultyNormal, 0).
		WithAIPlayerRace(ai.Races[9], cs.AIDifficultyNormal, 1))

	if err != nil {
		t.Errorf("host game %v", err)
	}

	// generate 100 turns
	for i := 0; i < 100; i++ {
		gr := gameRunner{
			dbConn: dbConn,
			client: cs.NewGamer(),
		}

		if _, err := gr.GenerateTurn(fullGame.ID); err != nil {
			t.Errorf("generate turn %v", err)
		}
	}
}

func Test_gameRunner_getGuestNum(t *testing.T) {
	tests := []struct {
		name     string
		username string
		want     int
		wantErr  bool
	}{
		{"1", "guest-1-1", 1, false},
		{"20", "guest-29-20", 20, false},
		{"fail", "bob", 0, true},
		{"fail", "bob-1-bob", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := &gameRunner{}
			u := cs.User{Username: tt.username}
			got, err := gr.getGuestNum(&u)
			if (err != nil) != tt.wantErr {
				t.Errorf("gameRunner.getGuestNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gameRunner.getGuestNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
