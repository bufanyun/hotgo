// Package casbin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package casbin

import (
	"context"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gogf/gf/v2/database/gdb"
	"math"
	"strings"
)

const (
	defaultTableName     = "hg_admin_role_casbin"
	dropPolicyTableSql   = `DROP TABLE IF EXISTS %s`
	createPolicyTableSql = `
CREATE TABLE IF NOT EXISTS %s (
  id bigint(20) NOT NULL AUTO_INCREMENT,
  p_type varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  v0 varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  v1 varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  v2 varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  v3 varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  v4 varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  v5 varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员_casbin权限表' ROW_FORMAT = Dynamic;
`
)

type (
	adapter struct {
		db    gdb.DB
		table string
	}

	policyColumns struct {
		ID    string // ID
		PType string // PType
		V0    string // V0
		V1    string // V1
		V2    string // V2
		V3    string // V3
		V4    string // V4
		V5    string // V5
	}

	// policy rule entity
	policyRule struct {
		ID    int64  `orm:"id" json:"id"`
		PType string `orm:"p_type" json:"p_type"`
		V0    string `orm:"v0" json:"v0"`
		V1    string `orm:"v1" json:"v1"`
		V2    string `orm:"v2" json:"v2"`
		V3    string `orm:"v3" json:"v3"`
		V4    string `orm:"v4" json:"v4"`
		V5    string `orm:"v5" json:"v5"`
	}
)

var (
	errInvalidDatabaseLink = errors.New("invalid database link")
	policyColumnsName      = policyColumns{
		ID:    "id",
		PType: "p_type",
		V0:    "v0",
		V1:    "v1",
		V2:    "v2",
		V3:    "v3",
		V4:    "v4",
		V5:    "v5",
	}
)

// NewAdapter Create a casbin adapter
func NewAdapter(link string) (adp *adapter, err error) {
	adp = &adapter{table: defaultTableName}
	config := strings.SplitN(link, ":", 2)

	if len(config) != 2 {
		err = errInvalidDatabaseLink
		return
	}

	if adp.db, err = gdb.New(gdb.ConfigNode{Type: config[0], Link: config[1]}); err != nil {
		return
	}

	err = adp.createPolicyTable()

	return
}

func (a *adapter) model() *gdb.Model {
	return a.db.Model(a.table).Safe().Ctx(context.TODO())
}

// create a policy table when it's not exists.
func (a *adapter) createPolicyTable() (err error) {
	_, err = a.db.Exec(context.TODO(), fmt.Sprintf(createPolicyTableSql, a.table))

	return
}

// drop policy table from the storage.
func (a *adapter) dropPolicyTable() (err error) {
	_, err = a.db.Exec(context.TODO(), fmt.Sprintf(dropPolicyTableSql, a.table))

	return
}

// LoadPolicy loads all policy rules from the storage.
func (a *adapter) LoadPolicy(model model.Model) (err error) {
	var rules []policyRule

	if err = a.model().Scan(&rules); err != nil {
		return
	}

	for _, rule := range rules {
		a.loadPolicyRule(rule, model)
	}

	return
}

// SavePolicy Saves all policy rules to the storage.
func (a *adapter) SavePolicy(model model.Model) (err error) {
	if err = a.dropPolicyTable(); err != nil {
		return
	}

	if err = a.createPolicyTable(); err != nil {
		return
	}

	policyRules := make([]policyRule, 0)

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			policyRules = append(policyRules, a.buildPolicyRule(ptype, rule))
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			policyRules = append(policyRules, a.buildPolicyRule(ptype, rule))
		}
	}

	if count := len(policyRules); count > 0 {
		if _, err = a.model().Insert(policyRules); err != nil {
			return
		}
	}

	return
}

// AddPolicy adds a policy rule to the storage.
func (a *adapter) AddPolicy(sec string, ptype string, rule []string) (err error) {
	_, err = a.model().Insert(a.buildPolicyRule(ptype, rule))

	return
}

