package main

import (
	"fmt"
	"math"
)

func main() {
	//TODO math包详解

	// 常量
	// 整型
	fmt.Printf("Int的最大值是:%d\n", math.MaxInt)
	fmt.Printf("Int的最小值是:%d\n", math.MinInt)
	//fmt.Printf("Uint的最大值是:%d\n", math.MaxUint)
	fmt.Printf("Int8的最大值是:%d\n", math.MaxInt8)
	fmt.Printf("Int8的最小值是:%d\n", math.MinInt8)
	fmt.Printf("Uint8的最大值是:%d\n", math.MaxUint8)
	fmt.Printf("Int16的最大值是:%d\n", math.MaxInt16)
	fmt.Printf("Int16的最小值是:%d\n", math.MinInt16)
	fmt.Printf("Uint16的最大值是:%d\n", math.MaxUint16)
	fmt.Printf("Int32的最大值是:%d\n", math.MaxInt32)
	fmt.Printf("Int32的最小值是:%d\n", math.MinInt32)
	fmt.Printf("Uint32的最大值是:%d\n", math.MaxUint32)
	fmt.Printf("Int64的最大值是:%d\n", math.MaxInt64)
	fmt.Printf("Int64的最小值是:%d\n", math.MinInt64)
	//fmt.Println("Uint64的最大值是:", math.MaxUint64)
	// 浮点型
	fmt.Printf("float64的最大值是:%.f\n", math.MaxFloat64)
	fmt.Printf("float64的最小值是:%.f\n", math.SmallestNonzeroFloat64)
	fmt.Printf("float32的最大值是:%.f\n", math.MaxFloat32)
	fmt.Printf("float32的最小值是:%.f\n", math.SmallestNonzeroFloat32)
	// 数学常数
	fmt.Printf("自然对数的底：%.60f\n", math.E)
	fmt.Printf("圆周率：%.60f\n", math.Pi)
	fmt.Printf("黄金分割数：%.60f\n", math.Phi)

	// 算数函数
	// 返回一个IEEE 754（这不是一个数字）值
	fmt.Println("非数字表达式：", math.NaN())
	// 判断是否为非数字（NaN）值
	fmt.Println("该变量是否不为数字：", math.IsNaN(math.NaN()))
	// 如果sign>=0函数返回正无穷大，否则返回负无穷大
	fmt.Println("正无穷大 or 负无穷大表达式：", math.Inf(0))
	// 判断该数是否为无穷大 如果sign>0，f是正无穷大时返回true；如果sign<0，f是负无穷大时返回true；sign==0则f是两种无穷大时都返回true
	fmt.Println(math.IsInf(math.Inf(0), 0))
	// 判断该数是否<0  如果x是一个负数，返回true
	fmt.Println(math.Signbit(0))
	// 根据y返回x的量值（绝对值） 如果y>=0时返回y的绝对值，如果y<0返回x负值
	fmt.Println(math.Copysign(-10, 0))
	// 向上取整
	fmt.Println(math.Ceil(1.1))
	// 向下取整
	fmt.Println(math.Floor(1.1))
	// 返回x的整数部分
	fmt.Println(math.Trunc(121.1))
	// 返回f的整数部分和小数部分（结果的正负号都和f相同，比如f为负数则结果都都有负号）
	n, f := math.Modf(121.1)
	fmt.Println("整数部分：", n, " 小数部分：", f)
	// 参数x到参数y的方向上，下一个可表示的数值；如果x==y将返回x
	fmt.Println("下一个可表示的数值：", math.Nextafter(1.52, 1.53))
	// 绝对值
	fmt.Println("绝对值：", math.Abs(-15))
	// 两数比较取最大值
	fmt.Println("最大值为：", math.Max(12, 13))
	// 两数比较取最小值
	fmt.Println("最小值为：", math.Min(12, 13))
	// 四舍五入
	fmt.Println("四舍五入：", math.Round(12.44))
	// 四舍五入 舍入为偶数
	fmt.Println("四舍五入：", math.RoundToEven(1.34))
	// 返回x-y的结果和0比较的最大值，如果x-y小于0则返回0
	fmt.Println("结果和0比较的最大值：", math.Dim(11, 12))
	// 取余 返回结果带有正负号与x相同
	fmt.Println("取余：", math.Mod(12, 8))
	// IEEE 754差数求值，即x减去最接近x/y的整数值（如果有两个整数与x/y距离相同，则取其中的偶数）与y的乘积（x-x/y)*y
	fmt.Println("差数值：", math.Remainder(12, 8))
	// 返回x的二次方根
	fmt.Println("二次方根：", math.Sqrt(2))
	// 返回x的三次方根
	fmt.Println("三次方根：", math.Cbrt(2))
	// 返回Sqrt(p*p + q*q)，注意要避免不必要的溢出或下溢
	fmt.Println("Sqrt(p*p + q*q)公式结果：", math.Hypot(2, 2))
	// 求正弦
	fmt.Println("正弦：", math.Sin(2))
	// 求余弦
	fmt.Println("余弦：", math.Cos(2))
	// 求正切
	fmt.Println("正切：", math.Tan(2))
	// 返回Sin(x), Cos(x)
	sin, cos := math.Sincos(2)
	fmt.Println("Sin：", sin, " cos：", cos)
	// 求反正弦（x是弧度）
	fmt.Println("反正弦：", math.Asin(2))
	// 求反余弦（x是弧度）
	fmt.Println("反余弦：", math.Acos(2))
	// 求反正切（x是弧度）
	fmt.Println("反正切：", math.Atan(2))
	// 类似Atan(y/x)，但会根据x，y的正负号确定象限
	fmt.Println("反正切2：", math.Atan2(2, 3))
	// 求双曲正弦
	fmt.Println("双曲正弦：", math.Sinh(2))
	// 求双曲余弦
	fmt.Println("双曲余弦：", math.Cosh(2))
	// 求双曲正切
	fmt.Println("双曲正切：", math.Tanh(2))
	// 求反双曲正弦
	fmt.Println("反双曲正弦：", math.Asinh(2))
	// 求反双曲余弦
	fmt.Println("反双曲余弦：", math.Acosh(2))
	// 求反双曲正切
	fmt.Println("反双曲正切：", math.Atanh(2))
	// 求自然对数
	fmt.Println("求自然对数：", math.Log(2))
	// 等价于Log(1+x)。但是在x接近0时，本函数更加精确
	fmt.Println("Log1p：", math.Log1p(2))
	// 求2为底的对数
	fmt.Println("2为底的对数：", math.Log2(2))
	// 求10为底的对数
	fmt.Println("10为底的对数：", math.Log10(2))
	// 返回x的二进制指数值，可以理解为Trunc(Log2(x))
	fmt.Println("二进制指数值：", math.Logb(2))
	// 类似Logb，但返回值是整型
	fmt.Println("Ilogb：", math.Ilogb(2))
	// 返回一个标准化小数frac和2的整型指数exp，满足f == frac * 2**exp，且0.5 <= Abs(frac) < 1
	frac, exp := math.Frexp(2)
	fmt.Println("标准化小数frac：", frac, " 2的整型指数exp:", exp)
	// Frexp的反函数，返回 frac * 2**exp
	fmt.Println("Frexp的反函数：", math.Ldexp(2, 1))
	// 返回E**x；x绝对值很大时可能会溢出为0或者+Inf，x绝对值很小时可能会下溢为1
	fmt.Println("Exp：", math.Exp(2))
	// 等价于Exp(x)-1，但是在x接近零时更精确；x绝对值很大时可能会溢出为-1或+Inf
	fmt.Println("Expm1：", math.Expm1(2))
	// 返回2**x
	fmt.Println("Exp2：", math.Exp2(2))
	// 返回x**y (X的Y次方)
	fmt.Println("Pow：", math.Pow(2, 4))
	// 返回10**e (10的N次方)
	fmt.Println("Pow10：", math.Pow10(2))
	// 返回Gamma(x)的自然对数和正负号
	lgamma, sign := math.Lgamma(2)
	fmt.Println("自然对数：", lgamma, "正负号：", sign)
	// 计算误差
	fmt.Println("误差：", math.Erf(2))
	// 计算余补误差
	fmt.Println("余补误差：", math.Erfc(2))
	// 第一类贝塞尔函数，0阶
	fmt.Println("第一类贝塞尔函数，0阶：", math.J0(2))
	// 第一类贝塞尔函数，1阶
	fmt.Println("第一类贝塞尔函数，1阶：", math.J1(2))
	// 第一类贝塞尔函数，n阶
	fmt.Println("第一类贝塞尔函数，n阶：", math.Jn(3, 2))
	// 第二类贝塞尔函数，0阶
	fmt.Println("第二类贝塞尔函数，0阶：", math.Y0(2))
	// 第一类贝塞尔函数，1阶
	fmt.Println("第二类贝塞尔函数，1阶：", math.Y1(2))
	// 第一类贝塞尔函数，n阶
	fmt.Println("第二类贝塞尔函数，n阶：", math.Yn(3, 2))
}
