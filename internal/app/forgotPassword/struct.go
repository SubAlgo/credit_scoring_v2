package forgotPassword

type forgotPasswordRequest struct {
}

type forgotPasswordQuestion struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
}

type forgotPasswordQuestionList struct {
	List []*forgotPasswordQuestion
}
