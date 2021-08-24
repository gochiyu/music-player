package vo

//表单验证
type CreatePostRequst struct {
	CategoryId uint   `json:"category_id" bining:"required"`
	Title      string `json:"title" bining:"required,max=10"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" bining:"required"`
}
