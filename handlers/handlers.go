package handlers

import (
	"fiber_sample/data"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var tracer oteltrace.Tracer

// ReadData handles reading all data
func ReadData(ctx *fiber.Ctx) error {
	tracer = otel.Tracer("fiber test server")
	startTime := time.Now()
	_, span := tracer.Start(ctx.UserContext(), "getUser")
	defer span.End()

	// Fiber-specific methods to get request data
	method := ctx.Method()
	scheme := ctx.Protocol()
	statusCode := fiber.StatusOK
	host := ctx.Hostname()
	url := ctx.OriginalURL()

	// Set span status
	span.SetStatus(codes.Ok, "")

	contentLengthStr := ctx.Get(fiber.HeaderContentLength)
	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		contentLength = 0
	}

	// Use semantic conventions for common attributes
	span.SetAttributes(
		semconv.HTTPMethodKey.String(method),
		semconv.HTTPSchemeKey.String(scheme),
		semconv.HTTPStatusCodeKey.Int(statusCode),
		semconv.HTTPTargetKey.String(ctx.Path()),
		semconv.HTTPURLKey.String(url),
		semconv.HTTPHostKey.String(host),
		semconv.NetHostPortKey.String(ctx.Port()),
		semconv.HTTPUserAgentKey.String(ctx.Get(fiber.HeaderUserAgent)),
		semconv.HTTPRequestContentLengthKey.Int(contentLength),
		semconv.NetPeerIPKey.String(ctx.IP()),
	)

	// Custom attributes
	span.SetAttributes(
		attribute.String("created_at", startTime.Format(time.RFC3339Nano)),
		attribute.Float64("duration_ns", float64(time.Since(startTime).Nanoseconds())),
		attribute.String("referer", ctx.Get(fiber.HeaderReferer)),
		attribute.String("request_type", "Incoming"),
		attribute.String("sdk_type", "go-fiber"),
		attribute.String("service_version", ""),
		attribute.StringSlice("tags", []string{}),
	)

	span.SetAttributes(
		attribute.String("path_params", fmt.Sprintf(`{"id":"%s"}`, ctx.Params("id"))),
		attribute.String("query_params", fmt.Sprintf("%v", ctx.Queries())),
		attribute.String("request_body", "{}"),
		attribute.String("request_headers", fmt.Sprintf("%v", ctx.GetReqHeaders())),
		attribute.String("response_body", "{}"),
		attribute.String("response_headers", "{}"),
	)

	dataService := data.InitData()
	return ctx.JSON(dataService.GetData())
}

// DeleteData handles deleting a record by ID
func DeleteData(ctx *fiber.Ctx) error {
	tracer = otel.Tracer("fiber test server")
	startTime := time.Now()
	_, span := tracer.Start(ctx.UserContext(), "deleteUser")
	defer span.End()

	id, _ := strconv.Atoi(ctx.Params("id"))
	dataService := data.InitData()

	span.SetAttributes(
		attribute.String("created_at", startTime.Format(time.RFC3339Nano)),
		attribute.Float64("duration_ns", float64(time.Since(startTime).Nanoseconds())),
		attribute.String("request_type", "Incoming"),
		attribute.String("id", strconv.Itoa(id)),
	)

	return ctx.JSON(dataService.DeleteData(id))
}

// InsertData handles inserting new data
func InsertData(ctx *fiber.Ctx) error {
	tracer = otel.Tracer("fiber test server")
	startTime := time.Now()
	_, span := tracer.Start(ctx.UserContext(), "insertUser")
	defer span.End()

	dataService := data.InitData()
	user := new(data.UserModel)

	if err := ctx.BodyParser(user); err != nil {
		span.SetStatus(codes.Error, "Body parsing error")
		return err
	}

	users := dataService.InsertData(data.UserModel{
		Name:   user.Name,
		Gender: user.Gender,
		Age:    user.Age,
	})
	// Set span status
	span.SetStatus(codes.Ok, "")

	span.SetAttributes(
		attribute.String("created_at", startTime.Format(time.RFC3339Nano)),
		attribute.Float64("duration_ns", float64(time.Since(startTime).Nanoseconds())),
		attribute.String("request_body", fmt.Sprintf("%v", user)),
	)

	return ctx.JSON(users)
}

// ReadDataById handles fetching a record by ID
func ReadDataById(ctx *fiber.Ctx) error {
	tracer = otel.Tracer("fiber test server")
	startTime := time.Now()
	_, span := tracer.Start(ctx.UserContext(), "getUserById")
	defer span.End()

	id, _ := strconv.Atoi(ctx.Params("id"))
	dataService := data.InitData()

	// Set span status
	span.SetStatus(codes.Ok, "")
	span.SetAttributes(
		attribute.String("created_at", startTime.Format(time.RFC3339Nano)),
		attribute.Float64("duration_ns", float64(time.Since(startTime).Nanoseconds())),
		attribute.String("id", strconv.Itoa(id)),
	)

	return ctx.JSON(dataService.GetDataById(id))
}

// PatchData handles updating a record by ID
func PatchData(ctx *fiber.Ctx) error {
	tracer = otel.Tracer("fiber test server")
	startTime := time.Now()
	_, span := tracer.Start(ctx.UserContext(), "updateUser")
	defer span.End()

	id, _ := strconv.Atoi(ctx.Params("id"))
	dataService := data.InitData()
	user := new(data.UserModel)

	if err := ctx.BodyParser(user); err != nil {
		span.SetStatus(codes.Error, "Body parsing error")
		return err
	}

	users := dataService.UpdateDataById(
		id,
		data.UserModel{
			Name:   user.Name,
			Gender: user.Gender,
			Age:    user.Age,
		},
	)
	// Set span status
	span.SetStatus(codes.Ok, "")
	span.SetAttributes(
		attribute.String("created_at", startTime.Format(time.RFC3339Nano)),
		attribute.Float64("duration_ns", float64(time.Since(startTime).Nanoseconds())),
		attribute.String("request_body", fmt.Sprintf("%v", user)),
	)

	return ctx.JSON(users)
}
