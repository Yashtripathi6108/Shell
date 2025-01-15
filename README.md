# Go Shell

![Go Shell](https://img.shields.io/badge/Go-Shell-green?style=for-the-badge&logo=go)

Go Shell is a **lightweight** and **fast** shell built using Go. It supports basic built-in commands like `exit`, `echo`, `pwd`, `cd`, `ls`, and `type`. It also allows you to execute external commands available in the system's `PATH`. The shell is simple, intuitive, and designed for performance, leveraging Go's concurrency model.

---

## Features

- **Built-in commands**: `exit`, `echo`, `pwd`, `cd`, `ls`, and `type`.
- **Execute external commands** available in the system's `PATH`.
- **Simple and intuitive interface**.
- **Lightweight** and **fast**, powered by Go’s concurrency model.

---

## Installation

### Prerequisites

To run the Go Shell, you’ll need:

- [Go](https://golang.org/doc/install) (version 1.19 or later)
- [Docker](https://docs.docker.com/get-docker/) (optional, for running inside a container)

---

## Running Locally

### Step 1: Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/Yashtripathi6108/Shell.git
cd Shell
```

### Step 2: Run Shell

Now run following command:

```bash
go run main.go
```

## Running with Docker

If you prefer to run the Go Shell using Docker, follow these steps.

### Step 1: Build the Docker Image

Make sure Docker is installed on your system. If not, follow the installation guide [here](https://docs.docker.com/get-docker/).

Once Docker is installed, build the Docker image from the project directory:

```bash
docker build -t Shell .
```

### Step 2: Run the Docker Image in interactive mode

```bash
docker run -it Shell
