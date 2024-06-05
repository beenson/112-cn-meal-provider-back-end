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

// /api/cart
type AddToCartRequest struct {
	UserId string `json:"uid"`
	FoodId string `json:"foodId"`
	Amount uint32 `json:"amount"`
}

type GetCartItemRequest struct {
	UserId string `json:"uid,omitempty"`
}

type UpdateCartAmountRequest struct {
	UserId string `json:"uid"`
	Amount uint32 `json:"amount"`
	Itemid string `json:"itemId"`
}

type ClearCartRequest struct {
	UserId string `json:"uid"`
}

// /api/food
type CreateFoodRequest struct {
	Name        string `json:"name"`
	Description string `json:"descrpition"`
	Price       int32  `json:"price"`
	URL         string `json:"url"`
}

type GetFoodRequest struct {
	Id string `json:"foodId"`
}

type UpdateFoodRequest struct {
	Id          string `json:"foodId"`
	Name        string `json:"name"`
	Description string `json:"descrpition"`
	Price       int32  `json:"price"`
	URL         string `json:"url"`
}

type DeleteFoodRequest struct {
	Id string `json:"foodId"`
}

type MakeOrderRequest struct {
	Uid string `json:"uid"`
}

type GetOrderRequest struct {
	OrderId string `json:"orderId"`
}

type GetOrdersRequest struct {
	Uid string `json:"uid"`
}
