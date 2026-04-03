package translation_test

import (
	"testing"

	"github.com/NalisRI/hello-api/translation"
)

func TestTranslate(t *testing.T) {
	//Arrange
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "German",
			Translation: "hallo",
		},
		{
			Word:        "Hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello ",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "bye",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "bye",
			Language:    "german",
			Translation: "",
		},
	}
	for _, tc := range tt {
		//Act
		res := translation.Translate(tc.Word, tc.Language)

		//Assert
		if res != tc.Translation {
			t.Errorf(
				`expected "%s" to be "%s" from "%s" but received "%s"`,
				tc.Word, tc.Language, tc.Translation, res)
		}
	}
}
