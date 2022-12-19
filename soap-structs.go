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

// users
type AddUserSoapModel struct {
	ProductId			int
	Descriptions		string
	MobileNumber		string
	EmailAddress		string
	NationalCode		string
	Name				string
	Family				string
	Corporation			string
	Phone				string
	Fax					string
	Address				string
	PostalCode			string
	CertificateNumber	string
}

type AddUserWithLocationSoapModel struct {
	ProductId			int
	Descriptions		string
	MobileNumber		string
	EmailAddress		string
	NationalCode		string
	Name				string
	Family				string
	Corporation			string
	Phone				string
	Fax					string
	Address				string
	PostalCode			string
	CertificateNumber	string
	Country				int
	Province			int
	City				int
}

type AddUserWithMobileNumberSoapModel struct {
	ProductId		int
	MobileNumber	string
	FirstName		string
	LastName		string
	Email			string
}

type AddUserWithMobileNumber2SoapModel struct {
	ProductId		int
	MobileNumber	string
	FirstName		string
	LastName		string
	Email			string
	TargetUsername	string
}

type AddUserWithUserNameAndPassSoapModel struct {
	TargetUsername		string
	TargetUserPassword	string
	ProductId			int
	Descriptions		string
	MobileNumber		string
	EmailAddress		string
	NationalCode		string
	Name				string
	Family				string
	Corporation			string
	Phone				string
	Fax					string
	Address				string
	PostalCode			string
	CertificateNumber	string
}

type ChangeUserCreditSoapModel struct {
	Amount			int
	Description		string
	TargetUsername	string
	GetTax			bool
}

type DeductUserCreditSoapModel struct {
	Amount		int
	Description	string
}

type ForgotPasswordSoapModel struct {
	MobileNumber	string
	EmailAddress	string
	TargetUsername	string
}

type GetCitiesSoapModel struct {
	ProvinceId	int
}

type GetUserBasePriceSoapModel struct {
	TargetUsername	string
}

type GetUserCreditSoapModel struct {
	TargetUsername	string
}

type GetUserDetailsSoapModel struct {
	TargetUsername	string
}

type GetUserIsExistSoapModel struct {
	TargetUsername	string
}

type GetUserTransactionsSoapModel struct {
	TargetUsername	string
	CreditType		string
	DateFrom		string
	DateTo			string
	Keyword			string
}

type GetUserWalletTransactionSoapModel struct {
	DateFrom		string
	DateTo			string
	Count			int
	StartIndex		int
	PayType			string
	PayLoc			string
}

type RemoveUserSoapModel struct {
	TargetUsername	string
}

// voice

type SendBulkSpeechTextSoapModel struct {
	Title		string
	Body		string
	Receivers	string
	DateToSend	string
	RepeatCount	int
}

type SendBulkVoiceSMSSoapModel struct {
	Title		string
	VoiceFileId int
	Receivers	string
	DateToSend	string
	RepeatCount	int
}

type UploadVoiceFileSoapModel struct {
	Title				string
	Base64StringFile	string
}

// Contacts

type AddContactSoapModel struct {
	GroupIds			string
	FirstName			string
	LastName			string
	Nickname			string
	Corporation			string
	MobileNumber		string
	Phone				string
	Fax					string
	Birthdate			string
	Email				string
	Gender				byte
	Province			int
	City				int
	Address				string
	PostalCode			string
	Additionaldate		string
	Additionaltext		string
	Descriptions		string
}

type AddContactEventsSoapModel struct {
	ContactId		int
	EventName		string
	EventType		byte
	EventDate		string
}

type AddGroupSoapModel struct {
	GroupName		string
	Descriptions	string
	ShowToChilds	bool
}

type ChangeContactSoapModel struct {
	ContactId			int
	MobileNumber		string
	FirstName			string
	LastName			string
	Nickname			string
	Corporation			string
	Phone				string
	Fax					string
	Email				string
	Gender				byte
	Province			int
	City				int
	Address				string
	PostalCode			string
	Additionaltext		string
	Descriptions		string
	ContactStatus		int
}

type ChangeGroupSoapModel struct {
	GroupId			int
	GroupName		string
	Descriptions	string
	ShowToChilds	bool
	GroupStatus		byte
}

type CheckMobileExistInContactSoapModel struct {
	MobileNumber	string
}

type GetContactEventsSoapModel struct {
	ContactId	int
}

type GetContactsSoapModel struct {
	GroupId		int
	Keyword		string
	From		int
	Count		int
}

type GetContactsByIDSoapModel struct {
	ContactId	int
	status		int
}

type MergeGroupsSoapModel struct {
	OriginGroupId		int
	DestinationGroupId	int
	DeleteOriginGroup	bool
}

type RemoveContactSoapModel struct {
	MobileNumber	string
}

type RemoveContactByContactIDSoapModel struct {
	ContactId	int
}

type RemoveGroupSoapModel struct {
	GroupId		int
}

// Schedule

type AddNewMultipleScheduleSoapModel struct {
	To					[]string
	From				string
	Text				[]string
	IsFlash				bool
	ScheduleDateTime	[]string
	Period				string
}

type AddNewUsanceSoapModel struct {
	To						string
	From					string
	Text					string
	IsFlash					bool
	ScheduleStartDateTime	string
	Countrepeat				int
	ScheduleEndDateTime		string
	PeriodType				string
}

type AddScheduleSoapModel struct {
	To						string
	From					string
	Text					string
	IsFlash					bool
	ScheduleDateTime		string
	Period					string
}

type GetScheduleDetailsSoapModel struct {
	ScheduleId	int
}

type GetScheduleStatusSoapModel struct {
	ScheduleId	int
}

type RemoveScheduleSoapModel struct {
	ScheduleId	int
}

// Bulk

type AddNumberBulkSoapModel struct {
	From		string
	Title		string
	Message 	string
	Receivers	string
	DateToSend	string
}

type BulkReceptionSoapModel struct {
	BulkId			int
	MaximumRows		int
	StartRowIndex	int
}

type BulkReceptionCountSoapModel struct {
	BulkId		int
}

type GetBulkDeliveriesSoapModel struct {
	RecIds		[]int64
}

type GetBulkDeliveries2SoapModel struct {
	RecId		string
}

type GetBulkDetailsSoapModel struct {
	BulkId		int
}