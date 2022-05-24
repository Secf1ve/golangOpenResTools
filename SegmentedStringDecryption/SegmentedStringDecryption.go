package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//测试输入值（明文）
	var valueString = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*(),.<>?;':[]{}|_"
	//测试输出值（密文）
	var keyString = "010003020504070609085053525554575659585b5a5d5c5f5e414043424544474649484b7073727574777679787b7a7d7c7f7e616063626564676669686b4f10711215146f171b19181d1f0d0f0e0a160b6a6c4a4c4d6e"
	uncrp := Append2stringToMap(keyString, 1, valueString, 0)

	fmt.Println(uncrp)
	earthUncry(uncrp, keyString)

	var queryStrings1 = "37090b59030f11060b0a1b4e0000000000004312170a1b0b0e4107174f1a0b044e0a000202134e0a161d17040359061d43370f15030b10414e340e1c0a0f0b0b061d430e0059220f11124059261ae281ba124e14001c06411a110e00435542495f5e430a0715000306150b0b1c4e4b5242495f5e430c07150a1d4a410216010943e281b54e1c0101160606591b0143121a0b0a1a00094e1f1d010e412d180307050e1c17060f43150159210b144137161d054d41270d4f0710410010010b431507140a1d43001d5903010d064e18010a4307010c1d4e1708031c1c4e02124e1d0a0b13410f0a4f2b02131a11e281b61d43261c18010a43220f1716010d40"
	var queryStrings2 = "3714171e0b0a550a1859101d064b160a191a4b0908140d0e0d441c0d4b1611074318160814114b0a1d06170e1444010b0a0d441c104b150106104b1d011b100e59101d0205591314170e0b4a552a1f59071a16071d44130f041810550a05590555010a0d0c011609590d13430a171d170c0f0044160c1e150055011e100811430a59061417030d1117430910035506051611120b45"
	var queryStrings3 = "2402111b1a0705070a41000a431a000a0e0a0f04104601164d050f070c0f15540d1018000000000c0c06410f0901420e105c0d074d04181a01041c170d4f4c2c0c13000d430e0e1c0a0006410b420d074d55404645031b18040a03074d181104111b410f000a4c41335d1c1d040f4e070d04521201111f1d4d031d090f010e00471c07001647481a0b412b1217151a531b4304001e151b171a4441020e030741054418100c130b1745081c541c0b0949020211040d1b410f090142030153091b4d150153040714110b174c2c0c13000d441b410f13080d12145c0d0708410f1d014101011a050d0a084d540906090507090242150b141c1d08411e010a0d1b120d110d1d040e1a450c0e410f090407130b5601164d00001749411e151c061e454d0011170c0a080d470a1006055a010600124053360e1f1148040906010e130c00090d4e02130b05015a0b104d0800170c0213000d104c1d050000450f01070b47080318445c090308410f010c12171a48021f49080006091a48001d47514c50445601190108011d451817151a104c080a0e5a"

	fmt.Println("queryString1", strings.Join(earthUncry(uncrp, queryStrings1), ""))
	fmt.Println("queryString2", strings.Join(earthUncry(uncrp, queryStrings2), ""))
	fmt.Println("queryString3", strings.Join(earthUncry(uncrp, queryStrings3), ""))
}

//将两个不同间隔但分组后总和相同的字符串转换为Key-Value形式
func Append2stringToMap(strKey string, strKeySpace int, strValue string, strValueSpace int) map[string]string {
	tmpRegKey := stringToSlice(strKeySpace, strKey)
	tmpRegValue := stringToSlice(strValueSpace, strValue)
	tmpMap := make(map[string]string)

	for i := 0; i < len(tmpRegKey); i++ {
		tmpMap[strings.Join(tmpRegKey[i], "")] = strings.Join(tmpRegValue[i], "")
	}
	return tmpMap
}

//根据选择的间隙将字符串分割为切片
func stringToSlice(space int, str string) [][]string {
	regStr := "\\S.{" + strconv.Itoa(space) + "}"
	//fmt.Println(regStr)
	reg0 := regexp.MustCompile(regStr)
	if reg0 == nil {
		fmt.Println("regexp err" + str)
		return nil
	}
	strTemp := reg0.FindAllStringSubmatch(str, -1)
	return strTemp[:]
}

//将字典中的字符串进行匹配
func earthUncry(map1 map[string]string, str1 string) []string {
	var strTemp []string
	keyStringSlice := stringToSlice(1, str1)
	for _, name := range keyStringSlice {
		mapValue, ok := map1[name[0]]
		if ok {
			strTemp = append(strTemp, mapValue)
			//fmt.Println(mapValue)
		}
	}
	//fmt.Println(strTemp)
	return strTemp
}
