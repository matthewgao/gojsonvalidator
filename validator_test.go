package validator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	type TVS struct {
		SS string
		BB bool
		II int
	}
	type VS struct {
		SS string
		BB bool
		II int
		TT *TVS
	}

	v := VS{}

	source := `{
		"SS": "test string",
		"BB": true,
		"II": 123,
		"TT": {
			"SS": "in test string",
			"BB": false,
			"II": 345
		}	
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	assert.Empty(t, err)
	assert.Equal(t, "test string", v.SS)
	assert.Equal(t, true, v.BB)
	assert.Equal(t, 123, v.II)
	assert.Equal(t, 345, v.TT.II)
	assert.Equal(t, false, v.TT.BB)
	assert.Equal(t, "in test string", v.TT.SS)
}

func TestParser1(t *testing.T) {
	type TVS struct {
		SS string
		BB bool
		II int
	}
	type VS struct {
		SS string `is_required:"true"`
		BB bool
		II int
		TT *TVS
	}

	v := VS{}

	source := `{
		"BB": true,
		"II": 123,
		"TT": {
			"SS": "in test string",
			"BB": false,
			"II": 345
		}	
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(err)
	assert.NotEmpty(t, err)
	// assert.Equal(t, "test string", v.SS)
	// assert.Equal(t, true, v.BB)
	// assert.Equal(t, 123, v.II)
	// assert.Equal(t, 345, v.TT.II)
	// assert.Equal(t, false, v.TT.BB)
	// assert.Equal(t, "in test string", v.TT.SS)
}

