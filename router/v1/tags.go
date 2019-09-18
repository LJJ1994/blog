package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"test2/models"
	"test2/pkg/e"
	"test2/pkg/setting"
	"test2/pkg/util"
)

func GetTags(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	name := c.Query("name")
	if name != "" {
		maps["name"] = name
	}

	var state = -1
	if args := c.Query("state");args!="" {
		state= com.StrTo(args).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagsTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg": e.GetMsg(code),
	})
}

func AddTags(c *gin.Context) {
	name := c.PostForm("name")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()
	createdBy := c.PostForm("created_by")

	fmt.Printf("name:%s state:%d createBy: %s", name, state, createdBy)

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if ! models.ExistsTagByName(name) {
			code = e.SUCCESS
			models.AddTags(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

func UpdateTags(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()

	name := c.PostForm("name")
	modifiedBy := c.PostForm("modified_by")

	valid := validation.Validation{}
	var state = -1
	if arg := c.PostForm("state");arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("id不能为空")
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy

			if state != -1 {
				data["state"] = state
			}

			if name != "" {
				data["name"] = name
			}

			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}

func DeleteTags(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if id < 0 {
		fmt.Println("Id 不存在!!")
	}
	if models.ExistById(id) {
		models.DeleteTag(id)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg": "操作成功",
			"data": make(map[string]interface{}),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg": "",
			"data": make(map[string]interface{}),
		})
	}
}
