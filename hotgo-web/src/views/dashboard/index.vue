<template>
  <div id="content" style="margin: -8px;">
    <common-use/>
    <about/>
    <to-do/>
  </div>
</template>

<script>
  import CommonUse from './portal/CommonUse.vue'
   import ToDo from './portal/ToDo.vue'
import About from './portal/About.vue'
  export default {
    components: {
     CommonUse,
     ToDo,
     About
    },
    data () {
      return {

      }
    },
    computed: {

    },
    created () {
    },
    methods: {
    }
  }
</script>
