package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDo(t *testing.T) {
	expectedTalkInfosList := `Marc Bouvier will be speaking in the following events:

- Tu ne le sais pas encore mais tu l'as déjà documenté!
  - Agile tour Strasbourg 2022
  - 20th may 2022 • 9am - 6pm
  - Epitech Strasbourg • 4 Rue du Dôme • Strasbourg • France
  - https://www.meetup.com/MugStrasbourg/events/281350303/`
	actual, err := Do()

	assert.Nil(t, err)
	assert.Equal(t, actual, expectedTalkInfosList)
}
