package main

import (
  "net/http"
  "github.com/martini-contrib/binding"
)

type NewLocationForm struct {
  Name        string `form:"name" binding:"required"`
  Description string `form:"description"`

  Visited bool `form:"visited"`

  Lat  float64 `form:"lat" binding:"required"`
  Long float64 `form:"long" binding:"required"`
}

func (form NewLocationForm) Validate(errors binding.Errors, req *http.Request) binding.Errors {
  if form.Lat < -90 || form.Lat > 90 {
      errors = append(errors, binding.Error{
          FieldNames:     []string{"lat"},
          Classification: "ComplaintError",
          Message:        "Invalid latitude.",
      })
  }

  if form.Long < -180 || form.Long > 180 {
      errors = append(errors, binding.Error{
          FieldNames:     []string{"long"},
          Classification: "ComplaintError",
          Message:        "Invalid longitude.",
      })
  }

  return errors
}
