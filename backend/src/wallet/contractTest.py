import PhimChainContract

class contractName(PhimChainContract.PhimChainContract):

    def initialization(self):

        self.list = ["ola", 1]

    def enter(self):

        self.list.append(self.interactorInfo.address)