package main

//go:generate go run ./internal/tools/gen-const -p convert -o convert/constant.go -use64flg
//go:generate go run ./internal/tools/gen-const -p convert_test -o convert/constant_test.go

//go:generate go run ./internal/tools/gen-int -t int8 -o convert/int8.go
//go:generate go run ./internal/tools/gen-int -t int16 -o convert/int16.go
//go:generate go run ./internal/tools/gen-int -t int32 -o convert/int32.go
//go:generate go run ./internal/tools/gen-int -t int64 -o convert/int64.go

//go:generate go run ./internal/tools/gen-int -t uint8 -o convert/uint8.go
//go:generate go run ./internal/tools/gen-int -t uint16 -o convert/uint16.go
//go:generate go run ./internal/tools/gen-int -t uint32 -o convert/uint32.go
//go:generate go run ./internal/tools/gen-int -t uint64 -o convert/uint64.go
