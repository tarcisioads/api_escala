package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateEscalaRequest struct {
  Ministry string `json:"ministry"`
  Month int64 `json:"month"`
  Year int64 `json:"year"`
  Data string `json:"data"`
}

func (r *CreateEscalaRequest) Validate() error {
	if r.Ministry == "" && r.Month <=0 && r.Year <= 0 && r.Data == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Ministry == "" {
		return errParamIsRequired("ministry", "string")
	}
	if r.Month<= 0 {
		return errParamIsRequired("month", "int64")
	}
	if r.Year <= 0 {
		return errParamIsRequired("year", "int64")
	}
	if r.Data == "" {
		return errParamIsRequired("data", "string")
	}

  return nil
}

type UpdateEscalaRequest struct {
  Ministry string `json:"ministry"`
  Month int64 `json:"month"`
  Year int64 `json:"year"`
  Data string `json:"data"`
}

func (r *UpdateEscalaRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Ministry != "" || r.Month > 0 || r.Year > 0 || r.Data != "" {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
