package main

import (
	"fmt"
	"os"
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
  	"github.com/tmc/langchaingo/llms/openai"
)


//your cool!

const OpenAIKey = ""
const BedrockKey = ""

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle the different types of messages and return the updated model and any commands you want to run.
	fmt.Println("hello"
	return m, nil
}
const GPT4Turbo = "gpt4-turbo"
const Titan = "Titan-1"


func main() {

	prompt := "What would be a good company name for a company that makes colorful socks?"

	completion := GenerateWithOpenAI(prompt, GPT4Turbo)
	
	fmt.Println(completion)

}



func GenerateWithBedrock(prompt string, model string) string {
	var input string
	return input
}

func GenerateWithOpenAI(prompt string, model string) string {	
	
	ctx := context.Background()
  	llm, err := openai.New()

  	if err != nil {
    		log.Fatal(err)
  	}
  	
  	
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
  	
	if err != nil {
    		log.Fatal(err)
  	}
  	

	return completion

}

func GetMechanicLogs(){

}
