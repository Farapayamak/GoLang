package farapayamak



type RestResponse struct {
	Value			string
	RetStatus		int
	StrRetStatus	string
}

type SmartRestResponse struct {
	ReqStatus		string
	Message			string
	Result			[]SmartRestResult
}

type SmartRestResult struct {
	Mobile		string
	ID			int64
}


type SendSMSRestModel struct {
	To		string
	From	string
	Text	string
}

type GetDeliveries2RestModel struct { 
	RecId	int		`json:"recID"`
}

type GetMessagesRestModel struct {
	Location	int
	From		string
	Index		int
	Count		int
}

type BaseServiceNumberRestModel struct {
	Text		string
	To			string
	BodyId		int
}

type SendSmartSMSRestModel struct {
	To				string
	Text			string
	From			string
	FromSupportOne	string
	FromSupportTwo	string
}

type SendMultipleSmartSMSRestModel struct {
	To				[]string
	Text			[]string
	From			string
	FromSupportOne	string
	FromSupportTwo	string
}

type GetSmartDeliveries2RestModel struct {
	Id				int64
}

type GetSmartDeliveriesRestModel struct {
	Ids				[]int64
}