package files_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("Should read file located at data/input.txt", func(t *testing.T) {
		pathname := "data/input.txt"
		expected := "4\n5\n32\n100\n867543"
		got, _ := files.ReadFile(pathname)

		assert.Equal(t, expected, got)
	})

	t.Run("Should read file located at data/input2.csv", func(t *testing.T) {
		pathname := "data/input2.csv"
		expected := "4,5,32,100,867543"
		got, _ := files.ReadFile(pathname)

		assert.Equal(t, expected, got)
	})

	t.Run("Should through the error if file not found", func(t *testing.T) {
		pathname := "data/input22.csv"
		_, err := files.ReadFile(pathname)

		assert.Error(t, err)
	})
}
