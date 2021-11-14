package common

type QuestionOptions int

const (
	Multichoice QuestionOptions = iota
	SingleChoice
	YesOrNo
	FreeText
)

type Poll struct {
	Id        int64      `json:"id"`
	Questions []Question `json:"questions"`
}

type PollResponse struct {
	Answers []Answer `json:"answer"`
}

type Answer struct {
	Id   int         `json:"id"`
	Data interface{} `json:"data"`
}

type Question struct {
	Type QuestionOptions `json:"type"`
	Id   int             `json:"id"`
	Data interface{}     `json:"data"`
}

type SingleChoiceQuestion struct {
	Header string `json:"header"`
}
