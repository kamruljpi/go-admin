package paginator

import (
	"testing"

	"github.com/kamruljpi/go-admin/modules/config"
	"github.com/kamruljpi/go-admin/plugins/admin/modules/parameter"
	_ "github.com/kamruljpi/themes/sword"
)

func TestGet(t *testing.T) {
	config.Initialize(&config.Config{Theme: "sword"})
	Get(Config{
		Size:         105,
		Param:        parameter.BaseParam().SetPage("7"),
		PageSizeList: []string{"10", "20", "50", "100"},
	})
}
