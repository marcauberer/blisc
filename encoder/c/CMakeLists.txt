project(ClientLib)
set(BINARY ${CMAKE_PROJECT_NAME})

set(SOURCES
	main.c
	config/config.c
	config/config.h)

add_executable(${BINARY}_run ${SOURCES})