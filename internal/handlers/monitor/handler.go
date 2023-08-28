/* 
 *  Copyright 2023 CPDS Author
 *  
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *  
 *       https://www.apache.org/licenses/LICENSE-2.0
 *  
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package monitor

import (
	"cpds/cpds-analyzer/internal/models/monitor"
	cpdserr "cpds/cpds-analyzer/internal/pkg/errors"
	"cpds/cpds-analyzer/internal/pkg/response"
	"cpds/cpds-analyzer/pkg/cpds-analyzer/config"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler interface {
	GetMonitorTargets() gin.HandlerFunc

	GetClusterResource() gin.HandlerFunc

	GetClusterContainerStatus() gin.HandlerFunc

	GetNodeStatus() gin.HandlerFunc

	GetNodeInfo() gin.HandlerFunc

	GetNodeResource() gin.HandlerFunc

	GetNodeContainerStatus() gin.HandlerFunc
}

type handler struct {
	logger   *zap.Logger
	operator monitor.Operator
}

func New(config *config.Config, logger *zap.Logger) Handler {
	return &handler{
		logger:   logger,
		operator: monitor.NewOperator(config.DetectorOptions.Host, config.DetectorOptions.Port),
	}
}

func (h *handler) GetMonitorTargets() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := h.operator.GetMonitorTargets()
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.MONITOR_GET_TARGET_ERROR, err))
			return
		}

		response.HandleOK(ctx, data)
	}
}

func (h *handler) GetClusterResource() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := parseClusterMonitorDataQueryParams(ctx)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.MONITOR_GET_CLUSTER_RESOURCES_ERROR, err))
			return
		}

		data, err := h.operator.GetClusterResource(time.Unix(p.StartTime, 0), time.Unix(p.EndTime, 0), p.StepSecond)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.MONITOR_GET_CLUSTER_RESOURCES_ERROR, err))
			return
		}

		response.HandleOK(ctx, data)
	}
}

func (h *handler) GetClusterContainerStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := parseClusterMonitorDataQueryParams(ctx)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.MONITOR_GET_CLUSTER_RESOURCES_ERROR, err))
			return
		}

		data, err := h.operator.GetClusterContainerStatus(time.Unix(p.StartTime, 0), time.Unix(p.EndTime, 0), p.StepSecond)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.MONITOR_GET_CLUSTER_CONTAINER_STATUS_ERROR, err))
			return
		}

		response.HandleOK(ctx, data)
	}
}

func (h *handler) GetNodeInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		instance := ctx.Query("instance")
		records, err := h.operator.GetNodeInfo(instance)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.MONITOR_GET_NODE_INFO_ERROR, err))
			return
		}

		response.HandleOK(ctx, records)

	}
}

func (h *handler) GetNodeStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		instance := ctx.Query("instance")
		records, err := h.operator.GetNodeStatus(instance)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.MONITOR_GET_NODE_STATUS_ERROR, err))
			return
		}

		response.HandleOK(ctx, records)
	}
}

func (h *handler) GetNodeResource() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := parseNodeMonitorDataQueryParams(ctx)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.MONITOR_GET_NODE_RESOURCES_ERROR, err))
			return
		}

		records, err := h.operator.GetNodeResources(p.Instance, time.Unix(p.StartTime, 0), time.Unix(p.EndTime, 0), p.StepSecond)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.MONITOR_GET_NODE_RESOURCES_ERROR, err))
			return
		}

		response.HandleOK(ctx, records)
	}
}

func (h *handler) GetNodeContainerStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		instance, err := parseInstanceFromParams(ctx)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.MONITOR_GET_CLUSTER_CONTAINER_STATUS_ERROR, err))
			return
		}

		records, err := h.operator.GetNodeContainerStatus(instance)
		if err != nil {
			response.HandleError(ctx, http.StatusInternalServerError, cpdserr.NewError(cpdserr.MONITOR_GET_CLUSTER_CONTAINER_STATUS_ERROR, err))
			return
		}

		response.HandleOK(ctx, records)
	}
}

func parseInstanceFromParams(ctx *gin.Context) (string, error) {
	instance, exist := ctx.GetQuery("instance")
	if !exist {
		return "", errors.New("params instance cannot be empty")
	}
	return instance, nil
}

func parseNodeMonitorDataQueryParams(ctx *gin.Context) (*nodeMonitorDataQueryParams, error) {
	var p nodeMonitorDataQueryParams
	var err error

	p.Instance = ctx.Query("instance")
	if p.Instance == "" {
		return nil, errors.New("instance cannot be empty")
	}

	p.StartTime, err = strconv.ParseInt(ctx.Query("start_time"), 10, 64)
	if err != nil {
		return nil, err
	}

	p.EndTime, err = strconv.ParseInt(ctx.Query("end_time"), 10, 64)
	if err != nil {
		return nil, err
	}

	p.StepSecond, err = strconv.ParseInt(ctx.Query("step"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func parseClusterMonitorDataQueryParams(ctx *gin.Context) (*clusterMonitorQueryParams, error) {
	var p clusterMonitorQueryParams
	var err error

	p.StartTime, err = strconv.ParseInt(ctx.Query("start_time"), 10, 64)
	if err != nil {
		return nil, err
	}

	p.EndTime, err = strconv.ParseInt(ctx.Query("end_time"), 10, 64)
	if err != nil {
		return nil, err
	}

	p.StepSecond, err = strconv.ParseInt(ctx.Query("step"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
