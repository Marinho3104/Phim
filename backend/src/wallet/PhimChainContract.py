import json

class InteractorInfo:

    def __init__(self, address, amount):

        self.address = address
        self.amount = amount

class SendTo:

    def __init__(self,addressFrom,  addressTo, amount):

        self.addressFrom =  addressFrom # fixInvalidCaracthers(addressFrom)
        self.addressTo = addressTo# fixInvalidCaracthers(addressTo)
        self.amount = amount

    def __repr__(self) -> str:
        return f"""{{ "addressFrom": "{self.addressFrom}", "addressTo": "{self.addressTo}", "amount": {self.amount} }}"""
 
    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__, 
            sort_keys=True, indent=4)

def fixInvalidCaracthers(string: str):
        
    return string.replace("\\", "\\\\").replace("\"", "\\\"")

class PhimChainContract():

    def __init__(self, address, contractBalance, fee, interactorInfo: InteractorInfo) -> None:
        
        self.address = address
        self.contractBalance = contractBalance
        self.fee = fee
        self.interactorInfo = interactorInfo
        self.__operations = []

    def availabeBalance(self):

        return self.contractBalance - self.fee

    def sendTo(self, addressTo, amount) -> bool:

        if self.contractBalance < amount:
            return False
        self.__operations.append(SendTo(self.address, addressTo, amount))
        return True

    def __del__(self):

        _var = vars(self)

        key_list = list(_var.keys())
        val_list = list(_var.values())

        #print(json.dumps(_var[3:]))

        print("{")

        for _ in range(4, len(_var)):

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



    