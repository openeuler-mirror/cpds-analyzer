package prometheus

import (
	"cpds/cpds-analyzer/internal/models/prometheus"
	cpdserr "cpds/cpds-analyzer/internal/pkg/errors"
	"cpds/cpds-analyzer/internal/pkg/response"
	"cpds/cpds-analyzer/pkg/cpds-analyzer/config"
	prometheusutil "cpds/cpds-analyzer/pkg/utils/prometheus"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler interface {
	Query() gin.HandlerFunc

	QueryRange() gin.HandlerFunc

	QueryValidate() gin.HandlerFunc
}

type handler struct {
	logger   *zap.Logger
	operator prometheus.Operator
}

func New(logger *zap.Logger, config *config.Config) Handler {
	return &handler{
		logger:   logger,
		operator: prometheus.NewOperator(config.DetectorOptions.Host, config.DetectorOptions.Port),
	}
}

func (h handler) Query() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := parseQueryParams(ctx)
		if err != nil {
			response.HandleError(ctx, http.StatusBadRequest, cpdserr.NewError(cpdserr.PROMETHEUS_QUERY_ERROR, err))
			return
		}

		metric, err := h.operator.Query(p.Query, p.Time)
		if err != nil {
			response.HandleError(ctx, http.StatusBadRequest, cpdserr.NewError(cpdserr.PROMETHEUS_QUERY_ERROR, err))
			return
		}

		response.HandleOK(ctx, metric)
	}
}

func (h handler) QueryRange() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := parseQueryRangeParams(ctx)
		if err != nil {
			response.HandleError(ctx, http.StatusBadRequest, cpdserr.NewError(cpdserr.PROMETHEUS_QUERY_RANGE_ERROR, err))
			return
		}

		metric, err := h.operator.QueryRange(p.Query, p.StartTime, p.EndTime, p.StepSecond)
		if err != nil {
			response.HandleError(ctx, http.StatusBadRequest, cpdserr.NewError(cpdserr.PROMETHEUS_QUERY_ERROR, err))
			return
		}

		response.HandleOK(ctx, metric)
	}
}

func (h handler) QueryValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		queryExpr := ctx.Query("query")
		if queryExpr == "" {
			response.HandleError(ctx, http.StatusBadRequest, cpdserr.NewError(cpdserr.PROMETHEUS_QUERY_VALIDATE_ERROR, errors.New("query cannot be empty")))
			return
		}

		if !prometheusutil.IsExprValid(queryExpr) {
			response.HandleError(ctx, http.StatusBadRequest, cpdserr.NewError(cpdserr.PROMETHEUS_QUERY_VALIDATE_ERROR, errors.New("illegal query expression")))
			return
		}

		response.HandleOK(ctx, nil)
	}
}

func parseQueryParams(ctx *gin.Context) (*queryParams, error) {
	var p queryParams

	p.Query = ctx.Query("query")
	if p.Query == "" {
		return nil, errors.New("query cannot be empty")
	}

	timeStr := ctx.Query("time")
	if timeStr == "" {
		p.Time = time.Now().Unix()
	} else {
		var err error
		p.Time, err = strconv.ParseInt(timeStr, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &p, nil
}

func parseQueryRangeParams(ctx *gin.Context) (*queryRangeParams, error) {
	query := ctx.Request.URL.Query().Get("query")
	if query == "" {
		return nil, errors.New("query cannot be empty")
	}

	startTime, err := strconv.ParseInt(ctx.Query("start_time"), 10, 64)
	if err != nil {
		return nil, err
	}

	endTime, err := strconv.ParseInt(ctx.Query("end_time"), 10, 64)
	if err != nil {
		return nil, err
	}

	step, err := strconv.ParseInt(ctx.Query("step"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &queryRangeParams{
		Query:      query,
		StartTime:  startTime,
		EndTime:    endTime,
		StepSecond: step,
	}, nil
}