func TestParser2(t *testing.T) {
	type TVS struct {
		SS string `is_required:"true"`
		BB bool
		II int
	}
	type VS struct {
		SS string
		BB bool
		II int
		TT *TVS
	}

	v := VS{}

	source := `{
		"SS" : "aa",
		"BB": true,
		"II": 123,
		"TT": {
			"BB": false,
			"II": 345
		}	
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(err)
	assert.NotEmpty(t, err)
	// assert.Equal(t, "test string", v.SS)
	// assert.Equal(t, true, v.BB)
	// assert.Equal(t, 123, v.II)
	// assert.Equal(t, 345, v.TT.II)
	// assert.Equal(t, false, v.TT.BB)
	// assert.Equal(t, "in test string", v.TT.SS)
}

func TestParser3(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool `is_required:"true"`
		II int
	}
	type VS struct {
		SS string
		BB bool
		II int
		TT *TVS
	}

	v := VS{}

	source := `{
		"SS" : "aa",
		"BB": true,
		"II": 123,
		"TT": {
			"SS": "gege",
			"II": 345
		}	
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(err)
	assert.NotEmpty(t, err)
	// assert.Equal(t, "test string", v.SS)
	// assert.Equal(t, true, v.BB)
	// assert.Equal(t, 123, v.II)
	// assert.Equal(t, 345, v.TT.II)
	// assert.Equal(t, false, v.TT.BB)
	// assert.Equal(t, "in test string", v.TT.SS)
}

func TestParser4(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II int
	}
	type VS struct {
		SS string
		BB *bool `is_required:"true"`
		II int
		TT *TVS
	}

	v := VS{}

	source := `{
		"SS" : "aa",

		"II": 123,
		"TT": {
			"SS": "gege",
			"II": 345
		}	
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(err)
	assert.NotEmpty(t, err)
	// assert.Equal(t, "test string", v.SS)
	// assert.Equal(t, true, v.BB)
	// assert.Equal(t, 123, v.II)
	// assert.Equal(t, 345, v.TT.II)
	// assert.Equal(t, false, v.TT.BB)
	// assert.Equal(t, "in test string", v.TT.SS)
}

func TestParser5(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II int
	}
	type VS struct {
		SS string
		BB *bool `default:"true"`
		II int
		TT *TVS
	}

	v := VS{}

	source := `{
		"SS" : "aa",
		"II": 123,
		"TT": {
			"SS": "gege",
			"II": 345
		}	
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(err)
	assert.Empty(t, err)
	// assert.Equal(t, "test string", v.SS)
	assert.Equal(t, true, *v.BB)
	// assert.Equal(t, 123, v.II)
	// assert.Equal(t, 345, v.TT.II)
	// assert.Equal(t, false, v.TT.BB)
	// assert.Equal(t, "in test string", v.TT.SS)
}

func TestParser6(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II *int `default:"888"`
	}
	type VS struct {
		SS string
		BB *bool `default:"true"`
		II int   `default:"777"`
		TT *TVS
	}

	v := VS{}

	source := `{
		"SS" : "aa",
		"TT": {
			"SS": "gege"
		}	
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(*v.TT)
	assert.Empty(t, err)
	// assert.Equal(t, "test string", v.SS)
	assert.Equal(t, true, *v.BB)
	assert.Equal(t, 777, v.II)
	assert.Equal(t, 888, *v.TT.II)
	// assert.Equal(t, false, v.TT.BB)
	// assert.Equal(t, "in test string", v.TT.SS)
}

func TestParser7(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II *int `default:"888"`
	}
	type VS struct {
		SS string
		BB *bool `default:"true"`
		II int   `default:"777"`
		TT *TVS  `is_required:"true"`
	}

	v := VS{}

	source := `{
		"SS" : "aa"
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	// fmt.Println(*v.TT)
	assert.NotEmpty(t, err)
	// assert.Equal(t, "test string", v.SS)
	// assert.Equal(t, true, *v.BB)
	// assert.Equal(t, 777, v.II)
	// assert.Equal(t, 888, *v.TT.II)
	// assert.Equal(t, false, v.TT.BB)
	// assert.Equal(t, "in test string", v.TT.SS)
}

func TestParser8(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II *int `default:"888"`
	}
	type VS struct {
		SS string
		BB *bool `default:"true"`
		II int   `default:"777"`
		TT *TVS
	}

	v := VS{}

	source := `{
		"SS" : "aa",
		"TT" :{
			
		}
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	// fmt.Println(*v.TT)
	assert.Empty(t, err)
	// assert.Equal(t, "test string", v.SS)
	// assert.Equal(t, true, *v.BB)
	// assert.Equal(t, 777, v.II)
	assert.Equal(t, 888, *v.TT.II)
	// assert.Equal(t, false, v.TT.BB)
	// assert.Equal(t, "in test string", v.TT.SS)
}

//FIXME:This test will throw a panic
func TestParser9(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II *int `default:"888"`
	}
	type VS struct {
		SS string
		BB *bool `default:"true"`
		II int   `default:"777"`
		TT *TVS
	}

	v := VS{}

	source := `{
		"SS" : "aa"
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(*v.TT)
	assert.Empty(t, err)
	// assert.Empty(t, v.TT)
	// assert.Equal(t, nil, v.TT)
	// assert.Equal(t, true, *v.BB)
	// assert.Equal(t, 777, v.II)
	assert.Equal(t, 888, *v.TT.II)
	// assert.Equal(t, false, v.TT.BB)
	// assert.Equal(t, "in test string", v.TT.SS)
}

//TODO: max len test

func TestParser10(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II *int `default:"888" max:"900"`
	}
	type VS struct {
		SS string
		BB *bool `default:"true"`
		II int   `default:"777"`
		TT *TVS
	}

	v := VS{}

	source := `{
		"SS" : "aa",
		"TT": {
			"SS": "gege",
			"II": 901
		}	
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(*v.TT)
	fmt.Println(err)

	assert.NotEmpty(t, err)
	// assert.Equal(t, "test string", v.SS)
	// assert.Equal(t, true, *v.BB)
	// assert.Equal(t, 777, v.II)
	// assert.Equal(t, 888, *v.TT.II)
	// assert.Equal(t, false, v.TT.BB)
	// assert.Equal(t, "in test string", v.TT.SS)
}

func TestParser11(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II int `default:"888" max:"900"`
	}
	type VS struct {
		SS string
		BB *bool `default:"true"`
		II int   `default:"777"`
		TT *TVS
	}

	v := VS{}

	source := `{
		"SS" : "aa",
		"TT": {
			"SS": "gege",
			"II": 901
		}	
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(*v.TT)
	fmt.Println(err)

	assert.NotEmpty(t, err)
	// assert.Equal(t, "test string", v.SS)
	// assert.Equal(t, true, *v.BB)
	// assert.Equal(t, 777, v.II)
	// assert.Equal(t, 888, *v.TT.II)
	// assert.Equal(t, false, v.TT.BB)
	// assert.Equal(t, "in test string", v.TT.SS)
}

func TestParser12(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II int `default:"888" max:"900"`
	}
	type VS struct {
		SS   string
		SiSi []int64 `is_required:"true"`
	}

	v := VS{}

	source := `{
		"SS" : "aa",
		"SiSi": [1]	
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(v.SiSi)
	fmt.Println(err)

	assert.Empty(t, err)
}

func TestParser13(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II int `default:"888" max:"900"`
	}
	type VS struct {
		SS   string
		SiSi []int64 `is_required:"true"`
	}

	v := VS{}

	source := `{
		"SS" : "aa"
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(v.SiSi)
	fmt.Println(err)

	assert.NotEmpty(t, err)
}

func TestParser14(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II int `max:"900" min:"1"`
	}
	type VS struct {
		SS   string
		SiSi []int64 `is_required:"true"`
		TT   *TVS
	}

	v := VS{}

	source := `{
		"SS" : "aa",
		"SiSi": [1]
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(v.SiSi)
	fmt.Println(err)

	assert.Empty(t, err)
	// assert.Equal(t, 888, v.TT.II)
}

func TestParser15(t *testing.T) {
	type TVS struct {
		SS string
		BB *bool
		II int `max:"900" min:"1"`
	}
	type VS struct {
		SS   string
		SiSi []int64 `is_required:"true"`
		TT   *TVS
	}

	v := VS{}

	source := `{
		"SS" : "aa",
		"SiSi": [1],
		"TT":{
			"II": -1
		}
	}`
	err := ValidateJson([]byte(source), &v)
	// err := v.ValidateJson([]byte(source))

	fmt.Println(v.SiSi)
	fmt.Println(err)

	assert.NotEmpty(t, err)
	assert.Equal(t, -1, v.TT.II)
}
