package service

import "gitlab.com/sannonthachai/find-the-hidden-backend/business/vocab"

type vocabService struct {
	vocabRepo vocab.Repository
}

func NewVocabService(vocabRepo vocab.Repository) vocab.Service {
	return &vocabService{
		vocabRepo: vocabRepo,
	}
}
