import time
import PhimChainContract

class contractName(PhimChainContract.PhimChainContract):

    def initialization(self):

        self.list = []

        self.addBlockChain()

        self.addBlock()

        self.addInfoToLastBlock()

        self.addInfoToLastBlock()

    def autoExecution(self):

        self.addBlock()



