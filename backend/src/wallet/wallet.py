import os
import time
import requests
import _thread

class WalletPhim:

    def __init__(self, address, privateAddress):

        self.url = "http://localhost:8000/"

        self.address = address
        self.privateAddress = privateAddress
        self.currentBalance = 0
        self.cTransaction = 0
        self.cContract = 0

    def sendTransactions(self, addressTo, amount):
        
        data = {"addressFrom": self.address, "addressTo": addressTo, "amount": amount, "c": self.cTransaction, "sign": "", "fee": 1}

        requests.post(self.url + "addTransation", data = data, timeout=100)

    def sendContract(self):

        data = {"address": self.address, "contractName": "contractName", "data": self.getContractData("contractTest.py"), "c": self.cContract, "sign": "signtest", "fee": 1, "autoExec": "false", "blockChain": "true"}

        requests.post(self.url + "addContract", data = data, timeout=100)

    def sendContractInteraction(self, amount, function, arguments):

        data = {"address": self.address, "contractAddress": (self.address + "ª0"), "amount": amount, "function": function, "args": arguments, "fee": 1, "sign": "testContractInteraction", "c": 0}

        requests.post(self.url + "addContractInteraction", data = data, timeout=100)

    def getBalance(self):
        
        data = {"address": self.address}

        resp = requests.post(self.url + "getBalance", data = data, timeout=100)

        self.currentBalance = int(resp.text)

        return resp.text

    def getContractData(self, contractFileName):

        with open(contractFileName, 'r') as file:

            return ''.join([_ for _ in file.readlines()])




wallet = WalletPhim("eu", "o")

while True:

    resp = input("1 - Send Transaction\n2 - Send Contract\n3 - Get Balance\n4 - Contract Interaction\n --> ")

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

    elif int(resp) == 4:

        function = input('Function -> ')
        arguments = input('arguments -> ')
        amount = input('amount -> ')

        wallet.sendContractInteraction(amount, function, arguments)

