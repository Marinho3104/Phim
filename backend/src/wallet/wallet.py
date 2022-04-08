import os
import time
import requests
import _thread

class WalletPhim:

    def __init__(self, address, privateAddress):

        self.address = address
        self.privateAddress = privateAddress
        self.currentBalance = 0
        self.cTransaction = 0
        self.cContract = 0

    def sendTransactions(self, addressTo, amount):
        
        data = {"addressFrom": self.address, "addressTo": addressTo, "amount": amount, "c": self.cTransaction, "sign": "", "fee": 1}

        requests.post("http://localhost:8000/addTransation", data = data, timeout=100)

    def sendContract(self):

        data = {"address": self.address, "contractName": "contractName", "data": self.getContractData("contractTest.py"), "c": self.cContract, "sign": "signtest", "fee": 1}

        requests.post("http://localhost:8000/addContract", data = data, timeout=100)

    def getBalance(self):
        
        data = {"address": self.address}

        resp = requests.post("http://localhost:8000/getBalance", data = data, timeout=100)

        self.currentBalance = int(resp.text)

        return resp.text

    def getContractData(self, contractFileName):

        with open(contractFileName, 'r') as file:

            return ''.join([_ for _ in file.readlines()])




wallet = WalletPhim("eu", "o")

while True:

    resp = input("1 - Send Transaction\n2 - Send Contract\n3 - Get Balance\n --> ")

    if resp == "":

        os.system('cls')

    elif int(resp) == 1:

        addressTo = input("Address to -> ")
        amount = input("Amount -> ")

        wallet.sendTransactions(addressTo, amount)

    elif int(resp) == 2:

        wallet.sendContract()

    elif int(resp) == 3:

        print(wallet.getBalance())

