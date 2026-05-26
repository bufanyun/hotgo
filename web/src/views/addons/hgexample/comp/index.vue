<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <n-p class="title">组件示例</n-p>
      <n-space vertical>
        <n-layout has-sider>
          <n-scrollbar :style="scrollbarStyle">
            <n-layout-sider
              bordered
              collapse-mode="width"
              :collapsed-width="64"
              :width="180"
              :collapsed="collapsed"
              show-trigger
              @collapse="collapsed = true"
              @expand="collapsed = false"
            >
              <n-menu
                v-model:value="activeKey"
                :collapsed="collapsed"
                :collapsed-width="64"
                :collapsed-icon-size="22"
                :options="menuOptions"
              />
            </n-layout-sider>
          </n-scrollbar>

          <n-layout>
            <n-card :bordered="false" class="proCard">
              <component :is="currentComponent" />
            </n-card>
          </n-layout>
        </n-layout>
      </n-space>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import type { MenuOption } from 'naive-ui';
  import {
    ProfileOutlined,
    BorderOutlined,
    NumberOutlined,
    BorderOuterOutlined,
    CheckCircleOutlined,
    ExclamationCircleOutlined,
    FundProjectionScreenOutlined,
    ReadOutlined,
    StarOutlined,
    PushpinOutlined,
    LayoutOutlined,
    PrinterOutlined,
    DownloadOutlined,
  } from '@vicons/antd';
  import {
    HourglassOutline,
    NotificationsOutline,
    LocationOutline,
    SwapVertical,
    EllipsisHorizontalCircleOutline,
    AlbumsOutline,
    LogoChrome,
  } from '@vicons/ionicons5';
  import { renderIcon } from '@/utils';
  import FormBasic from './form/basic.vue';
  import FormUseForm from './form/useForm.vue';
  import Modal from './modal/index.vue';
  import Drag from './drag/index.vue';
  import Directive from './directive/index.vue';
  import ResultSuccess from './result/success.vue';
  import ResultFail from './result/fail.vue';
  import ResultInfo from './result/info.vue';
  import Exception403 from '../../../exception/403.vue';
  import Exception404 from '../../../exception/404.vue';
  import Exception500 from '../../../exception/500.vue';
  import VisPPchart from './vis/ppchart.vue';
  import VisEcharts from './vis/echarts.vue';
  import VisMadeAPie from './vis/madeapie.vue';
  import TextPinyin from './text/pinyin/index.vue';
  import TextMint from './text/mint/index.vue';
  import TextGradient from './text/gradient/index.vue';
  import TextHigh from './text/high/index.vue';
  import IconsAntd from './icons/antd.vue';
  import IconsIonicons5 from './icons/ionicons5.vue';
  import IconsSelector from './icons/selector.vue';
  import Watermark from './watermark/index.vue';
  import Des from './des/index.vue';
  import Calendar from './calendar/index.vue';
  import Timeline from './timeline/index.vue';
  import Notice from './notice/index.vue';
  import MapGaode from './map/gaode.vue';
  import MapBaidu from './map/baidu.vue';
  import Print from './print/index.vue';
  import TagsView from './tagsView/index.vue';
  import MoreComponents from './moreComponents/index.vue';
  import Waterfall from './waterfall/index.vue';
  import ImportExcel from './import/excel.vue';
  import FingerPrintJs from './fingerprintjs/index.vue';

  const components = {
    formBasic: FormBasic,
    formUseForm: FormUseForm,
    modal: Modal,
    drag: Drag,
    directive: Directive,
    resultSuccess: ResultSuccess,
    resultFail: ResultFail,
    resultInfo: ResultInfo,
    exception403: Exception403,
    exception404: Exception404,
    exception500: Exception500,
    visPPchart: VisPPchart,
    visEcharts: VisEcharts,
    visMadeAPie: VisMadeAPie,
    textPinyin: TextPinyin,
    textMint: TextMint,
    textGradient: TextGradient,
    textHigh: TextHigh,
    iconsAntd: IconsAntd,
    iconsIonicons5: IconsIonicons5,
    iconsSelector: IconsSelector,
    watermark: Watermark,
    des: Des,
    calendar: Calendar,
    timeline: Timeline,
    notice: Notice,
    mapGaode: MapGaode,
    mapBaidu: MapBaidu,
    print: Print,
    tagsView: TagsView,
    moreComponents: MoreComponents,
    waterfall: Waterfall,
    importExcel: ImportExcel,
    fingerPrintJs: FingerPrintJs,
  };

  const activeKey = ref<string>('formBasic');
  const collapsed = ref(false);

  const scrollbarStyle = computed(() => {
    const height = '82vh';
    const width = collapsed.value ? '74px' : '188px';
    return {
      height: height,
      'max-height': height,
      'min-width': width,
      width: width,
    };
  });

  const currentComponent = computed(() => {
    return components[activeKey.value] || null;
  });

  const menuOptions: MenuOption[] = [
    {
      label: '表单',
      key: 'form',
      icon: renderIcon(ProfileOutlined),
      children: [
        {
          label: '基础使用',
          key: 'formBasic',
        },
        {
          label: 'useForm',
          key: 'formUseForm',
        },
      ],
    },
    {
      label: '图标',
      key: 'icons',
      icon: renderIcon(StarOutlined),
      children: [
        {
          label: 'Antd',
          key: 'iconsAntd',
        },
        {
          label: 'IonIcons5',
          key: 'iconsIonicons5',
        },
        {
          label: '选择器',
          key: 'iconsSelector',
        },
      ],
    },
    {
      label: '文字处理',
      key: 'text',
      icon: renderIcon(ReadOutlined),
      children: [
        {
          label: '汉字拼音',
          key: 'textPinyin',
        },
        {
          label: '敏感词汇',
          key: 'textMint',
        },
        {
          label: '渐变文字',
          key: 'textGradient',
        },
        {
          label: '文字高亮',
          key: 'textHigh',
        },
      ],
    },
    {
      label: '可视化',
      key: 'vis',
      icon: renderIcon(FundProjectionScreenOutlined),
      children: [
        {
          label: 'PPChart',
          key: 'visPPchart',
        },
        {
          label: 'Echarts',
          key: 'visEcharts',
        },
        {
          label: 'MadeAPie',
          key: 'visMadeAPie',
        },
      ],
    },
    {
      label: '异常页面',
      key: 'exception',
      icon: renderIcon(ExclamationCircleOutlined),
      children: [
        {
          label: '403',
          key: 'exception403',
        },
        {
          label: '404',
          key: 'exception404',
        },
        {
          label: '500',
          key: 'exception500',
        },
      ],
    },
    {
      label: '结果页面',
      key: 'result',
      icon: renderIcon(CheckCircleOutlined),
      children: [
        {
          label: '成功页',
          key: 'resultSuccess',
        },
        {
          label: '失败页',
          key: 'resultFail',
        },
        {
          label: '信息页',
          key: 'resultInfo',
        },
      ],
    },
    {
      label: '地图',
      key: 'map',
      icon: renderIcon(LocationOutline),
      children: [
        {
          label: '高德地图',
          key: 'mapGaode',
        },
        {
          label: '百度地图',
          key: 'mapBaidu',
        },
      ],
    },
    {
      label: '弹窗扩展',
      key: 'modal',
      icon: renderIcon(BorderOutlined),
    },
    {
      label: '消息提示',
      key: 'notice',
      icon: renderIcon(NotificationsOutline),
    },
    {
      label: '瀑布流',
      key: 'waterfall',
      icon: renderIcon(AlbumsOutline),
    },
    {
      label: '拖拽',
      key: 'drag',
      icon: renderIcon(NumberOutlined),
    },
    {
      label: '水印',
      key: 'watermark',
      icon: renderIcon(PushpinOutlined),
    },
    {
      label: '日历',
      key: 'calendar',
      icon: renderIcon(ReadOutlined),
    },
    {
      label: '打印',
      key: 'print',
      icon: renderIcon(PrinterOutlined),
    },

    {
      label: '时间线',
      key: 'timeline',
      icon: renderIcon(HourglassOutline),
    },
    {
      label: 'tagsView',
      key: 'tagsView',
      icon: renderIcon(SwapVertical),
    },
    {
      label: '导入Excel',
      key: 'importExcel',
      icon: renderIcon(DownloadOutlined),
    },
    {
      label: '卡片描述',
      key: 'des',
      icon: renderIcon(LayoutOutlined),
    },
    {
      label: '指令示例',
      key: 'directive',
      icon: renderIcon(BorderOuterOutlined),
    },
    {
      label: '浏览器指纹',
      key: 'fingerPrintJs',
      icon: renderIcon(LogoChrome),
    },
    {
      label: '更多组件',
      key: 'moreComponents',
      icon: renderIcon(EllipsisHorizontalCircleOutline),
    },
  ];
</script>

<style lang="less" scoped>
  .title {
    font-size: 18px;
    transition: color 0.3s var(--n-bezier);
    flex: 1;
    min-width: 0;
    color: var(--n-title-text-color);
  }
</style>
