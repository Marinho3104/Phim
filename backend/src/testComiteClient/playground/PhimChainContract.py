import json
import json
import os

from numpy import block

class BlockChain:

    def __init__(self, path):
        self.path = path
        self.blockChainData = []
        self.getBlockChain()

    def getBlockChain(self):

        try:

            itemsList = os.listdir(os.getcwd() + self.path)

        except:
            return

        for __ in itemsList:
            with open(os.getcwd() + self.path + "\\" + __, "rb") as file:

                try:
                    self.blockChainData.append(json.loads(''.join([_.decode("utf-8") for _ in file.readlines()])))
                except:
                  self.blockChainData.append({})
class InteractorInfo:

    def __init__(self, address, amount):

        self.address = address
        self.amount = amount

class DisableContract:

    def __init__(self):

        self.operationType = "DisableContract"

class SendTo:

    def __init__(self,  addressTo, amount, c, fee):

        self.operationType = "Transaction"
        self.addressTo = addressTo# fixInvalidCaracthers(addressTo)
        self.amount = amount


def fixInvalidCaracthers(string: str):
        
    return string.replace("\\", "\\\\").replace("\"", "\\\"")

class PhimChainContract():

    def __init__(self, blockChain, address, contractBalance, fee, c, interactorInfo: InteractorInfo) -> None:

        self.address = address
        self.contractBalance = contractBalance
        self.fee = fee
        self.interactorInfo = interactorInfo
        self.c = c
        self.blockChain = blockChain
        self.__operations = []

    def availabeBalance(self):

        return self.contractBalance - self.fee

    def addBlockChain(self):

        self.__operations.append({"operationType": "AddBlockChain"})


    def addBlock(self):

        self.__operations.append({"operationType": "AddNewBlock"})

    def addInfoToLastBlock(self):

        self.__operations.append({"operationType": "AddInfoToLastBlock", "info": "{\"ola\": \"meu\"}"})

    def sendTo(self, addressTo, amount) -> bool:

        if self.contractBalance < amount:
            return False
        self.__operations.append(SendTo(addressTo, amount))
        return True

    def disableContract(self):

        self.__operations.append(DisableContract())

    def __del__(self):

        _var = vars(self)

        key_list = list(_var.keys())
        val_list = list(_var.values())

        #print(json.dumps(_var[3:]))

        print("{")

        for _ in range(6, len(_var)):

            if type(val_list[_]) is str:
                _value = f"\"{val_list[_]}\""

            elif type(val_list[_]) is list:

                temp = []

                for __ in val_list[_]:
                    try:
                        temp.append(__.__dict__)
                    except:

                        if type(__) == str:
                            __ = f"{__}"
                        temp.append(__)

                _value = json.dumps(temp)
            
            else:
                _value = f"{val_list[_]}"

            print(f"\"{key_list[_]}\": {_value}", end='')

            print((len(_var) - 1 != _) and ", " or "")

        print("}")          



    