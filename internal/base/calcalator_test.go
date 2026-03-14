package base_test

import (
	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
	"testing"
)

func Test_Add(t *testing.T) {
	t.Parallel()

	t.Run("1 + 2 = 3", func(t *testing.T) {
		t.Parallel()

		rsl := base.Add(1, 2)
		expected := 3

		assert.Equal(t, rsl, expected)
	})

	t.Run("2 + 2 = 4", func(t *testing.T) {
		t.Parallel()

		rsl := base.Add(2, 2)
		expected := 4

		assert.Equal(t, rsl, expected)
	})

}
