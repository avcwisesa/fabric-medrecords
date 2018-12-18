package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Session struct {
	NIP        string    `json:"nip"`
	Treatment  string    `json:"treatment"`
	Medication string    `json:"medication"`
	Datetime   time.Time `json:"datetime"`
}

type MedicalRecord struct {
	NIK    string    `json:"nik"`
	Name   string    `json:"name"`
	Record []Session `json:"session"`
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	if function == "addSession" {
		return s.AddSession(stub, args)
	} else if function == "queryByNIK" {
		return s.QueryByNIK(stub, args)
	} else if function == "seed" {
		return s.Seed(stub)
	}

	return shim.Error("Invalid contact function name.")
}

func (s *SmartContract) AddSession(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	patientBytes, _ := stub.GetState(args[0])
	patient := MedicalRecord{}

	json.Unmarshal(patientBytes, &patient)

	session := Session{
		NIP:        args[1],
		Treatment:  args[2],
		Medication: args[3],
		Datetime:   time.Now(),
	}

	patient.Record = append(patient.Record, session)

	patientBytes, _ = json.Marshal(patient)
	stub.PutState(args[0], patientBytes)

	return shim.Success(nil)
}

func (s *SmartContract) AddPatient(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	patient := MedicalRecord{
		NIK:  args[0],
		Name: args[1],
	}

	patientBytes, _ = json.Marshal(patient)
	stub.PutState(args[0], patient)

	return shim.Success(nil)
}

func (s *SmartContract) QueryByNIK(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	patientBytes, _ := stub.GetState(args[0])

	return shim.Success(patientBytes)
}

func (s *SmartContract) Seed(stub shim.ChaincodeStubInterface) pb.Response {

	session1 := Session{
		NIP:        "123",
		Treatment:  "Totok hidung",
		Medication: "-",
		Datetime:   time.Now(),
	}

	session2 := Session{
		NIP:        "456",
		Treatment:  "-",
		Medication: "Jamu kuat",
		Datetime:   time.Now(),
	}

	patient0 := MedicalRecord{
		NIK:  "1234567890",
		Name: "Harry A. A. Munir",
	}

	patient0.Record = []Session{session1, session2}

	patient0JSON, _ := json.Marshal(patient0)
	stub.PutState("1234567890", patient0JSON)

	return shim.Success(nil)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
