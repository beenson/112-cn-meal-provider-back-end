package manager

import (
	context "context"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/rating/model"
	common "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/common"
	protocol "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/proto/gen/rating"
)

func (m *GRPCManager) CreateFeedback(ctx context.Context, fb *protocol.Feedback) (*protocol.CommentId, error) {
	feedback := model.Feedback{
		Foodid:  uint(*fb.FoodId),
		Uid:     *fb.Uid,
		Rating:  *fb.Rating,
		Comment: *fb.Comment,
	}
	if res := m.db.Create(&feedback); res.Error != nil {
		return nil, res.Error
	}
	ret := int32(feedback.ID)
	return &protocol.CommentId{
		Id: &ret,
	}, nil
}

func (m *GRPCManager) ReadFeedbackByUserId(ctx context.Context, uid *protocol.UserId) (*protocol.FeedbackList, error) {
	var reply protocol.FeedbackList
	var feedbacks []model.Feedback
	if res := m.db.Where("Uid = ?", uid.Id).Find(&feedbacks); res.Error != nil {
		return nil, res.Error
	}
	for i := 0; i < len(feedbacks); i++ {
		repId := int32(feedbacks[i].Foodid)
		repId2 := int32(feedbacks[i].ID)
		reply.List = append(reply.List, &protocol.Feedback{
			FoodId:  &repId,
			Uid:     &feedbacks[i].Uid,
			Rating:  &feedbacks[i].Rating,
			Comment: &feedbacks[i].Comment,
			Id:      &repId2,
		})
	}
	return &reply, nil
}

func (m *GRPCManager) ReadFeedbackByFoodId(ctx context.Context, foodId *protocol.FoodId) (*protocol.FeedbackList, error) {
	var reply protocol.FeedbackList
	var feedbacks []model.Feedback
	if res := m.db.Where("Foodid = ?", foodId.Id).Find(&feedbacks); res.Error != nil {
		return nil, res.Error
	}
	for i := 0; i < len(feedbacks); i++ {
		repId := int32(feedbacks[i].Foodid)
		repId2 := int32(feedbacks[i].ID)
		reply.List = append(reply.List, &protocol.Feedback{
			FoodId:  &repId,
			Uid:     &feedbacks[i].Uid,
			Rating:  &feedbacks[i].Rating,
			Comment: &feedbacks[i].Comment,
			Id:      &repId2,
		})
	}
	return &reply, nil
}

func (m *GRPCManager) ReadFeedbackByFoodRating(ctx context.Context, foodId *protocol.FoodId) (*protocol.FeedbackList, error) {
	var reply protocol.FeedbackList
	var feedbacks []model.Feedback
	if res := m.db.Where("Foodid = ? AND Rating = ?", foodId.Id, foodId.Rate).Find(&feedbacks); res.Error != nil {
		return nil, res.Error
	}
	for i := 0; i < len(feedbacks); i++ {
		repId := int32(feedbacks[i].Foodid)
		repId2 := int32(feedbacks[i].ID)
		reply.List = append(reply.List, &protocol.Feedback{
			FoodId:  &repId,
			Uid:     &feedbacks[i].Uid,
			Rating:  &feedbacks[i].Rating,
			Comment: &feedbacks[i].Comment,
			Id:      &repId2,
		})
	}
	return &reply, nil
}

func (m *GRPCManager) ReadFeedbackByCommentId(ctx context.Context, commentId *protocol.CommentId) (*protocol.FeedbackList, error) {
	var feedback model.Feedback
	if res := m.db.Where("ID = ?", commentId.Id).Find(&feedback); res.Error != nil {
		return nil, res.Error
	}
	var reply protocol.FeedbackList
	repId := int32(feedback.Foodid)
	repId2 := int32(feedback.ID)
	reply.List = append(reply.List, &protocol.Feedback{
		FoodId:  &repId,
		Uid:     &feedback.Uid,
		Rating:  &feedback.Rating,
		Comment: &feedback.Comment,
		Id:      &repId2,
	})
	return &reply, nil
}

func (m *GRPCManager) UpdateFeedback(ctx context.Context, fb *protocol.UpdateRequest) (*common.Status, error) {
	var feedback model.Feedback
	if res := m.db.Where("ID = ?", fb.Id.Id).Find(&feedback); res.Error != nil {
		return nil, res.Error
	}
	feedback.Comment = *fb.Fb.Comment
	feedback.Rating = *fb.Fb.Rating
	if res := m.db.Save(&feedback); res.Error != nil {
		return nil, res.Error
	}
	code := int32(0)
	return &common.Status{
		Code: &code,
	}, nil
}
func (m *GRPCManager) DeleteFeedback(ctx context.Context, commentId *protocol.CommentId) (*common.Status, error) {
	var feedback model.Feedback
	if res := m.db.Where("ID = ?", commentId.Id).Find(&feedback); res.Error != nil {
		return nil, res.Error
	}

	if res := m.db.Delete(&feedback); res.Error != nil {
		return nil, res.Error
	}
	code := int32(0)
	return &common.Status{
		Code: &code,
	}, nil
}
