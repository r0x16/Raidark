package repository

import "git.tnschile.com/sistemas/tnsgo/raidark/src/shared/domain/model"

type RandomRepository interface {
	GenerateRandomString(int) *model.Random
}