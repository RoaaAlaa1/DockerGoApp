# Go RPC Chatroom – Dockerized Application

This project contains a simple Go-based RPC chat application consisting of a server and a client.  
The server handles incoming RPC messages and stores chat history, while the client allows a user to interact with the server in real time.

---

## 🚀 Project Structure

.
├── server.go
├── client.go
├── Dockerfile (server)
├── Dockerfile.client (client)
└── README.md


## 🧩 How the Application Works

### **Server**
- Listens on TCP port **1234**
- Receives messages through `ChatService.SendMessage`
- Stores messages in memory and returns updated chat history
- Runs inside a Docker container

### **Client**
- Connects to the server using RPC over TCP
- Sends messages typed by the user
- Displays returned chat history
- Runs inside a Docker container

---

## 🐳 Docker Files:

### **Server Dockerfile:**
-Builds a Docker image for the server
-Installing dependencies
-Setting up the environment
-Running the server inside a container.

## **Docker Hub Image Link:**
>> https://hub.docker.com/r/roaaalaa1/chatroom-server

## To run this project:
### 1. Run the server from Docker Hub above:
docker pull username/rpc-chat-server:new
docker run -p 1234:1234 roo/chatroom-server:latest

### 2. Run client:
go run client.go

