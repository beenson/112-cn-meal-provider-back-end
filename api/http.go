package api

// login
type UserLoginMsg struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// /api/mail
type PaymentNotifyMsg struct {
	Uid    string `json:"uid"`
	Amount int32  `json:"amount"`
}

// /api/comment
type CreateCommentMsg struct {
	Uid     string `json:"uid"`
	FoodId  int32  `json:"foodId"`
	Rate    int32  `json:"rate"`
	Comment string `json:"comment"`
}

type ReadCommentMsg struct {
	Uid       string `json:"uid,omitempty"`
	FoodId    int32  `json:"foodId,omitempty"`
	Rate      int32  `json:"rate,omitempty"`
	CommentId int32  `json:"commentId,omitempty"`
}

type UpdateCommentMsg struct {
	CommentId int32  `json:"commentId"`
	Uid       string `json:"uid"`
	FoodId    int32  `json:"foodId"`
	Rate      int32  `json:"rate"`
	Comment   string `json:"comment"`
}

type DeleteCommentMsg struct {
	CommentId int32 `json:"commentId"`
}
