package types

/*
TODO::
1. Generate Service (Ex: UserService)
2. Generate Entities for service (Ex: Users, UserRoles, UsersHasRoles, ...)
3. Generate Protocols for service (Ex: Grpc, Http, ...)
4. Generate Repository for database (Ex: Postgres, Mysql, ...)
5. Implement Repository
6. Implement Service
*/

type IGoat interface {
	GenerateService(name string)
	GenerateGrpc()
	GenerateHttp()
	GenerateWebSocket()
	GenerateEntity(name string)
	GenerateProtobuf(name string)
	GenerateProtocol(name string)
	GenerateRepo(name string)
	GenerateDocker(name string)
	GenerateEntryPoint()
}
