package list

func Do() (string, error) {
	return plainTextFutureTalksList(), nil
}

func plainTextFutureTalksList() string {
	return `Marc Bouvier will be speaking in the following events:

- Tu ne le sais pas encore mais tu l'as déjà documenté!
  - Agile tour Strasbourg 2022
  - 20th may 2022 • 9am - 6pm
  - Epitech Strasbourg • 4 Rue du Dôme • Strasbourg • France
  - https://www.meetup.com/MugStrasbourg/events/281350303/`
}
