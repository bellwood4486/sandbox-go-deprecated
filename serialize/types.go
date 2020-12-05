//go:generate deep-copy -pointer-receiver -o types_deepcopy.go --type SampleLarge --type SampleSmall .

package serialize

import "encoding/json"

type SampleSmall struct {
	Values []*string
}

type SampleLarge struct {
	Glossary struct {
		Title    string
		GlossDiv struct {
			Title     string
			GlossList struct {
				GlossEntry struct {
					ID        string
					SortAs    string
					GlossTerm string
					Acronym   string
					Abbrev    string
					GlossDef  struct {
						Para         string
						GlossSeeAlso []string
					}
					GlossSee string
				}
			}
		}
	}
}

const data = `
{
    "glossary": {
        "title": "example glossary",
		"GlossDiv": {
            "title": "S",
			"GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": ["GML", "XML"]
                    },
					"GlossSee": "markup"
                }
            }
        }
    }
}
`

var largeObj = createLargeObj()

func createLargeObj() *SampleLarge {
	var obj SampleLarge
	_ = json.Unmarshal([]byte(data), &obj)
	return &obj
}
