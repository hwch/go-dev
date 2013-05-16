package main

func getShortFuncName(f string) string {
        iLen := len(f)
        v := make([]byte, iLen)

        iLen -= 1
        i := 0
        for iLen >= 0 {
                if f[iLen] == '.' {
                        iLen++
                        break
                }
                iLen--
                i++
        }
        byteCopy(v, []byte(f)[iLen:], i)

        return string(v[:i])
}

func getFuncName(f string) string {
        iLen := len(f)
        v := make([]byte, iLen)

        iLen -= 1
        i := 0
        for iLen >= 0 {
                if f[iLen] == '/' {
                        iLen++
                        break
                }
                iLen--
                i++
        }
        byteCopy(v, []byte(f)[iLen:], i)

        return string(v[:i])
}

func byteCopy(dest, src []byte, iLen int) {
        for i := 0; i < iLen; i++ {
                dest[i] = src[i]
        }

}

func main() {
        s := "_/fds/fdsa/fdsa/utils.123"
        println(getShortFuncName(s))
}
