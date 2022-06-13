package apiDT

import "time"

type ResponseUser struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	ImgURL string `json:"imgUrl"`
}

type ResponseUserAndBio struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	ImgURL string `json:"imgUrl"`
	Bio    string `json:"bio"`
}

type ResponseLogin struct {
	User ResponseUser `json:"user"`
}
type ResponseSignup struct {
	User ResponseUser `json:"user"`
}

type ResponseShortPost struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"desc"`
	CreatedAt    time.Time `json:"createdAt"`
	Slug         string    `json:"slug"`
	FeedImg      string    `json:"feedImg"`
	AuthorName   string    `json:"authorName"`
	AuthorImgURL string    `json:"authorImg"`
}

type ResponseCompletePost struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Description  string    `json:"desc"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	Published    bool      `json:"published"`
	Slug         string    `json:"slug"`
	FeedImg      string    `json:"feedImg"`
	ArticleImg   string    `json:"articleImg"`
	AuthorName   string    `json:"authorName"`
	AuthorImgURL string    `json:"authorImg"`
}

type ResponseCreatePost struct {
	ID int `json:"id"`
}
