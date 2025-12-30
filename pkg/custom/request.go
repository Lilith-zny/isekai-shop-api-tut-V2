package custom

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	EchoRequest interface {
		// ล้อเลียนเหมือนตัว echo
		Bind(obj any) error
	}

	customEchoRequest struct {
		ctx       echo.Context
		validator *validator.Validate
	}
)

// pattern singleton
var (
	once              sync.Once
	validatorInstance *validator.Validate
)

func NewCustomEchoRequest(echoRequest echo.Context) EchoRequest {
	once.Do(func() {
		validatorInstance = validator.New()
	})

	return &customEchoRequest{
		ctx: echoRequest,
		// validator: validator.New(), แก้เรื่องที่ตอน new ใหม่ทุกรอบ
		validator: validatorInstance,
	}
}

func (r *customEchoRequest) Bind(obj any) error {
	if err := r.ctx.Bind(obj); err != nil {
		return err
	}

	if err := r.validator.Struct(obj); err != nil {
		return err
	}

	return nil
}
