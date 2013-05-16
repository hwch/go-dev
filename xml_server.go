package main

import (
        "encoding/xml"
        "fmt"
        "io/ioutil"
)

type XmlBit struct {
        BitNo      string  `xml:"bitNo"`
        BitValue   string  `xml:"bitValue"`
        BitComment string  `xml:"bitComment"`
}

type XmlTran struct {
        TranName string   `xml:"tranName"`
        TranCode string   `xml:"tranNo"`
        TranBit  []XmlBit `xml:"bit"`
}

type XmlTranConfig struct {
        XMLName xml.Name  `xml:"tranconfig"`
        Version string    `xml:"version,attr"`
        Card    []XmlTran `xml:"tran"`
}

func main() {

        sc := new(XmlTranConfig)
        if data, err := ioutil.ReadFile("server.ini"); err != nil {
                fmt.Printf("Error:%v\n", err)
                return
        } else {
                if err := xml.Unmarshal(data, sc); err != nil {
                        fmt.Printf("Error:%v\n", err)
                        return
                }
        }
        fmt.Printf("%#v\n", *sc)
        if v, err := xml.MarshalIndent(sc, "", "\t"); err != nil {
                fmt.Printf("Error:%v\n", err)
        } else {
                fmt.Printf("%s\n", string(v))
        }

}
