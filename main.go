package main

import (
	"context"

	"gendemo/biz"
)

func main() {

	ctx := context.Background()

	_ = biz.Create(ctx)
}
