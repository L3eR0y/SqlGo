module internal/server

go 1.23.2

replace internal/handlers => ../handlers

replace internal/mdb => ../mdb

require internal/handlers v0.0.0-00010101000000-000000000000

require internal/mdb v0.0.0-00010101000000-000000000000 // indirect