// AddPolicies adds policy rules to the storage.
func (a *adapter) AddPolicies(sec string, ptype string, rules [][]string) (err error) {
	if len(rules) == 0 {
		return
	}

	policyRules := make([]policyRule, 0, len(rules))

	for _, rule := range rules {
		policyRules = append(policyRules, a.buildPolicyRule(ptype, rule))
	}

	_, err = a.model().Insert(policyRules)

	return
}

// RemovePolicy removes a policy rule from the storage.
func (a *adapter) RemovePolicy(sec string, ptype string, rule []string) (err error) {
	db := a.model()
	db = db.Where(policyColumnsName.PType, ptype)
	for index := 0; index < len(rule); index++ {
		db = db.Where(fmt.Sprintf("v%d", index), rule[index])
	}
	_, err = db.Delete()
	return err
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) (err error) {
	db := a.model()
	db = db.Where(policyColumnsName.PType, ptype)
	for index := 0; index <= 5; index++ {
		if fieldIndex <= index && index < fieldIndex+len(fieldValues) {
			db = db.Where(fmt.Sprintf("v%d", index), fieldValues[index-fieldIndex])
		}
	}
	_, err = db.Delete()
	return
}

// RemovePolicies removes policy rules from the storage (implements the persist.BatchAdapter interface).
func (a *adapter) RemovePolicies(sec string, ptype string, rules [][]string) (err error) {
	db := a.model()

	for _, rule := range rules {
		where := map[string]interface{}{policyColumnsName.PType: ptype}

		for i := 0; i <= 5; i++ {
			if len(rule) > i {
				where[fmt.Sprintf("v%d", i)] = rule[i]
			}
		}

		db = db.WhereOr(where)
	}

	_, err = db.Delete()

	return
}

// UpdatePolicy updates a policy rule from storage.
func (a *adapter) UpdatePolicy(sec string, ptype string, oldRule, newRule []string) (err error) {
	_, err = a.model().Update(a.buildPolicyRule(ptype, newRule), a.buildPolicyRule(ptype, oldRule))

	return
}

// UpdatePolicies updates some policy rules to storage, like db, redis.
func (a *adapter) UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) (err error) {
	if len(oldRules) == 0 || len(newRules) == 0 {
		return
	}

	err = a.db.Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
		for i := 0; i < int(math.Min(float64(len(oldRules)), float64(len(newRules)))); i++ {
			if _, err = tx.Model(a.table).Update(a.buildPolicyRule(ptype, newRules[i]), a.buildPolicyRule(ptype, oldRules[i])); err != nil {
				return err
			}
		}

		return nil
	})

	return
}

// 加载策略规则
func (a *adapter) loadPolicyRule(rule policyRule, model model.Model) {
	ruleText := rule.PType

	if rule.V0 != "" {
		ruleText += ", " + rule.V0
	}

	if rule.V1 != "" {
		ruleText += ", " + rule.V1
	}

	if rule.V2 != "" {
		ruleText += ", " + rule.V2
	}

	if rule.V3 != "" {
		ruleText += ", " + rule.V3
	}

	if rule.V4 != "" {
		ruleText += ", " + rule.V4
	}

	if rule.V5 != "" {
		ruleText += ", " + rule.V5
	}

	if err := persist.LoadPolicyLine(ruleText, model); err != nil {
		panic(err)
	}
}

// 构建策略规则
func (a *adapter) buildPolicyRule(ptype string, data []string) policyRule {
	rule := policyRule{PType: ptype}

	if len(data) > 0 {
		rule.V0 = data[0]
	}

	if len(data) > 1 {
		rule.V1 = data[1]
	}

	if len(data) > 2 {
		rule.V2 = data[2]
	}

	if len(data) > 3 {
		rule.V3 = data[3]
	}

	if len(data) > 4 {
		rule.V4 = data[4]
	}

	if len(data) > 5 {
		rule.V5 = data[5]
	}

	return rule
}
