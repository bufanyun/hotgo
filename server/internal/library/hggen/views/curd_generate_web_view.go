// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package views

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

func (l *gCurd) webViewTplData(ctx context.Context, in *CurdPreviewInput) (data g.Map, err error) {
	data = make(g.Map)
	data["item"] = l.generateWebViewItem(ctx, in)
	return
}

func (l *gCurd) generateWebViewItem(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	for _, field := range in.masterFields {
		if !field.IsEdit {
			continue
		}

		var (
			defaultComponent = fmt.Sprintf("<n-descriptions-item>\n          <template #label>%s</template>\n          {{ formValue.%s }}\n        </n-descriptions-item>", field.Dc, field.TsName)
			component        string
		)

		switch field.FormMode {

		case FormModeInputTextarea, FormModeInputEditor:
			component = fmt.Sprintf("<n-descriptions-item>\n          <template #label>%s</template>\n          <span v-html=\"formValue.%s\"></span></n-descriptions-item>", field.Dc, field.TsName)

		case FormModeInputDynamic:
			component = defaultComponent

		case FormModeDate:
			component = defaultComponent

		case FormModeTime:
			component = defaultComponent

		case FormModeRadio, FormModeCheckbox, FormModeSelect, FormModeSelectMultiple:
			component = fmt.Sprintf("<n-descriptions-item label=\"%s\">\n          <template v-for=\"(item, key) in formValue?.%s\" :key=\"key\">\n            <n-tag\n              :type=\"getOptionTag(options.%s, item)\"\n              size=\"small\"\n              class=\"min-left-space\"\n              >{{ getOptionLabel(options.%s, item) }}</n-tag\n            >\n          </template>\n        </n-descriptions-item>", field.Dc, field.TsName, in.options.dictMap[field.TsName], in.options.dictMap[field.TsName])

		case FormModeUploadImage:
			component = fmt.Sprintf("<n-descriptions-item>\n          <template #label>%s</template>\n          <n-image style=\"margin-left: 10px; height: 100px; width: 100px\" :src=\"formValue.%s\"\n        /></n-descriptions-item>", field.Dc, field.TsName)

		case FormModeUploadImages:
			component = fmt.Sprintf("<n-descriptions-item>\n          <template #label>%s</template>\n          <n-image-group>\n            <n-space>\n              <span v-for=\"(item, key) in formValue?.%s\" :key=\"key\">\n                <n-image style=\"margin-left: 10px; height: 100px; width: 100px\" :src=\"item\" />\n              </span>\n            </n-space>\n          </n-image-group>\n        </n-descriptions-item>", field.Dc, field.TsName)

		case FormModeUploadFile:
			component = fmt.Sprintf("<n-descriptions-item>\n          <template #label>%s</template>\n          <div\n            class=\"upload-card\"\n            v-show=\"formValue.%s !== ''\"\n            @click=\"download(formValue.%s)\"\n          >\n            <div class=\"upload-card-item\" style=\"height: 100px; width: 100px\">\n              <div class=\"upload-card-item-info\">\n                <div class=\"img-box\">\n                  <n-avatar :style=\"fileAvatarCSS\">{{ getFileExt(formValue.%s) }}</n-avatar>\n                </div>\n              </div>\n            </div>\n          </div>\n        </n-descriptions-item>", field.Dc, field.TsName, field.TsName, field.TsName)

		case FormModeUploadFiles:
			component = fmt.Sprintf("<n-descriptions-item>\n          <template #label>%s</template>\n          <div class=\"upload-card\">\n            <n-space style=\"gap: 0px 0px\">\n              <div\n                class=\"upload-card-item\"\n                style=\"height: 100px; width: 100px\"\n                v-for=\"(item, key) in formValue.%s\"\n                :key=\"key\"\n              >\n                <div class=\"upload-card-item-info\">\n                  <div class=\"img-box\">\n                    <n-avatar :style=\"fileAvatarCSS\" @click=\"download(item)\">{{\n                      getFileExt(item)\n                    }}</n-avatar>\n                  </div>\n                </div>\n              </div>\n            </n-space>\n          </div>\n        </n-descriptions-item>", field.Dc, field.TsName)

		case FormModeSwitch:
			component = fmt.Sprintf("<n-descriptions-item label=\"%s\">\n          <n-switch v-model:value=\"formValue.%s\" :unchecked-value=\"2\" :checked-value=\"1\" :disabled=\"true\"\n        /></n-descriptions-item>", field.Dc, field.TsName)

		case FormModeRate:
			component = fmt.Sprintf("<n-descriptions-item label=\"%s\"\n          ><n-rate readonly :default-value=\"formValue.%s\"\n        /></n-descriptions-item>", field.Dc, field.TsName)

		default:
			component = defaultComponent
		}

		buffer.WriteString("        " + component + "\n\n")
	}
	return buffer.String()
}
