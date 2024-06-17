package model

type Vocab struct {
	VocabEn      string `json:"vocabEn" gorm:"column:vocab_en"`
	VocabTh      string `json:"vocabTh" gorm:"column:vocab_th"`
	PartOfSpeech string `json:"part_of_speech" gorm:"column:part_of_speech"`
	Chapter      int    `json:"chapter" gorm:"column:chapter"`
	Sentence     string `json:"sentence" gorm:"column:sentence"`
}
