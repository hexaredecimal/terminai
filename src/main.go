package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
  "github.com/akamensky/argparse"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)


func Process(ctx context.Context, cs *genai.ChatSession, input string) {
  resp, err := cs.SendMessage(ctx, genai.Text(input))

  if err != nil {
    fmt.Printf("Error: %s", err.Error())
    return
  }

  candidates := resp.Candidates[0]
  for i := range candidates.Content.Parts {
    part := candidates.Content.Parts[i]
    str := strings.ReplaceAll(fmt.Sprintf("%s", part), "\t", " ")
    fmt.Printf("[🤖] %s\n", str)
  }
}

func main() {
  parser := argparse.NewParser("termini", "AI assistant on the terminal using Gemini AI")
  showHelp := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Print version info"})

  err := parser.Parse(os.Args)
	if err != nil {
    return
  }

  if *showHelp {
    fmt.Printf("Termini AI - version 0.1\nBy Gama Sibusiso\nUsing Gemini AI\n")
    return
  }

  buffer := bufio.NewReader(os.Stdin)
  ctx := context.Background()
  client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
  if err != nil {
    log.Fatal(err)
  }
  defer client.Close()

  model := client.GenerativeModel("gemini-pro")
  cs := model.StartChat()

  for {
    fmt.Printf("[⚡️] ")
    input, err := buffer.ReadString('\n')

    if err != nil {
      break
    }

    if len(input) > 0 {
      Process(ctx, cs, input)
    }
  }
}

