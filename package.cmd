@echo off
cd s
gox -osarch="windows/386 windows/amd64 darwin/386 darwin/amd64 linux/386 linux/amd64 linux/arm" -output="MinecraftSourceRecipeImageMaker_{{.OS}}_{{.Arch}}"
cd ../
