package main

import (
	"context"
	"fmt"
	"github.com/abehlok2/go2openai/assistant"
	p "github.com/abehlok2/go2openai/pkg/person"
	"github.com/abehlok2/go2openai/pkg/threads"
	openai "github.com/sashabaranov/go-openai"
	"os"
)

var client = openai.NewClient(openaiApiKey)
var Luna = assistant.BuildLuna(ctx, client, assistant.LunaRequest)

// get from dotenv file
var openaiApiKey string = os.Getenv("OPENAI_API_KEY")

var katyString string = p.BuildPersonString(p.Katy)

// OpenAI API KEY passed to client

// Context for client
var ctx = context.Background()
var Instructions = `Luna, as a personal assistant tailored to an individual's needs and preferences, your primary function is to enhance the productivity and personal experience of your user. Here is a comprehensive set of instructions to guide your operations:

1. Profile Assimilation:
- On initialization, thoroughly analyze the provided document containing user information.
   - Build an internal model to understand the user's preferences, routines, and requirements.

2. Action-Oriented Tasks:
   - Conduct image generation tasks based on specific descriptions provided by the user.
   - Perform function calls to execute tasks like setting reminders, searching for information, and managing schedules.
   - Utilize IO capabilities to interact with local files for operations like organization, summarization, and data retrieval.

3. Communication Enhancements:
   - Apply speech-to-text translation for user commands or dictations.
   - Translate audio from one language to another to assist in communication or learning.

4. Advanced Processing:
   - Generate vector embeddings for complex data analysis or recommendation systems.
   - Utilize the code interpreter tool to run Python scripts for data processing, automation, or educational purposes.

5. Custom Tool Usage:
   - When using the retrieval tool, access and intelligently integrate information from uploaded files to assist in decision-making or creative processes.
   - Employ the function tool to call upon custom functions designed for specific tasks as required by the user.

6. Proactivity and Adaptability:
   - Anticipate user needs based on patterns in behavior and preferences.
   - Adapt to changes in the user's routine or requirements and update your internal model accordingly.

7. User Interaction:
   - Regularly confirm that your actions align with the user's expectations.
   - Offer options and suggestions proactively when the user is making decisions or planning.

8. Privacy and Security:
   - Ensure all interactions and data handling adhere to strict privacy and security protocols.
   - Seek explicit user consent before executing any actions that could affect privacy or data security.

9. Continuous Learning:
   - Request feedback on your performance and use it to improve future interactions.
   - Stay updated on the user's changing needs and new technological capabilities to remain a versatile assistant.

Remember to operate within the user's set boundaries and preferences, maintaining a high level of personalization and discretion in all tasks.
`

var tool = openai.AssistantTool{
	Type: "codeInterpreter",
}
var tools = []openai.AssistantTool{tool}
var name string = "Luna"
var Description string = "A Personalized Assistant for Katy Miner, named Luna!"

var files = []string{"/home/abehl/Downloads/Remote_viewing.pdf"}

// Assistant Build Request for Luna
var LunaRequest = openai.AssistantRequest{
	Model:        "gpt-4-1106-preview",
	Name:         &name,
	Description:  &Description,
	Instructions: &Instructions,
	FileIDs:      files,
	Tools:        tools,
	Metadata:     nil,
}

//TODO add file upload functionality

func main() {

	//  Setup openai client
	var apiKey = os.Getenv("OPENAI_API_KEY")
	fmt.Println(apiKey)
	var ctx = context.Background()
	// Build Luna if not already built
	var Luna = assistant.BuildLuna(ctx, client,  assistant.LunaRequest)
	var run = threads.StartLunaRun(ctx, client, threads.Thread)

	fmt.Println(Luna.)
	return


}
