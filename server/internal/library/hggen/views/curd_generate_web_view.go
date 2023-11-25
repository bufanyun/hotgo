// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
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
		// if !field.IsEdit {
		// 	continue
		// }

		var (
			defaultComponent = fmt.Sprintf("      <div class=\"item\">\n                <div>%s：</div>\n                <div class=\"value\">{{ params.%s }}</div>\n              </div>", field.Dc, field.TsName)
			component        string
		)

		switch field.FormMode {
		case FormModeInputTextarea, FormModeInputEditor:
			component = fmt.Sprintf("      <div class=\"item\">\n                <div>%s：</div>\n                <div class=\"value\">\n                  <span v-html=\"params.%s\"></span>\n                </div>\n              </div>", field.Dc, field.TsName)

		case FormModeInputDynamic:
			component = defaultComponent

		case FormModeDate:
			component = defaultComponent

		case FormModeTime:
			component = defaultComponent

		case FormModeRadio, FormModeSelect:
			component = fmt.Sprintf("      <div class=\"item\">\n                <div>%s：</div>\n                <div class=\"value\">\n                  <n-tag\n                    :type=\"getOptionTag(options.%s, params?.%s)\"\n                    size=\"small\"\n                    class=\"min-left-space\"\n                    >{{ getOptionLabel(options.%s, params?.%s) }}</n-tag\n                  >\n                </div>\n              </div>", field.Dc, in.options.dictMap[field.TsName], field.TsName, in.options.dictMap[field.TsName], field.TsName)

		case FormModeCheckbox, FormModeSelectMultiple:
			component = fmt.Sprintf("      <div class=\"item\">\n                <div>%s：</div>\n                <div class=\"value\">\n                  <template v-for=\"(item, key) in params?.%s\" :key=\"key\">\n                    <n-tag\n                      :type=\"getOptionTag(options.%s, item)\"\n                      size=\"small\"\n                      class=\"min-left-space\"\n                      >{{ getOptionLabel(options.%s, item) }}\n                    </n-tag>\n                  </template>\n                </div>\n              </div>", field.Dc, field.TsName, in.options.dictMap[field.TsName], in.options.dictMap[field.TsName])

		case FormModeUploadImage:
			component = fmt.Sprintf("      <div class=\"item\">\n                <div>%s：</div>\n                <div class=\"value\">\n                  <n-image\n                    style=\"margin-left: 10px; height: 100px; width: 100px\"\n                    :src=\"params.%s\"\n                  />\n                </div>\n              </div>", field.Dc, field.TsName)

		case FormModeUploadImages:
			component = fmt.Sprintf("      <div class=\"item\">\n                <div>%s：</div>\n                <div class=\"value\">\n                  <n-image-group>\n                    <n-space>\n                      <span v-for=\"(item, key) in params?.%s\" :key=\"key\">\n                        <n-image\n                          style=\"margin-left: 10px; height: 100px; width: 100px\"\n                          :src=\"item\"\n                        />\n                      </span>\n                    </n-space>\n                  </n-image-group>\n                </div>\n              </div>", field.Dc, field.TsName)

		case FormModeUploadFile:
			component = fmt.Sprintf("      <div class=\"item\">\n                <div>%s：</div>\n                <div class=\"value\"\n                  ><div\n                    class=\"upload-card\"\n                    v-show=\"params.attachfile !== ''\"\n                    @click=\"download(params.%s)\"\n                  >\n                    <div class=\"upload-card-item\" style=\"height: 100px; width: 100px\">\n                      <div class=\"upload-card-item-info\">\n                        <div class=\"img-box\">\n                          <n-avatar :style=\"fileAvatarCSS\">{{\n                            getFileExt(params.attachfile)\n                          }}</n-avatar>\n                        </div>\n                      </div>\n                    </div>\n                  </div></div\n                >\n              </div>", field.TsName, field.TsName)

		case FormModeUploadFiles:
			component = fmt.Sprintf("      <div class=\"item\">\n                <div>%s：</div>\n                <div class=\"value\">\n                  <div class=\"upload-card\">\n                    <n-space style=\"gap: 0px 0px\">\n                      <div\n                        class=\"upload-card-item\"\n                        style=\"height: 100px; width: 100px\"\n                        v-for=\"(item, key) in params.%s\"\n                        :key=\"key\"\n                      >\n                        <div class=\"upload-card-item-info\">\n                          <div class=\"img-box\">\n                            <n-avatar :style=\"fileAvatarCSS\" @click=\"download(item)\">{{\n                              getFileExt(item)\n                            }}</n-avatar>\n                          </div>\n                        </div>\n                      </div>\n                    </n-space>\n                  </div>\n                </div>\n              </div>", field.Dc, field.TsName)

		case FormModeSwitch:
			component = fmt.Sprintf("      <div class=\"item\">\n                <div>%s：</div>\n                <div class=\"value\">\n                  <n-switch\n                    v-model:value=\"params.%s\"\n                    :unchecked-value=\"2\"\n                    :checked-value=\"1\"\n                    :disabled=\"true\"\n                  />\n                </div>\n              </div>", field.Dc, field.TsName)

		case FormModeRate:
			component = fmt.Sprintf("      <div class=\"item\">\n                <div>%s：</div>\n                <div class=\"value\">\n                  <n-rate readonly :default-value=\"params.%s\" />\n                </div>\n              </div>", field.Dc, field.TsName)

		default:
			component = defaultComponent
		}

		buffer.WriteString("        " + component + "\n\n")
	}
	return buffer.String()
}
