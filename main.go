/*
Copyright © 2024 Vladyslav Papidokha <vladyslavpapidokha@gmail.com>
*/
package main

import (
	"context"
	"log"

	"github.com/vpapidokha/email-validator/internal/delivery/cli"
)

func main() {
	if err := cli.Execute(context.Background()); err != nil {
		log.Fatal(err)
	}
}
