package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
)

// GetPostsByCategoryId 根据文章的分类id展示文章列表
func GetPostsByCategoryId(cid, page, pageSize int) (*models.CategoryResponse, error) {
	// 类似于service/index.go, 我们可以拷贝进来然后稍作改动
	categorys, err := dao.GetAllCategory()
	if err != nil {
		// 当返回nil的时候需要返回*models.CategoryResponse而不是models.CategoryResponse否则会报错
		return nil, err
	}
	// 分页查询
	posts, err := dao.GetPostPageByCategoryId(cid, page, pageSize)
	var postsMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postsMores = append(postsMores, postMore)
	}
	total := dao.CountGetAllPost()
	// 向上取整
	pageCount := (total + pageSize) / pageSize
	var pages = []int{}
	for i := 1; i <= pageCount; i++ {
		pages = append(pages, i)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postsMores,
		total,
		page,
		pages,
		page != pageCount,
	}
	categoryName := dao.GetCategoryNameById(cid)
	categoryResponse := &models.CategoryResponse{
		hr,
		categoryName,
	}
	return categoryResponse, nil
}
