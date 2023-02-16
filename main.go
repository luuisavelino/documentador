package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
)

type OAIChoices struct {
	Text         string
	Index        uint8
	Logprobs     uint8
	FinishReason string
}

type OAIResponse struct {
	Id      string
	Object  string
	Create  uint64
	Model   string
	Choices []OAIChoices
}

type OAIRequest struct {
	Prompt     string `json:"prompt"`
	Max_tokens uint32 `json:"max_tokens"`
}

var dir, file string

func main() {
	fmt.Println("\x1bc")
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	var userInput string

	for {
		fmt.Println("Escolha o arquivo a ser analisado:")
		fmt.Print("> ")

		fmt.Scanln(&userInput)

		dir, file = filepath.Split(userInput)
		fmt.Println(dir + file)

		s.Start()
		s.Suffix = " A OpenAI está realizando a documentação do seu código..."

		data, err := ioutil.ReadFile(dir + file)
		if err != nil {
			fmt.Println("Erro ao ler o arquivo:", err.Error())
			return
		}

		requestOpenAI(string(data))
		s.Stop()
	}
}

func requestOpenAI(docRequest string) {
	oaiToken := os.Getenv("OPENAI_KEY")
	bearer := "Bearer " + oaiToken
	preanble := `Desenvolva uma documentação em markdown (.md) do seguinte código. Utilize o README.md do projeto do kubernetes como base. Utilize também os subtópicos (##) a mais: Requerimentos; Instrução de uso; descrição do código.`
	uri := "https://api.openai.com/v1/engines/text-davinci-002/completions"

	OAIRequest := OAIRequest{
		Prompt:     fmt.Sprintf("%s %s", preanble, docRequest),
		Max_tokens: 500,
	}

	var payload bytes.Buffer
	err := json.NewEncoder(&payload).Encode(OAIRequest)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, uri, &payload)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response OAIResponse
	err = json.Unmarshal([]byte(bytes), &response)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(dir + "README.md")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(response.Choices[0].Text)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("")
}
