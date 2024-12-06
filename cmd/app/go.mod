module app

go 1.23.2

replace internal/server => ../../internal/server

replace internal/handlers => ../../internal/handlers

replace internal/mdb => ../../internal/mdb

require internal/server v0.0.0-00010101000000-000000000000

require (
	internal/handlers v0.0.0-00010101000000-000000000000 // indirect
	internal/mdb v0.0.0-00010101000000-000000000000 // indirect
)
