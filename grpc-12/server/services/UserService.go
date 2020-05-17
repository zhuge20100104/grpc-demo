package services

import context "context"

type UserService struct{}

func (*UserService) GetUserScore(ctx context.Context, req *UserScoreRequest) (*UserScoreResponse, error) {
	var score int32 = 100
	users := make([]*UserInfo, 0)
	for _, user := range req.Users {
		user.UserScore = score
		score++
		users = append(users, user)
	}

	return &UserScoreResponse{Users: users}, nil
}
