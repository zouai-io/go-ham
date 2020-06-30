package main

import (
	"context"
	"gopkg.zouai.io/ham"
	"log"
)

func main() {
	userInput := int64(15)
	ctx := ham.WithContext(context.Background())
	exponent := ham.Int64(ctx, 3)
	input := ham.Int64(ctx, userInput)
	result := input.Pow(ctx, exponent)
	exp2 := ham.Int64(ctx, 2)
	res2 := result.Pow(ctx, exp2)
	_ = result.Pow(ctx, exponent)
	log.Print(res2)
	ham.Print(ctx)
	ham.PrintSingleLine(ctx)
}
