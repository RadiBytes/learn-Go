package mps

import (
	"testing"
	pointrserrs "testing_test/7-pointrs_errs"
	"testing_test/utils"
)

func TestSearch(t *testing.T) {
	ent1 := pointrserrs.Entity{Name: "Symposium", Function: "Creates musical notes"}
	entityMap := Obj{"Entity_1": ent1}

	t.Run("Known word", func(t *testing.T) {
		ent, _ := entityMap.Search("Entity_1")
		got := ent.Name
		want := "Symposium"

		utils.AssertCorrectMessage(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := entityMap.Search("Entity_2")
		want := "No entity with the name: Entity_2"

		utils.AssertCorrectMessage(t, err.Error(), want)
	})

}
