package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
	"nginx/dto"
	"nginx/pkg/responser"
	"nginx/pkg/utils"
	"nginx/pkg/validator"
)

// Create godoc
// @Summary domain create
// @Description create new domain
// @Tags domain
// @Accept json
// @Produce json
// @Param DomainCreateRequest body dto.CreateDomainRequest true "necessary item for create new"
// @Success 200
// @Router /api/v1/domain [post]
func (h AdminHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		span, spannedCtx := opentracing.StartSpanFromContext(ctx, "Create[http.handler]")
		defer span.Finish()
		var req dto.CreateDomainRequest
		if err := c.Bind(&req); err != nil {
			return responser.NewErrorBuilder().SetStatusCode(http.StatusUnprocessableEntity).SetMessage(err.Error()).SetDetail("body", err.Error()).Build().Respond(c)
		}
		if errRess := validator.ValidateRequestDto(ctx, req); errRess != nil {
			return c.JSON(http.StatusInternalServerError, errRess.Respond(c))
		}

		addr, err := h.uc.Create(spannedCtx, req.MapToModel())
		if err != nil {
			log.Println("err", err)
			return responser.NewErrorBuilder().SetStatusCode(409).SetMessage(err.Error()).Build().Respond(c)
		}

		return c.JSON(http.StatusOK, addr)
	}
}

// Delete godoc
// @Summary domain delete
// @Description delete domain with id
// @Tags domain
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/domain/:id [delete]
func (h AdminHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		span, spannedCtx := opentracing.StartSpanFromContext(ctx, "Delete[http.handler]")
		defer span.Finish()
		var req dto.DeleteDomainRequest
		if err := c.Bind(&req); err != nil {
			return responser.NewErrorBuilder().SetStatusCode(422).SetMessage(err.Error()).SetDetail("body", err.Error()).Build().Respond(c)
		}
		if errRess := validator.ValidateRequestDto(ctx, req); errRess != nil {
			return c.JSON(http.StatusInternalServerError, errRess.Respond(c))
		}

		err := h.uc.Delete(spannedCtx, req.ID)
		if err != nil {
			return responser.NewErrorBuilder().SetMessage(err.Error()).Build().Respond(c)
		}

		return c.JSON(http.StatusOK, responser.NewResponse(201, "success", nil))
	}
}

// GetAll godoc
// @Summary get all domain
// @Description get all domains
// @Tags domain
// @Accept json
// @Produce json
// @Success 200 {object} []dto.Domain
// @Router /api/v1/domain [get]
func (h AdminHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		span, spannedCtx := opentracing.StartSpanFromContext(ctx, "GetAll[http.handler]")
		defer span.Finish()
		var resp dto.Domain
		domains, err := h.uc.GetAll(spannedCtx)
		if err != nil {
			return responser.NewErrorBuilder().SetMessage(err.Error()).Build().Respond(c)
		}

		d := utils.Map(domains, resp.MapFromEntity)

		return c.JSON(http.StatusOK, dto.GetDomainListResponse{Domains: d})
	}
}

// CreateCacheZone
// @Summary create zone
// @Description create new zone
// @Tags zone
// @Accept json
// @Produce json
// @Param CreateCacheZone body dto.AddCacheZoneRequest true "necessary item for create new"
// @Success 200
// @Router /api/v1/zone [post]
func (h AdminHandler) CreateCacheZone() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		span, spannedCtx := opentracing.StartSpanFromContext(ctx, "CreateCacheZone[http.handler]")
		defer span.Finish()
		var req dto.AddCacheZoneRequest
		if err := c.Bind(&req); err != nil {
			return responser.NewErrorBuilder().SetStatusCode(http.StatusUnprocessableEntity).SetMessage(err.Error()).Respond(c)
		}
		if errRess := validator.ValidateRequestDto(ctx, req); errRess != nil {
			return errRess.Respond(c)
		}

		m, err := h.uc.CreateZone(spannedCtx, req.MapToModel())
		if err != nil {
			return responser.NewErrorBuilder().SetMessage(err.Error()).SetStatusCode(http.StatusInternalServerError).Respond(c)
		}

		return c.JSON(http.StatusOK, m)
	}
}

// GetCacheZone godoc
// @Summary get all zones
// @Description get all zone
// @Tags zone
// @Accept json
// @Produce json
// @Success 200 {object} []models.CacheZone
// @Router /api/v1/zone [get]
func (h AdminHandler) GetCacheZone() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		span, spannedCtx := opentracing.StartSpanFromContext(ctx, "CreateCacheZone[http.handler]")
		defer span.Finish()
		zones, err := h.uc.GetAllZone(spannedCtx)
		if err != nil {
			return responser.NewErrorBuilder().SetMessage(err.Error()).SetStatusCode(http.StatusInternalServerError).Respond(c)
		}
		return c.JSON(http.StatusOK, zones)
	}
}
