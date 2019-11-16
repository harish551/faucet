package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
	"errors"

	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/joho/godotenv"
	"github.com/tomasen/realip"
)

type ErrorResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type SuccessResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type AccountQueryRes struct {
	//emcli query struct
	Type 		string 			`json:"type"`
	Value 		Value 			`json:"value"`

	//Faucet query struct
	//Account_query []Account_query `json:"account_query"`
	//Raw           []Raw           `json:"raw"`
}

type Value struct {
	Address 		string 		`json:"address"`
	Coins 			[]Coin 		`json:"coins"`
	Public_key 		Public_key 	`json:"public_key"`
	Account_number 	int64 		`json:"account_number"`
	Sequence 		int64 		`json:"sequence"`
}

type Coin struct {
	Denom 			string 		`json:"denom"`
	Amount 			int64 		`json:"amount"`
}

type Public_key struct {
	Type 		string 		`json:"type"`
	Value 		string 		`json:"value"`
}

type Raw struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
}

type Account_query struct {
	Balance            string `json:"balance"`
	Nonce              string `json:"nonce"`
	Public_key_address string `json:"public_key_address"`
}

var (
	DENOM = "x3ngm"
)

var recaptchaSecretKey string
var amountFaucet string
var amountSteak string
var key string
var pass string
var node string
var publicUrl string

type claim_struct struct {
	Address  string `json:"address"`
	Response string `json:"response"`
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(key, "=", value)
		return value
	} else {
		log.Fatal("Error loading environment variable: ", key)
		return ""
	}
}

func main() {
	err := godotenv.Load(".env.local", ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	recaptchaSecretKey = getEnv("FAUCET_RECAPTCHA_SECRET_KEY")
	amountFaucet = getEnv("FAUCET_AMOUNT_FAUCET")
	amountSteak = getEnv("FAUCET_AMOUNT_STEAK")
	key = getEnv("FAUCET_KEY")
	pass = getEnv("FAUCET_PASS")
	publicUrl = getEnv("FAUCET_PUBLIC_URL")

	recaptcha.Init(recaptchaSecretKey)

	http.HandleFunc("/claim", getCoinsHandler)

	if err := http.ListenAndServe(publicUrl, nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func executeCmd(command string, writes ...string) {
	cmd, wc, _ := goExecute(command)

	for _, write := range writes {
		wc.Write([]byte(write + "\n"))
	}
	cmd.Wait()
}

func goExecute(command string) (cmd *exec.Cmd, pipeIn io.WriteCloser, pipeOut io.ReadCloser) {
	cmd = getCmd(command)
	pipeIn, _ = cmd.StdinPipe()
	pipeOut, _ = cmd.StdoutPipe()
	go cmd.Start()
	time.Sleep(time.Second)
	return cmd, pipeIn, pipeOut
}

func getCmd(command string) *exec.Cmd {
	// split command into command and args
	split := strings.Split(command, " ")

	var cmd *exec.Cmd
	if len(split) == 1 {
		cmd = exec.Command(split[0])
	} else {
		cmd = exec.Command(split[0], split[1:]...)
	}

	return cmd
}

func CheckAccountBalance(address string, amountFaucet string, key string, chain string, node string) error {
	var queryRes AccountQueryRes

	command := fmt.Sprintf("gaiacli query account %s --chain-id %s --node %s -o json", address, chain, node)
	fmt.Println(" command ", command)

	out, accErr := exec.Command("bash", "-c", command).Output()

	if accErr == nil {
		if err := json.Unmarshal(out, &queryRes); err != nil {
			fmt.Printf("Error unmarshalling command line output %v", err)
			return err
		}
	}

	if &queryRes != nil && &queryRes.Value != nil && &queryRes.Value.Coins != nil && len(queryRes.Value.Coins)>0{
		for _, coin := range queryRes.Value.Coins {
			if coin.Denom == DENOM {
				if coin.Amount < 1000 {
					return  nil
				} else {
					return errors.New("You have enough tokens in your account")
				}
			}
		}
	}

	return nil
}

func getCoinsHandler(res http.ResponseWriter, request *http.Request) {
	address := request.FormValue("address")
	captchaResponse := request.FormValue("response")
	chain := request.FormValue("chain")

	// TODO: Loop over nodes to get chain rpc
	node := "http://34.82.17.52:26657"

	fmt.Println("No error", address, captchaResponse)

	(res).Header().Set("Access-Control-Allow-Origin", "*")

	if len(address) != 45 {
		panic("Invalid address")
	}

	// make sure captcha is valid
	clientIP := realip.FromRequest(request)
	captchaPassed, captchaErr := recaptcha.Confirm(clientIP, captchaResponse)
	if captchaErr != nil {
		panic(captchaErr)
	}

	fmt.Println("Captcha passed? ", captchaPassed)

	if !captchaPassed {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{
			Status: false,
			Message: "Invalid captcha",
		})
		return
	}

	if captchaPassed {

		//check account balance
		err := CheckAccountBalance(address, amountFaucet, key, chain, node)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(ErrorResponse{
				Status:  false,
				Message: err.Error(),
				Error:   err,
			})
			return
		}

		// send the coins!
		sendFaucet := fmt.Sprintf(
			"gaiacli tx send %v %v %v --chain-id %v --node %v",
			key, address, amountFaucet, chain, node)
		fmt.Println(time.Now().UTC().Format(time.RFC3339), address, "[1]")
		executeCmd(sendFaucet, pass)
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(SuccessResponse{
		Status: true,
		Data:   address,
	})

	return
}
