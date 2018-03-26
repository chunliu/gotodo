package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/urfave/cli.v2"
)

const baseURL = "http://localhost:8080/todo/"

func printResult(todo Todo) {
	fmt.Println("Id: ", todo.ID)
	fmt.Println("Name: ", todo.Name)
	fmt.Println("IsCompleled: ", todo.IsCompleted)
}

func get(c *cli.Context) error {
	id := c.Int("id")
	url := baseURL
	if id != 0 {
		url = fmt.Sprintf("%s%d", baseURL, id)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get: ", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return fmt.Errorf(resp.Status)
	}

	if id == 0 {
		var todos TodoItems

		if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
			fmt.Println("JSON decode: ", err)
			return err
		}

		for _, item := range todos {
			printResult(item)
		}
	} else {
		var todo Todo

		if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
			fmt.Println("JSON decode: ", err)
			return err
		}

		printResult(todo)
	}

	return nil
}

func create(c *cli.Context) error {
	n := c.String("name")
	todo := Todo{
		ID:          0,
		Key:         0,
		Name:        n,
		IsCompleted: false,
	}
	jsonValue, _ := json.Marshal(todo)
	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Post: ", err)
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Println(resp.Status)
		return fmt.Errorf(resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
		fmt.Println("JSON decode: ", err)
		return err
	}

	printResult(todo)

	return nil
}

func update(c *cli.Context) error {
	id := c.Int("id")
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
		fmt.Println("PUT: ", err)
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		fmt.Println(resp.Status)
		return fmt.Errorf(resp.Status)
	}

	fmt.Println(resp.Status)
	return nil
}

func delete(c *cli.Context) error {
	id := c.Int("id")
	url := fmt.Sprintf("%s%d", baseURL, id)
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Delete: ", err)
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		fmt.Println(resp.Status)
		return fmt.Errorf(resp.Status)
	}

	fmt.Println(resp.Status)
	return nil
}
