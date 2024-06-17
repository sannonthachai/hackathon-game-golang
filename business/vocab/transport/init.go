package transport

import "gitlab.com/sannonthachai/find-the-hidden-backend/business/vocab"

type Handler struct {
	vocabService vocab.Service
}

func NewVocabHandler(vocabService vocab.Service) Handler {
	return Handler{
		vocabService: vocabService,
	}
}
