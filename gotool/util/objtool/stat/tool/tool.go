// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tool

import (
	"fmt"
	"github.com/unknwon/i18n"
	"strings"
	"time"
)

const TIME_LIMIT_CODE_LENGTH = 12 + 6 + 40

// Seconds-based time units
const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

func computeTimeDiff(diff int64) (int64, string) {
	diffStr := ""
	switch {
	case diff <= 0:
		diff = 0
		diffStr = "now"
	case diff < 2:
		diff = 0
		diffStr = "1 second"
	case diff < 1*Minute:
		diffStr = fmt.Sprintf("%d seconds", diff)
		diff = 0

	case diff < 2*Minute:
		diff -= 1 * Minute
		diffStr = "1 minute"
	case diff < 1*Hour:
		diffStr = fmt.Sprintf("%d minutes", diff/Minute)
		diff -= diff / Minute * Minute

	case diff < 2*Hour:
		diff -= 1 * Hour
		diffStr = "1 hour"
	case diff < 1*Day:
		diffStr = fmt.Sprintf("%d hours", diff/Hour)
		diff -= diff / Hour * Hour

	case diff < 2*Day:
		diff -= 1 * Day
		diffStr = "1 day"
	case diff < 1*Week:
		diffStr = fmt.Sprintf("%d days", diff/Day)
		diff -= diff / Day * Day

	case diff < 2*Week:
		diff -= 1 * Week
		diffStr = "1 week"
	case diff < 1*Month:
		diffStr = fmt.Sprintf("%d weeks", diff/Week)
		diff -= diff / Week * Week

	case diff < 2*Month:
		diff -= 1 * Month
		diffStr = "1 month"
	case diff < 1*Year:
		diffStr = fmt.Sprintf("%d months", diff/Month)
		diff -= diff / Month * Month

	case diff < 2*Year:
		diff -= 1 * Year
		diffStr = "1 year"
	default:
		diffStr = fmt.Sprintf("%d years", diff/Year)
		diff = 0
	}
	return diff, diffStr
}

// TimeSincePro calculates the time interval and generate full user-friendly string.
func TimeSincePro(then time.Time) string {
	now := time.Now()
	diff := now.Unix() - then.Unix()

	if then.After(now) {
		return "future"
	}

	var timeStr, diffStr string
	for {
		if diff == 0 {
			break
		}

		diff, diffStr = computeTimeDiff(diff)
		timeStr += ", " + diffStr
	}
	return strings.TrimPrefix(timeStr, ", ")
}

func timeSince(then time.Time, lang string) string {
	now := time.Now()

	lbl := i18n.Tr(lang, "tool.ago")
	diff := now.Unix() - then.Unix()
	if then.After(now) {
		lbl = i18n.Tr(lang, "tool.from_now")
		diff = then.Unix() - now.Unix()
	}

	switch {
	case diff <= 0:
		return i18n.Tr(lang, "tool.now")
	case diff <= 2:
		return i18n.Tr(lang, "tool.1s", lbl)
	case diff < 1*Minute:
		return i18n.Tr(lang, "tool.seconds", diff, lbl)

	case diff < 2*Minute:
		return i18n.Tr(lang, "tool.1m", lbl)
	case diff < 1*Hour:
		return i18n.Tr(lang, "tool.minutes", diff/Minute, lbl)

	case diff < 2*Hour:
		return i18n.Tr(lang, "tool.1h", lbl)
	case diff < 1*Day:
		return i18n.Tr(lang, "tool.hours", diff/Hour, lbl)

	case diff < 2*Day:
		return i18n.Tr(lang, "tool.1d", lbl)
	case diff < 1*Week:
		return i18n.Tr(lang, "tool.days", diff/Day, lbl)

	case diff < 2*Week:
		return i18n.Tr(lang, "tool.1w", lbl)
	case diff < 1*Month:
		return i18n.Tr(lang, "tool.weeks", diff/Week, lbl)

	case diff < 2*Month:
		return i18n.Tr(lang, "tool.1mon", lbl)
	case diff < 1*Year:
		return i18n.Tr(lang, "tool.months", diff/Month, lbl)

	case diff < 2*Year:
		return i18n.Tr(lang, "tool.1y", lbl)
	default:
		return i18n.Tr(lang, "tool.years", diff/Year, lbl)
	}
}
