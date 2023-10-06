package translation

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslationSchemaSuccess(t *testing.T) {
	type (
		SubStruct struct {
			Foo string `admin:"multilingual"`
			Bar string `admin:"multilingual"`
		}
		validStruct struct {
			Foo  string `admin:"multilingual"`
			Bar  []SubStruct
			Test SubStruct
		}
	)

	var (
		langCode    = "uk"
		validSchema = []validStruct{
			{
				Foo: "{\"en\":\"decision\",\"uk\":\"положення\"}\n",
				Bar: []SubStruct{
					{
						Foo: "{\"en\":\"decision\",\"uk\":\"положення\"}\n",
						Bar: "{\"en\":\"decision\",\"uk\":\"положення\"}\n",
					},
				},
				Test: SubStruct{
					Foo: "{\"en\":\"decision\",\"uk\":\"положення\"}\n",
					Bar: "{\"en\":\"decision\",\"uk\":\"положення\"}\n",
				},
			},
		}
		translatedSchema = []validStruct{
			{
				Foo: "положення",
				Bar: []SubStruct{
					{
						Foo: "положення",
						Bar: "положення",
					},
				},
				Test: SubStruct{
					Foo: "положення",
					Bar: "положення",
				},
			},
		}
	)

	TranslateSchema(validSchema, langCode)
	assert.Equal(t, validSchema, translatedSchema, "The two schemas should be the same.")
}

func TestTranslationSchemaInvalidSchema(t *testing.T) {
	type invalidStruct struct {
		Foo string
		Bar string
	}
	var (
		langCode      = "uk"
		invalidSchema = invalidStruct{
			Foo: "{\"en\":\"decision\",\"uk\":\"положення\"}\n",
			Bar: "{\"en\":\"decision\",\"uk\":\"положення\"}\n",
		}
		expectedSchema = invalidStruct{
			Foo: "{\"en\":\"decision\",\"uk\":\"положення\"}\n",
			Bar: "{\"en\":\"decision\",\"uk\":\"положення\"}\n",
		}
	)

	TranslateSchema(invalidSchema, langCode)
	assert.Equal(t, invalidSchema, expectedSchema, "The two schemas should be the same.")
}

func TestGetSessionLang(t *testing.T) {
	reqWithLang := httptest.NewRequest("GET", "http://localhost:8080/api/personas/1?language=en", nil)
	reqWithoutLang := httptest.NewRequest("GET", "http://localhost:8080/api/personas/1", nil)
	w := httptest.NewRecorder()

	lang := GetSessionLang(w, reqWithLang)
	assert.Equal(t, lang, "en")
	// while UT run `uadmin.GetDefaultLanguage().Code` returns empty string
	lang = GetSessionLang(w, reqWithoutLang)
	assert.Equal(t, lang, "")
}
