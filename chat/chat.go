package chat

import (
	"context"
	gogpt "github.com/sashabaranov/go-gpt3"
	"log"
	"net/http"
	"net/url"
)

var GPTClient *gogpt.Client

func InitChatClient(token, proxyUrl string) {
	if len(token) > 0 {
		clientConfig := gogpt.DefaultConfig(token)
		if len(proxyUrl) > 0 {
			proxyURL, err := url.Parse(proxyUrl)
			if err != nil {
				panic("parse proxy url err.")
			}
			httpClient := http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyURL(proxyURL),
				},
			}
			clientConfig.HTTPClient = &httpClient
		}
		GPTClient = gogpt.NewClientWithConfig(clientConfig)

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
