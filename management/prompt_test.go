package management

import (
	"testing"
)

func TestPrompt(t *testing.T) {

	t.Run("Read", func(t *testing.T) {
		ps, err := m.Prompt.Read()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", ps)
	})

	falseV := false
	trueV := true

	t.Run("Update", func(t *testing.T) {
		t.Cleanup(func() {
			m.Prompt.Update(&Prompt{
				UniversalLoginExperience: "classic",
				IdentifierFirst:          &falseV,
			})
		})

		expected := &Prompt{
			UniversalLoginExperience: "new",
			IdentifierFirst:          &trueV,
		}

		err := m.Prompt.Update(&Prompt{
			UniversalLoginExperience: expected.UniversalLoginExperience,
			IdentifierFirst:          expected.IdentifierFirst,
		})
		if err != nil {
			t.Error(err)
		}

		ps, err := m.Prompt.Read()
		if err != nil {
			t.Error(err)
		}
		if ps.UniversalLoginExperience != expected.UniversalLoginExperience {
			t.Errorf("unexpected output. have %v, expected %v", ps.UniversalLoginExperience, expected.IdentifierFirst)
		}
		if ps.IdentifierFirst != expected.IdentifierFirst {
			t.Errorf("unexpected output. have %v, expected %v", ps.IdentifierFirst, expected.IdentifierFirst)
		}
		t.Logf("%v\n", ps)
	})
}
