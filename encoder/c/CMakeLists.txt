project(ClientLib)

set(SOURCES
	main.c
	config/config.c
	config/config.h
	encoder/encoder.c
	encoder/encoder.h
	internal/output.c
	internal/output.h
	internal/mask.c
	internal/mask.h)

add_executable(${CMAKE_PROJECT_NAME}_run ${SOURCES})