package transporthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"

	"go-postgres/common/chttp"

	"github.com/gorilla/mux"

	transport "go-postgres/services/user/transport"
)

func CommonHTTPRequestEncoder(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func CommonHTTPResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return chttp.EncodeResponse(ctx, w, response)
}

func _Decode_Add_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.AddRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_Get_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		_param string
	)
	var ok bool
	_vars := mux.Vars(r)
	_param, ok = _vars["id"]
	if !ok {
		return nil, errors.New("param id not found")
	}
	id, err := strconv.ParseInt(_param, 10, 64)
	if err != nil {
		return nil, errors.New("param id not valid")
	}
	return &transport.GetRequest{Id: id}, nil
}

// func _Decode_List_Request(_ context.Context, r *http.Request) (interface{}, error) {
// 	return &transport.ListRequest{}, nil
// }

// func _Decode_Update_Request(_ context.Context, r *http.Request) (interface{}, error) {
// 	var req transport.UpdateRequest
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	return &req, err
// }

func _Decode_Add_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.AddResponse
	err := chttp.DecodeResponse(ctx, r, &resp)
	return &resp, err
}

func _Decode_Get_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetResponse
	err := chttp.DecodeResponse(ctx, r, &resp)
	return &resp, err
}

// func _Decode_List_Response(ctx context.Context, r *http.Response) (interface{}, error) {
// 	var resp transport.ListResponse
// 	err := chttp.DecodeResponse(ctx, r, &resp)
// 	return &resp, err
// }

// func _Decode_Update_Response(ctx context.Context, r *http.Response) (interface{}, error) {
// 	var resp transport.UpdateResponse
// 	err := chttp.DecodeResponse(ctx, r, &resp)
// 	return &resp, err
// }

func _Encode_Add_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "brand")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_Get_Request(ctx context.Context, r *http.Request, request interface{}) error {
	req := request.(*transport.GetRequest)
	r.URL.Path = path.Join(r.URL.Path, "brand",
		string(req.Id),
	)
	fmt.Println("r.URL.Path", r.URL.Path)
	return nil
}

func _Encode_List_Request(ctx context.Context, r *http.Request, request interface{}) error {
	//req := request.(*transport.ListRequest)
	r.URL.Path = path.Join(r.URL.Path, "brand")
	return nil
}

func _Encode_Update_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "brand")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_Add_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_Get_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_List_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_Update_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}
