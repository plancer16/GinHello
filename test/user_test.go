package test

import (
	"GinHello/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestUserSave(t *testing.T) {
	username := "lisi"
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
}

func TestUserSaveQuery(t *testing.T) {
	username := "lisi"
	age := 18
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+",年龄:"+strconv.Itoa(age)+"已经保存", w.Body.String())
}

func TestUserSaveWithNotAge(t *testing.T) {
	username := "lisi"
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+",年龄:24已经保存", w.Body.String())
}