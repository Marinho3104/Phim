import time
import requests
import _thread

def sendTransactions(privateAddress, addressFrom, addressTo, amount, c):

    #textToSign = hashlib.sha256((addressFrom).encode('utf-8')).hexdigest() + hashlib.sha256((addressTo).encode('utf-8')).hexdigest() + hashlib.sha256((str(amount)).encode('utf-8')).hexdigest() + hashlib.sha256((str(c)).encode('utf-8')).hexdigest()
    
    data = {"addressFrom": addressFrom, "addressTo": addressTo, "amount": amount, "c": c, "sign": "", "fee": 1}

    requests.post("http://localhost:8000/addTransation", data = data, timeout=100)

def getBalance():
    
    data = {"address": "eu"}

    resp = requests.post("http://localhost:8000/getBalance", data = data, timeout=100)

    return resp.text

while True:

    input(getBalance())

    for _ in range(1):

        _thread.start_new_thread(sendTransactions, ("ola", "eu", "mim", 10, _,))

        print(_)
        print("----------")

