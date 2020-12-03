package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sendgrid/internal/echo/context"
	"testing"
)

func TestHealtz(t *testing.T) {
	type args struct {
		c *context.MyContext
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	res := httptest.NewRecorder()
	c := &context.MyContext{Context: e.NewContext(req, res)}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "正常",
			args:    args{c: c},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Healtz(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Healtz() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.Equal(t, http.StatusOK, c.Response().Status)
				assert.Equal(t, "hc", res.Body.String())
			}

		})
	}
}
