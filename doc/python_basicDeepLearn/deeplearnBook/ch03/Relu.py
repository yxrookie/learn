import numpy as np


class Relu:
    def __init__(self):
        # 布尔数组
        self.mask = None

    def forward(self, x):
        self.mask = (x <= 0)
        out = x.copy()
        out[self.mask] = 0

        return out

    def backward(self, dout):
        # 复用正向传播得到的 mask 进行操作数组
        dout[self.mask] = 0
        dx = dout

        return dx

r = Relu()
x = np.array([[1.0, -2.0], [2.3, -9.7]])
print(r.forward(x))
print(r.backward(x))