package titan

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type Request struct {
	InputText string `json:"inputText"`
}

type Response struct {
	Embedding           []float64 `json:"embedding"`
	InputTextTokenCount int       `json:"inputTextTokenCount"`
}

const defaultRegion = "eu-central-1"

var Client *bedrockruntime.Client

func init() {

	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = defaultRegion
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	Client = bedrockruntime.NewFromConfig(cfg)
}

const (
	titanEmbeddingModelID = "amazon.titan-embed-text-v1" //https://docs.aws.amazon.com/bedrock/latest/userguide/model-ids-arns.html
)

func FetchEmbedding(input string) ([]float32, error) {
	// See https://github.com/build-on-aws/amazon-bedrock-go-sdk-examples/blob/main/titan-text-embedding/main.go

	payload := Request{
		InputText: input,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	output, err := Client.InvokeModel(context.Background(), &bedrockruntime.InvokeModelInput{
		Body:        payloadBytes,
		ModelId:     aws.String(titanEmbeddingModelID),
		ContentType: aws.String("application/json"),
	})

	if err != nil {
		log.Fatal("failed to invoke model: ", err)
	}

	var resp Response

	err = json.Unmarshal(output.Body, &resp)

	if err != nil {
		log.Fatal("failed to unmarshal", err)
	}

	//fmt.Println("embedding vector from LLM\n", resp.Embedding)

	// fmt.Println("generated embedding for input -", input)
	fmt.Println("generated vector length -", len(resp.Embedding))

	var embeddings []float32
	for _, item := range resp.Embedding {

		embeddings = append(embeddings, float32(item))
	}
	return embeddings, nil
}
