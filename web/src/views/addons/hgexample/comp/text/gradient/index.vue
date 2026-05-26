<template>
  <div>
    <n-space vertical>
      <n-card
        :segmented="{ content: true, footer: true }"
        header-style="padding:10px"
        footer-style="padding:10px"
      >
        <template #header> 渐变文字 </template>
        <template #header-extra>
          <n-gradient-text type="success"> 渐变文字成了 </n-gradient-text>
        </template>
        <n-space vertical>
          <n-gradient-text style="white-space: normal" type="success">
            <p
              >苏轼，（1037年1月8日-1101年8月24日）字子瞻、和仲，号铁冠道人、东坡居士，世称苏东坡、苏仙，汉族，眉州眉山（四川省眉山市）人，祖籍河北栾城，北宋著名文学家、书法家、画家，历史治水名人。苏轼是北宋中期文坛领袖，在诗、词、散文、书、画等方面取得很高成就。文纵横恣肆；诗题材广阔，清新豪健，善用夸张比喻，独具风格，与黄庭坚并称“苏黄”；词开豪放一派，与辛弃疾同是豪放派代表，并称“苏辛”；散文著述宏富，豪放自如，与欧阳修并称“欧苏”，为“唐宋八大家”之一。苏轼善书，“宋四家”之一；擅长文人画，尤擅墨竹、怪石、枯木等。与韩愈、柳宗元和欧阳修合称“千古文章四大家”。作品有《东坡七集》《东坡易传》《东坡乐府》《潇湘竹石图卷》《古木怪石图卷》等</p
            >
          </n-gradient-text>
        </n-space>
      </n-card>
      <n-card
        :segmented="{ content: true, footer: true }"
        header-style="padding:10px"
        footer-style="padding:10px"
      >
        <template #header> 渐变文字 </template>
        <template #header-extra>
          <n-gradient-text type="warning"> 渐变文字成了 </n-gradient-text>
        </template>
        <n-space vertical>
          <n-gradient-text style="white-space: normal" type="warning">
            <p
              >白居易（772年－846年），字乐天，号香山居士，又号醉吟先生，祖籍太原，到其曾祖父时迁居下邽，生于河南新郑。是唐代伟大的现实主义诗人，唐代三大诗人之一。白居易与元稹共同倡导新乐府运动，世称“元白”，与刘禹锡并称“刘白”。白居易的诗歌题材广泛，形式多样，语言平易通俗，有“诗魔”和“诗王”之称。官至翰林学士、左赞善大夫。公元846年，白居易在洛阳逝世，葬于香山。有《白氏长庆集》传世，代表诗作有《长恨歌》《卖炭翁》《琵琶行》等。</p
            >
          </n-gradient-text>
        </n-space>
      </n-card>
      <n-card
        :segmented="{ content: true, footer: true }"
        header-style="padding:10px"
        footer-style="padding:10px"
      >
        <template #header> 自定义颜色 </template>
        <template #header-extra>
          <n-gradient-text gradient="linear-gradient(90deg, red 0%, green 50%, blue 100%)">
            渐变文字成了
          </n-gradient-text>
        </template>
        <n-space vertical>
          <n-gradient-text
            gradient="linear-gradient(90deg, red 0%, green 50%, blue 100%)"
            style="white-space: normal"
          >
            <p
              >杜甫（712－770），字子美，自号少陵野老，世称“杜工部”、“杜少陵”等，汉族，河南府巩县（今河南省巩义市）人，唐代伟大的现实主义诗人，杜甫被世人尊为“诗圣”，其诗被称为“诗史”。杜甫与李白合称“李杜”，为了跟另外两位诗人李商隐与杜牧即“小李杜”区别开来，杜甫与李白又合称“大李杜”。他忧国忧民，人格高尚，他的约1400余首诗被保留了下来，诗艺精湛，在中国古典诗歌中备受推崇，影响深远。759-766年间曾居成都，后世有杜甫草堂纪念</p
            >
          </n-gradient-text>
        </n-space>
      </n-card>
    </n-space>
  </div>
</template>
<script lang="ts" setup></script>
<style lang="less" scoped>
  p {
    font-size: 14px;
    line-height: 2;
    text-indent: 28px;
  }
</style>
