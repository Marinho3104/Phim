import time
import PhimChainContract

class contractName(PhimChainContract.PhimChainContract):

    def initialization(self):

        self.list = []

        self.currentTime = time.time()

    def autoExecution(self):

        if time.time() - self.currentTime >= 15.0:
            self.sendTo("eu", 1)
            self.currentTime = time.time()


