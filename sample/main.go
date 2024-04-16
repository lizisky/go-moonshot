package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lizisky/go-moonshot/kimi"
)

const (
	KIMI_API_URL            = "https://api.moonshot.cn/v1"
	KIMI_API_KEY            = "YOUR_API_KEY"
	YOUR_FILE_TO_BE_UPLODAD = "moonshot.pdf"
)

// 估算 token 数量
func EstimateTokenCount() {
	ctx := context.Background()

	client := kimi.NewClient[kimi.Moonshot](kimi.Moonshot{
		URL:    KIMI_API_URL,
		KEY:    KIMI_API_KEY,
		CLIENT: http.DefaultClient,
		// LOG: func(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration) {
		// 	log.Printf("[%s] %s %s", caller, request.URL, elapse)
		// },
	})

	estimateTokenCount, err := client.EstimateTokenCount(ctx, &kimi.EstimateTokenCountRequest{
		Messages: []*kimi.Message{
			// {
			// 	Role:    kimi.RoleSystem,
			// 	Content: &kimi.Content{Text: "你是 Kimi，由 Moonshot AI 提供的人工智能助手，你更擅长中文和英文的对话。你会为用户提供安全，有帮助，准确的回答。同时，你会拒绝一切涉及恐怖主义，种族歧视，黄色暴力等问题的回答。Moonshot AI 为专有名词，不可翻译成其他语言。"},
			// },
			{
				Role:    kimi.RoleUser,
				Content: &kimi.Content{Text: "你好，我是粒子，1+1等于多少？"},
			},
		},
		Model:       kimi.ModelMoonshot8K,
		MaxTokens:   4096,
		N:           1,
		Temperature: "0.3",
	})

	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}

	log.Printf("total_tokens=%d", estimateTokenCount.Data.TotalTokens)
}

// 获取余额
func getBalance() {
	ctx := context.Background()

	client := kimi.NewClient[kimi.Moonshot](kimi.Moonshot{
		URL:    KIMI_API_URL,
		KEY:    KIMI_API_KEY,
		CLIENT: http.DefaultClient,
		// LOG: func(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration) {
		// 	log.Printf("[%s] %s %s", caller, request.URL, elapse)
		// },
	})

	balance, err := client.CheckBalance(ctx)
	if err == nil {
		log.Printf("balance=%s", balance.Data.AvailableBalance)
	}
}

// 聊天
func chat() {
	ctx := context.Background()

	client := kimi.NewClient[kimi.Moonshot](kimi.Moonshot{
		URL:    KIMI_API_URL,
		KEY:    KIMI_API_KEY,
		CLIENT: http.DefaultClient,
		// LOG: func(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration) {
		// 	log.Printf("[%s] %s %s", caller, request.URL, elapse)
		// },
	})

	completion, err := client.CreateChatCompletion(ctx, &kimi.ChatCompletionRequest{
		Messages: []*kimi.Message{
			{
				Role:    kimi.RoleSystem,
				Content: &kimi.Content{Text: "你是 Kimi，由 Moonshot AI 提供的人工智能助手，你更擅长中文和英文的对话。你会为用户提供安全，有帮助，准确的回答。同时，你会拒绝一切涉及恐怖主义，种族歧视，黄色暴力等问题的回答。Moonshot AI 为专有名词，不可翻译成其他语言。"},
			},
			{
				Role:    kimi.RoleUser,
				Content: &kimi.Content{Text: "你好，我是粒子，1+1等于多少？"},
			},
		},
		Model:       kimi.ModelMoonshot8K,
		MaxTokens:   4096,
		N:           1,
		Temperature: "0.3",
	})

	if err == nil {
		fmt.Println(completion.GetMessageContent())
	}
}

