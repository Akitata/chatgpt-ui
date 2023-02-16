package chat

import (
	"context"
	gogpt "github.com/sashabaranov/go-gpt3"
	"log"
)

var GPTClient *gogpt.Client

func InitChatClient(token string) {
	if len(token) > 0 {
		GPTClient = gogpt.NewClient(token)
		models, err := GPTClient.ListModels(context.Background())
		if err != nil {
			log.Printf("get gpt models err. %s \n", err.Error())
		}
		if len(models.Models) > 0 {
			log.Println("init gpt client ok.")
		} else {
			panic("init gpt client err, please check token or network.")
		}
	} else {
		panic("token is empty.")
	}
}
