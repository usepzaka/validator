# validator
Validator Struct, Variable and Type for GO
validator

#### Installation
Make sure that Go is installed on your computer.
Type the following command in your terminal:

	go get github.com/usepzaka/validator


After it the package is ready to use.


#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/usepzaka/validator"
```
If you are unhappy to use long `validator`, you can do something like this:
```go
import (
  valid "github.com/usepzaka/validator"
)
```


#### Examples
###### IsURL
```go
println(validator.IsURL(`http://user@pass:domain.com/path/page`))
```
###### IsType
```go
println(validator.IsType("Bob", "string"))
println(validator.IsType(1, "int"))
i := 1
println(validator.IsType(&i, "*int"))
```

IsType can be used through the tag `type` which is essential for map validation:
```go
type User	struct {
  Name string      `validate:"type(string)"`
  Age  int         `validate:"type(int)"`
  Meta interface{} `validate:"type(string)"`
}
result, err := validator.ValidateStruct(User{"Bob", 20, "meta"})
if err != nil {
	println("error: " + err.Error())
}
println(result)
```
###### ToString
```go
type User struct {
	FirstName string
	LastName string
}

str := validator.ToString(&User{"John", "Juan"})
println(str)
```
###### Each, Map, Filter, Count for slices
Each iterates over the slice/array and calls Iterator for every item
```go
data := []interface{}{1, 2, 3, 4, 5}
var fn validator.Iterator = func(value interface{}, index int) {
	println(value.(int))
}
validator.Each(data, fn)
```
```go
data := []interface{}{1, 2, 3, 4, 5}
var fn validator.ResultIterator = func(value interface{}, index int) interface{} {
	return value.(int) * 3
}
_ = validator.Map(data, fn) // result = []interface{}{1, 6, 9, 12, 15}
```
```go
data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var fn validator.ConditionIterator = func(value interface{}, index int) bool {
	return value.(int)%2 == 0
}
_ = validator.Filter(data, fn) // result = []interface{}{2, 4, 6, 8, 10}
_ = validator.Count(data, fn) // result = 5
```
###### ValidateStruct [#2](https://github.com/usepzaka/validator/pull/2)
If you want to validate structs, you can use tag `validate` for any field in your structure. All validators used with this field in one tag are separated by comma. If you want to skip validation, place `-` in your tag. If you need a validator that is not on the list below, you can add it like this:
```go
validator.TagMap["duck"] = validator.Validator(func(str string) bool {
	return str == "duck"
})
```
For completely custom validators (interface-based), see below.

Here is a list of available validators for struct fields (validator - used function):
```go
"email":              IsEmail,
"url":                IsURL,
"dialstring":         IsDialString,
"requrl":             IsRequestURL,
"requri":             IsRequestURI,
"alpha":              IsAlpha,
"utfletter":          IsUTFLetter,
"alphanum":           IsAlphanumeric,
"utfletternum":       IsUTFLetterNumeric,
"numeric":            IsNumeric,
"utfnumeric":         IsUTFNumeric,
"utfdigit":           IsUTFDigit,
"hexadecimal":        IsHexadecimal,
"hexcolor":           IsHexcolor,
"rgbcolor":           IsRGBcolor,
"lowercase":          IsLowerCase,
"uppercase":          IsUpperCase,
"int":                IsInt,
"float":              IsFloat,
"null":               IsNull,
"uuid":               IsUUID,
"uuidv3":             IsUUIDv3,
"uuidv4":             IsUUIDv4,
"uuidv5":             IsUUIDv5,
"creditcard":         IsCreditCard,
"isbn10":             IsISBN10,
"isbn13":             IsISBN13,
"json":               IsJSON,
"multibyte":          IsMultibyte,
"ascii":              IsASCII,
"printableascii":     IsPrintableASCII,
"fullwidth":          IsFullWidth,
"halfwidth":          IsHalfWidth,
"variablewidth":      IsVariableWidth,
"base64":             IsBase64,
"datauri":            IsDataURI,
"ip":                 IsIP,
"port":               IsPort,
"ipv4":               IsIPv4,
"ipv6":               IsIPv6,
"dns":                IsDNSName,
"host":               IsHost,
"mac":                IsMAC,
"latitude":           IsLatitude,
"longitude":          IsLongitude,
"ssn":                IsSSN,
"semver":             IsSemver,
"rfc3339":            IsRFC3339,
"rfc3339WithoutZone": IsRFC3339WithoutZone,
"ISO3166Alpha2":      IsISO3166Alpha2,
"ISO3166Alpha3":      IsISO3166Alpha3,
"ulid":               IsULID,
```
Validators with parameters

```go
"range(min|max)": Range,
"length(min|max)": ByteLength,
"runelength(min|max)": RuneLength,
"stringlength(min|max)": StringLength,
"matches(pattern)": StringMatches,
"in(string1|string2|...|stringN)": IsIn,
"rsapub(keylength)" : IsRsaPub,
"minstringlength(int): MinStringLength,
"maxstringlength(int): MaxStringLength,
```
Validators with parameters for any type

```go
"type(type)": IsType,
```

And here is small example of usage:
```go
type Post struct {
	Title    string `valid:"alphanum,required"`
	Message  string `valid:"duck,ascii"`
	Message2 string `valid:"animal(dog)"`
	AuthorIP string `valid:"ipv4"`
	Date     string `valid:"-"`
}
post := &Post{
	Title:   "My Example Post",
	Message: "duck",
	Message2: "dog",
	AuthorIP: "123.234.54.3",
}