// 聊天，流式返回
func chat_stream_one() {
	ctx := context.Background()

	client := kimi.NewClient[kimi.Moonshot](kimi.Moonshot{
		URL:    KIMI_API_URL,
		KEY:    KIMI_API_KEY,
		CLIENT: http.DefaultClient,
		// LOG: func(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration) {
		// 	log.Printf("[%s] %s %s", caller, request.URL, elapse)
		// },
	})

	stream, err := client.CreateChatCompletionStream(ctx, &kimi.ChatCompletionStreamRequest{
		Messages: []*kimi.Message{
			// {
			// 	Role:    kimi.RoleSystem,
			// 	Content: &kimi.Content{Text: "你是 Kimi，由 Moonshot AI 提供的人工智能助手，你更擅长中文和英文的对话。你会为用户提供安全，有帮助，准确的回答。同时，你会拒绝一切涉及恐怖主义，种族歧视，黄色暴力等问题的回答。Moonshot AI 为专有名词，不可翻译成其他语言。"},
			// },
			{
				Role:    kimi.RoleUser,
				Content: &kimi.Content{Text: "扬沙天气，带什么口罩最合适"},
			},
		},
		Model:       kimi.ModelMoonshot8K,
		MaxTokens:   4096,
		N:           1,
		Temperature: "0.3",
	})

	if err != nil {
		return
	}

	defer stream.Close()
	for chunk := range stream.C {
		fmt.Printf("%s", chunk.GetDeltaContent())
	}
	fmt.Println("finish ")
}

// 聊天，流式返回
func chat_stream_two() {
	ctx := context.Background()

	client := kimi.NewClient[kimi.Moonshot](kimi.Moonshot{
		URL:    KIMI_API_URL,
		KEY:    KIMI_API_KEY,
		CLIENT: http.DefaultClient,
		// LOG: func(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration) {
		// 	log.Printf("[%s] %s %s", caller, request.URL, elapse)
		// },
	})

	stream, err := client.CreateChatCompletionStream(ctx, &kimi.ChatCompletionStreamRequest{
		Messages: []*kimi.Message{
			// {
			// 	Role:    kimi.RoleSystem,
			// 	Content: &kimi.Content{Text: "你是 Kimi，由 Moonshot AI 提供的人工智能助手，你更擅长中文和英文的对话。你会为用户提供安全，有帮助，准确的回答。同时，你会拒绝一切涉及恐怖主义，种族歧视，黄色暴力等问题的回答。Moonshot AI 为专有名词，不可翻译成其他语言。"},
			// },
			{
				Role:    kimi.RoleUser,
				Content: &kimi.Content{Text: "如何快速、准确的对一个细分行业进行分析、总结"},
			},
		},
		Model:       kimi.ModelMoonshot8K,
		MaxTokens:   4096,
		N:           1,
		Temperature: "0.3",
	})

	if err != nil {
		fmt.Println("----------------- 222", err)
		return
	}

	defer stream.Close()
	message := stream.CollectMessage()
	fmt.Println(message.Content.Text)
}

// 获取文件内容
func retrieveFileContent() {
	ctx := context.Background()

	client := kimi.NewClient[kimi.Moonshot](kimi.Moonshot{
		URL:    KIMI_API_URL,
		KEY:    KIMI_API_KEY,
		CLIENT: http.DefaultClient,
		// LOG: func(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration) {
		// 	log.Printf("[%s] %s %s", caller, request.URL, elapse)
		// },
	})

	pdf, err := os.Open(YOUR_FILE_TO_BE_UPLODAD)
	if err != nil {
		fmt.Println("----------------- 222", err)
		return
	}

	defer pdf.Close()

	file, err := client.UploadFile(ctx, &kimi.UploadFileRequest{
		File:    pdf,
		Purpose: "file-extract",
	})

	if err != nil {
		return
	}

	log.Printf("file_id=%q; status=%s", file.ID, file.Status)

	content, err := client.RetrieveFileContent(ctx, file.ID)
	if err != nil {
		return
	}

	fmt.Println(string(content))

}

func main() {
	// EstimateTokenCount()
	// EstimateTokenCount()
	// getBalance()
	chat_stream_one()
	// chat_stream_two()
	// retrieveFileContent()
}
