package main

type Task struct {
    Id     int  `json:"id"`
    Name  string  `json:"name"`
    Description string  `json:"description"`
    Status  string `json:"status"`
}
