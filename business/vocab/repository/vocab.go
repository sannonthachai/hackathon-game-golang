package repository

import (
	"fmt"

	model "gitlab.com/sannonthachai/find-the-hidden-backend/model/vocab"
)

func (r *vocabRepository) GetVocabByChapter(chapter int) ([]model.Vocab, error) {
	vocab := []model.Vocab{}

	if err := r.vocabDB.Table("vocab").Where("chapter = ?", chapter).Find(&vocab).Error; err != nil {
		fmt.Println("Error Repo GetVocabByChapter: ", err)
		return vocab, err
	}

	return vocab, nil
}
