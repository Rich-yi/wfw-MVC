#!/bin/bash

go run service/getArea/main.go &
go run service/getImage/main.go &
go run service/register/main.go &
go run main.go