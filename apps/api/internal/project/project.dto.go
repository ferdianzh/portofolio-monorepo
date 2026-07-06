package project

type CreateProjectDTO struct {
	Title      string  `json:"title" validate:"required,min=3,max=100"`
	IsFeatured bool    `json:"is_featured"`
	LiveUrl    *string `json:"live_url"`
	Content    *string `json:"content"`
}

func (d CreateProjectDTO) ToModel() Project {
	return Project{
		Title:      d.Title,
		IsFeatured: d.IsFeatured,
		LiveUrl:    d.LiveUrl,
		Content:    d.Content,
	}
}

type UpdateProjectDTO struct {
	Title      string  `json:"title" validate:"omitempty,min=3,max=100"`
	IsFeatured bool    `json:"is_featured"`
	LiveUrl    *string `json:"live_url"`
	Content    *string `json:"content"`
}

func (d UpdateProjectDTO) ToModel() Project {
	return Project{
		Title:      d.Title,
		IsFeatured: d.IsFeatured,
		LiveUrl:    d.LiveUrl,
		Content:    d.Content,
	}
}
