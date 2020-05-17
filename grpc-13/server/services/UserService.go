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

func (*UserService) GetUserScoreByServerStream(req *UserScoreRequest,
	stream UserService_GetUserScoreByServerStreamServer) error {
	var score int32 = 100
	users := make([]*UserInfo, 0)
	for index, user := range req.Users {
		user.UserScore = score
		score++
		users = append(users, user)
		if (index+1)%2 == 0 && index > 0 {
			err := stream.Send(&UserScoreResponse{Users: users})
			if err != nil {
				return err
			}
			users = users[0:0]
		}

	}
	// 发送最后一批
	if len(users) > 0 {
		err := stream.Send(&UserScoreResponse{Users: users})
		if err != nil {
			return err
		}
	}

	return nil
}
