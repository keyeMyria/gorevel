package controllers

import (
	"strconv"

	"github.com/revel/revel"

	"gorevel/app/models"
)

var (
	SUCCESS_JSON = map[string]bool{"status": true}
	ERROR_JSON   = map[string]bool{"status": false}
)

func init() {
	revel.OnAppStart(Init)
	revel.InterceptMethod((*Base).checkUser, revel.BEFORE)

	revel.TemplateFuncs["eqis"] = func(i int64, s string) bool {
		return strconv.FormatInt(i, 10) == s
	}
}

func Init() {
	engine = models.GetEngine()
	UPLOAD_PATH = revel.BasePath + "/public/upload/"
}
