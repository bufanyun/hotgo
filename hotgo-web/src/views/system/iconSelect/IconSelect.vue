<template>
  <icon-detail :svgIcons="iconList" :allIcon="allIcon"/>
</template>
<script>
import IconDetail from '@/components/IconSelector/IconDetail'
import allIcon from '@/core/icons'
import icons from '@/utils/requireIcons'
export default {
  name: 'IconSelect',
  components: {
    IconDetail
  },
  data () {
    return {
      allIcon,
       iconList: icons
    }
  }
}
</script>

<style></style>
