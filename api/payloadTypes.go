package api

// payload types
type PayloadSignup struct {
  Name string `json:"name" binding:"required"`
  ImgURL string `json:"imgUrl"`
  Password string `json:"password" binding:"require"`
}

type PayloadLogin struct{
  Name string `json:"name" binding:"required"`
  Password string `json:"password" binding:"require"`
}
