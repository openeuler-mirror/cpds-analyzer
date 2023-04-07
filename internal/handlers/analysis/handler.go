package analysis

import (
	"cpds/cpds-analyzer/internal/models/analysis"
	cpdserr "cpds/cpds-analyzer/internal/pkg/errors"
	"cpds/cpds-analyzer/internal/pkg/response"
	"cpds/cpds-analyzer/pkg/cpds-analyzer/config"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler interface {
	GetResult() gin.HandlerFunc

	DeleteResult() gin.HandlerFunc

	GetRawData() gin.HandlerFunc
}

type handler struct {
	logger   *zap.Logger
	operator analysis.Operator
}

func New(logger *zap.Logger, db *gorm.DB, config *config.Config) Handler {
	return &handler{
		logger:   logger,
		operator: analysis.NewOperator(config.DetectorOptions.Host, config.DetectorOptions.Port, db),
	}
}

func (h *handler) GetResult() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		opt, err := parseGetResultParams(ctx)
		if err != nil {
			response.HandleError(ctx, http.StatusBadRequest, cpdserr.NewError(cpdserr.ANALYSIS_GET_RESULT_ERROR, err))
			return
		}

		records, err := h.operator.GetAnalysisResult(opt.filter, opt.sortField, opt.sortOrder, opt.pageNo, opt.pageSize)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.ANALYSIS_GET_RESULT_ERROR, err))
			return
		}

		responseData := &getResultResponse{
			Records:   records,
			PageNo:    opt.pageNo,
			PageSize:  opt.pageSize,
			PageTotal: h.operator.GetTotalPages(opt.pageSize),
		}
		response.HandleOK(ctx, responseData)
	}
}

func (h *handler) DeleteResult() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req deleteResultRequest
		if err := ctx.BindJSON(&req); err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.ANALYSIS_DELETE_RESULT_ERROR, err))
			return
		}

		if err := h.operator.DeleteAnalysisResultByID(req.ID); err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.ANALYSIS_DELETE_RESULT_ERROR, err))
			return
		}

		response.HandleOK(ctx, nil)
	}
}

func (h *handler) GetRawData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req getRawDataRequset
		if err := ctx.BindJSON(&req); err != nil {
			response.HandleError(ctx, http.StatusBadRequest, cpdserr.NewError(cpdserr.ANALYSIS_GET_RAW_DATA_ERROR, err))
			return
		}

		rawData, err := h.operator.GetRawData(req.ID)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.ANALYSIS_GET_RAW_DATA_ERROR, err))
			return
		}

		responseData := &getRawDataResponse{
			Records: rawData,
		}

		response.HandleOK(ctx, responseData)
	}
}

func parseGetResultParams(ctx *gin.Context) (*getResultOptions, error) {
	pageNo, err := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	if err != nil {
		return nil, fmt.Errorf("invalid params")
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	if err != nil {
		return nil, fmt.Errorf("invalid params")
	}

	return &getResultOptions{
		filter:    ctx.Query("filter"),
		sortField: ctx.DefaultQuery("sort_field", "rule_name"),
		sortOrder: ctx.DefaultQuery("sort_order", "asc"),
		pageNo:    pageNo,
		pageSize:  pageSize,
	}, nil
}
