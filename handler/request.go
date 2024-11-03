package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

//Create Production

type CreateProductionRequest struct {
	Name        string `json:"name"`
	Producer    string `json:"producer"`
	Movie       *bool  `json:"movie"`
	Series      *bool  `json:"series"`
	Protagonist string `json:"protagonist"`
	Notice      int64  `json:"notice"`
	Assessment  string `json:"assessment"`
}

func (r *CreateProductionRequest) Validate() error {
	if r.Name == "" && r.Producer == "" && r.Movie == nil && r.Series == nil && r.Protagonist == "" && r.Notice <= 0 && r.Assessment == "" {
		return fmt.Errorf("request body is empty")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Producer == "" {
		return errParamIsRequired("producer", "string")
	}
	if r.Movie == nil {
		return errParamIsRequired("movie", "bool")
	}
	if r.Series == nil {
		return errParamIsRequired("series", "bool")
	}
	if r.Protagonist == "" {
		return errParamIsRequired("protagonist", "string")
	}
	if r.Notice <= 0 {
		return errParamIsRequired("notice", "int64")
	}
	if r.Assessment == "" {
		return errParamIsRequired("assessment", "string")
	}
	return nil
}

type UpdateProductionRequest struct {
	Name        string `json:"name"`
	Producer    string `json:"producer"`
	Movie       *bool  `json:"movie"`
	Series      *bool  `json:"series"`
	Protagonist string `json:"protagonist"`
	Notice      int64  `json:"notice"`
	Assessment  string `json:"assessment"`
}

func (r *UpdateProductionRequest) Validate() error {
	// If any field is provided, validation is true
	if r.Name != "" && r.Producer != "" && r.Movie != nil && r.Series != nil && r.Protagonist != "" && r.Notice > 0 && r.Assessment != "" {
		return nil
	}
	// If none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
