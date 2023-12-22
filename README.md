# bedrockembedding
Use embeddings with Amazon Bedrock titan


## Usage example

Have some text:
```go
  content := "This is some text"
```

Get vectors from titan:
```go
  singleEmbedding, err := be.FetchEmbedding(content)
```

Store the vector in postgres:

```go
_, err = conn.Exec(ctx, "INSERT INTO documents (content,  embedding) VALUES ($1, $2)", content, pgvector.NewVector(singleEmbedding))
```
