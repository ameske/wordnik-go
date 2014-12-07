package wordnik

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

var testWord = `
[
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [
      {
        "type": "fld",
        "text": "(Zoöl.)"
      }
    ],
    "citations": [],
    "word": "dog",
    "text": "A quadruped of the genus Canis, esp. the domestic dog (Canis familiaris).",
    "sequence": "0",
    "score": 0,
    "partOfSpeech": "noun",
    "notes": [
      {
        "pos": 0,
        "value": "The dog is distinguished above all others of the inferior animals for intelligence, docility, and attachment to man. There are numerous carefully bred varieties, as the akita, beagle, bloodhound, bulldog, coachdog, collie, Danish dog, foxhound, greyhound, mastiff, pointer, poodle, St. Bernard, setter, spaniel, spitz, terrier, German shepherd, pit bull, Chihuahua, etc. There are also many mixed breeds, and partially domesticated varieties, as well as wild dogs, like the dingo and dhole. (See these names in the Vocabulary.)",
        "appliesTo": []
      }
    ],
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English",
    "seqString": "1."
  },
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [],
    "citations": [
      {
        "source": " 2 Kings viii. 13 (Rev. Ver. )",
        "cite": "What is thy servant, which is but a dog, that he should do this great thing?"
      }
    ],
    "word": "dog",
    "text": "A mean, worthless fellow; a wretch.",
    "sequence": "1",
    "score": 0,
    "partOfSpeech": "noun",
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English",
    "seqString": "2."
  },
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [
      {
        "text": "a sly dog; a lazy dog."
      }
    ],
    "relatedWords": [],
    "labels": [
      {
        "type": "mark",
        "text": "colloq."
      }
    ],
    "citations": [],
    "word": "dog",
    "text": "A fellow; -- used humorously or contemptuously",
    "sequence": "2",
    "score": 0,
    "partOfSpeech": "noun",
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English",
    "seqString": "3."
  },
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [
      {
        "type": "fld",
        "text": "(Astron.)"
      }
    ],
    "citations": [],
    "word": "dog",
    "text": "One of the two constellations, Canis Major and Canis Minor, or the Greater Dog and the Lesser Dog. Canis Major contains the Dog Star (Sirius).",
    "sequence": "3",
    "score": 0,
    "partOfSpeech": "noun",
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English",
    "seqString": "4."
  },
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [],
    "citations": [],
    "word": "dog",
    "text": "An iron for holding wood in a fireplace; a firedog; an andiron.",
    "sequence": "4",
    "score": 0,
    "partOfSpeech": "noun",
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English",
    "seqString": "5."
  },
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [
      {
        "type": "fld",
        "text": "(Mech.)"
      }
    ],
    "citations": [],
    "word": "dog",
    "sequence": "5",
    "score": 0,
    "partOfSpeech": "noun",
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English",
    "seqString": "6."
  },
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [],
    "citations": [],
    "word": "dog",
    "text": "A grappling iron, with a claw or claws, for fastening into wood or other heavy articles, for the purpose of raising or moving them.",
    "sequence": "6",
    "score": 0,
    "partOfSpeech": "noun",
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English",
    "seqString": "6 (a)"
  },
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [],
    "citations": [],
    "word": "dog",
    "text": "An iron with fangs fastening a log in a saw pit, or on the carriage of a sawmill.",
    "sequence": "7",
    "score": 0,
    "partOfSpeech": "noun",
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English",
    "seqString": "6 (b)"
  },
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [],
    "citations": [],
    "word": "dog",
    "text": "A piece in machinery acting as a catch or clutch; especially, the carrier of a lathe, also, an adjustable stop to change motion, as in a machine tool.",
    "sequence": "8",
    "score": 0,
    "partOfSpeech": "noun",
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English",
    "seqString": "6 (c)"
  },
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [
      {
        "type": "mark",
        "text": "slang"
      }
    ],
    "citations": [],
    "word": "dog",
    "text": "an ugly or crude person, especially an ugly woman.",
    "sequence": "9",
    "score": 0,
    "partOfSpeech": "noun",
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English",
    "seqString": "7."
  },
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [
      {
        "type": "mark",
        "text": "slang"
      }
    ],
    "citations": [],
    "word": "dog",
    "text": "a hot dog.",
    "sequence": "10",
    "score": 0,
    "partOfSpeech": "noun",
    "notes": [
      {
        "pos": 0,
        "value": "☞ Dog is used adjectively or in composition, commonly in the sense of relating to, or characteristic of, a dog. It is also used to denote a male; as, dog fox or g-fox, a male fox; dog otter or dog-otter, dog wolf, etc.; -- also to denote a thing of cheap or mean quality; as, dog Latin.",
        "appliesTo": []
      }
    ],
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English",
    "seqString": "8."
  },
  {
    "textProns": [],
    "sourceDictionary": "gcide",
    "exampleUses": [],
    "relatedWords": [],
    "labels": [],
    "citations": [
      {
        "source": " Pope.",
        "cite": "I have been pursued, dogged, and waylaid."
      },
      {
        "source": "Burroughs.",
        "cite": "Your sins will dog you, pursue you."
      },
      {
        "source": " South.",
        "cite": "Eager ill-bred petitioners, who do not so properly supplicate as hunt the person whom they address to, dogging him from place to place, till they even extort an answer to their rude requests."
      }
    ],
    "word": "dog",
    "text": "To hunt or track like a hound; to follow insidiously or indefatigably; to chase with a dog or dogs; to worry, as if by dogs; to hound with importunity.",
    "sequence": "0",
    "score": 0,
    "partOfSpeech": "verb-transitive",
    "attributionText": "from the GNU version of the Collaborative International Dictionary of English"
  }
]`

func TestDefinitionByDictionaryCanonical(t *testing.T) {
	def, err := testClient.DefinitionByDictionaryCanonical("dog", "webster")
	if err != nil {
		t.Errorf(err.Error())
	}

	var testDef []Definition
	err = json.NewDecoder(bytes.NewBuffer([]byte(testWord))).Decode(&testDef)
	if err != nil {
		t.Fatalf("Couldn't decode testDef: %s", err.Error())
	}

	if len(testDef) != len(def) {
		t.Error("Definition response does not match expected: Sizes are not equal")
	}

	for i := 0; i < len(testDef); i++ {
		if !reflect.DeepEqual(def[i], testDef[i]) {
			t.Errorf("Definition response does not match expected: Array elements are not equal")
		}

	}

}
