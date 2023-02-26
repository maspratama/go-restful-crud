package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/maspratama/go-restful-crud/helper"
	"github.com/maspratama/go-restful-crud/models/web"
	"github.com/maspratama/go-restful-crud/services"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	// dari controller panggil service Create
	categoryResponse := c.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		ResponseCode:    200,
		ResponseMessage: "OK",
		Data:            categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := c.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		ResponseCode:    200,
		ResponseMessage: "OK",
		Data:            categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	c.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		ResponseCode:    200,
		ResponseMessage: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := c.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		ResponseCode:    200,
		ResponseMessage: "OK",
		Data:            categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponse := c.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		ResponseCode:    200,
		ResponseMessage: "OK",
		Data:            categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
