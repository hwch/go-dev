package main

import (
        "crypto/des"
        "encoding/hex"
        "fmt"
        "hlog"
)

// 加解密用
const (
        // 加密标志
        ENCRYPT = true
        // 解密标志
        DECRYPT = false
)

type Logger interface {
        WriteLog(uint, string, uint, []byte, int) error
}

// 字节拷贝
func byteCopy(dest, src []byte, iLen int) {
        for i := 0; i < iLen; i++ {
                dest[i] = src[i]
        }

}

// 设置字节
func byteSet(src []byte, s byte, iLen int) {
        for i := 0; i < iLen; i++ {
                src[i] = s
        }
}

/**************************************************************************
* 程序名称: hDes
* 程序功能: 提供PIN及密钥加密之DES算法
* 入口参数: iIn               需加密(解密)之数据( ASCII )
*           iDestKey                密钥                          ( ASCII )
*           bFlg                         模式( ENCRYPT: 加密 DECRYPT:解密 )
* 出口参数: iOut                需解密(加密)之数据( ASCII )
**************************************************************************/
func hDes(iIn, iOut, iDestKey []byte, bFlg bool) {
        // DES算法用
        var table6 = [16]uint16{
                1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1,
        }

        // DES算法用
        var table7 = [64]uint16{
                0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1,
                0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 1,
                1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 1, 1,
                1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1,
        }

        var s1 = [4][16]uint16{
                {14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
                {0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
                {4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
                {15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
        }

        var s2 = [4][16]uint16{
                {15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
                {3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
                {0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
                {13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
        }

        var s3 = [4][16]uint16{
                {10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
                {13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
                {13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
                {1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
        }
        var s4 = [4][16]uint16{
                {7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
                {13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
                {10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
                {3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
        }

        var s5 = [4][16]uint16{
                {2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
                {14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
                {4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
                {11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
        }

        var s6 = [4][16]uint16{
                {12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
                {10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
                {9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
                {4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
        }

        var s7 = [4][16]uint16{
                {4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
                {13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
                {1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
                {6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},
        }

        var s8 = [4][16]uint16{
                {13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
                {1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
                {7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
                {2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11},
        }

        var table [64]uint16
        var table8 [64]uint16
        var table3 [64]uint16
        var table1 [48]uint16
        var table5 [48]uint16
        var table2 [56]uint16

        for i := 0; i < 8; i++ {
                j := uint16(iDestKey[i])
                table8[8*i] = (j / 128) % 2
                table8[8*i+1] = (j / 64) % 2
                table8[8*i+2] = (j / 32) % 2
                table8[8*i+3] = (j / 16) % 2
                table8[8*i+4] = (j / 8) % 2
                table8[8*i+5] = (j / 4) % 2
                table8[8*i+6] = (j / 2) % 2
                table8[8*i+7] = j % 2
        }

        for i := 0; i < 8; i++ {
                j := uint16(iIn[i])
                table3[8*i] = (j / 128) % 2
                table3[8*i+1] = (j / 64) % 2
                table3[8*i+2] = (j / 32) % 2
                table3[8*i+3] = (j / 16) % 2
                table3[8*i+4] = (j / 8) % 2
                table3[8*i+5] = (j / 4) % 2
                table3[8*i+6] = (j / 2) % 2
                table3[8*i+7] = j % 2
        }

        table[0] = table3[57]
        table[1] = table3[49]
        table[2] = table3[41]
        table[3] = table3[33]
        table[4] = table3[25]
        table[5] = table3[17]
        table[6] = table3[9]
        table[7] = table3[1]
        table[8] = table3[59]
        table[9] = table3[51]
        table[10] = table3[43]
        table[11] = table3[35]
        table[12] = table3[27]
        table[13] = table3[19]
        table[14] = table3[11]
        table[15] = table3[3]
        table[16] = table3[61]
        table[17] = table3[53]
        table[18] = table3[45]
        table[19] = table3[37]
        table[20] = table3[29]
        table[21] = table3[21]
        table[22] = table3[13]
        table[23] = table3[5]
        table[24] = table3[63]
        table[25] = table3[55]
        table[26] = table3[47]
        table[27] = table3[39]
        table[28] = table3[31]
        table[29] = table3[23]
        table[30] = table3[15]
        table[31] = table3[7]
        table[32] = table3[56]
        table[33] = table3[48]
        table[34] = table3[40]
        table[35] = table3[32]
        table[36] = table3[24]
        table[37] = table3[16]
        table[38] = table3[8]
        table[39] = table3[0]
        table[40] = table3[58]
        table[41] = table3[50]
        table[42] = table3[42]
        table[43] = table3[34]
        table[44] = table3[26]
        table[45] = table3[18]
        table[46] = table3[10]
        table[47] = table3[2]
        table[48] = table3[60]
        table[49] = table3[52]
        table[50] = table3[44]
        table[51] = table3[36]
        table[52] = table3[28]
        table[53] = table3[20]
        table[54] = table3[12]
        table[55] = table3[4]
        table[56] = table3[62]
        table[57] = table3[54]
        table[58] = table3[46]
        table[59] = table3[38]
        table[60] = table3[30]
        table[61] = table3[22]
        table[62] = table3[14]
        table[63] = table3[6]

        table2[0] = table8[56]
        table2[1] = table8[48]
        table2[2] = table8[40]
        table2[3] = table8[32]
        table2[4] = table8[24]
        table2[5] = table8[16]
        table2[6] = table8[8]
        table2[7] = table8[0]
        table2[8] = table8[57]
        table2[9] = table8[49]
        table2[10] = table8[41]
        table2[11] = table8[33]
        table2[12] = table8[25]
        table2[13] = table8[17]
        table2[14] = table8[9]
        table2[15] = table8[1]
        table2[16] = table8[58]
        table2[17] = table8[50]
        table2[18] = table8[42]
        table2[19] = table8[34]
        table2[20] = table8[26]
        table2[21] = table8[18]
        table2[22] = table8[10]
        table2[23] = table8[2]
        table2[24] = table8[59]
        table2[25] = table8[51]
        table2[26] = table8[43]
        table2[27] = table8[35]
        table2[28] = table8[62]
        table2[29] = table8[54]
        table2[30] = table8[46]
        table2[31] = table8[38]
        table2[32] = table8[30]
        table2[33] = table8[22]
        table2[34] = table8[14]
        table2[35] = table8[6]
        table2[36] = table8[61]
        table2[37] = table8[53]
        table2[38] = table8[45]
        table2[39] = table8[37]
        table2[40] = table8[29]
        table2[41] = table8[21]
        table2[42] = table8[13]
        table2[43] = table8[5]
        table2[44] = table8[60]
        table2[45] = table8[52]
        table2[46] = table8[44]
        table2[47] = table8[36]
        table2[48] = table8[28]
        table2[49] = table8[20]
        table2[50] = table8[12]
        table2[51] = table8[4]
        table2[52] = table8[27]
        table2[53] = table8[19]
        table2[54] = table8[11]
        table2[55] = table8[3]

        for flage3 := 1; flage3 < 17; flage3++ {
                for i := 0; i < 32; i++ {
                        table3[i] = table[32+i]
                }

                table1[0] = table3[31]
                table1[1] = table3[0]
                table1[2] = table3[1]
                table1[3] = table3[2]
                table1[4] = table3[3]
                table1[5] = table3[4]
                table1[6] = table3[3]
                table1[7] = table3[4]
                table1[8] = table3[5]
                table1[9] = table3[6]
                table1[10] = table3[7]
                table1[11] = table3[8]
                table1[12] = table3[7]
                table1[13] = table3[8]
                table1[14] = table3[9]
                table1[15] = table3[10]
                table1[16] = table3[11]
                table1[17] = table3[12]
                table1[18] = table3[11]
                table1[19] = table3[12]
                table1[20] = table3[13]
                table1[21] = table3[14]
                table1[22] = table3[15]
                table1[23] = table3[16]
                table1[24] = table3[15]
                table1[25] = table3[16]
                table1[26] = table3[17]
                table1[27] = table3[18]
                table1[28] = table3[19]
                table1[29] = table3[20]
                table1[30] = table3[19]
                table1[31] = table3[20]
                table1[32] = table3[21]
                table1[33] = table3[22]
                table1[34] = table3[23]
                table1[35] = table3[24]
                table1[36] = table3[23]
                table1[37] = table3[24]
                table1[38] = table3[25]
                table1[39] = table3[26]
                table1[40] = table3[27]
                table1[41] = table3[28]
                table1[42] = table3[27]
                table1[43] = table3[28]
                table1[44] = table3[29]
                table1[45] = table3[30]
                table1[46] = table3[31]
                table1[47] = table3[0]

                switch {
                case bFlg:
                        flage2 := table6[flage3-1]
                        for i := 0; i < int(flage2); i++ {
                                temp1 := table2[0]
                                temp2 := table2[28]
                                for j := 0; j < 27; j++ {
                                        table2[j] = table2[j+1]
                                        table2[j+28] = table2[j+29]
                                }
                                table2[27] = temp1
                                table2[55] = temp2
                        }
                case flage3 > 1:
                        flage2 := table6[17-flage3]
                        for i := 0; i < int(flage2); i++ {
                                temp1 := table2[27]
                                temp2 := table2[55]
                                for j := 27; j > 0; j-- {
                                        table2[j] = table2[j-1]
                                        table2[j+28] = table2[j+27]
                                }
                                table2[0] = temp1
                                table2[28] = temp2
                        }
                }

                table5[0] = table2[13]
                table5[1] = table2[16]
                table5[2] = table2[10]
                table5[3] = table2[23]
                table5[4] = table2[0]
                table5[5] = table2[4]
                table5[6] = table2[2]
                table5[7] = table2[27]
                table5[8] = table2[14]
                table5[9] = table2[5]
                table5[10] = table2[20]
                table5[11] = table2[9]
                table5[12] = table2[22]
                table5[13] = table2[18]
                table5[14] = table2[11]
                table5[15] = table2[3]
                table5[16] = table2[25]
                table5[17] = table2[7]
                table5[18] = table2[15]
                table5[19] = table2[6]
                table5[20] = table2[26]
                table5[21] = table2[19]
                table5[22] = table2[12]
                table5[23] = table2[1]
                table5[24] = table2[40]
                table5[25] = table2[51]
                table5[26] = table2[30]
                table5[27] = table2[36]
                table5[28] = table2[46]
                table5[29] = table2[54]
                table5[30] = table2[29]
                table5[31] = table2[39]
                table5[32] = table2[50]
                table5[33] = table2[44]
                table5[34] = table2[32]
                table5[35] = table2[47]
                table5[36] = table2[43]
                table5[37] = table2[48]
                table5[38] = table2[38]
                table5[39] = table2[55]
                table5[40] = table2[33]
                table5[41] = table2[52]
                table5[42] = table2[45]
                table5[43] = table2[41]
                table5[44] = table2[49]
                table5[45] = table2[35]
                table5[46] = table2[28]
                table5[47] = table2[31]

                for i := 0; i < 48; i++ {
                        table1[i] = table1[i] ^ table5[i]
                }

                flage1 := s1[2*table1[0]+table1[5]][2*(2*(2*table1[1]+table1[2])+table1[3])+table1[4]]
                flage1 = flage1 * 4
                table5[0] = table7[0+flage1]
                table5[1] = table7[1+flage1]
                table5[2] = table7[2+flage1]
                table5[3] = table7[3+flage1]
                flage1 = s2[2*table1[6]+table1[11]][2*(2*(2*table1[7]+table1[8])+table1[9])+table1[10]]
                flage1 = flage1 * 4
                table5[4] = table7[0+flage1]
                table5[5] = table7[1+flage1]
                table5[6] = table7[2+flage1]
                table5[7] = table7[3+flage1]
                flage1 = s3[2*table1[12]+table1[17]][2*(2*(2*table1[13]+table1[14])+table1[15])+table1[16]]
                flage1 = flage1 * 4
                table5[8] = table7[0+flage1]
                table5[9] = table7[1+flage1]
                table5[10] = table7[2+flage1]
                table5[11] = table7[3+flage1]
                flage1 = s4[2*table1[18]+table1[23]][2*(2*(2*table1[19]+table1[20])+table1[21])+table1[22]]
                flage1 = flage1 * 4
                table5[12] = table7[0+flage1]
                table5[13] = table7[1+flage1]
                table5[14] = table7[2+flage1]
                table5[15] = table7[3+flage1]
                flage1 = s5[2*table1[24]+table1[29]][2*(2*(2*table1[25]+table1[26])+table1[27])+table1[28]]
                flage1 = flage1 * 4
                table5[16] = table7[0+flage1]
                table5[17] = table7[1+flage1]
                table5[18] = table7[2+flage1]
                table5[19] = table7[3+flage1]
                flage1 = s6[2*table1[30]+table1[35]][2*(2*(2*table1[31]+table1[32])+table1[33])+table1[34]]
                flage1 = flage1 * 4
                table5[20] = table7[0+flage1]
                table5[21] = table7[1+flage1]
                table5[22] = table7[2+flage1]
                table5[23] = table7[3+flage1]
                flage1 = s7[2*table1[36]+table1[41]][2*(2*(2*table1[37]+table1[38])+table1[39])+table1[40]]
                flage1 = flage1 * 4
                table5[24] = table7[0+flage1]
                table5[25] = table7[1+flage1]
                table5[26] = table7[2+flage1]
                table5[27] = table7[3+flage1]
                flage1 = s8[2*table1[42]+table1[47]][2*(2*(2*table1[43]+table1[44])+table1[45])+table1[46]]
                flage1 = flage1 * 4
                table5[28] = table7[0+flage1]
                table5[29] = table7[1+flage1]
                table5[30] = table7[2+flage1]
                table5[31] = table7[3+flage1]

                table1[0] = table5[15]
                table1[1] = table5[6]
                table1[2] = table5[19]
                table1[3] = table5[20]
                table1[4] = table5[28]
                table1[5] = table5[11]
                table1[6] = table5[27]
                table1[7] = table5[16]
                table1[8] = table5[0]
                table1[9] = table5[14]
                table1[10] = table5[22]
                table1[11] = table5[25]
                table1[12] = table5[4]
                table1[13] = table5[17]
                table1[14] = table5[30]
                table1[15] = table5[9]
                table1[16] = table5[1]
                table1[17] = table5[7]
                table1[18] = table5[23]
                table1[19] = table5[13]
                table1[20] = table5[31]
                table1[21] = table5[26]
                table1[22] = table5[2]
                table1[23] = table5[8]
                table1[24] = table5[18]
                table1[25] = table5[12]
                table1[26] = table5[29]
                table1[27] = table5[5]
                table1[28] = table5[21]
                table1[29] = table5[10]
                table1[30] = table5[3]
                table1[31] = table5[24]

                for i := 0; i < 32; i++ {
                        table[i+32] = table[i] ^ table1[i]
                        table[i] = table3[i]
                }
        }

        for i := 0; i < 32; i++ {
                j := table[i]
                table[i] = table[32+i]
                table[32+i] = j
        }

        table3[0] = table[39]
        table3[1] = table[7]
        table3[2] = table[47]
        table3[3] = table[15]
        table3[4] = table[55]
        table3[5] = table[23]
        table3[6] = table[63]
        table3[7] = table[31]
        table3[8] = table[38]
        table3[9] = table[6]
        table3[10] = table[46]
        table3[11] = table[14]
        table3[12] = table[54]
        table3[13] = table[22]
        table3[14] = table[62]
        table3[15] = table[30]
        table3[16] = table[37]
        table3[17] = table[5]
        table3[18] = table[45]
        table3[19] = table[13]
        table3[20] = table[53]
        table3[21] = table[21]
        table3[22] = table[61]
        table3[23] = table[29]
        table3[24] = table[36]
        table3[25] = table[4]
        table3[26] = table[44]
        table3[27] = table[12]
        table3[28] = table[52]
        table3[29] = table[20]
        table3[30] = table[60]
        table3[31] = table[28]
        table3[32] = table[35]
        table3[33] = table[3]
        table3[34] = table[43]
        table3[35] = table[11]
        table3[36] = table[51]
        table3[37] = table[19]
        table3[38] = table[59]
        table3[39] = table[27]
        table3[40] = table[34]
        table3[41] = table[2]
        table3[42] = table[42]
        table3[43] = table[10]
        table3[44] = table[50]
        table3[45] = table[18]
        table3[46] = table[58]
        table3[47] = table[26]
        table3[48] = table[33]
        table3[49] = table[1]
        table3[50] = table[41]
        table3[51] = table[9]
        table3[52] = table[49]
        table3[53] = table[17]
        table3[54] = table[57]
        table3[55] = table[25]
        table3[56] = table[32]
        table3[57] = table[0]
        table3[58] = table[40]
        table3[59] = table[8]
        table3[60] = table[48]
        table3[61] = table[16]
        table3[62] = table[56]
        table3[63] = table[24]

        j := 0
        for i := 0; i < 8; i++ {
                iOut[i] = 0x0
                for k := 0; k < 7; k++ {
                        iOut[i] = byte((uint16(iOut[i]) + table3[j+k]) * 2)
                }
                iOut[i] = byte(uint16(iOut[i]) + table3[j+7])
                j += 8
        }
}

/**************************************************************************
* 程序名称: hDes3
* 程序功能: 提供PIN及密钥加密之3DES算法
* 入口参数: iIn        需加密(解密)之数据( ASCII )
*           iDestKey       密钥              ( ASCII )
*           bFlg           模式( ENCRYPT: 加密 DECRYPT:解密 )
* 出口参数: iOutt       需解密(加密)之数据( ASCII )
**************************************************************************/
func hDes3(iIn, iOut, iDestKey []byte, iKeyLen int, bFlg bool) {
        myKey1 := make([]byte, 8)
        myKey2 := make([]byte, 8)
        myKey3 := make([]byte, 8)
        iOut1 := make([]byte, 8)
        iOut2 := make([]byte, 8)

        byteCopy(myKey1, iDestKey, 8)
        byteCopy(myKey2, iDestKey[8:], 8)
        if iKeyLen == 16 {
                byteCopy(myKey3, myKey1, 8)
        } else {
                byteCopy(myKey3, iDestKey[16:], 8)
        }

        if bFlg { /*加密*/
                hDes(iIn, iOut1, myKey1, ENCRYPT)
                hDes(iOut1, iOut2, myKey2, DECRYPT)
                hDes(iOut2, iOut, myKey3, ENCRYPT)
        } else { /*解密*/
                hDes(iIn, iOut1, myKey1, DECRYPT)
                hDes(iOut1, iOut2, myKey2, ENCRYPT)
                hDes(iOut2, iOut, myKey3, DECRYPT)
        }
}

// 解密密码
func decryptPinAP(strCardNo string, encPin, decPin []byte, pinKey string, als Logger) error {
        pinCardNo := make([]byte, 32)
        pinBlock := make([]byte, 32)

        byteCopy(pinCardNo, []byte("0000"), 4)
        iLen := len(strCardNo)
        byteCopy(pinCardNo[4:], []byte(strCardNo[iLen-13:]), 12)

        binKey, err2 := hex.DecodeString(pinKey)
        if err2 != nil {
                als.WriteLog(hlog.ERR_LEVEL, err2.Error(), hlog.RPT_TO_FILE, nil, 0)
                return err2
        }
        iLen = len(binKey)
        if iLen == 8 {
                hDes(encPin, decPin, binKey, DECRYPT)
        } else {
                hDes3(encPin, decPin, binKey, iLen, DECRYPT)
        }

        bcdCardNo, err := hex.DecodeString(string(pinCardNo[:16]))
        if err != nil {
                als.WriteLog(hlog.ERR_LEVEL, err.Error(), hlog.RPT_TO_FILE, nil, 0)
                return err
        }
        iLen = len(bcdCardNo)
        for i := 0; i < 8; i++ {
                pinBlock[i] = decPin[i] ^ bcdCardNo[i]
        }

        strPin := hex.EncodeToString(pinBlock[:iLen])
        byteCopy(decPin, []byte(strPin[2:]), 6)
        als.WriteLog(hlog.DEBUG_LEVEL, "解密前密码", hlog.RPT_TO_FILE, encPin, len(encPin))
        als.WriteLog(hlog.DEBUG_LEVEL, "解密后密码", hlog.RPT_TO_FILE, decPin, 8)
        return nil
}

func encryptPinAP(decPin []byte, passLen int, strCardNo string, cardNoLen int,
        encPin []byte, pinKey string, als Logger) error {
        strPin := make([]byte, 16)
        pinCardNo := make([]byte, 32)

        byteCopy(pinCardNo, []byte("0000"), 4)
        iLen := len(strCardNo)
        byteCopy(pinCardNo[4:], []byte(strCardNo[iLen-13:]), 12)
        myStr := fmt.Sprintf("CardPan:[%d][%s]", 16, string(pinCardNo[:16]))
        als.WriteLog(hlog.DEBUG_LEVEL, myStr, hlog.RPT_TO_FILE, nil, 0)

        byteSet(strPin, 'F', 16)
        if decPin[7] == 0 {
                passLen = 6
        }
        myStr = fmt.Sprintf("%02d%s", passLen, string(decPin[:passLen]))
        byteCopy(strPin, []byte(myStr), len(myStr))

        bcdCardNo, err := hex.DecodeString(string(pinCardNo[:16]))
        if err != nil {
                als.WriteLog(hlog.ERR_LEVEL, err.Error(), hlog.RPT_TO_FILE, nil, 0)
                return err
        }
        pinBlock, err1 := hex.DecodeString(string(strPin[:16]))
        if err1 != nil {
                als.WriteLog(hlog.ERR_LEVEL, err1.Error(), hlog.RPT_TO_FILE, nil, 0)
                return err1
        }
        binKey, err2 := hex.DecodeString(pinKey)
        if err2 != nil {
                als.WriteLog(hlog.ERR_LEVEL, err2.Error(), hlog.RPT_TO_FILE, nil, 0)
                return err2
        }
        iLen = len(binKey)
        for i := 0; i < 8; i++ {
                pinBlock[i] ^= bcdCardNo[i]
        }

        if iLen == 8 {
                hDes(pinBlock, encPin, binKey, ENCRYPT)
        } else {
                x, err := des.NewTripleDESCipher(binKey)
                if err != nil {
                        als.WriteLog(hlog.ERR_LEVEL, err.Error(), hlog.RPT_TO_FILE, nil, 0)
                        return err
                }
                x.Encrypt(encPin, pinBlock)
                // hDes3(pinBlock, encPin, binKey, iLen, ENCRYPT)
        }

        als.WriteLog(hlog.DEBUG_LEVEL, "加密前密码", hlog.RPT_TO_FILE, decPin, passLen)
        als.WriteLog(hlog.DEBUG_LEVEL, "加密后密码", hlog.RPT_TO_FILE, encPin, 8)

        return nil
}

func main() {
        pinKey := "075431b09e70fe8fd6fd8c466dfb9867d6fd8c466dfb9867"
        cardNo := "62242009000000017"
        // 初始化日志打印级别
        hlog.InitLogLevel(hlog.DEBUG_LEVEL)
        // 初始化日志
        myLog := hlog.InitLog("1.log")
        myLog.ChgLogFuncStyle(hlog.SHORT_FUNC)
        encPin := make([]byte, 8)

        decPin := []byte{'8', '8', '8', '8', '8', '8', 0x00, 0x00}
        err := encryptPinAP(decPin, len(decPin), cardNo,
                len(cardNo),
                encPin, pinKey, myLog)
        if err != nil {
                myLog.WriteLog(hlog.DEBUG_LEVEL, err.Error(), hlog.RPT_TO_FILE, nil, 0)
                return
        }
        decPinOut := make([]byte, 8)
        if err := decryptPinAP(cardNo, encPin,
                decPinOut, pinKey, myLog); err != nil {
                myLog.WriteLog(hlog.DEBUG_LEVEL, err.Error(), hlog.RPT_TO_FILE, nil, 0)
                return
        }
}
