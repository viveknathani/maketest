#!/usr/bin/env python3

# Build script for maketest
import subprocess
import os

PATH_BREAK = "/"

if os.name == "nt":
    PATH_BREAK = "\\"

if os.path.isdir("../bin") is False:
    subprocess.run("mkdir .." + PATH_BREAK + "bin", shell=True)

compile_main = "g++ .." + PATH_BREAK + "src" + PATH_BREAK + "main.cpp -o .." + PATH_BREAK + "bin" + PATH_BREAK + "maketest"
subprocess.run(compile_main, shell=True)