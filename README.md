# gojsonvalidator
Generic tool to validate input json string and check the parameters

## Introduce
In most of the micro-service application, we use json to deliver data from one to another, but validating the parameter in json is always boring, so this tool will give you a big help. It can check which parameters are required, and if the value are exceed the limitation, or you can set a default value if some optional key is not given.

## How to Use
Assuming you have this json:
~~~
{
    "name": "Matthew",
    "age" : 17,
    "married" : true,
    "hobby" : ["music","sport"],
    "major" : "cs",
    "contact_info" : {
        "tele" : 123456,
        "city" : ""
    } 
}
~~~

* name is required
* age is in 0-120 range, and required
* married is required
* hobby can not specify more then 4
* major can only be "cs" or "ee"
* contact_info is optional
* tele has a default value 1234
* city has a default value "shanghai"


Then you must have this struct to unmarshal the json, and check each of the parameter

~~~
type Person struct{
    Name string     `json:"name"`
    Age int         `json:"age"`
    Married bool    `json:"married"`
    Hobby []string  `json:"hobby"`
    Major string    `json:"major"`
    CI ContactInfo  `json:"contact_info"`
}

type ContactInfo struct{
    Tele int    `json:"tele"`
    City string `json:"city"`
}
~~~

use this tool you can done all of this checking work in one line, define the struct like this:

~~~
type Person struct{
    Name    string          `json:"name" is_required:"true"`
    Age     *int            `json:"age" is_required:"true" min:"0" max:"120"`
    Married *bool           `json:"married" is_required:"true"` 
    Hobby   []string        `json:"hobby" max_len:"4"`
    Major   string          `json:"major enum:"cs,ee""`
    CI      *ContactInfo    `json:"contact_info"`
}

type ContactInfo struct{
    Tele int    `json:"tele" default:"1234"`
    City string `json:"city" default:"shanghai"`
}
~~~
with all of this in hand, then run this:

~~~
p := Person{}
body := `
{
    "name": "Matthew",
    "age" : 17,
    "married" : true,
    "hobby" : ["music","sport"],
    "major" : "cs",
    "contact_info" : {
        "tele" : 123456,
        "city" : ""
    } 
}
`
err := ValidateJson([]byte(body), &p)
fmt.Println(err)
~~~

Cool! Isn't it?

## Note
* int and bool are two special type, because int takes 0 as a default value when doing unmarshal, bool take a false as a default value, so we have no way to know if 0 is a caller specified value or default value, so we use pointer, if it those key is not given, then it will be a nil pointer

* A embed struct must be define as a pointer, just like ContactInfo

## Unsupport Case
* set default value in array/slice is not support now, will supported later
* nested linked-list is not supported
