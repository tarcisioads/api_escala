package handler

import "fmt"

type SignupUserRequest struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

func (r *SignupUserRequest) Validate() error {
  if r.Email == "" {
    return errParamIsRequired("email", "string")
  }
  if r.Password == "" {
    return errParamIsRequired("password", "string")
  }
 
  return nil
}

type SigninUserRequest struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

func (r *SigninUserRequest) Validate() error {
  // If any field is provided, validation is truthy
  if r.Email != "" {
    return nil
  }
  if r.Password != "" {
    return nil
  }
  // If none of the fields were provided, return falsy
  return fmt.Errorf("at least one valid field must be provided")
}


