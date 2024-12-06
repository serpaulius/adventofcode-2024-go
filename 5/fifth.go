package fifth

import (
	"adventofcode/2024-go/util"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type RuleMap map[int64][]int64

func parseRule(line string) (int64, int64) {
	pages := strings.Split(line, "|")
	page1, _ := strconv.ParseInt(pages[0], 10, 64)
	page2, _ := strconv.ParseInt(pages[1], 10, 64)
	return page1, page2

}

func parseRules(lines []string) RuleMap {
	var rules = RuleMap{}
	for _, line := range lines {
		if strings.Contains(line, "|") {
			page, before := parseRule(line)
			rules[page] = append(rules[page], before)
		}
	}
	return rules
}

type Update struct {
	pages []int64
}

func parseUpdates(lines []string) []Update {
	var rules = []Update{}
	for _, line := range lines {
		if strings.Contains(line, ",") {
			updates := strings.Split(line, ",")
			var numbers = []int64{}
			for _, update := range updates {
				number, _ := strconv.ParseInt(update, 10, 64)
				numbers = append(numbers, number)
			}
			rules = append(rules, Update{numbers})
		}
	}
	return rules
}

func validateUpdate(update Update, rules RuleMap) (int64, error) {
	for i, currentPage := range update.pages {
		prevPages := update.pages[:i]
		for _, previousPage := range prevPages {
			ruleViolated := slices.Contains(rules[currentPage], previousPage)
			if ruleViolated {
				return 0, fmt.Errorf("update %v violated one of rules %v at %v", update, rules[currentPage], currentPage)
			}
		}
	}
	return update.pages[(len(update.pages)-1)/2], nil
}

func checkUpdates(lines []string) int64 {
	rules := parseRules(lines)
	updates := parseUpdates(lines)
	log.Println(rules, updates)

	var sumOfUpdateMidValues int64
	for _, update := range updates {
		result, err := validateUpdate(update, rules)
		log.Println(update, result, err)
		if err == nil {
			sumOfUpdateMidValues += result
		}
	}
	return sumOfUpdateMidValues
}

func Run() {
	file := util.OpenFileOrPanicPlz("./5/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	result := checkUpdates(lines)
	fmt.Println("5.1 - pages and rules", result)

}
