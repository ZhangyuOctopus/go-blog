package dao

import (
	"go-blog/models"
	"log"
)

// GetPostPage 处理文章列表的相关数据库操作, 查询文章的分页列表, 并且需要返回错误
func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize

	// 分页查询
	rows, err := DB.Query("select * from blog_post limit ?, ?", page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		// 将查询到的数据追加到posts列表中, 返回的数据就为查询到的文章列表
		posts = append(posts, post)
	}
	return posts, nil
}

// CountGetAllPost 查询博客的总数
func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post")
	_ = rows.Scan(&count)
	return
}

// CountGetAllPostBySlug 根据slug限制查询文章的总数
func CountGetAllPostBySlug(slug string) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where slug = ?", slug)
	_ = rows.Scan(&count)
	return
}

// GetPostPageByCategoryId 根据文章分类ID查分文章分类列表
func GetPostPageByCategoryId(cid, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	// 分页查询, 与获取文章列表的逻辑是一模一样的, 知识这里多了一个category_id的限制
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?, ?", cid, page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		// 将查询到的数据追加到posts列表中, 返回的数据就为查询到的文章列表
		posts = append(posts, post)
	}
	return posts, nil
}

// GetPostById 根据pid查询文章的详情信息
func GetPostById(pid int) (models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid = ?", pid)
	var post models.Post
	if row.Err() != nil {
		return post, row.Err()
	}
	err := row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt)
	if err != nil {
		return post, row.Err()
	}
	return post, nil
}

// GetPostPageBySlug 根据slug标签查询所有的文章信息
func GetPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	// 分页查询, 与获取文章列表的逻辑是一模一样的, 知识这里多了一个category_id的限制
	rows, err := DB.Query("select * from blog_post where slug = ? limit ?, ?", slug, page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		// 将查询到的数据追加到posts列表中, 返回的数据就为查询到的文章列表
		posts = append(posts, post)
	}
	return posts, nil
}

// GetPostAll 查询所有的文章
func GetPostAll() ([]models.Post, error) {
	var posts []models.Post
	rows, err := DB.Query("select * from blog_post")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		// 将查询到的数据追加到posts列表中, 返回的数据就为查询到的文章列表
		posts = append(posts, post)
	}
	return posts, nil
}

func SavePost(post *models.Post) {
	res, err := DB.Exec("insert into blog_post (title, content, markdown, category_id, "+
		"user_id, view_count, type, "+
		"slug, create_at, update_at) "+
		"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", post.Title, post.Content,
		post.Markdown, post.CategoryId, post.UserId, post.ViewCount,
		post.Type, post.Slug, post.CreateAt, post.UpdateAt)
	if err != nil {
		log.Println(err)
	}
	pid, _ := res.LastInsertId()
	post.Pid = int(pid)
}

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post set title = ?, content = ?, markdown = ?, category_id = ?, type = ?, "+
		"slug = ?, update_at = ? where pid = ?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid)
	if err != nil {
		log.Println("更新出错: ", err)
	}
}

// GetPostSearch 根据搜索框的内容查询所有文章信息
func GetPostSearch(condition string) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post where title like ? ", "%"+condition+"%")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		// 将查询到的数据追加到posts列表中, 返回的数据就为查询到的文章列表
		posts = append(posts, post)
	}
	return posts, nil
}
