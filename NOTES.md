# Notes

This is a new repo, I'm attempting to reach an MVP. A lot of it is where I want it.
I got lost in the minutia and need to finish up.
We are going to clearly define the goals of this project that will drive the mplementation (and cleanup).
The primary stated goal is a Go implementation of the dev.unmango.cmd API, which is a protobuf specification for process execution.
The API also includes a conversion service for converting protobuf messages into commandline arguments using reflection.
This project implements the specification using two tools, protocmd and argconv.
protocmd implements the command service, at a high level it is capable of executing a v1alpha1.Process.
argconv implements the conversion service, at a high level it can convert a protobuf message into an array of commandline arguments.
