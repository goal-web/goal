package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/models"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func TestModel(t *testing.T) {
	initApp()

	fmt.Println(strings.ToTitle("id"))

	models.Articles().Create(contracts.Fields{
		"title": "create-title",
	})

	article := models.Articles().FirstOrFail()

	article.Title = "new-title"
	err := article.Save()
	assert.NoError(t, err, err)

	err = article.Update(contracts.Fields{
		"title": "update-title",
	})
	assert.NoError(t, err, err)
	assert.True(t, article.Title == "update-title")

	models.Articles().Where("id", article.Id).Update(contracts.Fields{
		"title": "refresh-title",
	})
	err = article.Refresh()
	assert.NoError(t, err, err)
	assert.True(t, article.Title == "refresh-title")

	err = article.Delete()
	assert.NoError(t, err, err)
	assert.False(t, article.Exists(), "不存在")

	fmt.Println(article)
}

func TestReflect(t *testing.T) {
	article := models.Article{}

	value := reflect.ValueOf(&article)
	value.Elem().MethodByName("InitModel").Call([]reflect.Value{
		reflect.ValueOf(models.ArticleClass),
		reflect.ValueOf("articles"),
		reflect.ValueOf("id"),
		reflect.ValueOf(&article),
		reflect.ValueOf(reflect.ValueOf(article)),
	})
	value.Elem().FieldByName("Id").SetString("1")
	fmt.Println("success", article)
}
