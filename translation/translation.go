package translation

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/arbrix/uadmin"
)

// TranslateSchema looks for schema fields with multilingual tag
// and translates them using uadmin.Translate
// Supported types:
// []{struct type}, &{struct type}
// For examples please check TestTranslationSchemaSuccess
func TranslateSchema(input interface{}, lang string) {
	val := reflect.ValueOf(input)
	switch val.Kind() {
	case reflect.Ptr:
		val = val.Elem()
		if val.Type().Kind() == reflect.Struct {
			for i := 0; i < val.NumField(); i++ {
				if isMultilingField(val.Type().Field(i)) {
					r := translate(val.Field(i).Interface(), lang)
					val.Field(i).Set(reflect.ValueOf(r))
				}

				if isStructField(val.Type().Field(i)) {
					TranslateSchema(val.Field(i).Addr().Interface(), lang)
				}

				if isSliceField(val.Type().Field(i)) {
					s := reflect.ValueOf(val.Field(i).Interface())
					for i := 0; i < s.Len(); i++ {
						TranslateSchema(s.Index(i).Addr().Interface(), lang)
					}
				}
			}
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			TranslateSchema(val.Index(i).Addr().Interface(), lang)
		}
	default:
		return
	}
}

func translate(s interface{}, lang string) interface{} {
	val, ok := s.(string)
	if !ok {
		return s
	}

	return uadmin.Translate(val, lang)
}

func isMultilingField(f reflect.StructField) bool {
	pattern := "multilingual"
	return strings.Contains(string(f.Tag), pattern)
}

func isStructField(f reflect.StructField) bool {
	return f.Type.Kind() == reflect.Struct
}

func isSliceField(f reflect.StructField) bool {
	return f.Type.Kind() == reflect.Slice
}

// GetSessionLang extract language code from request/cookies
func GetSessionLang(w http.ResponseWriter, r *http.Request) string {
	lang := r.URL.Query().Get("language")

	if lang != "" {
		// Assign the name, value, and path to build the language cookie.
		langC := http.Cookie{
			Name:  "language",
			Value: lang,
			Path:  "/",
		}
		// Set the language cookie to the browser.
		http.SetCookie(w, &langC)

	} else {
		// If no language cookie has been found
		if langC, err := r.Cookie("language"); err != nil || (langC != nil && langC.Value == "") {
			// Get the default language code from the database instead and assign to this variable.
			lang = uadmin.GetDefaultLanguage().Code
		} else {
			// Assign the value from the language cookie in this variable.
			lang = langC.Value
		}
	}

	return lang
}

func ComposeSearchString(value, lang string) string {
	return "%" + fmt.Sprintf("\"%s\":\"%s\"", lang, value) + "%"
}
