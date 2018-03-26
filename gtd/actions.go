package main

import (
	"fmt"

	"gopkg.in/urfave/cli.v2"
)

func get(c *cli.Context) error {
	id := c.Int("id")
	if id == 0 {
		fmt.Println("Get All")
	} else {
		fmt.Printf("Get %d!\n", id)
	}
	return nil
}

func create(c *cli.Context) error {
	n := c.String("name")
	fmt.Printf("Hello, %s!\n", n)
	return nil
}

func update(c *cli.Context) error {
	id := c.Int("id")
	fmt.Printf("Hello, %d!\n", id)
	return nil
}

func delete(c *cli.Context) error {
	id := c.Int("id")
	fmt.Printf("Hello, %d!\n", id)
	return nil
}
