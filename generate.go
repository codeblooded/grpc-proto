package grpc_proto

// Build grpc/core package
//go:generate protoc -I . --go_out=paths=source_relative,grpc:./genproto grpc/core/stats.proto

// Build grpc/testing package
//go:generate protoc -I . --go_out=paths=source_relative,grpc:./genproto grpc/testing/benchmark_service.proto
//go:generate protoc -I . --go_out=paths=source_relative,grpc:./genproto grpc/testing/control.proto
//go:generate protoc -I . --go_out=paths=source_relative,grpc:./genproto grpc/testing/messages.proto
//go:generate protoc -I . --go_out=paths=source_relative,grpc:./genproto grpc/testing/payloads.proto
//go:generate protoc -I . --go_out=paths=source_relative,grpc:./genproto grpc/testing/report_qps_scenario_service.proto
//go:generate protoc -I . --go_out=paths=source_relative,grpc:./genproto grpc/testing/stats.proto
//go:generate protoc -I . --go_out=paths=source_relative,grpc:./genproto grpc/testing/worker_service.proto

