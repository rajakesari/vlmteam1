/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Trade Finance Use Case - WORK IN  PROGRESS
 */

package main


import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}


// Define the letter of credit
type VehicleManage struct {
	RequestId		string   	`json:"RequestId"`
	ChasisId		string		`json:"ChasisId"`
	ExpiryDate		string		`json:"ExpiryDate"`
	Customer                string   `json:"Customer"`
	Manufacturer		string		`json:"Manufacturer"`
	RegistrationID		string		`json:"RegationID"`
	RTAId                   string          `json:"RTAId"`
	VehicleType		string          `json:"VehicleType"`
//	Amount			int		`json:"amount"`
	Status			string		`json:"Status"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "placeOrder" {
		return s.placeOrder(APIstub, args)
	} else if function == "issueOrder" {
		return s.issueOrder(APIstub, args)
	} else if function == "acceptOrder" {
		return s.acceptOrder(APIstub, args)
	}else if function == "getVehicle" {
		return s.getVehicle(APIstub, args)
	}else if function == "getVehicleHistory" {
		return s.getVehicleHistory(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}



// This function is initiate by Customer 
func (s *SmartContract) placeOrder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

        RequestId := args[0];
        Customer := args[1];
        Manufacturer := args[2];
        VehicleType := args[3];
//      amount, err := strconv.Atoi(args[3]);
//        if err != nil {
//                return shim.Error("Not able to parse Amount")
        }


	VM := VehicleManage{RequestId: RequestId, Customer: Customer, Manufacturer: Manufacturer, VehicleType: VehicleType, Status: "Requested"}
        VMBytes, err := json.Marshal(VM)

    APIstub.PutState(RequestId,VMBytes)
        fmt.Println("Order Placed -> ", VM)


        return shim.Success(nil)
}



// This function is initiate by Seller
/*func (s *SmartContract) issueLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcId := args[0];
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	LCAsBytes, _ := APIstub.GetState(lcId)

	var lc LetterOfCredit

	err := json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with LC json unmarshaling")
	}


	LC := LetterOfCredit{LCId: lc.LCId, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, Bank: lc.Bank, Seller: lc.Seller, Amount: lc.Amount, Status: "Issued"}
	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with LC json marshaling")
	}

    APIstub.PutState(lc.LCId,LCBytes)
	fmt.Println("LC Issued -> ", LC)


	return shim.Success(nil)
}
*/
func (s *SmartContract) acceptOrder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

        RequestId := args[0];
        
        

        VMAsBytes, _ := APIstub.GetState(RequestId)

        var vm VehicleManage

        err := json.Unmarshal(VMAsBytes, &vm)

        if err != nil {
                return shim.Error("Issue with VM json unmarshaling")
        }


        VM := VehicleManage{RequestId: vm.RequestId, Customer: vm.Customer, Manufacturer: vm.Manufacturer, VehicleType: vm.VehicleType, Status: "Accepted"}
        LCBytes, err := json.Marshal(LC)

        if err != nil {
                return shim.Error("Issue with LC json marshaling")
        }

    APIstub.PutState(vm.RequestId,VMBytes)
        fmt.Println("Order Accepted -> ", VM)

        return shim.Success(nil)
}



/*func (s *SmartContract) getLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcId := args[0];
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	LCAsBytes, _ := APIstub.GetState(lcId)

	return shim.Success(LCAsBytes)
}

func (s *SmartContract) getLCHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcId := args[0];
	
	

	resultsIterator, err := APIstub.GetHistoryForKey(lcId)
	if err != nil {
		return shim.Error("Error retrieving LC history.")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Error retrieving LC history.")
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getLCHistory returning:\n%s\n", buffer.String())

	

	return shim.Success(buffer.Bytes())
}
*/
// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
