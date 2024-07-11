package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Comment struct {
	CommentID      int       `json:"commentId"`
	CommentContent string    `json:"commentContent"`
	Replies        []Comment `json:"replies,omitempty"`
}

func loadCommentsFromJSON(filePath string) ([]Comment, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var comments []Comment
	err = json.Unmarshal(bytes, &comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// recursive
func countComments(comments []Comment) int {
	total := 0
	for _, comment := range comments {
		total++
		if len(comment.Replies) > 0 {
			total += countComments(comment.Replies)
		}
	}
	return total
}

func case_2() {
	fmt.Println("\n\n\n Case 2:")
	comments, err := loadCommentsFromJSON("comments.json")
	if err != nil {
		panic(err)
	}

	totalComments := countComments(comments)
	fmt.Printf("Total number of comments (including replies): %d\n", totalComments)
}
