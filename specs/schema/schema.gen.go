// Package schema provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package schema

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// Список станций по адресу
type AddressStations struct {
	Access    int       `json:"access"`
	Address   string    `json:"address"`
	Icon      *string   `json:"icon,omitempty"`
	IconType  *string   `json:"icon_type,omitempty"`
	Id        string    `json:"id"`
	Latitude  float32   `json:"latitude"`
	Longitude float32   `json:"longitude"`
	Name      string    `json:"name"`
	Score     *float32  `json:"score,omitempty"`
	Stations  []Station `json:"stations"`
}

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Сущность разъема.
type Outlet struct {
	Connector int      `json:"connector"`
	Id        string   `json:"id"`
	Kilowatts *float32 `json:"kilowatts"`
	Power     int      `json:"power"`
}

// ResponseLocations defines model for ResponseLocations.
type ResponseLocations struct {
	// Результат запроса.
	Data []AddressStations `json:"data"`
}

// Сущность станции.
type Station struct {
	Id      string   `json:"id"`
	Outlets []Outlet `json:"outlets"`
}

// GetChargingStationsParams defines parameters for GetChargingStations.
type GetChargingStationsParams struct {
	LatitudeMin  *float32 `form:"latitudeMin,omitempty" json:"latitudeMin,omitempty"`
	LongitudeMin *float32 `form:"longitudeMin,omitempty" json:"longitudeMin,omitempty"`
	LatitudeMax  *float32 `form:"latitudeMax,omitempty" json:"latitudeMax,omitempty"`
	LongitudeMax *float32 `form:"longitudeMax,omitempty" json:"longitudeMax,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Проверка сервиса
	// (GET /healthz)
	HealthCheck(w http.ResponseWriter, r *http.Request)
	// Получение списка зарядных станций в пределах координат
	// (GET /v1/stations)
	GetChargingStations(w http.ResponseWriter, r *http.Request, params GetChargingStationsParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// HealthCheck operation middleware
func (siw *ServerInterfaceWrapper) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.HealthCheck(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetChargingStations operation middleware
func (siw *ServerInterfaceWrapper) GetChargingStations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetChargingStationsParams

	// ------------- Optional query parameter "latitudeMin" -------------
	if paramValue := r.URL.Query().Get("latitudeMin"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "latitudeMin", r.URL.Query(), &params.LatitudeMin)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "latitudeMin", Err: err})
		return
	}

	// ------------- Optional query parameter "longitudeMin" -------------
	if paramValue := r.URL.Query().Get("longitudeMin"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "longitudeMin", r.URL.Query(), &params.LongitudeMin)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "longitudeMin", Err: err})
		return
	}

	// ------------- Optional query parameter "latitudeMax" -------------
	if paramValue := r.URL.Query().Get("latitudeMax"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "latitudeMax", r.URL.Query(), &params.LatitudeMax)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "latitudeMax", Err: err})
		return
	}

	// ------------- Optional query parameter "longitudeMax" -------------
	if paramValue := r.URL.Query().Get("longitudeMax"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "longitudeMax", r.URL.Query(), &params.LongitudeMax)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "longitudeMax", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetChargingStations(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/healthz", wrapper.HealthCheck)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/stations", wrapper.GetChargingStations)
	})

	return r
}
