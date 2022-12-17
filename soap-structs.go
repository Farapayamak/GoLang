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