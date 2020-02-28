# Reaktor Junior 2020 | Pre-assignment
Here is my implementation of the Reaktor Junior 2020 pre-assignment. Live version at https://niklasprojects.me

> *On Debian and Ubuntu systems, there is a file /var/lib/dpkg/status, that holds information about software packages that the system knows about. Write a small program in a programming language of your choice that exposes some key information about packages in the file via an HTML interface.*

Read more about the assignment [**here**](https://www.reaktor.com/junior-dev-assignment/) (07.02.2020)

---

My intentions with this project was *(1)* to learn how to setup a web server and run a live app on it, *(2)* get started with Golang and *(3)* refresh my skills with React.

In brief, the project is divided into three parts. The server is running a [**React app**](https://github.com/nikunicke/reaktor2019-frontend/tree/9f8f250ef6f5fa2d51a1936c58dbcb3ae44da844)
that requests data from /api/dpkg/. The data is [**parsed and updated**](https://github.com/nikunicke/jsonify/tree/93d9aafdea6f5caab4830c988fd8fde867dcac4b)
internally whenever a new package is installed on the system.

To run and serve on your system
```console
	git clone https://github.com/nikunicke/reaktor-junior-2020.git reaktor-app
	cd reaktor-app
	git submodule init
```

In order to make this run on your system, update your path in main.go and recompile. You need to have Golang installed for this.
The frontend build is included here and does not
need to be rebuilt.

```console
	go build -o "binary_name"
```