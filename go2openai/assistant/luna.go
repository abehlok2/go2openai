package assistant

import (
	"context"
	"fmt"
	"github.com/abehlok2/go2openai/pkg/person"
	"github.com/sashabaranov/go-openai"
	"os"
)

var ctx = context.Background()

var client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))

var KatyString string = person.BuildPersonString(person.Katy)

var name string = "Luna"
var description string = "Personal assistant tailored to an individual user, Katy Behlok"
var instructions string = ""
var LunaRequest openai.AssistantRequest = openai.AssistantRequest{
	Model:        "gpt-4-1106-preview",
	Name:         &name,
	Description:  &description,
	Instructions: &instructions,
	Tools:        nil,
	FileIDs:      nil,
	Metadata:     nil,
}

func BuildLuna(context.Context, *openai.Client, openai.AssistantRequest) openai.Assistant {
	var Luna, err = client.CreateAssistant(ctx, request)
	if err != nil {
		panic("Something went wrong")
	}
	return Luna
}

func GetLuna(ctx context.Context, client openai.Client, LunaID string) openai.Assistant {
	var Luna, err = client.RetrieveAssistant(ctx, LunaID)
	if err != nil {
		fmt.Println("Something went wrong with the Luna get.")
	}
	return Luna
}
