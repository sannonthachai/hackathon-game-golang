package vocab

import model "gitlab.com/sannonthachai/find-the-hidden-backend/model/vocab"

type Service interface {
	GetVocabByChapter(chapter int) ([]model.Vocab, error)
}
