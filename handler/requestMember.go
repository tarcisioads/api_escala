package handler

import "fmt"

type CreateMemberRequest struct {
  Name string `json:"name"`
}

func (r *CreateMemberRequest) Validate() error {
  if r.Name == "" {
    return errParamIsRequired("name", "string")
  }
  return nil
}

type UpdateMemberRequest struct {
  Name string `json:"name"`
}

func (r *UpdateMemberRequest) Validate() error {
  // If any field is provided, validation is truthy
  if r.Name != "" {
    return nil
  }
  // If none of the fields were provided, return falsy
  return fmt.Errorf("at least one valid field must be provided")
}


