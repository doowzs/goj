# GOJ

Go Online Judge generater - generate OJ problems with Go!

## What is GOJ?

GOJ is an online judge problem generator, written in Golang. It helps to easily package problem into FPS (Free Problem Set) XML format. 

## What is GOJ designed for?

GOJ is a reinvented wheel to ease the pain of creating problems on HUST OJ (because the system really ~~sucks~~ is not great (both the UX and the backend)). 

## Do I have other choices?

Of course! If you are a .NET or a C# lover, don't miss [the generator made by StardustDL](https://github.com/StardustDL/generator-oj-problem). However, for GNU/Linux users, using .NET is not an ideal choice. 

## How can someone install GOJ?

It's easy as long as you are using AMD64 (x86_64) architecture! There are four methods:

1. Source: clone and get dependencies with `go get ./...` and then run it with `go run main.go`. However, you need a stable network to download the Go dependencies.
2. Debian/Ubuntu package: install with `dpkg -i goj_x.y.z-0_amd64.deb` and then use with `goj`.
3. Linux universal pre-built executable: extract the files and then use it directly in shell (or add it to `$PATH`).
4. Windows pre-build executable: extract the files and then use it directly in CMD/PowerShell or something else (or add it to `%PATH%`).
