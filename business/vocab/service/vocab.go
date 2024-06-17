package service

import (
	"math/rand"
	"time"

	model "gitlab.com/sannonthachai/find-the-hidden-backend/model/vocab"
)

func (s *vocabService) GetVocabByChapter(chapter int) ([]model.Vocab, error) {
	result, err := s.vocabRepo.GetVocabByChapter(chapter)
	if err != nil {
		return result, err
	}

	vocab := []model.Vocab{}

	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(len(result))
	for _, r := range p[:10] {
		vocab = append(vocab, result[r])
	}

	return vocab, nil
}
