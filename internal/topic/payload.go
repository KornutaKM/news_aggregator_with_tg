package topic

type CreateTopicRequest struct {
	Name string `json:"name" validate:"required"`
}
