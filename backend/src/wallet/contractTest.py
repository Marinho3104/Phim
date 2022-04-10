import PhimChainContract

class contractName(PhimChainContract.PhimChainContract):

    def initialization(self):

        self.list = []

    def autoExecution(self):

        self.sendTo("eu", 1)

    def enter(self):

        if self.interactorInfo.amount > 50:
            self.list.append(self.interactorInfo.address)

        if len(self.list) == 2:

            self.sendTo(self.list[0], self.availabeBalance())