package api

import "github.com/gofiber/fiber/v2"

//	@Summary		k8s readiness probe
//	@Description	a readiness probe for k8s
//	@Tags			ReadinessProbe
//	@Router			/healthz [get]
//	@Success		200
func ReadinessProbe(ctx *fiber.Ctx) error {
	return ctx.Send(nil)
}
