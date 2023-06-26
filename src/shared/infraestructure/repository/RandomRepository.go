package repository

import (
	"math/rand"
	"time"

	"git.tnschile.com/sistemas/tnsgo/raidark/src/domain/model"
	i "git.tnschile.com/sistemas/tnsgo/raidark/src/domain/repository"
	"gorm.io/gorm"
)

type RandomRepositoryBasic struct {
	db *gorm.DB
}

var _ i.RandomRepository = &RandomRepositoryBasic{}

func NewRandomRepositoryGorm(db *gorm.DB) *RandomRepositoryBasic {
	repo := &RandomRepositoryBasic{db: db}
	return repo
}

// Generates a random string of length n
func (*RandomRepositoryBasic) GenerateRandomString(n int) *model.Random {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    rand.Seed(time.Now().UnixNano())

    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }

	random := &model.Random{
		Value: string(b),
	}

    return random
}