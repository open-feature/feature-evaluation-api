package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/open-feature/feature-evaluation-api/models"
	"github.com/open-feature/feature-evaluation-api/restapi/operations/version"
)

var Version string

type versionImpl struct{}

func NewVersionHandler() version.VersionHandler {
	return &versionImpl{}
}

func (impl *versionImpl) Handle(params version.VersionParams) middleware.Responder {
	responseVal := &models.Version{
		APIVersion: Version,
	}

	return version.NewVersionOK().WithPayload(responseVal)
}
