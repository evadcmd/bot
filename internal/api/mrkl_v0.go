package api

import (
	"fmt"

	"github.com/evadcmd/bot/internal/mrkl"
	"github.com/evadcmd/bot/pkg/dto"
	"github.com/gofiber/fiber/v2"
)

//	@Summary		chat API
//	@Description	a MRKL API
//	@Tags			/api/v0
//	@Router			/api/v0/mrkl [post]
//	@Param			payload	body	dto.Q	true	"Q"
//	@Produce		json
//	@Success		200
func MRKL(ctx *fiber.Ctx) error {
	var q dto.Q
	if err := ctx.BodyParser(&q); err != nil {
		return fmt.Errorf("failed to deserialize request body to dto.Q struct")
	}
	res, err := mrkl.Induce(ctx.Context(), q.Content)
	if err != nil {
		return fmt.Errorf("failed on calling the internal MRKL package: %w", err)
	}
	return ctx.JSON(res)
}
