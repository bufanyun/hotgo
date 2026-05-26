<template>
  <div>
    <n-card title="高亮" style="margin-bottom: 10px">
      <n-text>
        <p style="line-height: 2" v-html="compData1.text"></p>
      </n-text>
    </n-card>
    <n-card title="自定义颜色" style="margin-bottom: 10px">
      <n-text>
        <p style="line-height: 2" v-html="compData2.text"></p>
      </n-text>
    </n-card>
    <n-card title="自定义样式">
      <n-text>
        <p style="line-height: 2" v-html="compData3.text"></p>
      </n-text>
    </n-card>
  </div>
</template>
<script lang="ts" setup>
  import highHtml from '@/utils/highHtml';
  import { reactive, onMounted } from 'vue';

  const text =
    '先帝创业未半而中道崩殂，今天下三分，益州疲弊，此诚危急存亡之秋也。然侍卫之臣不懈于内，忠志之士忘身于外者，盖追先帝之殊遇，欲报之于陛下也。诚宜开张圣听，以光先帝遗德，恢弘志士之气，不宜妄自菲薄，引喻失义，以塞忠谏之路也。\n' +
    '\n' +
    '　　宫中府中，俱为一体；陟罚臧否，不宜异同。若有作奸犯科及为忠善者，宜付有司论其刑赏，以昭陛下平明之理，不宜偏私，使内外异法也。\n' +
    '\n' +
    '　　侍中、侍郎郭攸之、费祎、董允等，此皆良实，志虑忠纯，是以先帝简拔以遗陛下。愚以为宫中之事，事无大小，悉以咨之，然后施行，必能裨补阙漏，有所广益。\n' +
    '\n' +
    '　　将军向宠，性行淑均，晓畅军事，试用于昔日，先帝称之曰能，是以众议举宠为督。愚以为营中之事，悉以咨之，必能使行阵和睦，优劣得所。\n' +
    '\n' +
    '　　亲贤臣，远小人，此先汉所以兴隆也；亲小人，远贤臣，此后汉所以倾颓也。先帝在时，每与臣论此事，未尝不叹息痛恨于桓、灵也。侍中、尚书、长史、参军，此悉贞良死节之臣，愿陛下亲之信之，则汉室之隆，可计日而待也。\n' +
    '\n' +
    '　　臣本布衣，躬耕于南阳，苟全性命于乱世，不求闻达于诸侯。先帝不以臣卑鄙，猥自枉屈，三顾臣于草庐之中，咨臣以当世之事，由是感激，遂许先帝以驱驰。后值倾覆，受任于败军之际，奉命于危难之间，尔来二十有一年矣。\n' +
    '\n' +
    '　　先帝知臣谨慎，故临崩寄臣以大事也。受命以来，夙夜忧叹，恐托付不效，以伤先帝之明；故五月渡泸，深入不毛。今南方已定，兵甲已足，当奖率三军，北定中原，庶竭驽钝，攘除奸凶，兴复汉室，还于旧都。此臣所以报先帝而忠陛下之职分也。至于斟酌损益，进尽忠言，则攸之、祎、允之任也。\n' +
    '\n' +
    '　　愿陛下托臣以讨贼兴复之效，不效，则治臣之罪，以告先帝之灵。若无兴德之言，则责攸之、祎、允等之慢，以彰其咎；陛下亦宜自谋，以咨诹善道，察纳雅言，深追先帝遗诏。臣不胜受恩感激。今当远离，临表涕零，不知所言';

  const compData1 = reactive({
    text: '',
    keys: ['先帝', '先汉', '苟'],
  });
  const compData2 = reactive({
    text: '',
    keys: [
      {
        keyword: '先帝',
        bgColor: 'red',
        color: '#fff',
        renderHighlightKeyword: (str) => {
          return `<span class="high" style="padding:2px 5px;margin:0 5px;border-radius:3px;background-color: #3a5ccc;color: #ffffff">${str}</span>`;
        },
      },
      {
        keyword: '中',
        bgColor: '#b4d5ff',
        color: '#fff',
        renderHighlightKeyword: (str) => {
          return `<span class="high" style="padding:2px 5px;margin:0 5px;border-radius:3px;background-color: #e96656;color: #ffffff">${str}</span>`;
        },
      },
      { keyword: '苟', bgColor: '#b4d5ff', color: '#fff' },
    ],
  });
  const compData3 = reactive({
    text: '',
    keys: [
      {
        keyword: '之',
        bgColor: 'red',
        color: '#fff',
        style: { padding: '2px 5px', margin: '0 5px', 'border-radius': '3px' },
      },
      { keyword: '先帝', bgColor: '#b4d5ff', color: '#fff' },
      { keyword: '苟', bgColor: '#b4d5ff', color: '#fff' },
    ],
  });

  onMounted(() => {
    compData1.text = highHtml(text, compData1.keys).highText;
    compData2.text = highHtml(text, compData2.keys).highText;
    compData3.text = highHtml(text, compData3.keys).highText;
  });
</script>
<style lang="less">
  .high {
    &:hover {
      background-color: seagreen !important;
      cursor: pointer;
    }
  }
</style>