// Add your own struct validation tags
validator.TagMap["duck"] = validator.Validator(func(str string) bool {
	return str == "duck"
})

// Add your own struct validation tags with parameter
validator.ParamTagMap["animal"] = validator.ParamValidator(func(str string, params ...string) bool {
    species := params[0]
    return str == species
})
validator.ParamTagRegexMap["animal"] = regexp.MustCompile("^animal\\((\\w+)\\)$")

result, err := validator.ValidateStruct(post)
if err != nil {
	println("error: " + err.Error())
}
println(result)
```
###### ValidateMap [#2](https://github.com/usepzaka/validator/pull/338)
If you want to validate maps, you can use the map to be validated and a validation map that contain the same tags used in ValidateStruct, both maps have to be in the form `map[string]interface{}`

So here is small example of usage:
```go
var mapTemplate = map[string]interface{}{
	"name":"required,alpha",
	"family":"required,alpha",
	"email":"required,email",
	"cell-phone":"numeric",
	"address":map[string]interface{}{
		"line1":"required,alphanum",
		"line2":"alphanum",
		"postal-code":"numeric",
	},
}

var inputMap = map[string]interface{}{
	"name":"Bob",
	"family":"Smith",
	"email":"foo@bar.baz",
	"address":map[string]interface{}{
		"line1":"",
		"line2":"",
		"postal-code":"",
	},
}

result, err := validator.ValidateMap(inputMap, mapTemplate)
if err != nil {
	println("error: " + err.Error())
}
println(result)
```

###### WhiteList
```go
// Remove all characters from string ignoring characters between "a" and "z"
println(validator.WhiteList("a3a43a5a4a3a2a23a4a5a4a3a4", "a-z") == "aaaaaaaaaaaa")
```

###### Custom validation functions
Custom validation using your own domain specific validators is also available - here's an example of how to use it:
```go
import "github.com/usepzaka/validator"

type CustomByteArray [6]byte // custom types are supported and can be validated

type StructWithCustomByteArray struct {
  ID              CustomByteArray `valid:"customByteArrayValidator,customMinLengthValidator"` // multiple custom validators are possible as well and will be evaluated in sequence
  Email           string          `valid:"email"`
  CustomMinLength int             `valid:"-"`
}

validator.CustomTypeTagMap.Set("customByteArrayValidator", func(i interface{}, context interface{}) bool {
  switch v := context.(type) { // you can type switch on the context interface being validated
  case StructWithCustomByteArray:
    // you can check and validate against some other field in the context,
    // return early or not validate against the context at all – your choice
  case SomeOtherType:
    // ...
  default:
    // expecting some other type? Throw/panic here or continue
  }

  switch v := i.(type) { // type switch on the struct field being validated
  case CustomByteArray:
    for _, e := range v { // this validator checks that the byte array is not empty, i.e. not all zeroes
      if e != 0 {
        return true
      }
    }
  }
  return false
})
validator.CustomTypeTagMap.Set("customMinLengthValidator", func(i interface{}, context interface{}) bool {
  switch v := context.(type) { // this validates a field against the value in another field, i.e. dependent validation
  case StructWithCustomByteArray:
    return len(v.ID) >= v.CustomMinLength
  }
  return false
})
```

###### Loop over Error()
By default .Error() returns all errors in a single String. To access each error you can do this:
```go
  if err != nil {
    errs := err.(validator.Errors).Errors()
    for _, e := range errs {
      fmt.Println(e.Error())
    }
  }
```

###### Custom error messages
Custom error messages are supported via annotations by adding the `~` separator - here's an example of how to use it:
```go
type Ticket struct {
  Id        int64     `json:"id"`
  FirstName string    `json:"firstname" valid:"required~First name is blank"`
}
```

#### Notes
Documentation is available here: [godoc.org](https://godoc.org/github.com/usepzaka/validator).
Full information about code coverage is also available here: [validator on gocover.io](http://gocover.io/github.com/usepzaka/validator).

