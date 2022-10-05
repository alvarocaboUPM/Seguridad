package algoritmos

import (
	"bytes"
	"strings"
	"unicode"
	_ "workspace/libs"
)

//Definimos la tabla en función de la letra de cifrado de la clave
const top_row string ="ABCDEFGHIJKLM"
const bottom_row string="NOPQRSTUVWXYZ"

type pos struct{
	p int
	row bool
}

var PbTable = map[string]string{
		"A": (top_row+"NOPQRSTUVWXYZ"),
	
		"B": (top_row+"NOPQRSTUVWXYZ"),
	
		"C": (top_row+"ZNOPQRSTUVWXY"),
	
		"D": (top_row+"ZNOPQRSTUVWXY"),
	
		"E": (top_row+"YZNOPQRSTUVWX"),
	
		"F": (top_row+"YZNOPQRSTUVWX"),
	
		"G": (top_row+"XYZNOPQRSTUVW"),
	
		"H": (top_row+"XYZNOPQRSTUVW"),
	
		"I": (top_row+"WXYZNOPQRSTUV"),
	
		"J": (top_row+"WXYZNOPQRSTUV"),
	
		"K": (top_row+"VWXYZNOPQRSTU"),
	
		"L": (top_row+"VWXYZNOPQRSTU"),
	
		"M": (top_row+"UVWXYZNOPQRST"),
	
		"N": (top_row+"UVWXYZNOPQRST"),
	
		"O": top_row+("TUVWXYZNOPQRS"),
	
		"P": (top_row+"TUVWXYZNOPQRS"),
	
		"Q": (top_row+"STUVWXYZNOPQR"),
	
		"R": (top_row+"STUVWXYZNOPQR"),
	
		"S": (top_row+"RSTUVWXYZNOPQ"),
	
		"T": (top_row+"RSTUVWXYZNOPQ"),
	
		"U": (top_row+"QRSTUVWXYZNOP"),
	
		"V": (top_row+"QRSTUVWXYZNOP"),
	
		"W": (top_row+"PQRSTUVWXYZNO"),
	
		"X": (top_row+"PQRSTUVWXYZNO"),
	
		"Y": (top_row+"OPQRSTUVWXYZN"),
	
		"Z": (top_row+"OPQRSTUVWXYZN"),
	}

func Get_PbTable(key string) []string{
	res := make([]string, len(key))
	for l := range key {
		//En mayúscula
		letter := string(unicode.ToUpper(rune(key[l])))	
		res = append(res, PbTable[letter])
	}
	return res
}

//Posición en un string y en qué 'cilindro' está
func Get_pos(txt string, letter string) pos{
	var res pos
	res.p= strings.Index(txt, letter)
	if(strings.Contains(top_row, letter)){
		res.row=true
	} else {res.row=false}
	return res
}

func Get_op(table string, letter byte) string{
	if(unicode.IsLetter(rune((letter)))==false){return " "}
	l := string(unicode.ToUpper(rune(letter)))
	
	//Encontramos el opuesto en la tabla dada
	p:= Get_pos(table, l); 
	//Si abajo, el opuesto está en top_row (-12)
	if(p.row==false){return string(table[p.p-13])}
	return string(table[p.p+13])
}

func Encrypt(text string, key string) string{
	//Aquí escribiremos el texto cifrado
	var buffer bytes.Buffer
	
	//Obtenemos las rotaciones de la tabla que utiliza la clave
	
	count:=0

	for i:=0; i<len(text);i++{
		letter:=string(key[count])
		// fmt.Println("Letra: "+string(text[i])+
		// 			" Cifrado: "+Get_op(PbTable[letter],text[i]))
		encryp:=Get_op(PbTable[letter],text[i])
		buffer.WriteString(encryp)
		//Si no 
		if(!(strings.EqualFold(encryp," "))){
		count = (count+1) % len(key) //Repetimos la clave
		}
	}
	
	return buffer.String()
}
