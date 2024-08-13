package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func main() {
	// AWSセッションの作成
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("http://localstack:4566"), // LocalStackのエンドポイント
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", ""), // ダミーの認証情報
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	// Secrets Managerクライアントの作成
	svc := secretsmanager.New(sess)

	// シークレットの取得
	secretName := "MySecret"
	result, err := svc.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		log.Fatalf("Failed to retrieve secret: %v", err)
	}

	// シークレットの内容を表示
	fmt.Printf("Secret: %s\n", *result.SecretString)
}

// // シークレットのJSONデータをパース
// var secretData map[string]string
// err = json.Unmarshal([]byte(*result.SecretString), &secretData)
// if err != nil {
// 	log.Fatalf("Failed to unmarshal secret string: %v", err)
// }

// // usernameフィールドを取得
// username := secretData["username"]
// fmt.Printf("Username: %s\n", username)