package main

import (
	"fmt"
	"gee"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "//static")

	stu1 := &student{
		Name: "Gee",
		Age:  20,
	}

	stu2 := &student{
		Name: "Jack",
		Age:  22,
	}
	r.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(context *gee.Context) {
		context.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(context *gee.Context) {
		context.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2023, 8, 19, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")

}
