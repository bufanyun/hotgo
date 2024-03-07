<template>
  <n-card
    :segmented="{ content: true, footer: true }"
    header-style="padding:10px"
    footer-style="padding:10px"
  >
    <template #header> 瀑布流组件 </template>
    <template #header-extra>
      <n-tag type="success" @click="handleClick">vue-waterfall-plugin-next</n-tag>
    </template>
    <Waterfall
      :row-key="options.rowKey"
      :gutter="options.gutter"
      :has-around-gutter="options.hasAroundGutter"
      :width="options.width"
      :breakpoints="options.breakpoints"
      :img-selector="options.imgSelector"
      :background-color="options.backgroundColor"
      :animation-effect="options.animationEffect"
      :animation-duration="options.animationDuration"
      :animation-delay="options.animationDelay"
      :lazyload="options.lazyload"
      :list="compData.items"
    >
      <template #item="{ item }">
        <div class="waterfall-item" :style="item.style">
          <span>君问归期未有期，巴山夜雨涨秋池。何当共剪西窗烛，却话巴山夜雨时。</span>
        </div>
      </template>
    </Waterfall>
  </n-card>
</template>

<script lang="ts" setup>
  import { reactive } from 'vue';
  import { Waterfall } from 'vue-waterfall-plugin-next';
  import 'vue-waterfall-plugin-next/dist/style.css';
  import { rdmLightRgbColor } from '@/utils/hotgo';

  const options = reactive({
    // 唯一key值
    rowKey: 'id',
    // 卡片之间的间隙
    gutter: 20,
    // 是否有周围的gutter
    hasAroundGutter: true,
    // 卡片在PC上的宽度
    width: 250,
    // 自定义行显示个数，主要用于对移动端的适配
    breakpoints: {
      1200: {
        // 当屏幕宽度小于等于1200
        rowPerView: 5,
      },
      800: {
        // 当屏幕宽度小于等于800
        rowPerView: 3,
      },
      500: {
        // 当屏幕宽度小于等于500
        rowPerView: 1,
      },
    },
    // 动画效果
    animationEffect: 'animate__fadeInUp',
    // 动画时间
    animationDuration: 1000,
    // 动画延迟
    animationDelay: 300,
    // 背景色
    backgroundColor: '',
    // imgSelector
    imgSelector: 'src.original',
    // 是否懒加载
    lazyload: true,
  });

  const items: any = [];
  const genBetweenRight = (m, n) => Math.floor(Math.random() * (n - m) + 1) + m;
  for (let i = 0; i < 90; i++) {
    items.push({
      style: {
        height: genBetweenRight(100, 300) + 'px',
        'background-color': rdmLightRgbColor(),
      },
    });
  }
  const compData = reactive({
    items,
  });
  const handleClick = () => {
    window.open('https://vue-waterfall.netlify.app/');
  };
</script>

<style lang="less" scoped>
  .waterfall-item {
    border: 2px solid rgb(244, 244, 248);
    height: 100px;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 10px;
    border-radius: 4px;
  }
</style>
