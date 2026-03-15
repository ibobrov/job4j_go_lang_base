package base_test

import (
	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
	"testing"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("{ Title: '', Description: '123', UserID: 1  } - true", func(t *testing.T) {
		t.Parallel()

		invalid := base.Validate(&base.ValidateRequest{
			UserID:      "",
			Title:       "",
			Description: "",
		})

		assert.Equal(t,
			[]string{
				"UserID is empty",
				"Title is empty",
				"Description is empty",
			},
			invalid)
	})

	t.Run("nil - true", func(t *testing.T) {
		t.Parallel()

		errors := base.Validate(nil)

		assert.Equal(t, []string{"ValidateRequest is nil"}, errors)
	})
}
