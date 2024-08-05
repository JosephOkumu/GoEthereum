# GoEthereum
This program uses Golang to interact with the Ethereum blockchain.
## Documentation
This section illustrates how to make use of this program.

### Installation
To run this program, download and install the latest version of Go from [here](https://go.dev/doc/install) and download Ganache [here](https://archive.trufflesuite.com/ganache/)

### Usage
1. Clone this repository on to the terminal by using the following command:
```bash
git clone https://github.com/JosephOkumu/GoEthereum
```
2. Navigate into the GoEthereum directory by using the command:
```bash
cd GoEthereum
```
3. Log on to ganache, and pick the private key of sender and address of recepient. In main.go :
    - Replace the placeholder private key in the transferETH function with a valid private key from your Ethereum client. (Exclude the first two characters)
    - Update the toAddress in the transferETH function with the address to which you want to send ETH.
    - Update the address in the checkBalance function call with the address whose balance you want to check after the transaction.


4. To run the program, execute the command below:
```bash
go run . 
```
The following message will be printed out on the terminal:
```bash
We have a connection
```
Once you see the message above check ganache to track the transactions.

## Features
- This program can connect to the Ethereum client.
- Can generate a new Ethereum wallet and print its address.
- Can transfer some amount of ETH from a specified address to a given address.
- This program prints the transaction status and the balance of a specified address.

## Contributions
Pull requests are welcome where users of this program are allowed to contribute to this project in terms of adding features, or fixing bugs.

For major changes, please open an issue first to discuss what you would like to change.
## Authors
[JosephOkumu](https://github.com/JosephOkumu)

## Licence
[MIT](https://choosealicense.com/licenses/mit/)
## Credits
[Zone01Kisumu](https://www.zone01kisumu.ke/)
