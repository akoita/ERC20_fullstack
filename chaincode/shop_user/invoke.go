package main

import "fmt"
import "github.com/hyperledger/fabric/core/chaincode/shim"
import "github.com/hyperledger/fabric/protos/peer"

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func gethistory(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetHistoryForKey(args[0])

	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}

	var history string
	history = "\n"

	for value.HasNext() {
		history = fmt.Sprintf("%s%s", history, fmt.Sprintln(value.Next()))
	}

	return string(history), nil
}

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	var fct string
	var argv []string
	var ret string
	var err error

	fct, argv = stub.GetFunctionAndParameters()
	STUB = stub
	LOG = shim.NewLogger("Pcoin")
	LOG.SetLevel(shim.LogInfo)

	switch fct {
	// Temporary, not in production
	case "get":
		ret, err = _get(argv)
	case "set":
		ret, err = _set(argv)
	case "new":
		ret, err = userNew(argv)
	case "history":
		ret, err = gethistory(stub, argv)
	default:
		err = fmt.Errorf("Illegal function called \n")
	}

	if err != nil {
		LOG.Error(err)
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(ret))
}
