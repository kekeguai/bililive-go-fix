package missevan

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bililive-go/bililive-go/src/live"
)

func TestParseRoomInfoResponse(t *testing.T) {
	for _, test := range []struct {
		name       string
		statusCode int
		body       string
		wantErr    error
	}{
		{name: "正常房间", statusCode: http.StatusOK, body: `{"code":0}`},
		{name: "HTTP 404", statusCode: http.StatusNotFound, body: `not found`, wantErr: live.ErrRoomNotExist},
		{name: "房间不存在", statusCode: http.StatusOK, body: `{"code":500030004,"info":"无法找到该聊天室"}`, wantErr: live.ErrRoomNotExist},
		{name: "房间被封禁", statusCode: http.StatusOK, body: `{"code":500030013,"info":"直播间被封禁"}`, wantErr: live.ErrRoomBanned},
		{name: "风控", statusCode: http.StatusOK, body: `{"code":100010017,"info":"403"}`},
		{name: "HTTP 418 风控", statusCode: http.StatusTeapot, body: `blocked`},
		{name: "服务器错误", statusCode: http.StatusBadGateway, body: `error`},
		{name: "无效 JSON", statusCode: http.StatusOK, body: `not json`},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := parseRoomInfoResponse(test.statusCode, []byte(test.body))
			if test.wantErr != nil {
				if !errors.Is(err, test.wantErr) {
					t.Fatalf("error = %v, want errors.Is(_, %v)", err, test.wantErr)
				}
				return
			}
			if test.statusCode == http.StatusOK && test.body == `{"code":0}` {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatal("expected an error")
			}
			if errors.Is(err, live.ErrRoomNotExist) || errors.Is(err, live.ErrRoomBanned) {
				t.Fatalf("transient/API error was misclassified as a permanent room error: %v", err)
			}
		})
	}
}
