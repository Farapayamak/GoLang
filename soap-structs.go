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