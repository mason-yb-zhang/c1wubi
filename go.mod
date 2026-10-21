module c1wubi

go 1.26.0

require c1device v0.0.0-00010101000000-000000000000

require (
	golang.org/x/image v0.45.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
)

replace c1device => ../c1device
