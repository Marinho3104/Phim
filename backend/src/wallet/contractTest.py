import PhimChainContract

class contractName(PhimChainContract.PhimChainContract):

    def initialization(self):

        self.test = "olaa"

        self.list = ["ola", 1]

        self.sendTo("eu", 100000)

    def enter(self):

        self.list.append(self.interactorInfo.address)