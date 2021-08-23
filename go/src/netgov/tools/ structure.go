package tools

// Telephone work order structure

type ResponseContext struct {
	Total int `json:"total"`
	Rows  []struct {
		FKint       int    `json:"FKint"`
		SAttachPath string `json:"SAttachPath"`
		SBDeptName  string `json:"SBDeptName"`
		SDeptName   string `json:"SDeptName"`
		BkLimit     string `json:"bkLimit"`
		DealName    string `json:"dealName"`
		MyPkVal     int    `json:"myPkVal"`
		MyTwfNO     string `json:"myTwfNO"`
		MyTwfTopic  string `json:"myTwfTopic"`
		OpDeal      int    `json:"opDeal"`
		SdDate      string `json:"sdDate"`
		SendRemark  string `json:"sendRemark"`
		SeqName     string `json:"seqName"`
		SetProID    int    `json:"setProId"`
		SortName    string `json:"sortName"`
		WfID        int    `json:"wfId"`
	} `json:"rows"`
}

// Send order choice time

type OrderTime struct {
	Assignment struct {
		NextVertexID       int           `json:"next_vertex_id"`
		NextReviewerIds    []interface{} `json:"next_reviewer_ids"`
		DurationThresholds `json:"duration_thresholds"`
		CarbonCopyUserIds  []interface{} `json:"carbon_copy_user_ids"`
		Esignature         string        `json:"esignature"`
		TransferToUserID   interface{}   `json:"transfer_to_user_id"`
		Operation          string        `json:"operation"`
	} `json:"assignment"`
}

type DurationThresholds []struct {
	Gid   string `json:"gid"`
	Value string `json:"value"`
}

type Order struct {
	Assignment struct {
		Operation          string `json:"operation"`
		ResponseAttributes struct {
			EntriesAttributes `json:"entries_attributes"`
		} `json:"response_attributes"`
	} `json:"assignment"`
}

type EntriesAttributes []struct {
	FieldID  int         `json:"field_id"`
	Value    interface{} `json:"value"`
	OptionID interface{} `json:"option_id,omitempty"`
}

type OptionID []struct {
	Gid   string
	Id    int
	Value string
}
