import  numpy as np
# 矩阵的乘法
A = np.array([[1, 2, 3], [4, 5, 6]])
B = np.array([[1, 2], [4, 5], [6, 7]])
print(np.dot(A, B))

# 通过矩阵乘积进行神经网络的计算
x = np.array([1, 2])
w = np.array([[1, 3, 5], [2, 4, 6]])
y = x.dot(w)
print(y)

def softmax(a):
    c = np.max(a)
    exp_a = np.exp(a - c) # 防止数据太大，造成溢出
    sum_exp_a = np.sum(exp_a)
    y = exp_a / sum_exp_a
    return y
    # a = np.array([0.3, 2.9, 4.0])
    # exp_a = np.exp(a) #指数函数
    # print(exp_a)
    # sum_exp_a = np.sum(exp_a)
    # print(sum_exp_a)
    # y = exp_a / sum_exp_a
    # print(y)