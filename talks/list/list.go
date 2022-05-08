package list

import (
	"bytes"
	"text/template"
)

func Do() (string, error) {
	return plainTextFutureTalksList(), nil
}

type talkInfos struct {
	Title                 string
	EventName             string
	EventDay              string
	EventTime             string
	EventAddressPlaceName string
	EventAddressStreet    string
	EventAddressCity      string
	EventAddressCountry   string
	EventUrl              string
}

func plainTextFutureTalksList() string {
	var talk = talkInfos{
		Title:                 "Tu ne le sais pas encore mais tu l'as déjà documenté!",
		EventName:             "Agile tour Strasbourg 2022",
		EventDay:              "20th may 2022",
		EventTime:             "9am - 6pm",
		EventAddressPlaceName: "Epitech Strasbourg",
		EventAddressStreet:    "4 Rue du Dôme",
		EventAddressCity:      "Strasbourg",
		EventAddressCountry:   "France",
		EventUrl:              "https://www.meetup.com/MugStrasbourg/events/281350303/",
	}
	t, parseErr := template.New("talksList").Parse(
		`Marc Bouvier will be speaking in the following events:

- {{ .Title}}
  - {{ .EventName}}
  - {{ .EventDay}} • {{ .EventTime}}
  - {{ .EventAddressPlaceName}} • {{ .EventAddressStreet}} • {{ .EventAddressCity}} • {{ .EventAddressCountry}}
  - {{ .EventUrl}}`)

	if parseErr != nil {
		panic(parseErr)
	}

	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, talk)
	if err != nil {
		panic(err)
	}
	return tmplBytes.String()

}
