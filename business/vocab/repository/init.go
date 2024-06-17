package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/sannonthachai/find-the-hidden-backend/business/vocab"
)

type vocabRepository struct {
	vocabDB *gorm.DB
}

func NewVocabRepository(vocabDB *gorm.DB) vocab.Repository {
	return &vocabRepository{
		vocabDB: vocabDB,
	}
}
