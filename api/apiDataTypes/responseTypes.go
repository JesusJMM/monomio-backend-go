package apiDT

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	ImgURL string `json:"imgUrl"`
}

type ResponseLogin struct {
	User User `json:"user"`
}
type ResponseSignup struct {
	User User `json:"user"`
}

type ResponseShortPost struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"desc"`
	CreatedAt    string `json:"createdAt"`
	AuthorName   string `json:"authorName"`
	AuthorImgURL string `json:"authorImg"`
}

type ResponseCompletePost struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Description  string `json:"desc"`
	CreatedAt    string `json:"createdAt"`
	AuthorName   string `json:"authorName"`
	AuthorImgURL string `json:"authorImg"`
}

type ResponseCreatePost struct {
  ID int `json:"id"`
}
