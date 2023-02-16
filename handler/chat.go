package handler

import (
	"github.com/akitata/chatgpt-ui/chat"
	"github.com/gofiber/fiber/v2"
	gogpt "github.com/sashabaranov/go-gpt3"
	"net/http"
)

func Chat(ctx *fiber.Ctx) error {
	session, err := Sessions.Get(ctx)
	if err != nil {
		panic("get session err.")
	}

	uid := session.Get("uid")
	uidStr := uid.(string)
	prompt := ctx.Query("prompt")

	if len(prompt) <= 0 {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"type": "text",
			"content": fiber.Map{
				"code": -1,
				"text": "params: prompt is empty.",
			},
		})
	}

	response, err := chat.GPTClient.CreateCompletion(ctx.UserContext(), gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		Prompt:    prompt,
		MaxTokens: 2000,
		N:         1,
		User:      uidStr,
	})

	if err != nil {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"type": "text",
			"content": fiber.Map{
				"code": -2,
				"text": "error: " + err.Error(),
			},
		})
	}

	answer := response.Choices[0].Text
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"type": "text",
		"content": fiber.Map{
			"code": 0,
			"text": answer,
		},
	})
}
