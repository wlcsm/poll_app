package common

type QuestionOptions int

const (
	MultiChoice QuestionOptions = iota
	SingleChoice
	YesOrNo
	FreeText
)

type Poll struct {
	Id        int64      `json:"id"`
	Questions []Question `json:"questions"`
}

type Response struct {
	Responder string
	Answers   []Answer `json:"answers"`
}

type Answer struct {
	Id   int64       `json:"id"`
	Data interface{} `json:"data"`
}

type Question struct {
	Type QuestionOptions `json:"type"`
	Id   int64           `json:"id"`
	Data interface{}     `json:"data"`
}

type SingleChoiceQuestion struct {
	Header string `json:"header"`
}

type HTTPResponse struct {
	Code int
	Data interface{}
	Err  error
}
