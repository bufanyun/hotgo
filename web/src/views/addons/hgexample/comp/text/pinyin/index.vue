<template>
  <div>
    <n-space vertical>
      <n-card
        :segmented="{ content: true, footer: true }"
        header-style="padding:10px"
        footer-style="padding:10px"
      >
        <template #header> 字符拼音 </template>
        <template #header-extra> 陋室铭，唐代：刘禹锡 </template>
        <div class="pinyin">
          <template v-for="item in compData.items1">
            <div class="pinyin-item" :key="idx" v-for="(todo, idx) in item">
              <span class="py">{{ todo.py }}</span>
              <span class="hz">{{ todo.hz }}</span> </div
            ><br />
          </template>
        </div>
      </n-card>
      <n-card
        :segmented="{ content: true, footer: true }"
        header-style="padding:10px"
        footer-style="padding:10px"
      >
        <template #header> 字符拼音 </template>
        <template #header-extra> 李贺小传，李商隐〔唐代〕 </template>
        <div class="pinyin">
          <template v-for="item in compData.items2">
            <div class="pinyin-item" :key="idx" v-for="(todo, idx) in item">
              <span class="py">{{ todo.py }}</span>
              <span class="hz">{{ todo.hz }}</span> </div
            ><br />
          </template>
        </div>
      </n-card>
    </n-space>
  </div>
</template>
<script lang="ts" setup>
  import { reactive } from 'vue';
  import { pinyin } from 'pinyin-pro';

  interface WordInfo {
    hz: string;
    py: string;
  }

  const text1 = [
    '  山不在高，有仙则名。水不在深，有龙则灵。斯是陋室，惟吾德馨。苔痕上阶绿，草色入帘青。谈笑有鸿儒，往来无白丁。可以调素琴，阅金经。无丝竹之乱耳，无案牍之劳形。南阳诸葛庐，西蜀子云亭。孔子云：何陋之有？',
  ];
  const text2 = [
    '  京兆杜牧为李长吉集序，状长吉之奇甚尽，世传之。长吉姊嫁王氏者，语长吉之事尤备。',
    '  长吉细瘦，通眉，长指爪，能苦吟疾书。最先为昌黎韩愈所知。所与游者，王参元、杨敬之、权璩、崔植辈为密，每旦日出与诸公游，未尝得题然后为诗，如他人思量牵合，以及程限为意。恒从小奚奴，骑距驉，背一古破锦囊，遇有所得，即书投囊中。及暮归．太夫人使婢受囊出之，见所书多．辄曰：“是儿要当呕出心乃已尔。”上灯，与食。长吉从婢取书，研墨叠纸足成之，投他囊中。非大醉及吊丧日率如此，过亦不复省。王、杨辈时复来探取写去。长吉往往独骑往还京、洛，所至或时有著，随弃之，故沈子明家所余四卷而已。',
    '  长吉将死时，忽昼见一绯衣人，驾赤虬，持一板，书若太古篆或霹雳石文者，云当召长吉。长吉了不能读，欻下榻叩头，言：“阿㜷老且病，贺不愿去。”绯衣人笑曰：“帝成白玉楼，立召君为记。天上差乐，不苦也。”长吉独泣，边人尽见之。少之，长吉气绝。常所居窗中，勃勃有烟气，闻行车嘒管之声。太夫人急止人哭，待之如炊五斗黍许时，长吉竟死。王氏姊非能造作谓长吉者，实所见如此。',
    '  呜呼，天苍苍而高也，上果有帝耶?帝果有苑囿、宫室、观阁之玩耶?苟信然，则天之高邈，帝之尊严，亦宜有人物文采愈此世者，何独眷眷于长吉而使其不寿耶?噫，又岂世所谓才而奇者，不独地上少，即天上亦不多耶?长吉生二十七年，位不过奉礼太常，时人亦多排摈毁斥之，又岂才而奇者，帝独重之，而人反不重耶?又岂人见会胜帝耶?',
  ];

  const createHzPy = (text: string[]): WordInfo[][] => {
    const items: WordInfo[][] = [];
    text.forEach((item) => {
      const todo: WordInfo[] = [];
      for (let i = 0; i < item.length; i++) {
        const tg = item.charAt(i);
        todo.push({ hz: tg, py: pinyin(tg)[0] });
      }
      items.push(todo);
    });
    return items;
  };

  const compData = reactive({
    items1: createHzPy(text1),
    items2: createHzPy(text2),
  });
</script>
<style lang="less" scoped>
  .pinyin {
    &-item {
      line-height: 100%;
      width: 42px;
      text-align: center;
      display: inline-block;
      .py {
        clear: both;
        font-size: 12px;
        font-weight: normal;
        float: left;
        width: 42px;
      }
      .hz {
        clear: both;
        margin-bottom: 10px;
        text-align: center;
        float: left;
        font-size: 16px;
        height: 36px;
        width: 42px;
        line-height: 36px;
      }
    }
  }
</style>
