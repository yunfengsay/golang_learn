from ctypes import *  
  
class StructPointer(Structure):  
    _fields_ = [("p", c_char_p), ("n", c_longlong)]  
  
if __name__ == "__main__":  
    lib = cdll.LoadLibrary("./libhello.so")  
    lib.Hello.restype = StructPointer
    str = lib.Hello()  
    print(str.n)
    #str.n是GoString返回字符的长度，没有截取的话后面会跟着一大串字符串
    print(str.p[:str.n])
    lib.Test()

    # go build -x -v -ldflags "-s -w" -buildmode=c-shared  -o libhello.so test
