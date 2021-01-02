module main

go 1.13

replace github/maruimm/myGoLearning/tcpServer => ../tcpServer

replace github/maruimm/myGoLearning/tcpClient => ../tcpClient

replace github/maruimm/myGoLearning/selfProto => ../selfProto

require (
	github/maruimm/myGoLearning/selfProto v0.0.0-00010101000000-000000000000 // indirect
	github/maruimm/myGoLearning/tcpClient v0.0.0-00010101000000-000000000000 // indirect
	github/maruimm/myGoLearning/tcpServer v0.0.0-00010101000000-000000000000
)
