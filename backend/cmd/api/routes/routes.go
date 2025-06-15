package routes

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"net/http"

	"backend/internal/diplomamanager"
	"backend/internal/models"
	"backend/internal/staterecords"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	sepoliaRPC               = "https://sepolia.infura.io/v3/719b60d26f514a07a9ea38294ea3e1cc"
	diplomaContractAddress   = "0x23c8eC31EA95A47b37bDa3D6a004693883e97976"
	stateRecordsContractAddr = "0x3566d868931d5945A63AA2207435c4A60400D3f1"
	sepoliaChainID           = 11155111
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	client, err := ethclient.Dial(sepoliaRPC)
	if err != nil {
		panic(err)
	}

	// Initialize DiplomaManager contract
	diplomaContract, err := diplomamanager.NewDiplomamanager(common.HexToAddress(diplomaContractAddress), client)
	if err != nil {
		panic(err)
	}

	// Initialize StateRecords contract
	stateContract, err := staterecords.NewStaterecords(common.HexToAddress(stateRecordsContractAddr), client)
	if err != nil {
		panic(err)
	}

	// Diploma endpoints
	r.POST("/diplomas", func(c *gin.Context) {
		var req struct {
			Credential models.Credential `json:"diplomaData" binding:"required"`
			IssuerDID  string            `json:"issuerDID" binding:"required"`
			PrivateKey string            `json:"privateKey" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		credentialJSON, err := json.Marshal(req.Credential)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to serialize credential"})
			return
		}
		hash := sha256.Sum256(credentialJSON)
		key, err := crypto.HexToECDSA(req.PrivateKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid private key"})
			return
		}
		auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(sepoliaChainID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
			return
		}
		tx, err := diplomaContract.StoreDiploma(auth, hash, req.IssuerDID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"txHash": tx.Hash().Hex(), "hash": hex.EncodeToString(hash[:])})
	})

	r.GET("/diplomas/:hash", func(c *gin.Context) {
		hashHex := c.Param("hash")
		hashBytes, err := hex.DecodeString(hashHex)
		if err != nil || len(hashBytes) != 32 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hash format"})
			return
		}
		var hash [32]byte
		copy(hash[:], hashBytes)
		isStored, err := diplomaContract.IsDiplomaStored(nil, hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !isStored {
			c.JSON(http.StatusNotFound, gin.H{"error": "Diploma not found"})
			return
		}
		issuerDID, err := diplomaContract.GetIssuerDID(nil, hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"verified": true, "issuerDID": issuerDID})
	})

	r.POST("/documents", func(c *gin.Context) {
		// Bind the request to the StateVerificationRequest struct
		var req models.StateVerificationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Convert StateInfo to JSON for hashing
		stateInfoJSON, err := json.Marshal(req.StateInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize state info"})
			return
		}

		// Hash the state information
		hash := sha256.Sum256(stateInfoJSON)

		// Generate a random recordId
		randomBytes := make([]byte, 32)
		if _, err := rand.Read(randomBytes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate recordId"})
			return
		}
		recordId := sha256.Sum256(randomBytes)

		// Convert the private key from hex to ECDSA
		key, err := crypto.HexToECDSA(req.PrivateKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid private key"})
			return
		}

		// Create a transactor with the private key and chain ID
		auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(sepoliaChainID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
			return
		}

		// Call the issueRecord function on the smart contract
		tx, err := stateContract.IssueRecord(auth, recordId, hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Respond with the transaction hash and the recordId
		c.JSON(http.StatusOK, gin.H{
			"txHash":   tx.Hash().Hex(),
			"recordId": hex.EncodeToString(recordId[:]),
			"hash":     hex.EncodeToString(hash[:]),
		})
	})

	r.GET("/documents/:recordId", func(c *gin.Context) {
		recordIdHex := c.Param("recordId")
		recordIdBytes, err := hex.DecodeString(recordIdHex)
		if err != nil || len(recordIdBytes) != 32 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recordId format"})
			return
		}

		var recordId [32]byte
		copy(recordId[:], recordIdBytes)

		// Get the hash from the request query parameter
		hashHex := c.Query("hash")
		if hashHex == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Hash parameter is required"})
			return
		}

		hashBytes, err := hex.DecodeString(hashHex)
		if err != nil || len(hashBytes) != 32 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hash format"})
			return
		}

		var hash [32]byte
		copy(hash[:], hashBytes)

		// Verify the record
		isVerified, err := stateContract.VerifyRecord(nil, recordId, hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !isVerified {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found or hash mismatch"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"verified": true,
			"recordId": recordIdHex,
			"hash":     hashHex,
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}
