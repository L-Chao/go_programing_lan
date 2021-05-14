package tempconv

import "fmt"

type Celsius float64	//摄氏温度
type Fahrenheit float64	//华氏温度
type Kelvin float64		//开氏温度

const (
	AbslouteZeroC Celsius = -273.5 //绝对零度
	FreeZingC     Celsius = 0      //冰点
	Boiling       Celsius = 100    //沸点
)
func CToK(c Celsius) Kelvin {
	return Kelvin(c - 273.5)
}
func KToC(k Kelvin) Celsius {
	return Celsius(k + 273.5)
}
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func FTok(f Fahrenheit) Kelvin {
	return Kelvin(float64(FToC(f)) - 273.5)
}
func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}
//类型的方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}