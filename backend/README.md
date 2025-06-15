# Blockchain Public Sector API

This project provides a Go-based API for interacting with smart contracts deployed on the Ethereum Sepolia test network. The API handles two use cases:
- Diploma verification
- State document verification

## Project Structure

### Backend Components
- `internal/diplomamanager/`: Diploma verification implementation
- `internal/staterecords/`: State document verification implementation
- `internal/models/`: Data models and structures
- `cmd/`: Application entry points

## Backend API

The backend is built with Go and uses the Gin framework for the REST API. It provides endpoints for:

- Verifying diploma credentials
- Storing diploma credentials
- Verifying state documents
- Verifying state documents

## Features

- **Verify Diploma**: Checks if a diploma hash exists on the blockchain and retrieves the associated issuer's DID
- **Verify State Document**: Checks if a state document hash exists on the blockchain and retrieves the associated issuer's DID
- **Smart Contract Integration**: Interacts with smart contracts deployed on the Sepolia test network
- **Firebase Integration**: Uses Firebase Admin SDK for authentication
- **Blockchain Integration**: Connects to Ethereum Sepolia test network for verification
- **Postman Collection**: Includes API testing collection (postman.json)

## API Endpoints

#### Health Check
- `GET /health` - Check API status

#### Diplomas
- `POST /diplomas` - Store a new diploma credential
  - Request body:
    ```json
    {
      "diplomaData": {
        // Credential object
      },
      "issuerDID": "string",
      "privateKey": "string"
    }
    ```
  - Response: `{ "txHash": "string", "hash": "string" }`

- `GET /diplomas/:hash` - Verify a diploma credential
  - Response: `{ "verified": boolean, "issuerDID": "string" }`

#### State Documents
- `POST /documents` - Store a new state document
  - Request body:
    ```json
    {
      "stateInfo": "string",
      "privateKey": "string"
    }
    ```
  - Response: `{ "txHash": "string", "recordId": "string", "hash": "string" }`

- `GET /documents/:recordId?hash=<hash>` - Verify a state document
  - Response: `{ "verified": boolean, "recordId": "string" }`

## Running the Application

### Prerequisites

- Go 1.19 or higher
- Access to an Ethereum node (Infura or Alchemy) connected to the Sepolia test network
- Firebase Admin SDK credentials
### Setup

1. Create a `serviceAccountKey.json` file with your Firebase Admin SDK credentials

2. Set environment variables:
   ```
   FIREBASE_PROJECT_ID=your_project_id
   FIREBASE_PRIVATE_KEY=your_private_key
   FIREBASE_CLIENT_EMAIL=your_client_email
   ```

3. Run the server:
   ```bash
   cd backend
   go mod download
   go run main.go
   ```

The API will be available at http://localhost:8080