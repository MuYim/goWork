package vo

//binding:"required" 表示Name不能为空
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
