package farapayamak



type RestResponse struct {
	Value			string
	RetStatus		int
	StrRetStatus	string
}


type SendSMSRestModel struct {
	To		string
	From	string
	Text	string
}