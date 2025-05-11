package server

import (
	"net/http"

	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"
	kratosStatus "github.com/go-kratos/kratos/v2/transport/http/status"
	"google.golang.org/grpc/status"
)

// 统一返回格式
type httpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 成功统一返回格式
func responseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if v == nil {
		return nil
	}
	if rd, ok := v.(kratosHttp.Redirector); ok {
		url, code := rd.Redirect()
		http.Redirect(w, r, url, code)
		return nil
	}
	response := httpResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    v,
	}
	codec, _ := kratosHttp.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(response)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/"+codec.Name()+"; charset=utf-8")
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// 错误统一返回格式
func responseErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	resp := new(httpResponse)
	se, ok := status.FromError(err)
	if !ok {
		resp = &httpResponse{
			Code:    http.StatusInternalServerError,
			Message: "unknown error",
			Data:    nil,
		}
	} else {
		resp = &httpResponse{
			Code:    kratosStatus.FromGRPCCode(se.Code()),
			Message: se.Message(),
			Data:    nil,
		}
	}
	codec, _ := kratosHttp.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name()+"; charset=utf-8")
	w.WriteHeader(resp.Code)
	_, _ = w.Write(body)
}
