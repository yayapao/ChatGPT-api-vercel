package handler

import (
	"chat/utils"
	"context"
	"errors"
	"fmt"
	openapi "github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"os"
)

type ChatReq struct {
	Content string `json:"content" binding:"required"`
	Stream  bool   `json:"stream"`
}

func ProxyChatGPT(w http.ResponseWriter, r *http.Request) {
	client := openapi.NewClient(os.Getenv("OPENAPI_KEY"))
	var chatReq ChatReq
	err := utils.ExtractFromBody(r, &chatReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !chatReq.Stream {
		resp, err := client.CreateChatCompletion(
			context.Background(), openapi.ChatCompletionRequest{
				Model: openapi.GPT3Dot5Turbo,
				Messages: []openapi.ChatCompletionMessage{
					{
						Role:    openapi.ChatMessageRoleUser,
						Content: chatReq.Content,
					},
				},
			})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write([]byte(resp.Choices[0].Message.Content))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		fetchStreamOpenAPI(w, client, chatReq)
	}
}

func fetchStreamOpenAPI(w http.ResponseWriter, client *openapi.Client, chatReq ChatReq) {
	ctx := context.Background()
	req := openapi.ChatCompletionRequest{
		Model: openapi.GPT3Dot5Turbo,
		Messages: []openapi.ChatCompletionMessage{
			{
				Role:    openapi.ChatMessageRoleUser,
				Content: chatReq.Content,
			},
		},
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stream.Close()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}
		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}
		io.WriteString(w, response.Choices[0].Delta.Content)
		_, err = fmt.Fprintf(w, "event: message\ndata: %s\n", response.Choices[0].Delta.Content)
		if err != nil {
			return
		}
		w.(http.Flusher).Flush()
	}
}
