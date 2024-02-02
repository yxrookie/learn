import numpy as np

def numerical_gradient(f, x):
    """
    计算函数 f 在点 x 处的数值梯度
    f: 需要计算梯度的函数
    x: 梯度计算点
    """
    h = 1e-4  # 微小的数值，通常使用 1e-4 或更小
    grad = np.zeros_like(x)

    for idx in range(x.size):
        tmp_val = x[idx]
        # 计算 f(x + h)
        x[idx] = tmp_val + h
        fxh1 = f(x)

        # 计算 f(x - h)
        x[idx] = tmp_val - h
        fxh2 = f(x)

        # 数值梯度
        grad[idx] = (fxh1 - fxh2) / (2 * h)
        x[idx] = tmp_val  # 还原值

    return grad


def gradient_descent(f, init_x, lr=0.01, step_num=100):
    x = init_x

    for i in range(step_num):
        grad = numerical_gradient(f, x)
        x -= lr * grad

    return x

def function_2(x):
    return x[0]**2 + x[1]**2

init_x = np.array([-3.0, 4.0])
print(gradient_descent(function_2, init_x=init_x, lr=0.01, step_num=100))

