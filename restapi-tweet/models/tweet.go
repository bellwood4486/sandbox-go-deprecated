package models

type Tweet struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	UserName   string `json:"user_name"`
	CommentNum int    `json:"comment_num"`
	StarNum    int    `json:"star_num"`
	ReTweetNum int    `json:"re_tweet_num"`
}
