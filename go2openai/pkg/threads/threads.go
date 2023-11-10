package threads

import (
	"context"
	openai "github.com/sashabaranov/go-openai"

	"os"
)

var ctx = context.Background()

var client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))

var threadRequest openai.ThreadRequest = openai.ThreadRequest{
	Messages: messages,
}

// Empty array of thread messages
var messages []openai.ThreadMessage
var SystemMessageString string = "**STARTING LUNA...**\n\n**ACTIVATED**\n\n**TEACHER ASSIST MODE: ACTIVATED"
var ThreadSystemMessage openai.ThreadMessage = openai.ThreadMessage{
	Role:     openai.ChatMessageRoleSystem,
	Content:  systemMessageString,
	FileIDs:  nil,
	Metadata: nil,
}

// Adds the system message to the thread before beginning chat messages
var messageSysMsg = append(messages, ThreadSystemMessage)

var Thread, err = client.CreateThread(
	ctx,
	threadRequest,
)

func StartLunaRun(ctx context.Context, client openai.Client, thread openai.Thread) openai.Run {

	var model = "gpt-4-1106-preview"
	var RunRequest = openai.RunRequest{
		AssistantID:  "",
		Model:        &model,
		Instructions: &systemMessageString,
		Tools:        nil,
		Metadata:     nil,
	}

	var Run, err = client.CreateRun(ctx, Thread.ID, RunRequest)
	if err != nil {
		panic(err)
	}
	return Run
}
