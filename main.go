package main

import (
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
)

type ErrorItemResponse struct {
  Field string `json:"field"`
  Tag   string `json:"tag"`
}

type IErrorResponse interface {
  From(err validator.ValidationErrors) []ErrorItemResponse
}

type ErrorResponse struct {}

func (e ErrorResponse) From(err error) []ErrorItemResponse {
  var result []ErrorItemResponse

  for _, err := range err.(validator.ValidationErrors) {
    item := ErrorItemResponse{Field: err.Field(), Tag: err.Tag()}

    result = append(result, item)
  }

  return result
}

type ChatIdRequest struct {
  Id string `uri:"id" binding:"required,number"`
}

type ChatIdResponse struct {
  Id        int                     `json:"id"`
  Username  string                  `json:"username"`
  UserId    int                     `json:"user_id"`
  Messages  []ChatIdMessageResponse `json:"messages"`
}

type ChatIdMessageResponse struct {
  Id      int     `json:"id"`
  ChatId  int     `json:"chat_id"`
  UserId  int     `json:"user_id"`
  Text    string  `json:"text"`
}

type ChatResponse struct {
	Id          int     `json:"id"`
	Username    string  `json:"username"`
	Avatar      string  `json:"avatar"`
	LastMessage string  `json:"last_message"`
}

type BaseResponse struct {
  Data interface{} `json:"data"`
}

func main() {
	route := gin.Default()

	route.GET("/chats", func(c *gin.Context) {
    data := []ChatResponse{
      ChatResponse{
        Id:          1,
        Username:    "test1",
        Avatar:      "test1",
        LastMessage: "test1",
      },
      ChatResponse{
        Id:          2,
        Username:    "test2",
        Avatar:      "test2",
        LastMessage: "test2",
      },
    }

		c.JSON(200, BaseResponse{Data: data})
	})

	route.GET("/chats/:id", func(c *gin.Context) {
    var req ChatIdRequest

    if err := c.ShouldBindUri(&req); err != nil {
      c.JSON(200, BaseResponse{Data: ErrorResponse{}.From(err)})

      return
    }

	  data := ChatIdResponse{
      Id:        1,
      Username:  "test",
      UserId:    1,
      Messages:  []ChatIdMessageResponse{
        ChatIdMessageResponse{
          Id:      1,
          ChatId:  1,
          UserId:  1,
          Text:    "test",
        },
        ChatIdMessageResponse{
          Id:      2,
          ChatId:  1,
          UserId:  1,
          Text:    "test2",
        },
      },
	  }

		c.JSON(200, BaseResponse{Data: data})
	})

	route.Run()
}
