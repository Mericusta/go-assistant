package goass

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Mericusta/go-stp"
)

const (
	ATTRIBUTE_ATK = iota + 1
	ATTRIBUTE_DEF
	ATTRIBUTE_HP
	ATTRIBUTE_MP
)

type data_Logic struct {
	exp               int64
	level             int64
	attributesLevel   map[int64]int64
	idleTS            int64
	cumulativeRewards int64
	saveDataPath      string
}

type json_Logic struct {
	Exp             int64           `json:"exp"`
	Level           int64           `json:"level"`
	AttributesLevel map[int64]int64 `json:"attributes_level"`
	IdleTS          int64           `json:"idle_ts"`
}

func newLogicData(saveDataPath string) *data_Logic {
	dataLogic := &data_Logic{
		exp:   0,
		level: 1,
		attributesLevel: map[int64]int64{
			ATTRIBUTE_ATK: 0,
			ATTRIBUTE_DEF: 0,
			ATTRIBUTE_HP:  0,
			ATTRIBUTE_MP:  0,
		},
		idleTS:       time.Now().Unix(),
		saveDataPath: saveDataPath,
	}
	// dataLogic.LoadJSON()
	dataLogic.cumulativeRewards = time.Now().Unix() - dataLogic.idleTS
	return dataLogic
}

func (d *data_Logic) AddExp(delta int64) {
	d.exp += delta
	d.UpdateLevel()
}

func (d *data_Logic) UpdateLevel() {
	// 当前等级升级所需经验 = Σ(0, Level)
LEVEL_UP:
	levelUpNeedExp := int64(float64(0+d.level)*float64(d.level+1)/2) * 10
	if levelUpNeedExp > d.exp {
		return
	}
	d.level++
	d.exp -= levelUpNeedExp
	goto LEVEL_UP
}

func (d *data_Logic) IncreaseAttributeLevel(k int64) {
	if _, has := d.attributesLevel[k]; !has {
		panic(fmt.Sprintf("attribute key %v not exists", k))
	}
	release := d.level
	for _, level := range d.attributesLevel {
		release -= level
	}
	if release <= 0 {
		return
	}
	d.attributesLevel[k]++
}

func (d *data_Logic) ToJSON() []byte {
	jsonBytes, err := json.MarshalIndent(&json_Logic{
		Exp:             d.exp,
		Level:           d.level,
		AttributesLevel: d.attributesLevel,
		IdleTS:          d.idleTS,
	}, "", "  ")
	if err != nil {
		panic(err)
	}

	return jsonBytes
}

func (d *data_Logic) LoadJSON() {
	jsonLogic, err := stp.ReadFile(d.saveDataPath, func(b []byte) (*json_Logic, error) {
		jsonLogic := &json_Logic{}
		err := json.Unmarshal(b, jsonLogic)
		if err != nil {
			return nil, err
		}
		return jsonLogic, nil
	})

	if err != nil {
		panic(err)
	}

	if jsonLogic.Level == 0 || jsonLogic.IdleTS == 0 {
		panic("empty json data")
	}

	d.exp = jsonLogic.Exp
	d.level = jsonLogic.Level
	for attribute, level := range jsonLogic.AttributesLevel {
		if _, has := d.attributesLevel[attribute]; has {
			d.attributesLevel[attribute] = level
		}
	}
	d.idleTS = jsonLogic.IdleTS
}
