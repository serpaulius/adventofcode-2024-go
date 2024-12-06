package fifth

import (
	"adventofcode/2024-go/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type RuleMap map[string][]string

type PageList []string

type ViolatedRule struct {
	page           string
	shouldBeBefore string
}

type FilteredUpdates struct {
	valid   []PageList
	invalid []PageList
}

func parseRules(lines []string) RuleMap {
	var rules = RuleMap{}
	for _, line := range lines {
		if strings.Contains(line, "|") {
			rule := strings.Split(line, "|")
			page, shouldBeBefore := rule[0], rule[1]
			rules[page] = append(rules[page], shouldBeBefore)
		}
	}
	return rules
}

func parseUpdates(lines []string) []PageList {
	var pageLists []PageList
	for _, line := range lines {
		if strings.Contains(line, ",") {
			list := strings.Split(line, ",")
			pageLists = append(pageLists, list)
		}
	}
	return pageLists
}

// todo: not the nicest
func validateList(update PageList, rules RuleMap) (*ViolatedRule, error) {
	for i, currentPage := range update {
		prevPages := update[:i]
		for _, previousPage := range prevPages {
			ruleViolated := slices.Contains(rules[currentPage], previousPage)
			if ruleViolated {
				return &ViolatedRule{currentPage, previousPage}, fmt.Errorf("update %v violated one of rules %v at %v", update, rules[currentPage], currentPage)
			}
		}
	}
	return nil, nil
}

func filterValidUpdates(rules RuleMap, pageUpdates []PageList) FilteredUpdates {
	var results = FilteredUpdates{}
	for _, list := range pageUpdates {
		_, err := validateList(list, rules)

		if err == nil {
			results.valid = append(results.valid, list)
		}
		if err != nil {
			results.invalid = append(results.invalid, list)
		}
	}
	return results
}

func sumOfMiddleValues(validPageUpdates []PageList) int64 {
	var sumOfValidMidValues int64
	for _, list := range validPageUpdates {
		middleItem := list[(len(list)-1)/2]
		num, _ := strconv.ParseInt(middleItem, 10, 64)
		sumOfValidMidValues += num
	}
	return sumOfValidMidValues
}

// fixme: todo: pass values vs ref for everything?
func fixViolation(list PageList, violation *ViolatedRule) PageList {
	moveBeforeIndex := slices.Index(list, violation.shouldBeBefore)
	indexToDelete := slices.Index(list, violation.page)
	list = slices.Delete(list, indexToDelete, indexToDelete+1)
	list = slices.Insert(list, moveBeforeIndex, violation.page)
	return list
}

func fixInvalidUpdates(rules RuleMap, updates []PageList) []PageList {
	for _, list := range updates {
		violation, err := validateList(list, rules)
		for err != nil {
			list = fixViolation(list, violation)
			violation, err = validateList(list, rules)
		}
	}
	return updates
}

func Run() {
	file := util.OpenFileOrPanicPlz("./5/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	rules := parseRules(lines)
	pageUpdates := parseUpdates(lines)
	filteredUpdates := filterValidUpdates(rules, pageUpdates)

	result1 := sumOfMiddleValues(filteredUpdates.valid)
	fmt.Println("5.1 - pages and rules", result1)
	fixed := fixInvalidUpdates(rules, filteredUpdates.invalid)
	result2 := sumOfMiddleValues(fixed)
	fmt.Println("5.2 - fixed pages and rules", result2)

}
