package postgres

import (
	"gopkg.in/doug-martin/goqu.v5"
)

const placeholder_rune = '$'

func newDatasetAdapter(ds *goqu.Dataset) goqu.Adapter {
	def := goqu.NewDefaultAdapter(ds).(*goqu.DefaultAdapter)
	def.PlaceHolderRune = placeholder_rune
	def.IncludePlaceholderNum = true
	def.JoinOnDeleteSupported = false
	def.JoinOnUpdateSupported = false
	return def
}

func init() {
	goqu.RegisterAdapter("postgres", newDatasetAdapter)
}
