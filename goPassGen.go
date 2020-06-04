package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func exit() int {
	return 0
}

func main() {
	myFigure := figure.NewFigure("goPassGen", "big", true)
	myFigure.Print()
	fmt.Println("                                                     [bigb0ss]")
	fmt.Println("")
	fmt.Println("[+] Quick Easily-guessable Password List Generator")
	fmt.Println("")

	// Collecting the current Year information (*Required)
	fmt.Print("[*] Enther the current year (ex. 2020)    : ")
	var year int
	fmt.Scanln(&year)
	if year != 0 {
		// fmt.Println("[+] You Entered:", year)
	} else {
		fmt.Print("[*] Enther the current year (ex. 2020)    : ")
		var year int
		fmt.Scanln(&year)
		if year != 0 {
			// fmt.Println("[+] You Entered:", year)
		} else {
			fmt.Println("[-] Cannot Skip Year!")
			os.Exit(exit())
		}
	}

	// Collecting the current Season information (*Required)
	fmt.Print("[*] Enther the current season (ex. Spring): ")
	var season string
	fmt.Scanln(&season)
	if season != "" {
		// fmt.Println("[+] You Entered:", season)
	} else {
		fmt.Print("[*] Enther the current season (ex. Spring): ")
		var season string
		fmt.Scanln(&season)
		if season != "" {
			// fmt.Println("[+] You Entered:", season)
		} else {
			fmt.Println("[-] Cannot Skip Season!")
			os.Exit(exit())
		}
	}

	// Collecting the Company Name (*Required)
	fmt.Print("[*] Enther the company name (ex. Google)  : ")
	var compName string
	fmt.Scanln(&compName)
	if compName != "" {
		// fmt.Println("[+] You Entered:", compName)
	} else {
		fmt.Print("[*] Enther the company name (ex. Google)  : ")
		var compName string
		fmt.Scanln(&compName)
		if compName != "" {
			// fmt.Println("[+] You Entered:", compName)
		} else {
			fmt.Println("[-] Cannot Skip Company Name!")
			os.Exit(exit())
		}
	}

	// Collecting one additional word
	fmt.Print("[*] Enther additional word or hit 'Enter' : ")
	var addWord string
	fmt.Scanln(&addWord)
	if addWord != "" {
		// fmt.Println("[+] You Entered:", addWord)
	} else {
		fmt.Println("[-] Cool. No additional word entered")
	}

	//
	// Generating password list
	//
	fmt.Println("[+] Generating easily-guessable password list...")
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("________________________________________________________________")
	fmt.Println("")

	//---------------------------------------------------------------------------//
	//
	// [Massaging Season]
	//
	//---------------------------------------------------------------------------//
	titleSeason := strings.Title(strings.ToLower(season))
	// // upperSeason := strings.ToUpper(season)
	lowerSeason := strings.ToLower(season)

	// Season + Year Combination
	seasonTitleYear := fmt.Sprintf("%s%d", titleSeason, year)
	// // seasonUpperYear := fmt.Sprintf("%s%d", upperSeason, year)
	seasonLowerYear := fmt.Sprintf("%s%d", lowerSeason, year)
	fmt.Println(seasonTitleYear)
	// fmt.Println(seasonUpperYear)
	fmt.Println(seasonLowerYear)

	// Massaging Year
	lastYear := (year - 1)
	nextYear := (year + 1)

	// Season + Prev Year Combination
	seasonTitleLastYear := fmt.Sprintf("%s%d", titleSeason, lastYear)
	// // seasonUpperLastYear := fmt.Sprintf("%s%d", upperSeason, lastYear)
	seasonLowerLastYear := fmt.Sprintf("%s%d", lowerSeason, lastYear)
	fmt.Println(seasonTitleLastYear)
	// fmt.Println(seasonUpperLastYear)
	fmt.Println(seasonLowerLastYear)

	// Season + Next Year Combination
	seasonTitleNextYear := fmt.Sprintf("%s%d", titleSeason, nextYear)
	// // seasonUpperNextYear := fmt.Sprintf("%s%d", upperSeason, nextYear)
	seasonLowerNextYear := fmt.Sprintf("%s%d", lowerSeason, nextYear)
	fmt.Println(seasonTitleNextYear)
	// fmt.Println(seasonUpperNextYear)
	fmt.Println(seasonLowerNextYear)

	// Special Character
	spe1 := "!"
	spe2 := "@"

	//
	// [Special Character Gansger #1]
	//
	// Season + Year + spe1 Combination
	seasonTitleYearSpe1 := fmt.Sprintf("%s%d%s", titleSeason, year, spe1)
	// // seasonUpperYearSpe1 := fmt.Sprintf("%s%d%s", upperSeason, year, spe1)
	seasonLowerYearSpe1 := fmt.Sprintf("%s%d%s", lowerSeason, year, spe1)
	fmt.Println(seasonTitleYearSpe1)
	// fmt.Println(seasonUpperYearSpe1)
	fmt.Println(seasonLowerYearSpe1)

	// Season + Prev Year + spe1 Combination
	seasonTitleLastYearSpe1 := fmt.Sprintf("%s%d%s", titleSeason, lastYear, spe1)
	// // seasonUpperLastYearSpe1 := fmt.Sprintf("%s%d%s", upperSeason, lastYear, spe1)
	seasonLowerLastYearSpe1 := fmt.Sprintf("%s%d%s", lowerSeason, lastYear, spe1)
	fmt.Println(seasonTitleLastYearSpe1)
	// fmt.Println(seasonUpperLastYearSpe1)
	fmt.Println(seasonLowerLastYearSpe1)

	// Season + Next Year + spe1 Combination
	seasonTitleNextYearSpe1 := fmt.Sprintf("%s%d%s", titleSeason, nextYear, spe1)
	// // seasonUpperNextYearSpe1 := fmt.Sprintf("%s%d%s", upperSeason, nextYear, spe1)
	seasonLowerNextYearSpe1 := fmt.Sprintf("%s%d%s", lowerSeason, nextYear, spe1)
	fmt.Println(seasonTitleNextYearSpe1)
	// fmt.Println(seasonUpperNextYearSpe1)
	fmt.Println(seasonLowerNextYearSpe1)

	//
	// [Special Character Gangster #2]
	//
	// Season + spe2 + Year Combination
	seasonTitleSpe2Year := fmt.Sprintf("%s%s%d", titleSeason, spe2, year)
	// // seasonUpperSpe2Year := fmt.Sprintf("%s%s%d", upperSeason, spe2, year)
	seasonLowerSpe2Year := fmt.Sprintf("%s%s%d", lowerSeason, spe2, year)
	fmt.Println(seasonTitleSpe2Year)
	// fmt.Println(seasonUpperSpe2Year)
	fmt.Println(seasonLowerSpe2Year)

	// Season + spe2 + Prev Year Combination
	seasonTitleSpe2LastYear := fmt.Sprintf("%s%s%d", titleSeason, spe2, lastYear)
	// // seasonUpperSpe2LastYear := fmt.Sprintf("%s%s%d", upperSeason, spe2, lastYear)
	seasonLowerSpe2LastYear := fmt.Sprintf("%s%s%d", lowerSeason, spe2, lastYear)
	fmt.Println(seasonTitleSpe2LastYear)
	// fmt.Println(seasonUpperSpe2LastYear)
	fmt.Println(seasonLowerSpe2LastYear)

	// Season + spe2 + Next Year Combination
	seasonTitleSpe2NextYear := fmt.Sprintf("%s%s%d", titleSeason, spe2, nextYear)
	// // seasonUpperSpe2NextYear := fmt.Sprintf("%s%s%d", upperSeason, spe2, nextYear)
	seasonLowerSpe2NextYear := fmt.Sprintf("%s%s%d", lowerSeason, spe2, nextYear)
	fmt.Println(seasonTitleSpe2NextYear)
	// fmt.Println(seasonUpperSpe2NextYear)
	fmt.Println(seasonLowerSpe2NextYear)

	//---------------------------------------------------------------------------//
	//
	// [Massaging CompanyName]
	//
	//---------------------------------------------------------------------------//
	titleCompName := strings.Title(strings.ToLower(compName))
	// // upperCompName := strings.ToUpper(compName)
	lowerCompName := strings.ToLower(compName)

	// CompanyName + Year Combination
	compTitleYear := fmt.Sprintf("%s%d", titleCompName, year)
	// // compUpperYear := fmt.Sprintf("%s%d", upperCompName, year)
	compLowerYear := fmt.Sprintf("%s%d", lowerCompName, year)
	fmt.Println(compTitleYear)
	// fmt.Println(compUpperYear)
	fmt.Println(compLowerYear)

	// CompanyName + Prev Year Combination
	compTitleLastYear := fmt.Sprintf("%s%d", titleCompName, lastYear)
	// // compUpperLastYear := fmt.Sprintf("%s%d", upperCompName, lastYear)
	compLowerLastYear := fmt.Sprintf("%s%d", lowerCompName, lastYear)
	fmt.Println(compTitleLastYear)
	// fmt.Println(compUpperLastYear)
	fmt.Println(compLowerLastYear)

	// CompanyName + Next Year Combination
	compTitleNextYear := fmt.Sprintf("%s%d", titleCompName, nextYear)
	// // compUpperNextYear := fmt.Sprintf("%s%d", upperCompName, nextYear)
	compLowerNextYear := fmt.Sprintf("%s%d", lowerCompName, nextYear)
	fmt.Println(compTitleNextYear)
	// fmt.Println(compUpperNextYear)
	fmt.Println(compLowerNextYear)

	//
	// [Special Character Gansger #1]
	//
	// CompanyName + Year + spe1 Combination
	compTitleYearSpe1 := fmt.Sprintf("%s%d%s", titleCompName, year, spe1)
	// // compUpperYearSpe1 := fmt.Sprintf("%s%d%s", upperCompName, year, spe1)
	compLowerYearSpe1 := fmt.Sprintf("%s%d%s", lowerCompName, year, spe1)
	fmt.Println(compTitleYearSpe1)
	// fmt.Println(compUpperYearSpe1)
	fmt.Println(compLowerYearSpe1)

	// CompanyName + Prev Year + spe1 Combination
	compTitleLastYearSpe1 := fmt.Sprintf("%s%d%s", titleCompName, lastYear, spe1)
	// // compUpperLastYearSpe1 := fmt.Sprintf("%s%d%s", upperCompName, lastYear, spe1)
	compLowerLastYearSpe1 := fmt.Sprintf("%s%d%s", lowerCompName, lastYear, spe1)
	fmt.Println(compTitleLastYearSpe1)
	// fmt.Println(compUpperLastYearSpe1)
	fmt.Println(compLowerLastYearSpe1)

	// CompanyName + Next Year + spe1 Combination
	compTitleNextYearSpe1 := fmt.Sprintf("%s%d%s", titleCompName, nextYear, spe1)
	// // compUpperNextYearSpe1 := fmt.Sprintf("%s%d%s", upperCompName, nextYear, spe1)
	compLowerNextYearSpe1 := fmt.Sprintf("%s%d%s", lowerCompName, nextYear, spe1)
	fmt.Println(compTitleNextYearSpe1)
	// fmt.Println(compUpperNextYearSpe1)
	fmt.Println(compLowerNextYearSpe1)

	//
	// [Special Character Gangster #2]
	//
	// CompanyName + spe2 + Year Combination
	compTitleSpe2Year := fmt.Sprintf("%s%s%d", titleCompName, spe2, year)
	// // compUpperSpe2Year := fmt.Sprintf("%s%s%d", upperCompName, spe2, year)
	compLowerSpe2Year := fmt.Sprintf("%s%s%d", lowerCompName, spe2, year)
	fmt.Println(compTitleSpe2Year)
	// fmt.Println(compUpperSpe2Year)
	fmt.Println(compLowerSpe2Year)

	// CompanyName + spe2 + Prev Year Combination
	compTitleSpe2LastYear := fmt.Sprintf("%s%s%d", titleCompName, spe2, lastYear)
	// // compUpperSpe2LastYear := fmt.Sprintf("%s%s%d", upperCompName, spe2, lastYear)
	compLowerSpe2LastYear := fmt.Sprintf("%s%s%d", lowerCompName, spe2, lastYear)
	fmt.Println(compTitleSpe2LastYear)
	// fmt.Println(compUpperSpe2LastYear)
	fmt.Println(compLowerSpe2LastYear)

	// CompanyName + spe2 + Next Year Combination
	compTitleSpe2NextYear := fmt.Sprintf("%s%s%d", titleCompName, spe2, nextYear)
	// // compUpperSpe2NextYear := fmt.Sprintf("%s%s%d", upperCompName, spe2, nextYear)
	compLowerSpe2NextYear := fmt.Sprintf("%s%s%d", lowerCompName, spe2, nextYear)
	fmt.Println(compTitleSpe2NextYear)
	// fmt.Println(compUpperSpe2NextYear)
	fmt.Println(compLowerSpe2NextYear)

	//---------------------------------------------------------------------------//
	//
	// [Massaging AddWord]
	//
	//---------------------------------------------------------------------------//
	if addWord != "" {
		titleAddWord := strings.Title(strings.ToLower(addWord))
		// // upperAddWord := strings.ToUpper(addWord)
		lowerAddWord := strings.ToLower(addWord)

		// AddWord + Year Combination
		addWordTitleYear := fmt.Sprintf("%s%d", titleAddWord, year)
		// // addWordUpperYear := fmt.Sprintf("%s%d", upperAddWord, year)
		addWordLowerYear := fmt.Sprintf("%s%d", lowerAddWord, year)
		fmt.Println(addWordTitleYear)
		// fmt.Println(addWordUpperYear)
		fmt.Println(addWordLowerYear)

		// AddWord + Prev Year Combination
		addWordTitleLastYear := fmt.Sprintf("%s%d", titleAddWord, lastYear)
		// // addWordUpperLastYear := fmt.Sprintf("%s%d", upperAddWord, lastYear)
		addWordLowerLastYear := fmt.Sprintf("%s%d", lowerAddWord, lastYear)
		fmt.Println(addWordTitleLastYear)
		// fmt.Println(addWordUpperLastYear)
		fmt.Println(addWordLowerLastYear)

		// AddWord + Next Year Combination
		addWordTitleNextYear := fmt.Sprintf("%s%d", titleAddWord, nextYear)
		// // addWordUpperNextYear := fmt.Sprintf("%s%d", upperAddWord, nextYear)
		addWordLowerNextYear := fmt.Sprintf("%s%d", lowerAddWord, nextYear)
		fmt.Println(addWordTitleNextYear)
		// fmt.Println(addWordUpperNextYear)
		fmt.Println(addWordLowerNextYear)

		//
		// [Special Character Gansger #1]
		//
		// AddWord + Year + spe1 Combination
		addWordTitleYearSpe1 := fmt.Sprintf("%s%d%s", titleAddWord, year, spe1)
		// // addWordUpperYearSpe1 := fmt.Sprintf("%s%d%s", upperAddWord, year, spe1)
		addWordLowerYearSpe1 := fmt.Sprintf("%s%d%s", lowerAddWord, year, spe1)
		fmt.Println(addWordTitleYearSpe1)
		// fmt.Println(addWordUpperYearSpe1)
		fmt.Println(addWordLowerYearSpe1)

		// AddWord + Prev Year + spe1 Combination
		addWordTitleLastYearSpe1 := fmt.Sprintf("%s%d%s", titleAddWord, lastYear, spe1)
		// // addWordUpperLastYearSpe1 := fmt.Sprintf("%s%d%s", upperAddWord, lastYear, spe1)
		addWordLowerLastYearSpe1 := fmt.Sprintf("%s%d%s", lowerAddWord, lastYear, spe1)
		fmt.Println(addWordTitleLastYearSpe1)
		// fmt.Println(addWordUpperLastYearSpe1)
		fmt.Println(addWordLowerLastYearSpe1)

		// AddWord + Next Year + spe1 Combination
		addWordTitleNextYearSpe1 := fmt.Sprintf("%s%d%s", titleAddWord, nextYear, spe1)
		// // addWordUpperNextYearSpe1 := fmt.Sprintf("%s%d%s", upperAddWord, nextYear, spe1)
		addWordLowerNextYearSpe1 := fmt.Sprintf("%s%d%s", lowerAddWord, nextYear, spe1)
		fmt.Println(addWordTitleNextYearSpe1)
		// fmt.Println(addWordUpperNextYearSpe1)
		fmt.Println(addWordLowerNextYearSpe1)

		//
		// [Special Character Gangster #2]
		//
		// AddWord + spe2 + Year Combination
		addWordTitleSpe2Year := fmt.Sprintf("%s%s%d", titleAddWord, spe2, year)
		// // addWordUpperSpe2Year := fmt.Sprintf("%s%s%d", upperAddWord, spe2, year)
		addWordLowerSpe2Year := fmt.Sprintf("%s%s%d", lowerAddWord, spe2, year)
		fmt.Println(addWordTitleSpe2Year)
		// fmt.Println(addWordUpperSpe2Year)
		fmt.Println(addWordLowerSpe2Year)

		// AddWord + spe2 + Prev Year Combination
		addWordTitleSpe2LastYear := fmt.Sprintf("%s%s%d", titleAddWord, spe2, lastYear)
		// // addWordUpperSpe2LastYear := fmt.Sprintf("%s%s%d", upperAddWord, spe2, lastYear)
		addWordLowerSpe2LastYear := fmt.Sprintf("%s%s%d", lowerAddWord, spe2, lastYear)
		fmt.Println(addWordTitleSpe2LastYear)
		// fmt.Println(addWordUpperSpe2LastYear)
		fmt.Println(addWordLowerSpe2LastYear)

		// AddWord + spe2 + Next Year Combination
		addWordTitleSpe2NextYear := fmt.Sprintf("%s%s%d", titleAddWord, spe2, nextYear)
		// // addWordUpperSpe2NextYear := fmt.Sprintf("%s%s%d", upperAddWord, spe2, nextYear)
		addWordLowerSpe2NextYear := fmt.Sprintf("%s%s%d", lowerAddWord, spe2, nextYear)
		fmt.Println(addWordTitleSpe2NextYear)
		// fmt.Println(addWordUpperSpe2NextYear)
		fmt.Println(addWordLowerSpe2NextYear)
	} else {

	}

	//---------------------------------------------------------------------------//
	//
	// Printing Previous Season
	//
	//---------------------------------------------------------------------------//
	if lowerSeason == "spring" {
		strWin := "winter"

		titleSeason := strings.Title(strings.ToLower(strWin))
		// // upperSeason := strings.ToUpper(strWin)
		lowerSeason := strings.ToLower(strWin)

		// Season + Year Combination
		seasonTitleYear := fmt.Sprintf("%s%d", titleSeason, year)
		// // seasonUpperYear := fmt.Sprintf("%s%d", upperSeason, year)
		seasonLowerYear := fmt.Sprintf("%s%d", lowerSeason, year)

		fmt.Println(seasonTitleYear)
		// fmt.Println(seasonUpperYear)
		fmt.Println(seasonLowerYear)

		// Special Character
		spe1 := "!"
		spe2 := "@"

		seasonTitleYearSpe1 := fmt.Sprintf("%s%d%s", titleSeason, year, spe1)
		// // seasonUpperYearSpe1 := fmt.Sprintf("%s%d%s", upperSeason, year, spe1)
		seasonLowerYearSpe1 := fmt.Sprintf("%s%d%s", lowerSeason, year, spe1)
		fmt.Println(seasonTitleYearSpe1)
		// fmt.Println(seasonUpperYearSpe1)
		fmt.Println(seasonLowerYearSpe1)

		seasonTitleSpe2Year := fmt.Sprintf("%s%s%d", titleSeason, spe2, year)
		// // seasonUpperSpe2Year := fmt.Sprintf("%s%s%d", upperSeason, spe2, year)
		seasonLowerSpe2Year := fmt.Sprintf("%s%s%d", lowerSeason, spe2, year)
		fmt.Println(seasonTitleSpe2Year)
		// fmt.Println(seasonUpperSpe2Year)
		fmt.Println(seasonLowerSpe2Year)

	} else if lowerSeason == "summer" {
		strSpr := "spring"

		titleSeason := strings.Title(strings.ToLower(strSpr))
		// // upperSeason := strings.ToUpper(strSpr)
		lowerSeason := strings.ToLower(strSpr)

		// Season + Year Combination
		seasonTitleYear := fmt.Sprintf("%s%d", titleSeason, year)
		// // seasonUpperYear := fmt.Sprintf("%s%d", upperSeason, year)
		seasonLowerYear := fmt.Sprintf("%s%d", lowerSeason, year)

		fmt.Println(seasonTitleYear)
		// fmt.Println(seasonUpperYear)
		fmt.Println(seasonLowerYear)

		// Special Character
		spe1 := "!"
		spe2 := "@"

		seasonTitleYearSpe1 := fmt.Sprintf("%s%d%s", titleSeason, year, spe1)
		// // seasonUpperYearSpe1 := fmt.Sprintf("%s%d%s", upperSeason, year, spe1)
		seasonLowerYearSpe1 := fmt.Sprintf("%s%d%s", lowerSeason, year, spe1)
		fmt.Println(seasonTitleYearSpe1)
		// fmt.Println(seasonUpperYearSpe1)
		fmt.Println(seasonLowerYearSpe1)

		seasonTitleSpe2Year := fmt.Sprintf("%s%s%d", titleSeason, spe2, year)
		// // seasonUpperSpe2Year := fmt.Sprintf("%s%s%d", upperSeason, spe2, year)
		seasonLowerSpe2Year := fmt.Sprintf("%s%s%d", lowerSeason, spe2, year)
		fmt.Println(seasonTitleSpe2Year)
		// fmt.Println(seasonUpperSpe2Year)
		fmt.Println(seasonLowerSpe2Year)

	} else if lowerSeason == "fall" || lowerSeason == "autumn" {
		strSum := "summer"

		titleSeason := strings.Title(strings.ToLower(strSum))
		// // upperSeason := strings.ToUpper(strSum)
		lowerSeason := strings.ToLower(strSum)

		// Season + Year Combination
		seasonTitleYear := fmt.Sprintf("%s%d", titleSeason, year)
		// // seasonUpperYear := fmt.Sprintf("%s%d", upperSeason, year)
		seasonLowerYear := fmt.Sprintf("%s%d", lowerSeason, year)

		fmt.Println(seasonTitleYear)
		// fmt.Println(seasonUpperYear)
		fmt.Println(seasonLowerYear)

		// Special Character
		spe1 := "!"
		spe2 := "@"

		seasonTitleYearSpe1 := fmt.Sprintf("%s%d%s", titleSeason, year, spe1)
		// // seasonUpperYearSpe1 := fmt.Sprintf("%s%d%s", upperSeason, year, spe1)
		seasonLowerYearSpe1 := fmt.Sprintf("%s%d%s", lowerSeason, year, spe1)
		fmt.Println(seasonTitleYearSpe1)
		// fmt.Println(seasonUpperYearSpe1)
		fmt.Println(seasonLowerYearSpe1)

		seasonTitleSpe2Year := fmt.Sprintf("%s%s%d", titleSeason, spe2, year)
		// // seasonUpperSpe2Year := fmt.Sprintf("%s%s%d", upperSeason, spe2, year)
		seasonLowerSpe2Year := fmt.Sprintf("%s%s%d", lowerSeason, spe2, year)
		fmt.Println(seasonTitleSpe2Year)
		// fmt.Println(seasonUpperSpe2Year)
		fmt.Println(seasonLowerSpe2Year)

	} else if lowerSeason == "winter" {
		strFal := "fall"

		titleSeason := strings.Title(strings.ToLower(strFal))
		// // upperSeason := strings.ToUpper(strFal)
		lowerSeason := strings.ToLower(strFal)

		// Season + Year Combination
		seasonTitleYear := fmt.Sprintf("%s%d", titleSeason, year)
		// // seasonUpperYear := fmt.Sprintf("%s%d", upperSeason, year)
		seasonLowerYear := fmt.Sprintf("%s%d", lowerSeason, year)

		fmt.Println(seasonTitleYear)
		// fmt.Println(seasonUpperYear)
		fmt.Println(seasonLowerYear)

		// Special Character
		spe1 := "!"
		spe2 := "@"

		seasonTitleYearSpe1 := fmt.Sprintf("%s%d%s", titleSeason, year, spe1)
		// // seasonUpperYearSpe1 := fmt.Sprintf("%s%d%s", upperSeason, year, spe1)
		seasonLowerYearSpe1 := fmt.Sprintf("%s%d%s", lowerSeason, year, spe1)
		fmt.Println(seasonTitleYearSpe1)
		// fmt.Println(seasonUpperYearSpe1)
		fmt.Println(seasonLowerYearSpe1)

		seasonTitleSpe2Year := fmt.Sprintf("%s%s%d", titleSeason, spe2, year)
		// // seasonUpperSpe2Year := fmt.Sprintf("%s%s%d", upperSeason, spe2, year)
		seasonLowerSpe2Year := fmt.Sprintf("%s%s%d", lowerSeason, spe2, year)
		fmt.Println(seasonTitleSpe2Year)
		// fmt.Println(seasonUpperSpe2Year)
		fmt.Println(seasonLowerSpe2Year)

		strAut := "autumn"

		titleSeason2 := strings.Title(strings.ToLower(strAut))
		// // upperSeason2 := strings.ToUpper(strAut)
		lowerSeason2 := strings.ToLower(strAut)

		// Season + Year Combination
		seasonTitleYear2 := fmt.Sprintf("%s%d", titleSeason2, year)
		// // seasonUpperYear2 := fmt.Sprintf("%s%d", upperSeason2, year)
		seasonLowerYear2 := fmt.Sprintf("%s%d", lowerSeason2, year)

		fmt.Println(seasonTitleYear2)
		// fmt.Println(seasonUpperYear2)
		fmt.Println(seasonLowerYear2)

		seasonTitleYearSpe12 := fmt.Sprintf("%s%d%s", titleSeason2, year, spe1)
		// // seasonUpperYearSpe12 := fmt.Sprintf("%s%d%s", upperSeason2, year, spe1)
		seasonLowerYearSpe12 := fmt.Sprintf("%s%d%s", lowerSeason2, year, spe1)
		fmt.Println(seasonTitleYearSpe12)
		// fmt.Println(seasonUpperYearSpe12)
		fmt.Println(seasonLowerYearSpe12)

		seasonTitleSpe2Year2 := fmt.Sprintf("%s%s%d", titleSeason2, spe2, year)
		// // seasonUpperSpe2Year2 := fmt.Sprintf("%s%s%d", upperSeason2, spe2, year)
		seasonLowerSpe2Year2 := fmt.Sprintf("%s%s%d", lowerSeason2, spe2, year)
		fmt.Println(seasonTitleSpe2Year2)
		// fmt.Println(seasonUpperSpe2Year2)
		fmt.Println(seasonLowerSpe2Year2)
	}

}
