// Copyright 2023 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package specscanner provides a generic way to scan a swagger spec.
// It provides a Handler interface that can be implemented to receive callbacks
package specscanner

import (
	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"go.uber.org/zap"
)

const (
	opGet     = "GET"
	opPost    = "POST"
	opPut     = "PUT"
	opPatch   = "PATCH"
	opDelete  = "DELETE"
	opHead    = "HEAD"
	opOptions = "OPTIONS"
)

// Handler is the interface that should be implemented to receive callbacks when scanning
// a swagger spec.
type Handler interface {
	StartScan(apiSpec *spec.Swagger) error
	EndScan() error
}

// PathHandler is an optional interface that can be implemented to receive callbacks when
// the scanner encounters a new path.
type PathHandler interface {
	StartPath(path string, pathSpec spec.PathItem) error
	EndPath() error
}

// OperationHandler is an optional interface that can be implemented to receive callbacks when
// the scanner encounters a new operation.
type OperationHandler interface {
	StartOp(op string, opSpec *spec.Operation) error
	EndOp() error
}

// ResponseHandler is an optional interface that can be implemented to receive callbacks when
// the scanner encounters a new response.
type ResponseHandler interface {
	StartResponse(statusCode int, getSpec *spec.Response) error
	EndResponse() error
}

// ParamHandler is an optional interface that can be implemented to receive callbacks when
// the scanner encounters a new parameter.
type ParamHandler interface {
	StartParam(getSpec *spec.Parameter) error
	EndParam() error
}

type scanner struct {
	logger *zap.Logger
}

// Scan will scan the given spec data and call the implemented methods on the given handler
// specData should be the raw bytes of the swagger spec JSON
func Scan(specData []byte, handler Handler, logger *zap.Logger) error {
	if logger == nil {
		logger = zap.NewNop()
	}
	scanner := &scanner{logger: logger}
	return scanner.scan(specData, handler)
}

func (s *scanner) scan(specData []byte, handler Handler) error {
	doc, err := loads.Analyzed(specData, "")
	if err != nil {
		return err
	}

	err = analysis.Flatten(analysis.FlattenOpts{
		Spec:   doc.Analyzer,
		Expand: true,
	})
	if err != nil {
		return err
	}

	s.logger.Info("starting scan")
	if err := handler.StartScan(doc.Spec()); err != nil {
		return err
	}

	for path, pathItem := range doc.Analyzer.AllPaths() {
		if ph, ok := handler.(PathHandler); ok {
			s.logger.Info("scanning path", zap.String("path", path))
			if err := ph.StartPath(path, pathItem); err != nil {
				return err
			}

			if err = s.scanOperations(pathItem, handler); err != nil {
				return err
			}

			s.logger.Info("finished scanning path", zap.String("path", path))
			if err := ph.EndPath(); err != nil {
				return err
			}
		}
	}
	s.logger.Info("finished scan")
	return handler.EndScan()
}

func (s *scanner) scanOperations(pathItem spec.PathItem, handler Handler) error {
	if oh, ok := handler.(OperationHandler); ok {
		ops := map[string]*spec.Operation{
			opGet:     pathItem.Get,
			opPost:    pathItem.Post,
			opPut:     pathItem.Put,
			opPatch:   pathItem.Patch,
			opDelete:  pathItem.Delete,
			opHead:    pathItem.Head,
			opOptions: pathItem.Options,
		}
		for op, opSpec := range ops {
			if opSpec != nil {
				s.logger.Info("scanning operation", zap.String("op", op))
				if err := oh.StartOp(op, opSpec); err != nil {
					return err
				}

				if err := s.scanParameters(opSpec, handler); err != nil {
					return err
				}

				if err := s.scanResponses(opSpec, handler); err != nil {
					return err
				}

				s.logger.Info("finished scanning operation", zap.String("op", op))
				if err := oh.EndOp(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *scanner) scanParameters(op *spec.Operation, handler Handler) error {
	if ph, ok := handler.(ParamHandler); ok {
		for _, param := range op.Parameters {
			s.logger.Info("scanning param", zap.String("param", param.Name))
			if err := ph.StartParam(&param); err != nil {
				return err
			}

			s.logger.Info("finished scanning param", zap.String("param", param.Name))
			if err := ph.EndParam(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *scanner) scanResponses(op *spec.Operation, handler Handler) error {
	if rh, ok := handler.(ResponseHandler); ok {
		for statusCode, resp := range op.Responses.StatusCodeResponses {
			s.logger.Info("scanning response", zap.String("response", resp.Description))
			if err := rh.StartResponse(statusCode, &resp); err != nil {
				return err
			}

			s.logger.Info("finished scanning response")
			if err := rh.EndResponse(); err != nil {
				return err
			}
		}
	}
	return nil
}
