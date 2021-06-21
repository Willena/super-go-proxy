package main

import (
	"github.com/willena/super-go-mod-proxy/errors"
	"github.com/willena/super-go-mod-proxy/fetchMethods"
	"github.com/willena/super-go-mod-proxy/runner"
	"github.com/willena/super-go-mod-proxy/types"
	"go.uber.org/zap"
	"net/http"
)

func ZipVersionHandler(writer http.ResponseWriter, request *http.Request) {
	module, err := moduleFromRequest(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(errors.GenerateError(err))
	}

	logger.Debug("Getting zip file for module ", zap.String("moduleFromRequest", module.Path), zap.String("version", module.Version.String()))

	err = runner.NewRunner(&types.RunnerContext{
		GoModule:    module,
		FetchMethod: &fetchMethods.GoProxy{Url: mainConfig.General.DefaultRelayProxy},
		Action:      types.ActionGetModuleZip,
	}, pluginsInstances).Run(writer)

	if err != nil {
		logger.Error("Error while zip file for module", zap.String("module", module.Path), zap.Error(err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
