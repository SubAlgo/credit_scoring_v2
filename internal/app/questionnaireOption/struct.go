package questionnaireOption

type getQuestionnaireOptionRequest struct {
}

type option struct {
	Value string `json:"value"`
	Text  string `json:"text"`
}
