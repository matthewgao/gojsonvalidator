package validator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name    string       `json:"name" is_required:"true"`
	Age     *int         `json:"age" is_required:"true" min:"0" max:"120"`
	Married *bool        `json:"married" is_required:"true"`
	Hobby   []string     `json:"hobby" max_len:"4"`
	Major   string       `json:"major" enum:"cs,ee"`
	CI      *ContactInfo `json:"contact_info"`
}

type ContactInfo struct {
	Tele int    `json:"tele" default:"1234"`
	City string `json:"city" default:"shanghai"`
}

func TestExample(t *testing.T) {
	p := Person{}
	body := `
	{
		"name": "Matthew",
		"age" : 17,
		"married" : true,
		"hobby" : ["music","sport"],
		"major" : "cs",
		"contact_info" : {
			"city" : ""
		} 
	}
	`
	err := ValidateJson([]byte(body), &p)
	fmt.Println(err)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", p.CI)
}

func TestExampleMissingRequired(t *testing.T) {
	p := Person{}
	body := `
	{
		"age" : 17,
		"married" : true,
		"hobby" : ["music","sport"],
		"major" : "cs",
		"contact_info" : {
			"city" : ""
		} 
	}
	`
	err := ValidateJson([]byte(body), &p)
	fmt.Println(err)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", p.CI)

	assert.NotEmpty(t, err)
}

func TestExampleAgeExceed(t *testing.T) {
	p := Person{}
	body := `
	{
		"name": "Matthew",
		"age" : 130,
		"married" : true,
		"hobby" : ["music","sport"],
		"major" : "cs",
		"contact_info" : {
			"city" : ""
		} 
	}
	`
	err := ValidateJson([]byte(body), &p)
	fmt.Println(err)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", p.CI)
	assert.NotEmpty(t, err)
}

func TestExampleHobbyExceed4(t *testing.T) {
	p := Person{}
	body := `
	{
		"name": "Matthew",
		"age" : 17,
		"married" : true,
		"hobby" : ["music","sport","1","2","5"],
		"major" : "cs",
		"contact_info" : {
			"city" : ""
		} 
	}
	`
	err := ValidateJson([]byte(body), &p)
	fmt.Println(err)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", p.CI)
	assert.NotEmpty(t, err)
}

func TestExampleMajorNotInEnum(t *testing.T) {
	p := Person{}
	body := `
	{
		"name": "Matthew",
		"age" : 17,
		"married" : true,
		"hobby" : ["music","sport"],
		"major" : "csx",
		"contact_info" : {
			"city" : ""
		} 
	}
	`
	err := ValidateJson([]byte(body), &p)
	fmt.Println(err)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", p.CI)
	assert.NotEmpty(t, err)
}

func TestExampleEmptyCI(t *testing.T) {
	p := Person{}
	body := `
	{
		"name": "Matthew",
		"age" : 17,
		"married" : true,
		"hobby" : ["music","sport"],
		"major" : "cs"
	}
	`
	err := ValidateJson([]byte(body), &p)
	fmt.Println(err)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", p.CI)
	assert.Empty(t, err)
	assert.Equal(t, "shanghai", p.CI.City)
	assert.Equal(t, 1234, p.CI.Tele)
}
