package messages

type Message struct {
	From    string
	To      []string
	CC      []string
	BCC     []string
	ReplyTo string

	Subject string
	Content struct {
		Plain *string
		HTML  *string
	}
	Attachments []Attachment
}

type Attachment struct {
	Content     []byte
	Type        string
	Filename    string
	Disposition string
	ContentID   string
}
