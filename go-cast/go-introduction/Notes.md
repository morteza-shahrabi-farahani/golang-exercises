# Go

## Workspace
Workspace is a folder. Usually all the projects are located in one workspace. By specifying GOPATH environment variable, workspace location is determined.

in workspace folder, we have three subfolders:

* src: For managing your codes (either your codes and imported packages codes that downloaded and installed by you).

* pkg: For saving shared libraries

* bin: For saving executable files of projects (both your projects and installed projects using go install).