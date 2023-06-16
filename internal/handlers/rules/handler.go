package rules

import (
	"cpds/cpds-analyzer/internal/models/rules"
	cpdserr "cpds/cpds-analyzer/internal/pkg/errors"
	"cpds/cpds-analyzer/internal/pkg/response"
	"cpds/cpds-analyzer/pkg/cpds-analyzer/config"
	prometheusutil "cpds/cpds-analyzer/pkg/utils/prometheus"
	stringutil "cpds/cpds-analyzer/pkg/utils/string"
	timeutil "cpds/cpds-analyzer/pkg/utils/time"
	"errors"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler interface {
	Get() gin.HandlerFunc

	Create() gin.HandlerFunc

	Delete() gin.HandlerFunc

	Update() gin.HandlerFunc
}

type handler struct {
	config   *config.Config
	logger   *zap.Logger
	operator rules.Operator
}

func New(config *config.Config, logger *zap.Logger, db *gorm.DB) Handler {
	return &handler{
		logger:   logger,
		operator: rules.NewOperator(config.DetectorOptions.Host, config.DetectorOptions.Port, db),
	}
}

func (h *handler) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		opt, err := parseGetParams(ctx)
		if err != nil {
			response.HandleError(ctx, http.StatusBadRequest, cpdserr.NewError(cpdserr.RULES_GET_ERROR, err))
			return
		}
		records, err := h.operator.GetRules(opt.filter, opt.sortField, opt.sortOrder, opt.pageNo, opt.pageSize)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.RULES_GET_ERROR, err))
			return
		}

		responseData := &getResponse{
			Records:   records,
			PageNo:    opt.pageNo,
			PageSize:  opt.pageSize,
			PageTotal: h.operator.GetTotalPages(opt.filter),
		}
		response.HandleOK(ctx, responseData)
	}
}

func (h *handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.RULES_CREATE_ERROR, err))
			return
		}

		if err := validateRule(req.Rule); err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.RULES_CREATE_ERROR, err))
			return
		}

		if err := h.operator.CreateRule(req.Rule); err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.RULES_CREATE_ERROR, err))
			return
		}

		if err := h.operator.SendRuleUpdatedRequset(); err != nil {
			response.HandleError(
				ctx,
				http.StatusInternalServerError,
				cpdserr.NewError(cpdserr.DETECTOR_ERROR, fmt.Errorf("create rule success but unable to start analysis: %s", err)),
			)
			return
		}

		response.HandleOK(ctx, nil)
	}
}

func (h *handler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req updateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.RULES_UPDATE_ERROR, err))
			return
		}

		if err := validateRule(req.Rule); err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.RULES_UPDATE_ERROR, err))
			return
		}

		if err := h.operator.UpdateRule(req.Rule); err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.RULES_UPDATE_ERROR, err))
			return
		}

		if err := h.operator.SendRuleUpdatedRequset(); err != nil {
			response.HandleError(
				ctx,
				http.StatusInternalServerError,
				cpdserr.NewError(cpdserr.DETECTOR_ERROR, fmt.Errorf("update rule success but unable to start analysis: %s", err)),
			)
			return
		}

		response.HandleOK(ctx, nil)
	}
}

func (h *handler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req deleteRequest
		if err := ctx.BindJSON(&req); err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.RULES_DELETE_ERROR, err))
			return
		}

		if err := h.operator.DeleteRuleByID(req.ID); err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.RULES_DELETE_ERROR, err))
			return
		}

		if err := h.operator.SendRuleUpdatedRequset(); err != nil {
			response.HandleError(
				ctx,
				http.StatusInternalServerError,
				cpdserr.NewError(cpdserr.DETECTOR_ERROR, fmt.Errorf("delete rule success but unable to stop analysis: %s", err)),
			)
			return
		}

		response.HandleOK(ctx, nil)
	}
}

func parseGetParams(p *gin.Context) (*getOptions, error) {
	pageNo, err := strconv.Atoi(p.DefaultQuery("page_no", "1"))
	if err != nil {
		return nil, fmt.Errorf("invalid params")
	}

	pageSize, err := strconv.Atoi(p.DefaultQuery("page_size", "10"))
	if err != nil {
		return nil, fmt.Errorf("invalid params")
	}

	return &getOptions{
		filter:    p.Query("filter"),
		sortField: p.DefaultQuery("sort_field", "name"),
		sortOrder: p.DefaultQuery("sort_order", "asc"),
		pageNo:    pageNo,
		pageSize:  pageSize,
	}, nil
}

func validateRule(rule *rules.Rule) error {
	// check name
	re := regexp.MustCompile("^[A-Za-z0-9_]{1,20}$")
	if !re.MatchString(rule.Name) {
		return errors.New("invalid rule name")
	}

	// check expression
	if !prometheusutil.IsExprValid(rule.Expression) {
		return errors.New("invalid expression")
	}

	//check rule
	if (rule.SubhealthConditionType != "" && math.IsNaN(rule.SubhealthThresholds)) ||
		(rule.FaultConditionType != "" && math.IsNaN(rule.FaultThresholds)) {
		return errors.New("invalid rule")
	}

	// check severity
	if !stringutil.IsStringInArray(rule.Severity, []string{"warning", "error", "critical"}) {
		return errors.New("invalid severity")
	}

	// check duration
	if !timeutil.IsValidDuration(rule.Duration) {
		return errors.New("invalid duration")
	}

	return nil
}
