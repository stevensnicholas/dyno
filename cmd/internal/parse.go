package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

func ParseFuzz(file string) {
	// TODO change the reading of a file to read from AWS SQS
	internalServerErrors := 0
	useAfterFreeChecker := 0
	nameSpaceRuleChecker := 0
	resourceHierarchyChecker := 0
	leakageRuleChecker := 0
	invalidDynamicObjectChecker := 0
	payloadBodyChecker := 0
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	// Categorizing Errors
	for scanner.Scan() {
		
		line := scanner.Text()
		value, _ := strconv.Atoi(line[len(line) - 1:])
		if line[:len(line) - 1] == "Total Buckets: " {
			break;
		}
		s := strings.Split(scanner.Text(), "_")
		
		// Assigning all counts of categorized errors 
		internalServerErrors, 
		useAfterFreeChecker, 
		nameSpaceRuleChecker, 
		resourceHierarchyChecker, 
		leakageRuleChecker, 
		invalidDynamicObjectChecker, 
		payloadBodyChecker = errorCheck(s[0], 
																		value, 
																		internalServerErrors, 
																		useAfterFreeChecker, 
																		nameSpaceRuleChecker, 
																		resourceHierarchyChecker, 
																		leakageRuleChecker, 
																		invalidDynamicObjectChecker, 
																		payloadBodyChecker, 
																	)
	}

	// Formatting Issues to categorized errors 
	for scanner.Scan() {
		println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	} 
	
}

func errorCheck(s string, value int, 
								internalServerErrors int, 
								useAfterFreeChecker int, 
								nameSpaceRuleChecker int, 
								resourceHierarchyChecker int, 
								leakageRuleChecker int, 
								invalidDynamicObjectChecker int, 
								payloadBodyChecker int, 
							) (int, int, int, int, int, int, int) {
	if s == "InternalServerErrors" {
		internalServerErrors = value
	}

	if s == "UseAfterFreeChecker" {
		useAfterFreeChecker = value
	}

	if s == "NameSpaceRuleChecker" {
		nameSpaceRuleChecker = value
	}

	if s == "ResourceHierarchyChecker" {
		resourceHierarchyChecker = value
	}

	if s == "LeakageRuleChecker" {
		leakageRuleChecker = value
	}

	if s == "InvalidDynamicObjectChecker" {
		invalidDynamicObjectChecker = value
	}

	if s == "PayloadBodyChecker" {
		payloadBodyChecker = value
	}
	return internalServerErrors, 
				 useAfterFreeChecker, 
				 nameSpaceRuleChecker, 
				 resourceHierarchyChecker, 
				 leakageRuleChecker, 
				 invalidDynamicObjectChecker, 
				 payloadBodyChecker
}