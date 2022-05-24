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
