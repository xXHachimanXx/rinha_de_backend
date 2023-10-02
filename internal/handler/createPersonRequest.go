package handler

type CreatePersonRequest struct {
	ID        string   `valid:"uuidv4"`
	Nickname  string   `valid:"required" json:"apelido"`
	Name      string   `valid:"required" json:"nome"`
	Birthdate string   `valid:"required" json:"nascimento"`
	Stack     []string `valid:"required" json:"stack"`
}

func (c *CreatePersonRequest) Validate() error {
	return nil
}
