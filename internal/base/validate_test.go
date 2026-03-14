package base_test

import (
	"testing"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	//t.Run("{Title: '', Description: '123', UserID: 1  } - true", func(t *testing.T) {
	//	t.Parallel()
	//
	//	invalid := base.Validate(&base.ValidateRequest{
	//		Title:       "",
	//		Description: "",
	//		UserID:      "",
	//	})
	//
	//	assert.Equal(t,
	//		[]string{
	//			"title is empty",
	//			"description is empty",
	//			"userID is empty",
	//		},
	//		invalid)
	//})
	//
	//t.Run("nil - true", func(t *testing.T) {
	//	t.Parallel()
	//
	//	error := base.Validate(nil)
	//
	//	assert.Equal(t, 1, len(error))
	//})
}
