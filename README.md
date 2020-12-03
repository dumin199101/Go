# Go
Go语言语法学习

参考书目：《Go程序设计语言》、《Go语言学习笔记》


笔记：

第3章：基本数据

Go的数据类型分为四大类：

1.基础类型

        number(数字),string（字符串）,boolean（布尔）
        
        number:

           int8,int16,int32,int64,uint8,uint16,uint32,uint64
           int(32位系统：int32,64位系统int64,默认),uint,uintptr
           rune(int32),byte(uint8)
        
           取模余数的正负号总是取决于被除数，除法运算，整数相除，舍弃小数部分，浮点相除，保留小数部分
        
           float32,float64(默认)
           
           complex64,complex128
           
        boolean:
        
           true,false        
           
        string:不可变的字节序列
        
           "Hello World"  `Hello \r\n World` 
           
        常量使用const关键字声明
        
        iota是golang语言的常量计数器,只能在常量的表达式中使用。
        iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
        使用iota能简化定义，在定义枚举时很有用。
        
  
2.聚合类型

        array（数组）,struct（结构体）

3.引用类型
      
        pointer（指针）,slice（切片）,map（散列表）,channel（通道）,function（函数）
        
        slice属性:指针，长度，容量（起始元素到底层数组的最后一个元素）

4.接口类型







