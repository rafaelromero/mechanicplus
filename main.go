package main

import (

	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
  	"github.com/tmc/langchaingo/llms/openai"
)


//your cool!

const OpenAIKey = ""
const BedrockKey = ""

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
