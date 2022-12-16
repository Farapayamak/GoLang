package farapayamak



type RestResponse struct {
	Value			string
	RetStatus		int
	StrRetStatus	string
}


type SendSMSRestModel struct {
	to		string
	from	string
	text	string
}