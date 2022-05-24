package apiDT

// payload types
type PayloadSignup struct {
  Name string `json:"name" binding:"required"`
  ImgURL string `json:"imgUrl"`
  Password string `json:"password" binding:"required"`
}

type PayloadLogin struct{
  Name string `json:"name" binding:"required"`
  Password string `json:"password" binding:"required"`
}

type PayloadPost struct{
  Title string `json:"title" binding:"required"`
  Description string `json:"desc" binding:"required"`
  Content string `json:"content" binding:"required"`
}
