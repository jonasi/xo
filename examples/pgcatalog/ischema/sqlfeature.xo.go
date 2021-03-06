// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/knq/xo/examples/pgcatalog/pgtypes"

// GENERATED BY XO. DO NOT EDIT.

// SqlFeature represents a row from information_schema.sql_features.
type SqlFeature struct {
	Tableoid       pgtypes.Oid           // tableoid
	Cmax           pgtypes.Cid           // cmax
	Xmax           pgtypes.Xid           // xmax
	Cmin           pgtypes.Cid           // cmin
	Xmin           pgtypes.Xid           // xmin
	Ctid           pgtypes.Tid           // ctid
	FeatureID      pgtypes.CharacterData // feature_id
	FeatureName    pgtypes.CharacterData // feature_name
	SubFeatureID   pgtypes.CharacterData // sub_feature_id
	SubFeatureName pgtypes.CharacterData // sub_feature_name
	IsSupported    pgtypes.YesOrNo       // is_supported
	IsVerifiedBy   pgtypes.CharacterData // is_verified_by
	Comments       pgtypes.CharacterData // comments
}
