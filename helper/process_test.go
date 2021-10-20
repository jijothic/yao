package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/gou"
	"github.com/yaoapp/kun/maps"
)

func TestProcessArrayPluck(t *testing.T) {
	args := []interface{}{
		[]interface{}{"城市", "行业", "计费"},
		map[string]interface{}{
			"行业": map[string]interface{}{"key": "city", "value": "数量", "items": []map[string]interface{}{{"city": "北京", "数量": 32}, {"city": "上海", "数量": 20}}},
			"计费": map[string]interface{}{"key": "city", "value": "计费种类", "items": []map[string]interface{}{{"city": "北京", "计费种类": 6}, {"city": "西安", "计费种类": 3}}},
		},
	}
	process := gou.NewProcess("xiang.helper.ArrayPluck", args...)
	response := ProcessArrayPluck(process)
	assert.NotNil(t, response)
	items, ok := response.([]map[string]interface{})
	assert.True(t, ok)

	assert.Equal(t, 3, len(items))
	for _, item := range items {
		maps.Of(item).Has("城市")
		maps.Of(item).Has("行业")
		maps.Of(item).Has("计费")
	}
}