import json

class InteractorInfo:

    def __init__(self, address, amount):

        self.address = address
        self.amount = amount

class DisableContract:

    def __init__(self):

        self.operationType = "DisableContract"

class SendTo:

    def __init__(self,addressFrom,  addressTo, amount, c, fee):

        self.operationType = "Transaction"
        self.addressFrom =  addressFrom # fixInvalidCaracthers(addressFrom)
        self.addressTo = addressTo# fixInvalidCaracthers(addressTo)
        self.amount = amount
        self.c = c
        self.fee = fee



def fixInvalidCaracthers(string: str):
        
    return string.replace("\\", "\\\\").replace("\"", "\\\"")

class PhimChainContract():

    def __init__(self, address, contractBalance, fee, c, interactorInfo: InteractorInfo) -> None:
        
        self.address = address
        self.contractBalance = contractBalance
        self.fee = fee
        self.interactorInfo = interactorInfo
        self.c = c
        self.__operations = []

    def availabeBalance(self):

        return self.contractBalance - self.fee

    def sendTo(self, addressTo, amount) -> bool:

        if self.contractBalance < amount:
            return False
        self.__operations.append(SendTo(self.address, addressTo, amount, self.c, self.fee))
        return True

    def disableContract(self):

        self.__operations.append(DisableContract())

    def __del__(self):

        _var = vars(self)

        key_list = list(_var.keys())
        val_list = list(_var.values())

        #print(json.dumps(_var[3:]))

        print("{")

        for _ in range(5, len(_var)):

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



    