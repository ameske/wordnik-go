package wordnik

const (
	WORD_BASE = "/word.json/"
)

type Word struct {
	ID            int64
	Word          string
	OriginalWord  string
	Suggetsions   []string
	CanonicalForm string
	Vulgar        string
}

type Definition struct {
	ExtendedText     string
	Text             string
	SourceDictionary string
	Citations        []Citation
	Labels           []Label
	Score            float64
	ExampleUses      []ExampleUsage
	AttributionUrl   string
	SeqString        string
	AttributionText  string
	RelatedWords     []Related
	Sequence         string
	Word             string
	Notes            []Note
	TextProns        []TextPron
	PartOfSpeecth    string
}

type Citation struct {
	Cite   string
	Source string
}

type Label struct {
	Text string
	Type string
}

type ExampleUsage struct {
	Text string
}

type Related struct {
	Label1           string
	RelationshipType string
	Label2           string
	Label3           string
	Words            []string
	Gram             string
	Label4           string
}

type Note struct {
	NoteType  string
	AppliesTo []string
	Value     string
	Pos       int
}

type TextPron struct {
	Raw     string
	Seq     int
	RawType string
}
