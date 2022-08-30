package utils

type Response struct {
	Code int         `json:"code"`
	Desc string      `json:"desc"`
	Data interface{} `json:"data"`
}

func (r Response) Success(data interface{}) Response {
	return Response{
		Code: 0,
		Desc: "success",
		Data: data,
	}
}

func (r Response) Failure() Response {
	return Response{
		Code: 1,
		Desc: "failure",
		Data: nil,
	}
}

func (r Response) WithDesc(desc string) Response {
	r.Desc = desc
	return r
}
