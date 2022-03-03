project(ClientLib)
set(BINARY ${CMAKE_PROJECT_NAME})

set(SOURCES
	main.c
	config/config.c
	config/config.h
	encoder/encoder.c
	encoder/encoder.h
	internal/output.c
	internal/output.h)

add_executable(${BINARY}_run ${SOURCES})