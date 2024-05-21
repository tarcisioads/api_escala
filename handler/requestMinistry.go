package handler

import "fmt"

type CreateMinistryRequest struct {
  Name string `json:"name"`
  Members []*string `json:"members"`
}

func (r *CreateMinistryRequest) Validate() error {
  if r.Name == "" {
    return errParamIsRequired("name", "string")
  }
  if r.Members == nil {
    return errParamIsRequired("name", "string")
  }
 
  return nil
}

type UpdateMinistryRequest struct {
  Name string `json:"name"`
  Members []*string `json:"members"`
}

func (r *UpdateMinistryRequest) Validate() error {
  // If any field is provided, validation is truthy
  if r.Name != "" {
    return nil
  }
  if r.Members != nil {
    return nil
  }
  // If none of the fields were provided, return falsy
  return fmt.Errorf("at least one valid field must be provided")
}


