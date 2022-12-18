package farapayamak


type SendSimpleSMS2SoapModel struct {
	To		string
	From	string
	Text 	string
	IsFlash bool
}

type SendSimpleSMSSoapModel struct {
	To		[]string
	From	string
	Text	string
	IsFlash bool
}

type GetDeliveriesSoapModel struct {
	RecIds	[]int64
}

type GetDeliveries3SoapModel struct {
	RecId	[]string
}

type GetSmsPriceSoapModel struct {
	IrancellCount	int
	MtnCount		int
	From			string
	Text			string
}

type SendByBaseNumberSoapModel struct {
	Text	[]string
	To		string
	BodyId	int
}

type SendByBaseNumber2SoapModel struct {
	Text	string
	To		string
	BodyId	int
}

type SendByBaseNumber3SoapModel struct {
	Text	string
	To		string
}

type SendSmsSoapModel struct {
	To		[]string
	From	string
	Text	string
	IsFlash bool
	Udh		string
	RecId	[]int64
	Status	[]byte
}

type SendSms2SoapModel struct {
	To			[]string
	From		string
	Text		string
	IsFlash 	bool
	Udh			string
	RecId		[]int64
	Status		[]byte
	FilterId	int
}

type SendMultipleSMSSoapModel struct {
	To			[]string
	From		string
	Text		[]string
	IsFlash 	bool
	Udh			string
	RecId		[]int64
	Status		[]byte
}

type SendMultipleSMS2SoapModel struct {
	To			[]string
	From		[]string
	Text		[]string
	IsFlash 	bool
	Udh			string
	RecId		[]int64
	Status		[]byte
}

// receive
type ChangeMessageIsReadSoapModel struct {
	MsgIds	string
}

type GetInboxCountSoapModel struct {
	IsRead	bool
}

type GetLatestReceiveMsgSoapModel struct {
	Sender		string
	Receiver	string
}

type GetMessagesSoapModel struct {
	Location	int
	From		string
	Index		int
	Count		int
}

type GetMessagesAfterIDSoapModel struct {
	Location	int
	From		string
	Count		int
	MsgId		int
}

type GetMessagesFilterByDateSoapModel struct {
	Location	int
	From		string
	Index		int
	Count		int
	DateFrom	string
	DateTo		string
	IsRead		bool
}

type GetMessagesReceptionsSoapModel struct {
	MsgId		int
	FromRows	int
}

type GetMessagesWithChangeIsReadSoapModel struct {
	Location		int
	From			string
	Index			int
	Count			int
	IsRead			bool
	changeIsRead	bool
}

type RemoveMessagesSoapModel struct {
	Location	string
	MsgIds		string
}