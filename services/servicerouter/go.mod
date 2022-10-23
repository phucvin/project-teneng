module github.com/phucvin/project-teneng/services/servicerouter

go 1.16

require (
	github.com/phucvin/project-teneng/services/servicerouter/servicerouter/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.50.1
)

replace github.com/phucvin/project-teneng/services/servicerouter/servicerouter/proto => ./proto
