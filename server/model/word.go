package model

type Word struct {
	Model
	Uuid       string       `gorm:"index;size:32;comment:唯一标识符;" json:"uuid"`
	Name       string       `gorm:"size:32;" json:"name"`
	Characters []*Character `gorm:"foreignKey:pid;references:uuid;comment:关联词性表;" json:"characters"`
	Phrases    []*Phrase    `gorm:"foreignKey:pid;references:uuid;comment:关联短语表;" json:"phrases"`
	Sentences  []*Sentence  `gorm:"foreignKey:pid;references:uuid;comment:关联句子表;" json:"sentences"`
}

type Character struct {
	Model
	Uuid     string     `gorm:"index;size:32;comment:唯一标识符;" json:"uuid"`
	Pid      string     `gorm:"index;size:32;" json:"pid"`
	Ename    string     `gorm:"size:20;" json:"ename"`
	Cname    string     `gorm:"size:30;" json:"cname"`
	Meanings []*Meaning `gorm:"foreignKey:pid;references:uuid;comment:关联词义表;" json:"meanings"`
}

type Meaning struct {
	Model
	Uuid    string `gorm:"index;size:32;comment:唯一标识符;" json:"uuid"`
	Pid     string `gorm:"index;size:32;" json:"pid"`
	Explain string `gorm:"text;" json:"explain"`
}

type Phrase struct {
	Model
	Uuid    string `gorm:"index;size:32;comment:唯一标识符;" json:"uuid"`
	Pid     string `gorm:"index;size:32;" json:"pid"`
	Phrase  string `gorm:"siez:100;" json:"phrase"`
	Explain string `gorm:"text;" json:"explain"`
}

type Sentence struct {
	Model
	Uuid     string `gorm:"index;size:32;comment:唯一标识符;" json:"uuid"`
	Pid      string `gorm:"index;size:32;" json:"pid"`
	Sentence string `gorm:"text;" json:"sentence"`
	Explain  string `gorm:"text;" json:"explain"`
}
