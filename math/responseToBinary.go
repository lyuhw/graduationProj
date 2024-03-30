package math

type Response struct {
	Data string `json:"data"`
}

type Binary []byte

func (r *Response) ToBinary() Binary {
	return Binary(r.Data)
}

func ResponseToBinary(re string) Binary {
	response := &Response{
		Data: re,
	}

	binaryData := response.ToBinary()
	return binaryData
}
