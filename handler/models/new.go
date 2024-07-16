package models

type CreateNewRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  int    `json:"author"`
}

type UpdateNewRequest struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  int    `json:"author"`
}

type CreateNewResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  int    `json:"author"`
}

type GetNewResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  int    `json:"author"`
}
