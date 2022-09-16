package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
)

// GetAllIndexInfo 需要处理一下具体的错误, 所以返回值为两个, 第二个属于返回的错误, 因为涉及到分页所以在下面的方法中传递两个参数
func GetAllIndexInfo(slug string, page, pageSize int) (*models.HomeResponse, error) {
	// 页面上涉及到的所有模板数据必须有定义
	//页面上涉及到的所有的数据，必须有定义
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	// 因为post和post结构体中有几个字段是不一样的所以需要处理一下(主要是时间还有多了几个字段)
	var (
		posts []models.Post
		total int
	)
	if slug == "" {
		// 为空说明查询页面的所有信息
		posts, err = dao.GetPostPage(page, pageSize)
		total = dao.CountGetAllPost()
	} else {
		posts, err = dao.GetPostPageBySlug(slug, page, pageSize)
		total = dao.CountGetAllPostBySlug(slug)
	}

	var postsMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		// 这里转为rune类型这样可以按照中文字节来判断
		content := []rune(post.Content)
		if len(content) > 100 {
			// 当长度大于100的时候先切割一下, 这样显示的时候太长了显示的时候可以短一点
			content = content[:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			// 显示的Content需要是html格式所以需要转换一下
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
	return hr, nil
}
