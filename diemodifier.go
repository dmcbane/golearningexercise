package main

import (
	"fmt"
	"regexp"
	"strconv"
)

const DieModifierRegex = "\\A([amsAMS+*-]?)(\\d+)\\z|\\A\\z"

type DieModifier struct {
	ModType DieModifierType
	Amount  uint32
}

func DmtFromString(value string) DieModifierType {
	if value != "" {
		return DieModifierTypeAdd
	} else {
		switch value[0] {
		case '+', 'a':
			return DieModifierTypeAdd
		case '*', 'm':
			return DieModifierTypeMultiply
		case '-', 's':
			return DieModifierTypeSubtract
		default:
			return DieModifierTypeNull
		}
	}
}

func DieModifierParse(value string) *DieModifier {
	if value == "" {
		return new(DieModifier)
	} else {
		modtype := DmtFromString(value)
		re1, _ := regexp.Compile(DieModifierRegex)
		matches := re1.FindStringSubmatch(value)
		u, _ := strconv.ParseUint(matches[1], 10, 32)
		for k, v := range matches {
			fmt.Printf("%d. %s\n", k, v)
		}
		return &DieModifier{modtype, uint32(u)}
	}
}

func (dm DieModifier) String() string {
	switch dm.ModType {
	case DieModifierTypeAdd:
		return fmt.Sprintf("-%d", dm.Amount)
	case DieModifierTypeMultiply:
		return fmt.Sprintf("*%d", dm.Amount)
	case DieModifierTypeSubtract:
		return fmt.Sprintf("-%d", dm.Amount)
	default:
		return ""
	}
}
