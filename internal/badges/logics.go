package badges

import (
	"fmt"
	"strconv"
)

type Logicer interface {
	GetParams() map[string]string
}

type colorConf struct {
	background string
	text       string
}

type LogicMoreWorst struct {
	Value     int64
	Good      int64
	Bad       int64
	goodColor colorConf
	warnColor colorConf
	badColor  colorConf
}

func NewLogicMoreWorstFromValues(values map[string]string) *LogicMoreWorst {

	value, err := strconv.Atoi(values["value"])
	if err != nil {
		panic(fmt.Sprintf(`Could not convert "%s" as int`, values["value"]))
	}
	good, err := strconv.Atoi(values["good"])
	if err != nil {
		panic(fmt.Sprintf(`Could not convert "%s" as int`, values["good"]))
	}
	bad, err := strconv.Atoi(values["bad"])
	if err != nil {
		panic(fmt.Sprintf(`Could not convert "%s" as int`, values["bad"]))
	}

	return &LogicMoreWorst{
		Value: int64(value),
		Good:  int64(good),
		Bad:   int64(bad),
		goodColor: colorConf{
			background: values["goodBackgroundColor"],
			text:       values["goodTextColor"],
		},
		badColor: colorConf{
			background: values["badBackgroundColor"],
			text:       values["badTextColor"],
		},
		warnColor: colorConf{
			background: values["warningBackgroundColor"],
			text:       values["warningTextColor"],
		},
	}
}

func NewLogicMoreWorst(value, good, bad int64) *LogicMoreWorst {
	return &LogicMoreWorst{
		Value: value,
		Good:  good,
		Bad:   bad,
	}
}

func (lmw *LogicMoreWorst) GetParams() map[string]string {

	m := make(map[string]string)

	m["value"] = fmt.Sprintf("%d", lmw.Value)

	switch {
	case lmw.Value <= lmw.Good:
		m["background-value-color"] = lmw.goodColor.background
		m["text-value-color"] = lmw.goodColor.text
	case lmw.Value >= lmw.Bad:
		m["background-value-color"] = lmw.badColor.background
		m["text-value-color"] = lmw.badColor.text
	default:
		m["background-value-color"] = lmw.warnColor.background
		m["text-value-color"] = lmw.warnColor.text
	}
	return m
}
