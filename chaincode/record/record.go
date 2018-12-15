package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Session struct {
	NIP string `json:"nip"`
	Treatment string `json:"treatment"`
	Medication string `json:"medication"`
}

type MedicalRecord struct {
	NIK string `json:"nik"`
	Record []Session `json:"session"`
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	if function == "addSession" {
		return AddSession(stub, args)
	} else if function == "queryByNIK" {
		return QueryByNIK(stub, args)
	} else if function == "seed" {
		return Seed(stub)
	}

	return shim.Error("Invalid contact function name.")
}

func (s *SmartContract) AddSession(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) QueryByNIK(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Seed(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}


