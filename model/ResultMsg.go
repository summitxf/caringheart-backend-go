package model

type ResultMsg struct {
	Code string `json:"code" bson:"code"`
	Msg  string `json:"msg" bson:"msg"`
}

func RsMsg(resultIn string) ResultMsg {
	switch resultIn {
	case "0", "Success":
		return ResultMsg{Code: "0", Msg: "Success"}
	case "1", "Results":
		return ResultMsg{Code: "1", Msg: "No Results."}
	case "2", "Request":
		return ResultMsg{Code: "1", Msg: "Request ResultMsg."}
	case "3", "Delete":
		return ResultMsg{Code: "1", Msg: "Can't delete document."}
	case "4", "Formdata":
		return ResultMsg{Code: "1", Msg: "Incomplete form data or wrong params."}
	case "5", "Database":
		return ResultMsg{Code: "1", Msg: "Database ResultMsg."}
	case "6", "Duplicated":
		return ResultMsg{Code: "1", Msg: "Duplicated field in Database."}
	case "7", "Login":
		return ResultMsg{Code: "1", Msg: "Login wrong user or password."}
	case "8", "Session":
		return ResultMsg{Code: "1", Msg: "Session not match."}
	default:
		return ResultMsg{Code: "1", Msg: "ResultMsg."}
	}
}
