package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	res := trulyMostPopular(
		[]string{"Fcclu(70)", "Ommjh(63)", "Dnsay(60)", "Qbmk(45)", "Unsb(26)", "Gauuk(75)", "Wzyyim(34)", "Bnea(55)", "Kri(71)", "Qnaakk(76)", "Gnplfi(68)", "Hfp(97)", "Qoi(70)", "Ijveol(46)", "Iidh(64)", "Qiy(26)", "Mcnef(59)", "Hvueqc(91)", "Obcbxb(54)", "Dhe(79)", "Jfq(26)", "Uwjsu(41)", "Wfmspz(39)", "Ebov(96)", "Ofl(72)", "Uvkdpn(71)", "Avcp(41)", "Msyr(9)", "Pgfpma(95)", "Vbp(89)", "Koaak(53)", "Qyqifg(85)", "Dwayf(97)", "Oltadg(95)", "Mwwvj(70)", "Uxf(74)", "Qvjp(6)", "Grqrg(81)", "Naf(3)", "Xjjol(62)", "Ibink(32)", "Qxabri(41)", "Ucqh(51)", "Mtz(72)", "Aeax(82)", "Kxutz(5)", "Qweye(15)", "Ard(82)", "Chycnm(4)", "Hcvcgc(97)", "Knpuq(61)", "Yeekgc(11)", "Ntfr(70)", "Lucf(62)", "Uhsg(23)", "Csh(39)", "Txixz(87)", "Kgabb(80)", "Weusps(79)", "Nuq(61)", "Drzsnw(87)", "Xxmsn(98)", "Onnev(77)", "Owh(64)", "Fpaf(46)", "Hvia(6)", "Kufa(95)", "Chhmx(66)", "Avmzs(39)", "Okwuq(96)", "Hrschk(30)", "Ffwni(67)", "Wpagta(25)", "Npilye(14)", "Axwtno(57)", "Qxkjt(31)", "Dwifi(51)", "Kasgmw(95)", "Vgxj(11)", "Nsgbth(26)", "Nzaz(51)", "Owk(87)", "Yjc(94)", "Hljt(21)", "Jvqg(47)", "Alrksy(69)", "Tlv(95)", "Acohsf(86)", "Qejo(60)", "Gbclj(20)", "Nekuam(17)", "Meutux(64)", "Tuvzkd(85)", "Fvkhz(98)", "Rngl(12)", "Gbkq(77)", "Uzgx(65)", "Ghc(15)", "Qsc(48)", "Siv(47)"},
		[]string{"(Gnplfi,Qxabri)", "(Uzgx,Siv)", "(Bnea,Lucf)", "(Qnaakk,Msyr)", "(Grqrg,Gbclj)", "(Uhsg,Qejo)", "(Csh,Wpagta)", "(Xjjol,Lucf)", "(Qoi,Obcbxb)", "(Npilye,Vgxj)", "(Aeax,Ghc)", "(Txixz,Ffwni)", "(Qweye,Qsc)", "(Kri,Tuvzkd)", "(Ommjh,Vbp)", "(Pgfpma,Xxmsn)", "(Uhsg,Csh)", "(Qvjp,Kxutz)", "(Qxkjt,Tlv)", "(Wfmspz,Owk)", "(Dwayf,Chycnm)", "(Iidh,Qvjp)", "(Dnsay,Rngl)", "(Qweye,Tlv)", "(Wzyyim,Kxutz)", "(Hvueqc,Qejo)", "(Tlv,Ghc)", "(Hvia,Fvkhz)", "(Msyr,Owk)", "(Hrschk,Hljt)", "(Owh,Gbclj)", "(Dwifi,Uzgx)", "(Iidh,Fpaf)", "(Iidh,Meutux)", "(Txixz,Ghc)", "(Gbclj,Qsc)", "(Kgabb,Tuvzkd)", "(Uwjsu,Grqrg)", "(Vbp,Dwayf)", "(Xxmsn,Chhmx)", "(Uxf,Uzgx)"},
	)
	fmt.Println(res)
}

var parent []int
var nameToIndex map[string]int
var indexToNum map[int]int
var indexToName map[int]string

func initParent(n int) {
	parent = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
}

func trulyMostPopular(names []string, synonyms []string) (res []string) {
	mp := make(map[string]bool)
	nameToIndex = make(map[string]int)
	indexToNum = make(map[int]int)
	indexToName = make(map[int]string)
	for i := 0; i < len(synonyms); i++ {
		s1, s2 := parseSynonyms(synonyms[i])
		if !mp[s1] {
			mp[s1] = true
		}
		if !mp[s2] {
			mp[s2] = true
		}
	}
	for i := 0; i < len(names); i++ {
		name, _ := parseName(names[i])
		if !mp[name] {
			mp[name] = true
		}
	}
	count := 0
	for key, _ := range mp {
		indexToName[count] = key
		nameToIndex[key] = count
		count++
	}
	n := len(names)
	initParent(len(mp))
	indexToNum = make(map[int]int)
	for i := 0; i < n; i++ {
		name, num := parseName(names[i])
		indexToNum[nameToIndex[name]] = num
	}
	//for i := 0; i < len(synonyms); i++ {
	//	s1, s2 := parseSynonyms(synonyms[i])
	//	if t1, ok := nameToIndex[s1]; ok {
	//		if t2, is := nameToIndex[s2]; is {
	//			union(t1, t2)
	//		}
	//	}
	//}
	for i := 0; i < len(synonyms); i++ {
		s1, s2 := parseSynonyms(synonyms[i])
		t1, t2 := nameToIndex[s1], nameToIndex[s2]
		union(t1, t2)
	}
	fmt.Println(parent)
	result := make([]int, n)
	for i := 0; i < len(parent); i++ {
		num := indexToNum[i]
		pid := parent[i]
		result[pid] += num
	}
	for i := 0; i < len(result); i++ {
		if result[i] > 0 {
			name := indexToName[i]
			num := result[i]
			res = append(res, fmt.Sprintf("%s(%d)", name, num))
		}
	}
	sort.Strings(res)
	return
}

func union(t1, t2 int) {
	p1 := parent[t1]
	p2 := parent[t2]
	if strings.Compare(indexToName[p1], indexToName[p2]) > 0 {
		p1, p2 = p2, p1
	}
	for i := 0; i < len(parent); i++ {
		if parent[i] == p2 {
			parent[i] = p1
		}
	}
}

func parseName(name string) (str string, num int) {
	i := 0
	var extra string
	for ; i < len(name); i++ {
		if name[i] == '(' {
			str = name[:i]
			extra = name[i+1 : len(name)-1]
			break
		}
	}
	num, _ = strconv.Atoi(extra)
	return
}

func parseSynonyms(synonyms string) (string, string) {
	split := strings.Split(synonyms, ",")
	return split[0][1:], split[1][:len(split[1])-1]
}
