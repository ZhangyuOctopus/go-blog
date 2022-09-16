package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index     TemplateBlog
	Category  TemplateBlog
	Custom    TemplateBlog
	Detail    TemplateBlog
	Login     TemplateBlog
	PigOnHole TemplateBlog // 文章归档
	Writing   TemplateBlog
}

type IndexData struct {
	// 下面是`json:title`表示展示的时候按照小写来展示
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func isODD(num int) bool {
	return num%2 == 0
}

// 获取下一个路径
func getNextName(strs []string, index int) string {
	return strs[index+1]
}

// 获取对应的时间
func getDate(t string) string {
	return time.Now().Format(t)
}

// WriteData 空接口可以接受任意类型, 所以第二个数据类型定义为空接口比较方便
func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte("error"))
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

// WriteError 将错误数据写入到error中
func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}
}

// DateDay 方法名需要大写外部文件才可以获取
func DateDay(date time.Time) string {
	return date.Format("2006-01-10 21:53:34")
}

// 优化的第二个点, 遇到错误需要暴露出去所以返回值需要加上error
func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		// 访问首页博客模板的时候,因为已经有多个模板的嵌套,解析文件的时候, 需要将其涉及到的所有模板文件进行解析, 需要解析很多个
		_ = templateDir + "index.html"
		homePg := templateDir + "home.html"
		footerPg := templateDir + "layout/footer.html"
		headerPg := templateDir + "layout/header.html"
		personalPg := templateDir + "layout/personal.html"
		postPg := templateDir + "layout/post-list.html"
		paginationPg := templateDir + "layout/pagination.html"

		// 映射自定义的isOdd方法, 后面才可以解析
		t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": getNextName, "date": getDate, "dateDay": DateDay})

		// ParseFiles解析多个模板文件, 之前运行的时候出现t为nil的情况, 我们可以打印出具体的错误
		t, err := t.ParseFiles(templateDir+viewName, homePg, footerPg, headerPg, personalPg, postPg, paginationPg)
		if err != nil {
			log.Println("解析模板出错: ", err)
			return nil, err
		}
		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	return tbs, nil
}

// InitTemplate 传递一个模板路径
func InitTemplate(templateDir string) (HtmlTemplate, error) {
	tp, err := readTemplate([]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"}, templateDir)
	var htmlTemplate HtmlTemplate
	if err != nil {
		return htmlTemplate, err
	}
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.PigOnHole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate, nil
}
