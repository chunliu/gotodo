package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/urfave/cli.v2"
)

func printResult(todo Todo) {
	fmt.Println("Id: ", todo.ID)
	fmt.Println("Name: ", todo.Name)
	fmt.Println("IsCompleled: ", todo.IsCompleted)
}

func get(c *cli.Context) error {
	id := c.Int("id")
	url := baseURL.String()
	if id < 0 {
		return cli.Exit(fmt.Errorf("id is invalid"), 11)
	}

	if id > 0 {
		url = fmt.Sprintf("%s%d", baseURL, id)
	}

	resp, err := http.Get(url)
	if err != nil {
		return cli.Exit(fmt.Errorf("GET: %v", err), 12)
	}

	if resp.StatusCode != http.StatusOK {
		return cli.Exit(fmt.Errorf(resp.Status), 12)
	}

	if id == 0 {
		var todos TodoItems

		if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
			return cli.Exit(fmt.Errorf("JSON decode: %v", err), 13)
		}

		for _, item := range todos {
			printResult(item)
		}
	} else {
		var todo Todo

		if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
			return cli.Exit(fmt.Errorf("JSON decode: %v", err), 13)
		}

		printResult(todo)
	}

	return nil
}

func create(c *cli.Context) error {
	n := c.String("name")
	if n == "" {
		return cli.Exit(fmt.Errorf("name should not be empty"), 11)
	}

	todo := Todo{
		ID:          0,
		Key:         0,
		Name:        n,
		IsCompleted: false,
	}
	jsonValue, _ := json.Marshal(todo)
	resp, err := http.Post(baseURL.String(), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return cli.Exit(fmt.Errorf("POST: %v", err), 12)
	}

	if resp.StatusCode != http.StatusCreated {
		return cli.Exit(fmt.Errorf(resp.Status), 12)
	}

	if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
		return cli.Exit(fmt.Errorf("JSON decode: %v", err), 13)
	}

	printResult(todo)

	return nil
}

func update(c *cli.Context) error {
	id := c.Int("id")
	if id <= 0 {
		return cli.Exit(fmt.Errorf("id is invalid"), 11)
	}

	url := fmt.Sprintf("%s%d", baseURL, id)
	todo := Todo{
		ID:          id,
		Key:         id,
		Name:        c.String("name"),
		IsCompleted: c.Bool("completed"),
	}
	jsonValue, _ := json.Marshal(todo)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return cli.Exit(fmt.Errorf("PUT: %v", err), 12)
	}

	if resp.StatusCode != http.StatusNoContent {
		return cli.Exit(fmt.Errorf(resp.Status), 12)
	}

	fmt.Println(resp.Status)
	return nil
}

func delete(c *cli.Context) error {
	id := c.Int("id")
	if id <= 0 {
		return cli.Exit(fmt.Errorf("id is invalid"), 11)
	}

	url := fmt.Sprintf("%s%d", baseURL, id)
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return cli.Exit(fmt.Errorf("DELETE: %v", err), 12)
	}

	if resp.StatusCode != http.StatusNoContent {
		return cli.Exit(fmt.Errorf(resp.Status), 12)
	}

	fmt.Println(resp.Status)
	return nil
}
