package vocab

import model "gitlab.com/sannonthachai/find-the-hidden-backend/model/vocab"

type Repository interface {
	GetVocabByChapter(chapter int) ([]model.Vocab, error)
}
