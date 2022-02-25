<template>
  <global-footer class="footer custom-render">
    <template v-slot:links>
      <!-- <a target="_blank">Gitee</a>
      <a target="_blank">Github</a> -->
    </template>
    <template v-slot:copyright>
      <a target="_blank">Hot Go!</a>
    </template>
  </global-footer>
</template>
<script>
import { GlobalFooter } from '@/components/ProLayout'

export default {
  name: 'ProGlobalFooter',
  components: {
    GlobalFooter
  }
}
</script>
