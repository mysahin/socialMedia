package Models

type Follow struct {
	FollowedUserName string `json:"followed_user_name"`
	FollowerUserName string `json:"follower_user_name"`
	IsFollow         bool   `json:"is_follow"`
}
