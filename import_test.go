package gifs

import "testing"

func TestImport(t *testing.T) {
	var r Response
	var err error

	// p := ImportPayload{
	// 	Source: "https://www.youtube.com/watch?v=CND87JyaIOs",
	// 	Title:  "goku kaioken blue",
	// 	Attribution: Attribution{
	// 		Site: "www.youtube.com",
	// 		User: "Maxence",
	// 	},
	// }

	p := ImportPayload{
		Source: "https://vine.co/v/ibAU6OH2I0K",
		Title:  "2015 Craziness",
		Attribution: Attribution{
			Site: "vine",
			User: "Maxence",
		},
	}

	if r, err = Import(p); err != nil {
		t.Fatal(err)
	}

	t.Logf("response: %+v", r)
}

func TestImportError(t *testing.T) {
	var r Response
	var err error

	p := ImportPayload{}

	if r, err = Import(p); err == nil {
		t.Fatal("sould error")
	}

	t.Logf("response: %+v", r)
}
