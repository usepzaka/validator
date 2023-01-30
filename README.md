# validator
===========
[![GoDoc](https://godoc.org/github.com/usepzaka/validator?status.png)](https://godoc.org/github.com/usepzaka/validator)

Validator Struct, Variable and Type for GO


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
  somethingname "github.com/usepzaka/validator"
)
```


#### List of functions:
```go
func Abs(value float64) float64
func BlackList(str, chars string) string
func ByteLength(str string, params ...string) bool
func CamelCaseToUnderscore(str string) string
func Contains(str, substring string) bool
func Count(array []interface{}, iterator ConditionIterator) int
func Each(array []interface{}, iterator Iterator)
func ErrorByField(e error, field string) string
func ErrorsByField(e error) map[string]string
func Filter(array []interface{}, iterator ConditionIterator) []interface{}
func Find(array []interface{}, iterator ConditionIterator) interface{}
func GetLine(s string, index int) (string, error)
func GetLines(s string) []string
func HasLowerCase(str string) bool
func HasUpperCase(str string) bool
func HasWhitespace(str string) bool
func HasWhitespaceOnly(str string) bool
func InRange(value interface{}, left interface{}, right interface{}) bool
func InRangeFloat32(value, left, right float32) bool
func InRangeFloat64(value, left, right float64) bool
func InRangeInt(value, left, right interface{}) bool
func IsASCII(str string) bool
func IsAlpha(str string) bool
func IsAlphaSpace(str string) bool
func IsAlphanumeric(str string) bool
func IsAlphanumericSpace(str string) bool
func IsPhone(str string) bool
func IsIndoMPhone(str string) bool
func IsNik(str string) bool
func IsBase64(str string) bool
func IsByteLength(str string, min, max int) bool
func IsCIDR(str string) bool
func IsCRC32(str string) bool
func IsCRC32b(str string) bool
func IsCreditCard(str string) bool
func IsDNSName(str string) bool
func IsDataURI(str string) bool
func IsDialString(str string) bool
func IsDivisibleBy(str, num string) bool
func IsEmail(str string) bool
func IsExistingEmail(email string) bool
func IsFilePath(str string) (bool, int)
func IsFloat(str string) bool
func IsFullWidth(str string) bool
func IsHalfWidth(str string) bool
func IsHash(str string, algorithm string) bool
func IsHexadecimal(str string) bool
func IsHexcolor(str string) bool
func IsHost(str string) bool
func IsIP(str string) bool
func IsIPv4(str string) bool
func IsIPv6(str string) bool
func IsISBN(str string, version int) bool
func IsISBN10(str string) bool
func IsISBN13(str string) bool
func IsISO3166Alpha2(str string) bool
func IsISO3166Alpha3(str string) bool
func IsISO4217(str string) bool
func IsISO693Alpha2(str string) bool
func IsISO693Alpha3b(str string) bool
func IsIn(str string, params ...string) bool
func IsInRaw(str string, params ...string) bool
func IsInt(str string) bool
func IsJSON(str string) bool
func IsLatitude(str string) bool
func IsLongitude(str string) bool
func IsLowerCase(str string) bool
func IsMAC(str string) bool
func IsMD4(str string) bool
func IsMD5(str string) bool
func IsMagnetURI(str string) bool
func IsMongoID(str string) bool
func IsMultibyte(str string) bool
func IsNatural(value float64) bool
func IsNegative(value float64) bool
func IsNonNegative(value float64) bool
func IsNonPositive(value float64) bool
func IsNotNull(str string) bool
func IsNull(str string) bool
func IsNumeric(str string) bool
func IsPort(str string) bool
func IsPositive(value float64) bool
func IsPrintableASCII(str string) bool
func IsRFC3339(str string) bool
func IsRFC3339WithoutZone(str string) bool
func IsRGBcolor(str string) bool
func IsRegex(str string) bool
func IsRequestURI(rawurl string) bool
func IsRequestURL(rawurl string) bool
func IsRipeMD128(str string) bool
func IsRipeMD160(str string) bool
func IsRsaPub(str string, params ...string) bool
func IsRsaPublicKey(str string, keylen int) bool
func IsSHA1(str string) bool
func IsSHA256(str string) bool
func IsSHA384(str string) bool
func IsSHA512(str string) bool
func IsSSN(str string) bool
func IsSemver(str string) bool
func IsTiger128(str string) bool
func IsTiger160(str string) bool
func IsTiger192(str string) bool
func IsTime(str string, format string) bool
func IsType(v interface{}, params ...string) bool
func IsURL(str string) bool
func IsUTFDigit(str string) bool
func IsUTFLetter(str string) bool
func IsUTFLetterNumeric(str string) bool
func IsUTFNumeric(str string) bool
func IsUUID(str string) bool
func IsUUIDv3(str string) bool
func IsUUIDv4(str string) bool
func IsUUIDv5(str string) bool
func IsULID(str string) bool
func IsUnixTime(str string) bool
func IsUpperCase(str string) bool
func IsVariableWidth(str string) bool
func IsWhole(value float64) bool
func LeftTrim(str, chars string) string
func Map(array []interface{}, iterator ResultIterator) []interface{}
func Matches(str, pattern string) bool
func MaxStringLength(str string, params ...string) bool
func MinStringLength(str string, params ...string) bool
func NormalizeEmail(str string) (string, error)
func PadBoth(str string, padStr string, padLen int) string
func PadLeft(str string, padStr string, padLen int) string
func PadRight(str string, padStr string, padLen int) string
func PrependPathToErrors(err error, path string) error
func Range(str string, params ...string) bool
func RemoveTags(s string) string
func ReplacePattern(str, pattern, replace string) string
func Reverse(s string) string
func RightTrim(str, chars string) string
func RuneLength(str string, params ...string) bool
func SafeFileName(str string) string
func SetFieldsRequiredByDefault(value bool)
func SetNilPtrAllowedByRequired(value bool)
func Sign(value float64) float64
func StringLength(str string, params ...string) bool
func StringMatches(s string, params ...string) bool
func StripLow(str string, keepNewLines bool) string
func ToBoolean(str string) (bool, error)
func ToFloat(str string) (float64, error)
func ToInt(value interface{}) (res int64, err error)
func ToJSON(obj interface{}) (string, error)
func ToString(obj interface{}) string
func Trim(str, chars string) string
func Truncate(str string, length int, ending string) string
func TruncatingErrorf(str string, args ...interface{}) error
func UnderscoreToCamelCase(s string) string
func ValidateMap(inputMap map[string]interface{}, validationMap map[string]interface{}) (bool, error)
func ValidateStruct(s interface{}) (bool, error)
func WhiteList(str, chars string) string
type ConditionIterator
type CustomTypeValidator
type Error
func (e Error) Error() string
type Errors
func (es Errors) Error() string
func (es Errors) Errors() []error
type ISO3166Entry
type ISO693Entry
type InterfaceParamValidator
type Iterator
type ParamValidator
type ResultIterator
type UnsupportedTypeError
func (e *UnsupportedTypeError) Error() string
type Validator
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
  Name string      `valid:"type(string)"`
  Age  int         `valid:"type(int)"`
  Meta interface{} `valid:"type(string)"`
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
If you want to validate structs, you can use tag `valid` for any field in your structure. All validators used with this field in one tag are separated by comma. If you want to skip validation, place `-` in your tag. If you need a validator that is not on the list below, you can add it like this:
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

